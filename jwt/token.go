package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Token struct {
	ExpiryDate int64
	RawToken   string
}

// Returns true when the token expiry date is met
func (t *Token) Expired() bool {
	return time.Now().Unix() > t.ExpiryDate
}

// returns the authentication header value value including the Bearer prefix
func (t *Token) GetAuthenticationHeaderValue() string {
	return fmt.Sprintf("Bearer %s", t.RawToken)
}

type TokenRequest struct {
	Login          string `json:"login:omitempty"`
	Nonce          string `json:"nonce:omitempty"`
	ReadOnly       string `json:"read_only:omitempty"`
	ExpirationTime int64  `json:"expiration_time:omitempty"`
	Label          string `json:"label:omitempty"`
	Global_key     string `json:"global_key:omitempty"`
}

type TokenBody struct {
	ExpirationTime int64 `json:"exp"`
}

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

	var tokenRequest TokenBody
	err = json.Unmarshal(jsonBody, &tokenRequest)
	if err != nil {
		return Token{}, errors.New("Could not read token body, invalid json")
	}

	return Token{
		RawToken:   token,
		ExpiryDate: tokenRequest.ExpirationTime,
	}, nil
}
