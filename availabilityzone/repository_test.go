package availabilityzone

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	apiResponse   = `{ "availabilityZones": [ { "name": "ams0", "country": "nl", "isDefault": true } ] }`
	errorResponse = `{ "error": "errortest" }`
)

func getMockServer(t *testing.T, url string, method string, statusCode int, response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, url, req.URL.String()) // check if right url is called
		assert.Equal(t, method, req.Method)    // check if the right request method is used
		rw.WriteHeader(statusCode)             // respond with given status code
		rw.Write([]byte(response))
	}))
}

func getRepository(t *testing.T, responseStatusCode int, response string) (Repository, func()) {
	server := getMockServer(t, "/availability-zones", "GET", responseStatusCode, response)
	config := gotransip.DemoClientConfiguration
	config.URL = server.URL
	client, err := gotransip.NewClient(config)
	require.NoError(t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		server.Close()
	}

	return Repository{Client: client}, tearDown
}

func TestRepository_GetAll(t *testing.T) {
	repo, tearDown := getRepository(t, 200, apiResponse)
	defer tearDown()

	all, err := repo.GetAll()
	require.NoError(t, err)

	assert.Equal(t, 1, len(all))
	assert.Equal(t, "ams0", all[0].Name)
	assert.Equal(t, true, all[0].IsDefault)
	assert.Equal(t, "nl", all[0].Country)
}

func TestRepository_GetAllError(t *testing.T) {
	repo, tearDown := getRepository(t, 406, errorResponse)
	defer tearDown()

	obj, err := repo.GetAll()
	require.Error(t, err)
	assert.Nil(t, obj)
	assert.Equal(t, &rest.Error{Message: "errortest", StatusCode: 406}, err)
}
