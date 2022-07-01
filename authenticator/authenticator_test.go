package authenticator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/jwt"
)

const amountOfNoncesToGet = 10

func getMockServer(t *testing.T) *httptest.Server {
	tokenAsJSON := fmt.Sprintf(`{"token":"%s"}`, DemoToken)

	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check if right url is called
		assert.Equal(t, req.URL.String(), "/auth")
		// check if a signature is set
		assert.NotEmpty(t, req.Header.Get("Signature"), "Signature not set")
		// check if the request has the right content-type
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		// check if the request has the right content-type
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		// check if the request is a POST request
		assert.Equal(t, req.Method, "POST")
		// send a Token as response
		_, err := rw.Write([]byte(tokenAsJSON))
		require.NoError(t, err, "error when writing mock response")
	}))
}

func getFailedMockServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check if right url is called
		assert.Equal(t, req.URL.String(), "/auth")
		// check if a signature is set
		assert.NotEmpty(t, req.Header.Get("Signature"), "Signature not set")
		// respond with a 409 error
		rw.WriteHeader(409)
		// send a Token as response
		_, err := rw.Write([]byte(`{"error":"Authentication failed, API is not enabled for customer"}`))
		require.NoError(t, err, "error when writing mock response")
	}))
}

func TestAuthenticatorGetToken(t *testing.T) {
	token := jwt.Token{ExpiryDate: time.Now().Unix() + 3600, RawToken: "123"}
	authenticator := Authenticator{
		Token: token,
	}

	returnedToken, err := authenticator.GetToken()
	assert.NoError(t, err)
	assert.Equal(t, token, returnedToken)
}

func TestRequestANewToken(t *testing.T) {
	server := getMockServer(t)
	defer server.Close()
	key, err := ioutil.ReadFile("../testdata/signature.key")
	require.NoError(t, err)

	authenticator := Authenticator{
		PrivateKeyBody: key,
		BasePath:       server.URL,
		Login:          "test-user",
		Whitelisted:    false,
		ReadOnly:       false,
		HTTPClient:     http.DefaultClient,
	}

	token, err := authenticator.requestNewToken()
	assert.NoError(t, err)
	assert.Equal(t, DemoToken, token.RawToken)
}

func TestAuthenticationErrorIsReturned(t *testing.T) {
	server := getFailedMockServer(t)
	defer server.Close()

	key, err := ioutil.ReadFile("../testdata/signature.key")
	require.NoError(t, err)

	authenticator := Authenticator{
		PrivateKeyBody: key,
		BasePath:       server.URL,
		Login:          "test-user",
		Whitelisted:    false,
		ReadOnly:       false,
		HTTPClient:     http.DefaultClient,
	}

	_, err = authenticator.requestNewToken()
	if assert.Errorf(t, err, "auth failed error not returned") {
		err = errors.Unwrap(err)
		assert.Equal(t, "Authentication failed, API is not enabled for customer", err.Error())
	}
}

func TestAuthenticator_ReturnsSigningError(t *testing.T) {
	authenticator := Authenticator{
		PrivateKeyBody: []byte{0x00},
		Login:          "test-user",
		HTTPClient:     http.DefaultClient,
	}

	_, err := authenticator.requestNewToken()
	if assert.Errorf(t, err, "private key decode error not returned") {
		assert.Equal(t, err, ErrDecodingPrivateKey)
	}
}

func TestAuthenticator_HttpRequestMarshalingError(t *testing.T) {
	authenticator := Authenticator{
		PrivateKeyBody: []byte{0x00},
		Login:          "test-user",
		HTTPClient:     http.DefaultClient,
	}

	_, err := authenticator.requestNewToken()
	if assert.Errorf(t, err, "decode private key error not returned") {
		assert.Equal(t, err, ErrDecodingPrivateKey)
	}
}

func TestAuthenticator_GetTokenNoPrivateKey(t *testing.T) {
	authenticator := Authenticator{}
	_, err := authenticator.GetToken()

	if assert.Errorf(t, err, "token expired error not returned") {
		assert.Equal(t, err, ErrTokenExpired)
	}
}

func TestNonceIsNotStatic(t *testing.T) {
	nonces := getNoncesFromAuthenticator()

	for idx, nonce := range nonces {
		for jdx, previousNonce := range nonces {
			if idx == jdx {
				continue
			}
			require.NotEqual(t, nonce, previousNonce, "duplicate nonce found")
		}
	}
}

func TestAuthenticator_getAuthRequest(t *testing.T) {
	authenticator := Authenticator{
		BasePath:        "http://api.transip.nl/v6",
		Login:           "test-user1",
		Whitelisted:     true,
		ReadOnly:        true,
		TokenExpiration: 30 * time.Second,
	}

	authRequest, err := authenticator.getAuthRequest()
	require.NoError(t, err)
	body, err := authRequest.GetJSONBody()

	require.NoError(t, err)
	stringBody := string(body)

	assert.Contains(t, stringBody, `{"login":"test-user1",`)
	assert.Contains(t, stringBody, "gotransip-client-")
	assert.Contains(t, stringBody, `"read_only":true,`)
	assert.Contains(t, stringBody, `"global_key":false}`)
	assert.Contains(t, stringBody, `"expiration_time":"30 seconds",`)
	assert.Contains(t, stringBody, `"nonce":"`)
}

func getNoncesFromAuthenticator() [amountOfNoncesToGet]string {
	authenticator := Authenticator{}
	var nonces [amountOfNoncesToGet]string

	for i := 0; i < amountOfNoncesToGet; i++ {
		nonces[i], _ = authenticator.getNonce()
	}

	return nonces
}

func TestAuthenticator_getTokenCacheKey(t *testing.T) {
	authenticator := Authenticator{Login: "test"}

	assert.Equal(t, "gotransip-client-test-token", authenticator.getTokenCacheKey())
}

func TestAuthenticator_getTokenExpirationString(t *testing.T) {
	// Default expiration
	authenticator := Authenticator{}
	assert.Equal(t, defaultTokenExpiration, authenticator.getTokenExpirationString())

	// Custom expiration
	authenticator.TokenExpiration = 60 * time.Minute
	assert.Equal(t, "3600 seconds", authenticator.getTokenExpirationString())
}
