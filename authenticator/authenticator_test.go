package authenticator

import (
	"errors"
	"fmt"
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

func getMockServer(t *testing.T) *httptest.Server {
	tokenAsJson := fmt.Sprintf(`{"token":"%s"}`, DemoToken)

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
		rw.Write([]byte(tokenAsJson))
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
	assert.Equal(t, DemoToken, token.RawToken)
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

func TestAuthenticator_ReturnsSigningError(t *testing.T) {
	authenticator := TransipAuthenticator{
		PrivateKeyBody: []byte{0x00},
		Login:          "test-user",
		HTTPClient:     http.DefaultClient,
	}

	_, err := authenticator.requestNewToken()
	require.Error(t, err)
	assert.Equal(t, err, errors.New("could not decode private key"))
}

func TestAuthenticator_HttpRequestMarshalingError(t *testing.T) {
	authenticator := TransipAuthenticator{
		PrivateKeyBody: []byte{0x00},
		Login:          "test-user",
		HTTPClient:     http.DefaultClient,
	}

	_, err := authenticator.requestNewToken()
	require.Error(t, err)
	assert.Equal(t, err, errors.New("could not decode private key"))
}

func TestAuthenticator_GetTokenNoPrivateKey(t *testing.T) {
	authenticator := TransipAuthenticator{}
	_, err := authenticator.GetToken()

	require.Error(t, err)
	assert.Equal(t, err, errors.New("token expired and no private key is set"))
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
	authenticator := TransipAuthenticator{
		BasePath:       "http://api.transip.nl/v6",
		Login:          "test-user",
		Whitelisted:    true,
		ReadOnly:       true,
	}

	authRequest := authenticator.getAuthRequest()
	body, err := authRequest.GetJsonBody()

	require.NoError(t, err)
	stringBody := string(body)
	assert.Contains(t, stringBody, `{"login":"test-user",`)
	assert.Contains(t, stringBody, fmt.Sprintf(`"label":"gotransip-client-%d"`, time.Now().Unix()))
	// "read_only":true,"expiration_time":1584115100,"global_key":true
	assert.Contains(t, stringBody, `"read_only":true,`)
	assert.Contains(t, stringBody, `"global_key":true}`)
	assert.Contains(t, stringBody, fmt.Sprintf(`"expiration_time":%d,`, time.Now().Unix() + 86400), fmt.Sprintf(`"expiration_time":%d,`, time.Now().Unix() + 86400))
	assert.Contains(t, stringBody, `"nonce":"`)
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

