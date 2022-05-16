package openstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestProjectRepository_GetAll(t *testing.T) {
	const apiResponse = `{"projects":[{"id":"7a7a3bcb46c6450f95c53edb8dcebc7b","name": "example-datacenter","description": "This is an example project","isLocked": true,"isBlocked":false}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	err := repo.Handover("7a7a3bcb46c6450f95c53edb8dcebc7b", "example2")
	require.NoError(t, err)
}

func TestProjectRepository_Cancel(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects/7a7a3bcb46c6450f95c53edb8dcebc7b", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := ProjectRepository{Client: *client}

	err := repo.Cancel("7a7a3bcb46c6450f95c53edb8dcebc7b")
	require.NoError(t, err)
}
