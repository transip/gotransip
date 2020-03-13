package gotransip

import (
	"bytes"
	"errors"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/rest"
	"net/http"
	"net/http/httptest"
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
	server := getMockServer(t)
	clientConfig := ClientConfiguration{Token: authenticator.DemoToken, URL: server.URL}
	client, err := newClient(clientConfig)
	require.NoError(t, err)

	restRequest := rest.RestRequest{Endpoint: "/domains"}
	type domainResponse struct {
		Name string `json:"name"`
	}
	var domainsResponse struct {
		Domains []domainResponse `json:"domains"`
	}

	err = client.Get(restRequest, &domainsResponse)
	require.NoError(t, err)
	require.Equal(t, 1, len(domainsResponse.Domains))
	assert.Equal(t, "test.nl", domainsResponse.Domains[0].Name)
}

func TestEmptyBodyPostDoesPostWithoutBody(t *testing.T) {
	server := getPostMockServer(t)
	clientConfig := ClientConfiguration{Token: authenticator.DemoToken, URL: server.URL}
	client, err := newClient(clientConfig)
	require.NoError(t, err)

	restRequest := rest.RestRequest{Endpoint: "/test"}
	err = client.Post(restRequest)
	require.NoError(t, err)
}

// Test if we can connect to the api server using the demo token
func TestClientCallToApiServer(t *testing.T) {
	clientConfig := ClientConfiguration{
		Token: authenticator.DemoToken,
	}

	client, err := NewClient(clientConfig)
	require.NoError(t, err)

	request := rest.RestRequest{Endpoint: "/products"}
	var responseObject product.ProductsResponse

	err = client.Get(request, &responseObject)
	require.NoError(t, err)
	assert.Equal(t, 6, len(responseObject.Products.Vps))
}

func getPostMockServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check if right url is called
		assert.Equal(t, req.URL.String(), "/test")
		// check if the right method is used
		assert.Equal(t, req.Method, "POST")
		// check if body is empty
		assert.Equal(t, int64(0), req.ContentLength)
		// check if a signature is set
		assert.NotEmpty(t, req.Header.Get("Authorization"), "Authentication header not set")
		// check if the request has the right content-type
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		// check if the request has the right content-type
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		// respond with 200
		rw.WriteHeader(200)
	}))
}

func getMockServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check if right url is called
		assert.Equal(t, req.URL.String(), "/domains")
		// check if the right method is used
		assert.Equal(t, req.Method, "GET")
		// check if a signature is set
		assert.NotEmpty(t, req.Header.Get("Authorization"), "Authentication header not set")
		// check if the request has the right content-type
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		// check if the request has the right content-type
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		// respond with 200
		rw.WriteHeader(200)
		// send a token as response
		rw.Write([]byte(`{"domains":[{"name":"test.nl"}]}`))
	}))
}
