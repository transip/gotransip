package authenticator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/jwt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuthenticatorGetToken(t *testing.T) {
	token := jwt.Token{ExpiryDate: time.Now().Unix() + 3600, RawToken: "123"}
	authenticator := Authenticator{
		Token:    token,
		BasePath: "https://api.transip.nl",
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

	token, err := authenticator.RequestNewToken()
	assert.NoError(t, err)
	assert.Equal(t, "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw", token.RawToken)
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

	_, err = authenticator.RequestNewToken()
	require.Error(t, err)
	assert.Equal(t, errors.New("Authentication failed, API is not enabled for customer"), err)
}

func TestNonceIsNotStatic(t *testing.T) {
	authenticator := Authenticator{}
	previousNonce := authenticator.GetNonce()

	for i := 0; i < 100; i++ {
		nonce := authenticator.GetNonce()
		assert.NotEqual(t, nonce, previousNonce)
		previousNonce = nonce
	}
}

func getMockServer(t *testing.T) *httptest.Server {
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
		// send a token as response
		rw.Write([]byte(`{"token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw"}`))
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
		// send a token as response
		rw.Write([]byte(`{"error":"Authentication failed, API is not enabled for customer"}`))
	}))
}
