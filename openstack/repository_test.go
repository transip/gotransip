package openstack

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
)

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t               *testing.T
	expectedURL     string
	expectedMethod  string
	statusCode      int
	expectedRequest string
	response        string
	skipRequestBody bool
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedURL, req.URL.String()) // check if right expectedURL is called

		if m.skipRequestBody == false && req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := ioutil.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequest, string(body))
		}

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequest expectedMethod is used
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

func TestProjectRepository_GetAll(t *testing.T) {
	const apiResponse = `{"projects":[{"id":"7a7a3bcb46c6450f95c53edb8dcebc7b","name": "example-datacenter","description": "This is an example project","isLocked": true,"isBlocked":false}]}`
	server := mockServer{t: t, expectedURL: "/openstack/projects", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "7a7a3bcb46c6450f95c53edb8dcebc7b", all[0].ID)
	assert.Equal(t, "example-datacenter", all[0].Name)
	assert.Equal(t, "This is an example project", all[0].Description)
	assert.Equal(t, true, all[0].IsLocked)
	assert.Equal(t, false, all[0].IsBlocked)
}

func TestProjectRepository_GetByID(t *testing.T) {
	const apiResponse = `{"project":{"id":"7a7a3bcb46c6450f95c53edb8dcebc7b","name":"example-datacenter","description":"This is an example project","isLocked": true,"isBlocked":false}}`
	server := mockServer{t: t, expectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	project, err := repo.GetByID("7a7a3bcb46c6450f95c53edb8dcebc7b")
	require.NoError(t, err)

	assert.Equal(t, "7a7a3bcb46c6450f95c53edb8dcebc7b", project.ID)
	assert.Equal(t, "example-datacenter", project.Name)
	assert.Equal(t, "This is an example project", project.Description)
	assert.Equal(t, true, project.IsLocked)
	assert.Equal(t, false, project.IsBlocked)
}

func TestProjectRepository_Create(t *testing.T) {
	const expectedRequestBody = `{"name":"example-datacenter","description":"This is an example project","isLocked":false,"isBlocked":false}`
	server := mockServer{t: t, expectedURL: "/openstack/projects", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	p := Project{
		Name:        "example-datacenter",
		Description: "This is an example project",
	}
	err := repo.Create(p)
	require.NoError(t, err)
}

func TestProjectRepository_Update(t *testing.T) {
	const expectedRequestBody = `{"project":{"id":"7a7a3bcb46c6450f95c53edb8dcebc7b","name":"example-datacenter","description":"This is an example project","isLocked":false,"isBlocked":false}}`
	server := mockServer{t: t, expectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	p := Project{
		ID:          "7a7a3bcb46c6450f95c53edb8dcebc7b",
		Name:        "example-datacenter",
		Description: "This is an example project",
	}
	err := repo.Update(p)
	require.NoError(t, err)
}

func TestProjectRepository_Handover(t *testing.T) {
	const expectedRequestBody = `{"action":"handover","targetCustomerName":"example2"}`
	server := mockServer{t: t, expectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	err := repo.Handover("7a7a3bcb46c6450f95c53edb8dcebc7b", "example2")
	require.NoError(t, err)
}

func TestProjectRepository_Cancel(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	err := repo.Cancel("7a7a3bcb46c6450f95c53edb8dcebc7b")
	require.NoError(t, err)
}
