package gotransip

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

func TestNewClient(t *testing.T) {
	var cc ClientConfiguration
	var err error

	// empty ClientConfig should raise error about missing AccountName
	_, err = NewClient(cc)
	if assert.Errorf(t, err, "accountname error not returned") {
		assert.Equal(t, errors.New("AccountName is required"), err)
	}

	// ... unless a token is provided
	cc.Token = authenticator.DemoToken
	_, err = NewClient(cc)
	require.NoError(t, err, "No error when correct token is provided")
	cc.Token = ""

	// no error should be thrown when enabling demo mode
	client, err := newClient(DemoClientConfiguration)
	require.NoError(t, err, "No error should be thrown upon enabling demo mode")
	assert.Equal(t, authenticator.DemoToken, client.GetConfig().Token, "Token should be demo token")

	cc.AccountName = "foobar"
	// ClientConfig with only AccountName set should raise error about private keys
	_, err = NewClient(cc)
	if assert.Errorf(t, err, "expecting private key, token required error") {
		assert.EqualError(t, err, "PrivateKeyReader, PrivateKeyPath, token or KeyManager is required")
	}

	cc.PrivateKeyReader = iotest.TimeoutReader(bytes.NewReader([]byte{0, 1}))
	_, err = NewClient(cc)
	if assert.Errorf(t, err, "expecting private key reader error returned") {
		assert.EqualError(t, err, "error while reading private key: timeout")
	}

	cc.PrivateKeyReader = nil

	// Override PrivateKeyBody with PrivateKeyReader
	pkBody := []byte{2, 3, 4, 5}
	cc.PrivateKeyReader = bytes.NewReader(pkBody)

	client, err = newClient(cc)
	clientAuthenticator := client.GetAuthenticator()
	config := client.GetConfig()
	assert.NoError(t, err)
	// Test if private key body is passed to the authenticator
	assert.Equal(t, pkBody, clientAuthenticator.PrivateKeyBody)
	// Test if the base url is passed to the authenticator
	assert.Equal(t, config.URL, clientAuthenticator.BasePath)
	// Test if the account name is passed on to the authenticator
	assert.Equal(t, "foobar", clientAuthenticator.Login)

	// Test if private key from path is read and passed to the authenticator
	privateKeyFile, err := os.Open("testdata/signature.key")
	require.NoError(t, err)
	privateKeyBody, err := io.ReadAll(privateKeyFile)
	require.NoError(t, err)

	// Test that a tokencache is passed to the authenticator
	cacheFile := filepath.Join(os.TempDir(), "gotransip_test_token_cache")
	defer os.Remove(cacheFile)
	cache, err := authenticator.NewFileTokenCache(cacheFile)
	require.NoError(t, err)
	client, err = newClient(ClientConfiguration{PrivateKeyPath: "testdata/signature.key", AccountName: "example-user", TokenCache: cache})
	require.NoError(t, err)
	clientAuthenticator = client.GetAuthenticator()
	require.NotNil(t, clientAuthenticator.TokenCache)
	assert.Equal(t, cache, clientAuthenticator.TokenCache)

	// Check if private key read from file is the same as the key body on the authenticator
	assert.Equal(t, privateKeyBody, clientAuthenticator.PrivateKeyBody)

	// Test that the default expiration time is set on the authenticator
	assert.Equal(t, time.Duration(0), clientAuthenticator.TokenExpiration)

	// Test that the default whitelisted value is set on the authenticator
	assert.False(t, clientAuthenticator.Whitelisted)

	// Override TokenExpiration to 30 seconds
	cc.TokenExpiration = 30 * time.Second
	// Override TokenWhitelisted to true
	cc.TokenWhitelisted = true
	client, err = newClient(cc)
	clientAuthenticator = client.GetAuthenticator()
	assert.NoError(t, err)

	// Test that the new expiration time is set on the authenticator
	assert.Equal(t, cc.TokenExpiration, clientAuthenticator.TokenExpiration)

	// Test that the new whitelisted value is set on the authenticator
	assert.True(t, clientAuthenticator.Whitelisted)

	// Also, with no mode set, it should default to APIModeReadWrite
	assert.Equal(t, APIModeReadWrite, config.Mode)
	// Check if the base path is set by default
	assert.Equal(t, "https://api.transip.nl/v6", config.URL)
	cc.PrivateKeyReader = nil

	// Override API mode to APIModeReadOnly
	cc.Mode = APIModeReadOnly
	cc.Token = authenticator.DemoToken
	client, err = newClient(cc)
	clientAuthenticator = client.GetAuthenticator()
	assert.NoError(t, err)

	// Assert that the api mode is set on the authenticator
	assert.True(t, clientAuthenticator.ReadOnly)
}

