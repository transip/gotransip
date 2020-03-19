package traffic

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t                   *testing.T
	expectedURL         string
	expectedMethod      string
	statusCode          int
	expectedRequestBody string
	response            string
	skipRequestBody     bool
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedURL, req.URL.String()) // check if right expectedURL is called

		if m.skipRequestBody == false && req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := ioutil.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequestBody, string(body))
		}

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequestBody expectedMethod is used
		rw.WriteHeader(m.statusCode)                    // respond with given status code

		if m.response != "" {
			_, err := rw.Write([]byte(m.response))
			require.NoError(m.t, err, "error when writing mock response")
		}
	}))
}

func (m *mockServer) getClient() (*repository.Client, func()) {
	httpServer := m.getHTTPServer()
	config := gotransip.DemoClientConfiguration
	config.URL = httpServer.URL
	client, err := gotransip.NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return &client, tearDown
}

func TestRepository_GetTrafficInformationForVps(t *testing.T) {
	const apiResponse = `{ "trafficInformation": { "startDate": "2019-06-22", "endDate": "2019-07-22", "usedInBytes": 7860253754, "usedTotalBytes": 11935325369, "maxInBytes": 1073741824000 } }`
	server := mockServer{t: t, expectedURL: "/traffic", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}
	pool, err := repo.GetTrafficPool()
	require.NoError(t, err)

	assert.Equal(t, "2019-06-22", pool.StartDate.Format("2006-01-02"))
	assert.Equal(t, "2019-07-22", pool.EndDate.Format("2006-01-02"))
	assert.EqualValues(t, 7860253754, pool.UsedInBytes)
	assert.EqualValues(t, 11935325369, pool.UsedTotalBytes)
	assert.EqualValues(t, 1073741824000, pool.MaxInBytes)
}

func TestRepository_GetTrafficPool(t *testing.T) {
	const apiResponse = `{ "trafficInformation": { "startDate": "2019-06-22", "endDate": "2019-07-22", "usedInBytes": 7860253754, "usedTotalBytes": 11935325369, "maxInBytes": 1073741824000 } }`
	server := mockServer{t: t, expectedURL: "/traffic/test-vps", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}
	pool, err := repo.GetTrafficInformationForVps("test-vps")
	require.NoError(t, err)

	assert.Equal(t, "2019-06-22", pool.StartDate.Format("2006-01-02"))
	assert.Equal(t, "2019-07-22", pool.EndDate.Format("2006-01-02"))
	assert.EqualValues(t, 7860253754, pool.UsedInBytes)
	assert.EqualValues(t, 11935325369, pool.UsedTotalBytes)
	assert.EqualValues(t, 1073741824000, pool.MaxInBytes)
}
