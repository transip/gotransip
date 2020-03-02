package authenticator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/jwt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

const amountOfNoncesToGet = 10

func TestAuthenticatorGetToken(t *testing.T) {
	token := jwt.Token{ExpiryDate: time.Now().Unix() + 3600, RawToken: "123"}
	authenticator := TransipAuthenticator{
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

	authenticator := TransipAuthenticator{
		PrivateKeyBody: key,
		BasePath:       server.URL,
		Login:          "test-user",
		Whitelisted:    false,
		ReadOnly:       false,
		HTTPClient:     http.DefaultClient,
	}

	token, err := authenticator.requestNewToken()
	assert.NoError(t, err)
	assert.Equal(t, "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw", token.RawToken)
}

func TestAuthenticationErrorIsReturned(t *testing.T) {
	server := getFailedMockServer(t)
	defer server.Close()

	key, err := ioutil.ReadFile("../testdata/signature.key")
	require.NoError(t, err)

	authenticator := TransipAuthenticator{
		PrivateKeyBody: key,
		BasePath:       server.URL,
		Login:          "test-user",
		Whitelisted:    false,
		ReadOnly:       false,
		HTTPClient:     http.DefaultClient,
	}

	_, err = authenticator.requestNewToken()
	require.Error(t, err)
	assert.Equal(t, errors.New("Authentication failed, API is not enabled for customer"), err)
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

func TestIfGetNonceIsThreadSafe(t *testing.T) {
	const amountOfThreadSafeNonceThreads = 100
	var nonces [amountOfThreadSafeNonceThreads][amountOfNoncesToGet]string
	var wg sync.WaitGroup

	// get a list of nonces N=amountOfThreadSafeNonceThreads times
	for i := 0; i < amountOfThreadSafeNonceThreads; i++ {
		wg.Add(1)
		go func(idx int) {
			nonces[idx] = getNoncesFromAuthenticator()

			defer wg.Done()
		}(i)
	}

	wg.Wait()
	var combinedNonces [amountOfThreadSafeNonceThreads * amountOfNoncesToGet]string
	counter := 0
	for i := 0; i < amountOfThreadSafeNonceThreads; i++ {
		for _, nonce := range nonces[i] {
			combinedNonces[counter] = nonce
			counter++
		}
	}

	// check if nonces are unique
	for i, nonce := range combinedNonces {
		for j, previousNonce := range combinedNonces {
			if i == j {
				continue
			}

			require.NotEqual(t, nonce, previousNonce, "duplicate nonce found")
		}
	}
}

func getNoncesFromAuthenticator() [amountOfNoncesToGet]string {
	authenticator := TransipAuthenticator{}
	var nonces [amountOfNoncesToGet]string

	for i := 0; i < amountOfNoncesToGet; i++ {
		nonces[i] = authenticator.getNonce()
	}

	return nonces
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
		// send a Token as response
		rw.Write([]byte(`{"Token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw"}`))
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
		rw.Write([]byte(`{"error":"Authentication failed, API is not enabled for customer"}`))
	}))
}
