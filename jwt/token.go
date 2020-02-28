package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Token is the jwt that will be used by the client in the Authorization header
// to send in every request to the api server, every request except for the auth request
// which is used to request a new token
//
// for more information see: https://jwt.io/
type Token struct {
	ExpiryDate int64
	RawToken   string
}

// Expired returns true when the token expiry date is reached
// todo: race cond check +1h
func (t *Token) Expired() bool {
	return time.Now().Unix() > t.ExpiryDate
}

// GetAuthenticationHeaderValue returns the authentication header value value
// including the Bearer prefix
func (t *Token) GetAuthenticationHeaderValue() string {
	return fmt.Sprintf("Bearer %s", t.RawToken)
}

// tokenPayload is used to unpack the payload from the jwt
type tokenPayload struct {
	// This ExpirationTime is a 64 bit epoch that will be put into the token struct
	// that will later be used to validate if the token is expired or not
	// once expired we need to request a new token
	ExpirationTime int64 `json:"exp"`
}

// New expects a raw token as string
// It will try to decode it and return an error on error
// Once decoded it will retrieve the expiry date and
// return a Token struct with the RawToken and ExpiryDate set
func New(token string) (Token, error) {
	if len(token) == 0 {
		return Token{}, errors.New("No token given, a token should be set")
	}

	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return Token{}, errors.New("Invalid token given, token should exist at least of 3 parts")
	}

	jsonBody, err := base64.RawStdEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return Token{}, errors.New("Could not decode token, invalid base64")
	}

	var tokenRequest tokenPayload
	err = json.Unmarshal(jsonBody, &tokenRequest)
	if err != nil {
		return Token{}, errors.New("Could not read token body, invalid json")
	}

	return Token{
		RawToken:   token,
		ExpiryDate: tokenRequest.ExpirationTime,
	}, nil
}
