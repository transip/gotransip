package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
)

// MockServer struct is used to test the how the client sends a request
// and responds to a servers response
type MockServer struct {
	T               *testing.T
	ExpectedURL     string
	ExpectedMethod  string
	StatusCode      int
	ExpectedRequest string
	Response        string
	SkipRequestBody bool
}

// GetHTTPServer returns the server part of the MockServer
func (m *MockServer) GetHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.T, m.ExpectedURL, req.URL.String()) // check if right expectedURL is called

		if !m.SkipRequestBody && req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := io.ReadAll(req.Body)
			require.NoError(m.T, err)
			assert.Equal(m.T, m.ExpectedRequest, string(body))
		}

		assert.Equal(m.T, m.ExpectedMethod, req.Method) // check if the right expectedRequest expectedMethod is used
		rw.WriteHeader(m.StatusCode)                    // respond with given status code

		if m.Response != "" {
			_, err := rw.Write([]byte(m.Response))
			require.NoError(m.T, err, "error when writing mock response")
		}
	}))
}

// GetClient returns a client to the MockServer
func (m *MockServer) GetClient() (*repository.Client, func()) {
	httpServer := m.GetHTTPServer()
	config := gotransip.DemoClientConfiguration
	config.URL = httpServer.URL
	client, err := gotransip.NewClient(config)
	require.NoError(m.T, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return &client, tearDown
}

// GetMockServer is a wrapper for returning a more generic MockServer's HTTPServer.
// If you want more control, construct a MockServer manually
func GetMockServer(t *testing.T, url string, method string, statusCode int, response string) *httptest.Server {
	m := MockServer{
		T:              t,
		ExpectedURL:    url,
		ExpectedMethod: method,
		StatusCode:     statusCode,
		Response:       response,
	}
	return m.GetHTTPServer()
}
