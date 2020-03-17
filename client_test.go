package gotransip

import (
	"bytes"
	"errors"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	var cc ClientConfiguration
	var err error

	// empty ClientConfig should raise error about missing AccountName
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("AccountName is required"), err)

	// ... unless a token is provided
	cc.Token = authenticator.DemoToken
	_, err = NewClient(cc)
	require.NoError(t, err, "No error when correct token is provided")
	cc.Token = ""

	// no error should be thrown when enabling demo mode
	cc.DemoMode = true
	client, err := newClient(cc)
	require.NoError(t, err, "No error should be thrown upon enabling demo mode")
	assert.Equal(t, authenticator.DemoToken, client.GetConfig().Token, "Token should be demo token")
	cc.DemoMode = false

	cc.AccountName = "foobar"
	// ClientConfig with only AccountName set should raise error about private keys
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("PrivateKeyReader, token or PrivateKeyReader is required"), err)

	cc.PrivateKeyReader = iotest.TimeoutReader(bytes.NewReader([]byte{0, 1}))
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("error while reading private key: timeout"), err)

	cc.PrivateKeyReader = nil

	// Override PrivateKeyBody with PrivateKeyReader
	pkBody := []byte{2, 3, 4, 5}
	cc.PrivateKeyReader = bytes.NewReader(pkBody)

	client, err = newClient(cc)
	clientAuthenticator := client.GetAuthenticator()
	config := client.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, pkBody, clientAuthenticator.GetPrivateKeyBody())

	// Also, with no mode set, it should default to APIModeReadWrite
	assert.Equal(t, APIModeReadWrite, config.Mode)
	// Check if the base path is set by default
	assert.Equal(t, "https://api.transip.nl/v6", config.URL)
	cc.PrivateKeyReader = nil

	// override API mode to APIModeReadOnly
	cc.Mode = APIModeReadOnly
	cc.Token = authenticator.DemoToken
	client, err = newClient(cc)
	clientAuthenticator = client.GetAuthenticator()
	config = client.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, APIModeReadOnly, config.Mode)
}

func TestClientCallReturnsObject(t *testing.T) {
	apiResponse := `{"domains":[{"name":"test.nl"}]}`
	server := mockServer{t: t, expectedMethod: "GET", expectedUrl: "/domains", statusCode: 200, response: apiResponse}
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
	assert.Equal(t, "test.nl", domainsResponse.Domains[0].Name)
}

func TestEmptyBodyPostDoesPostWithoutBody(t *testing.T) {
	apiResponse := `{"domains":[{"name":"test.nl"}]}`
	server := mockServer{t: t, expectedMethod: "POST", expectedUrl: "/test", statusCode: 201, response: apiResponse}
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

	request := rest.Request{Endpoint: "/products"}
	var responseObject product.ProductsResponse

	err = client.Get(request, &responseObject)
	require.NoError(t, err)
	assert.NotZero(t, len(responseObject.Products.Vps))
}

func TestClient_TestMode(t *testing.T) {
	apiResponse := `{"ping":"pong"}`
	params := url.Values{"test": []string{"1"}}

	server := mockServer{t: t, expectedMethod: "POST", expectedUrl: "/test?test=1", statusCode: 200, response: apiResponse, expectedParams: params}
	httpServer := server.getHTTPServer()
	defer httpServer.Close()

	// setup a client with test mode enabled
	clientConfig := ClientConfiguration{DemoMode: true, TestMode: true, URL: httpServer.URL}
	client, err := NewClient(clientConfig)

	restRequest := rest.Request{Endpoint: "/test"}
	err = client.Post(restRequest)
	require.NoError(t, err)
}

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t               *testing.T
	expectedUrl     string
	expectedMethod  string
	statusCode      int
	expectedRequest string
	response        string
	expectedParams  url.Values
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedUrl, req.URL.String()) // check if right expectedUrl is called

		if req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := ioutil.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequest, string(body))
		}

		// expect http query strings to be equal
		assert.Equal(m.t, m.expectedParams.Encode(), req.URL.RawQuery)

		// check if a signature is set
		assert.NotEmpty(m.t, req.Header.Get("Authorization"), "Authentication header not set")
		// check if the request has the right content-type
		assert.Equal(m.t, req.Header.Get("Accept"), "application/json")
		// check if the request has the right content-type
		assert.Equal(m.t, req.Header.Get("Content-Type"), "application/json")

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequest expectedMethod is used
		rw.WriteHeader(m.statusCode)                    // respond with given status code

		if m.response != "" {
			rw.Write([]byte(m.response))
		}
	}))
}

func (m *mockServer) getClient() (repository.Client, func()) {
	httpServer := m.getHTTPServer()
	config := ClientConfiguration{DemoMode: true, URL: httpServer.URL}
	client, err := NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return client, tearDown
}
