package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getMockServer(t *testing.T, url string, method string, statusCode int, response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, url, req.URL.String()) // check if right url is called
		assert.Equal(t, method, req.Method)    // check if the right request method is used
		rw.WriteHeader(statusCode)             // respond with given status code
		rw.Write([]byte(response))
	}))
}

func getRepository(t *testing.T, url string, responseStatusCode int, response string) (Repository, func()) {
	server := getMockServer(t, url, "GET", responseStatusCode, response)
	config := gotransip.ClientConfiguration{DemoMode: true, URL: server.URL}
	client, err := gotransip.NewClient(config)
	require.NoError(t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		server.Close()
	}

	return Repository{Client: client}, tearDown
}

func TestRepository_Test(t *testing.T) {
	const apiResponse = `{ "ping":"pong" }`

	repo, tearDown := getRepository(t, "/api-test", 200, apiResponse)
	defer tearDown()

	require.NoError(t, repo.Test())
}

func TestRepository_TestError(t *testing.T) {
	const apiResponse = `{ "error":"blablabla" }`

	repo, tearDown := getRepository(t, "/api-test", 409, apiResponse)
	defer tearDown()

	err := repo.Test()
	require.Error(t, err)
	assert.Equal(t, errors.New("blablabla"), err)
}