func TestClientCallReturnsObject(t *testing.T) {
	apiResponse := `{"domains":[{"name":"testje.nl"}]}`
	server := mockServer{t: t, expectedMethod: "GET", expectedURL: "/domains", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()

	restRequest := rest.Request{Endpoint: "/domains"}
	type domainResponse struct {
		Name string `json:"name"`
	}
	var domainsResponse struct {
		Domains []domainResponse `json:"domains"`
	}

	err := client.Get(restRequest, &domainsResponse)
	require.NoError(t, err)
	require.Equal(t, 1, len(domainsResponse.Domains))
	assert.Equal(t, "testje.nl", domainsResponse.Domains[0].Name)
}

func TestEmptyBodyPostDoesPostWithoutBody(t *testing.T) {
	apiResponse := `{"domains":[{"name":"test.nl"}]}`
	server := mockServer{t: t, expectedMethod: "POST", expectedURL: "/test", statusCode: 201, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()

	restRequest := rest.Request{Endpoint: "/test"}
	err := client.Post(restRequest)
	require.NoError(t, err)
}

// Test if we can connect to the api server using the demo token
func TestClient_CallToLiveApiServer(t *testing.T) {
	clientConfig := ClientConfiguration{
		Token: authenticator.DemoToken,
	}

	client, err := NewClient(clientConfig)
	require.NoError(t, err)

	request := rest.Request{Endpoint: "/api-test"}
	var responseObject struct {
		Response string `json:"ping"`
	}

	err = client.Get(request, &responseObject)
	require.NoError(t, err)
	assert.NotZero(t, len(responseObject.Response))
}

func TestClient_TestMode(t *testing.T) {
	apiResponse := `{"ping":"pong"}`
	params := url.Values{"test": []string{"1"}}

	server := mockServer{t: t, expectedMethod: "POST", expectedURL: "/test?test=1", statusCode: 200, response: apiResponse, expectedParams: params}
	httpServer := server.getHTTPServer()
	defer httpServer.Close()

	// setup a client with test mode enabled
	clientConfig := DemoClientConfiguration
	clientConfig.URL = httpServer.URL
	clientConfig.TestMode = true

	client, err := NewClient(clientConfig)
	require.NoError(t, err)

	restRequest := rest.Request{Endpoint: "/test"}
	err = client.Post(restRequest)
	require.NoError(t, err)
}

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t               *testing.T
	expectedURL     string
	expectedMethod  string
	statusCode      int
	expectedRequest string
	response        string
	expectedParams  url.Values
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedURL, req.URL.String()) // check if right expectedURL is called

		if req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := io.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequest, string(body))
		}

		// expect http query strings to be equal
		assert.Equal(m.t, m.expectedParams.Encode(), req.URL.RawQuery)

		// check if a signature is set
		assert.NotEmpty(m.t, req.Header.Get("Authorization"), "Authentication header not set")
		// check if the request has the right content-type
		assert.Equal(m.t, req.Header.Get("Accept"), "application/json")
		// check if the request has the right user-agent
		assert.Equal(m.t, req.Header.Get("User-Agent"), userAgent)
		// check if the request has the right content-type
		assert.Equal(m.t, req.Header.Get("Content-Type"), "application/json")

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequest expectedMethod is used
		rw.WriteHeader(m.statusCode)                    // respond with given status code

		if m.response != "" {
			_, err := rw.Write([]byte(m.response))
			require.NoError(m.t, err, "error when writing mock response")
		}
	}))
}

func (m *mockServer) getClient() (repository.Client, func()) {
	httpServer := m.getHTTPServer()
	config := DemoClientConfiguration
	config.URL = httpServer.URL

	client, err := NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return client, tearDown
}
