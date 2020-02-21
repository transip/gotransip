package authenticator

import (
	"fmt"
	"github.com/transip/gotransip/v6/jwt"
	"github.com/transip/gotransip/v6/rest"
	"github.com/transip/gotransip/v6/rest/request"
	"github.com/transip/gotransip/v6/rest/response"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	// this is the header key we will add the signature to
	SignatureHeader string = "Signature"
	// this prefix will be used to name tokens we requested
	// customers are able to see this in their control panel
	LabelPrefix = "gotransip-client"
	// default a requested token expires after a day
	defaultExpirationTime = 86400
)

type Authenticator struct {
	// this contains a []byte representation of the the private key of the customer
	// this key will be used to sign a new token request
	PrivateKeyBody []byte
	// this is token, that is filled with a static token that a customer provides
	// or a token that we got from a token request
	Token jwt.Token
	// this is the http client to do auth requests with
	HTTPClient *http.Client
	// this would be the auth path, thus where we will get new tokens from
	BasePath string
	// this would be the acount name of customer
	Login string
	// When this is set to true the requested tokens can only be used with the 'ip' we are requesting with
	Whitelisted bool
	// Whether or not we want to request read only Tokens, that can only only be used to retrieve information
	// not to create, modify or delete it
	ReadOnly bool
}

// This is the auth request struct we fill and
// send in order to retrieve a token
// for more information, see: https://api.transip.nl/rest/docs.html#header-authentication
type AuthRequest struct {
	// Account name
	Login string `json:"login"`
	// Unique number for this request
	Nonce string `json:"nonce"`
	// Custom name to give this token, you can see your tokens in the transip control panel
	Label string `json:"label,omitempty"`
	// Enable read only mode
	ReadOnly bool `json:"read_only,omitempty"`
	// Unix time stamp of when this token should expire
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Whether this key can be used from everywhere, e.g should not be whitelisted to the current requesting ip
	GlobalKey bool `json:"global_key,omitempty"`
}

// todo: implement token caching to filesystem
func (a *Authenticator) GetToken() (jwt.Token, error) {
	if a.Token.Expired() {
		var err error
		a.Token, err = a.RequestNewToken()

		if err != nil {
			return jwt.Token{}, err
		}
	}

	return a.Token, nil
}

// this method will request a new token with the http client
// creating a new AuthRequest
func (a *Authenticator) RequestNewToken() (jwt.Token, error) {
	restRequest := a.GetRestAuthRequest()
	getMethod := rest.PostRestMethod

	httpRequest, err := restRequest.GetHttpRequest(a.BasePath, getMethod.Method)
	if err != nil {
		return jwt.Token{}, nil
	}
	bodyToSign, err := restRequest.GetBody()
	if err != nil {
		return jwt.Token{}, nil
	}
	signature, err := signWithKey(bodyToSign, a.PrivateKeyBody)
	if err != nil {
		return jwt.Token{}, err
	}
	httpRequest.Header.Add(SignatureHeader, signature)

	httpResponse, err := a.HTTPClient.Do(httpRequest)
	if err != nil {
		return jwt.Token{}, fmt.Errorf("request error:\n%s", err.Error())
	}

	defer httpResponse.Body.Close()

	// read entire response body
	b, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return jwt.Token{}, err
	}

	restResponse := response.RestResponse{
		Body:       b,
		StatusCode: httpResponse.StatusCode,
		Method:     getMethod,
	}

	var tokenToReturn TokenResponse
	err = restResponse.ParseResponse(&tokenToReturn)
	if err != nil {
		return jwt.Token{}, err
	}

	return jwt.New(tokenToReturn.Token)
}

type TokenResponse struct {
	Token string `json:"token"`
}

// Returns a random nonce each time it is called
func (a *Authenticator) GetNonce() string {
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)

	// convert to hex
	return fmt.Sprintf("%02x", randomBytes)
}

// this method generates the rest authentication request
func (a *Authenticator) GetRestAuthRequest() request.RestRequest {
	labelPostFix := time.Now().Unix()

	authRequest := AuthRequest{
		Login:          a.Login,
		Nonce:          a.GetNonce(),
		Label:          fmt.Sprintf("%s-%d", LabelPrefix, labelPostFix),
		ReadOnly:       a.ReadOnly,
		ExpirationTime: time.Now().Unix() + defaultExpirationTime,
		GlobalKey:      a.Whitelisted,
	}

	return request.RestRequest{
		Endpoint: "/auth",
		Body:     authRequest,
	}
}
