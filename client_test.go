package gotransip

import (
	"errors"
	"github.com/transip/gotransip/v6/domain"
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
	realToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw"
	var err error

	// empty ClientConfig should raise error about missing AccountName
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("AccountName is required"), err)

	// ... unless a Token is provided
	cc.Token = realToken
	_, err = NewClient(cc)
	require.NoError(t, err, "No error when correct token is provided")
	cc.Token = ""

	cc.AccountName = "foobar"
	// ClientConfig with only AccountName set should raise error about private keys
	_, err = NewClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("PrivateKeyPath, Token or PrivateKeyBody is required"), err)

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
	client, err := NewClient(cc)
	authenticator := client.GetAuthenticator()
	config := client.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, pkBody, authenticator.PrivateKeyBody)

	// Also, with no mode set, it should default to APIModeReadWrite
	assert.Equal(t, APIModeReadWrite, config.Mode)
	// Check if the base path is set by default
	assert.Equal(t, "https://api.transip.nl/v6", config.URL)
	os.Remove(tmpFile.Name())
	cc.PrivateKeyPath = ""

	// override API mode to APIModeReadOnly
	cc.Mode = APIModeReadOnly
	cc.Token = realToken
	client, err = NewClient(cc)
	authenticator = client.GetAuthenticator()
	config = client.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, APIModeReadOnly, config.Mode)
}

func TestClientCallReturnsObject(t *testing.T) {
	server := getMockServer(t)
	clientConfig := ClientConfiguration{
		Token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw",
		URL:   server.URL,
	}

	client, err := NewClient(clientConfig)
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
		Token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cyIVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhwIjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw",
	}

	client, err := NewClient(clientConfig)
	require.NoError(t, err)

	request := request.RestRequest{Endpoint: "/domains"}
	var domainsResponse domain.DomainsResponse

	err = client.call(rest.GetRestMethod, request, &domainsResponse)
	require.NoError(t, err)
	assert.Equal(t, len(domainsResponse.Domains), 5)
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
