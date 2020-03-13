package authenticator

import (
	"errors"
	"fmt"
	"github.com/transip/gotransip/v6/jwt"
	"github.com/transip/gotransip/v6/rest"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	// this is the header key we will add the signature to
	signatureHeader = "Signature"
	// this prefix will be used to name tokens we requested
	// customers are able to see this in their control panel
	labelPrefix = "gotransip-client"
	// a requested Token expires after a day
	tokenExpirationDuration = time.Hour * 24
	// DemoToken can be used to test with the api without using your own account
	DemoToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw"
)

// Authenticator is used by the client to retrieve a token to use as authentication mechanism in its requests
type Authenticator interface {
	GetToken() (jwt.Token, error)
	GetPrivateKeyBody() []byte
}

// TransipAuthenticator is used to store,retrieve and request new tokens during every request
// it checks the expiry date of a Token and if it is expired it will request a new one
type TransipAuthenticator struct {
	// this contains a []byte representation of the the private key of the customer
	// this key will be used to sign a new Token request
	PrivateKeyBody []byte
	// this is Token, that is filled with a static Token that a customer provides
	// or a Token that we got from a Token request
	Token jwt.Token
	// this is the http client to do auth requests with
	HTTPClient *http.Client
	// this would be the auth path, thus where we will get new tokens from
	// todo: add a test to check if this is actually set
	BasePath string
	// this would be the acount name of customer
	Login string
	// When this is set to true the requested tokens can only be used with the 'ip' we are requesting with
	// todo: check if token request is correct when this is set
	Whitelisted bool
	// Whether or not we want to request read only Tokens, that can only only be used to retrieve information
	// not to create, modify or delete it
	ReadOnly bool
}

// AuthRequest will be transformed and send in order to request a new Token
// for more information, see: https://api.transip.nl/rest/docs.html#header-authentication
type AuthRequest struct {
	// Account name
	Login string `json:"login"`
	// Unique number for this request
	Nonce string `json:"nonce"`
	// Custom name to give this Token, you can see your tokens in the transip control panel
	Label string `json:"label,omitempty"`
	// Enable read only mode
	ReadOnly bool `json:"read_only,omitempty"`
	// Unix time stamp of when this Token should expire
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Whether this key can be used from everywhere, e.g should not be whitelisted to the current requesting ip
	GlobalKey bool `json:"global_key,omitempty"`
}

// GetToken will return the current Token if it is not expired
// if it is expired it will try to request a new Token, set and return that
// on error it passes this back
// todo: implement Token caching to io.Writer/Reader/Closer
func (a *TransipAuthenticator) GetToken() (jwt.Token, error) {
	if a.Token.Expired() && a.PrivateKeyBody == nil {
		return jwt.Token{}, errors.New("token expired and no private key is set")
	}
	if a.Token.Expired() {
		var err error
		a.Token, err = a.requestNewToken()

		if err != nil {
			return jwt.Token{}, err
		}
	}

	return a.Token, nil
}

// requestNewToken will request a new Token using the http client
// creating a new AuthRequest, converting it to json and sending that to the api auth url
// on error it will pass this back
func (a *TransipAuthenticator) requestNewToken() (jwt.Token, error) {
	restRequest := a.getAuthRequest()
	getMethod := rest.PostMethod

	httpRequest, err := restRequest.GetHTTPRequest(a.BasePath, getMethod.Method)
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error constructing token http request: %w", err)
	}
	bodyToSign, err := restRequest.GetJsonBody()
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error marshalling token request: %w", err)
	}
	signature, err := signWithKey(bodyToSign, a.PrivateKeyBody)
	if err != nil {
		return jwt.Token{}, err
	}
	httpRequest.Header.Add(signatureHeader, signature)

	httpResponse, err := a.HTTPClient.Do(httpRequest)
	if err != nil {
		return jwt.Token{}, fmt.Errorf("request error: %w", err)
	}

	defer httpResponse.Body.Close()

	// read entire response body
	b, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return jwt.Token{}, err
	}

	restResponse := rest.Response{
		Body:       b,
		StatusCode: httpResponse.StatusCode,
		Method:     getMethod,
	}

	var tokenToReturn tokenResponse
	err = restResponse.ParseResponse(&tokenToReturn)
	if err != nil {
		return jwt.Token{}, err
	}

	return jwt.New(tokenToReturn.Token)
}

// tokenResponse is used to extract a Token from the api server response
type tokenResponse struct {
	Token string `json:"Token"`
}

// getNonce returns a random 16 character length string nonce
// each time it is called
func (a *TransipAuthenticator) getNonce() string {
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)

	// convert to hex
	return fmt.Sprintf("%02x", randomBytes)
}

// getAuthRequest returns a rest.RestRequest filled with a new AuthRequest
func (a *TransipAuthenticator) getAuthRequest() rest.RestRequest {
	labelPostFix := time.Now().Unix()

	authRequest := AuthRequest{
		Login:          a.Login,
		Nonce:          a.getNonce(),
		Label:          fmt.Sprintf("%s-%d", labelPrefix, labelPostFix),
		ReadOnly:       a.ReadOnly,
		ExpirationTime: time.Now().Add(tokenExpirationDuration).Unix(),
		GlobalKey:      a.Whitelisted,
	}

	return rest.RestRequest{
		Endpoint: "/auth",
		Body:     authRequest,
	}
}

// GetPrivateKeyBody returns the private key body so we can
// todo: remove this and test in a different way, no need to get this from the outside
func (a *TransipAuthenticator) GetPrivateKeyBody() []byte {
	return a.PrivateKeyBody
}
