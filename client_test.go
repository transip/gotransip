package gotransip

import (
	"errors"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/rest"
	"github.com/transip/gotransip/v6/rest/request"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

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

	cc.AccountName = "foobar"
	// ClientConfig with only AccountName set should raise error about private keys
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("PrivateKeyPath, token or PrivateKeyBody is required"), err)

	cc.PrivateKeyPath = "/file/not/found"
	// ClientConfig with PrivateKeyPath set to nonexisting file should raise error
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("^Could not open private key"), err.Error())

	// ClientConfig with PrivateKeyPath that does exist but is unreadable should raise
	// error
	// prepare tmpfile
	tmpFile, err := ioutil.TempFile("", "gotransip")
	assert.NoError(t, err)
	err = os.Chmod(tmpFile.Name(), 0000)
	assert.NoError(t, err)

	cc.PrivateKeyPath = tmpFile.Name()
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("permission denied"), err.Error())

	os.Remove(tmpFile.Name())
	cc.PrivateKeyPath = ""

	// Override PrivateKeyBody with PrivateKeyPath
	pkBody := []byte{2, 3, 4, 5}
	// prepare tmpfile
	tmpFile, err = ioutil.TempFile("", "gotransip")
	assert.NoError(t, err)
	err = ioutil.WriteFile(tmpFile.Name(), []byte(pkBody), 0)
	assert.NoError(t, err)

	cc.PrivateKeyPath = tmpFile.Name()
	client, err := newClient(cc)
	clientAuthenticator := client.GetAuthenticator()
	config := client.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, pkBody, clientAuthenticator.PrivateKeyBody)

	// Also, with no mode set, it should default to APIModeReadWrite
	assert.Equal(t, APIModeReadWrite, config.Mode)
	// Check if the base path is set by default
	assert.Equal(t, "https://api.transip.nl/v6", config.URL)
	os.Remove(tmpFile.Name())
	cc.PrivateKeyPath = ""

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
	clientConfig := ClientConfiguration{
		Token: authenticator.DemoToken,
		URL:   server.URL,
	}

	client, err := newClient(clientConfig)
	require.NoError(t, err)

	request := request.RestRequest{Endpoint: "/domains"}
	var domainsResponse domain.DomainsResponse

	err = client.call(rest.GetRestMethod, request, &domainsResponse)
	require.NoError(t, err)
	require.Equal(t, 1, len(domainsResponse.Domains))
	assert.Equal(t, "test.nl", domainsResponse.Domains[0].Name)
}

// Test if we can connect to the api server using the demo token
func TestClientCallToApiServer(t *testing.T) {
	clientConfig := ClientConfiguration{
		Token: authenticator.DemoToken,
	}

	client, err := NewClient(clientConfig)
	require.NoError(t, err)

	request := request.RestRequest{Endpoint: "/products"}
	var responseObject product.ProductsResponse

	err = client.Get(request, &responseObject)
	require.NoError(t, err)
	assert.Equal(t,  6, len(responseObject.Products.Vps))
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
