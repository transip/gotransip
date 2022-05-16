package openstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestUserRepository_GetAll(t *testing.T) {
	const apiResponse = `{"users":[{"id":"6322872d9c7e445dbbb49c1f9ca28adc","username":"example-support","description":"Supporter account","email":"support@example.com"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/users", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "6322872d9c7e445dbbb49c1f9ca28adc", all[0].ID)
	assert.Equal(t, "example-support", all[0].Username)
	assert.Equal(t, "Supporter account", all[0].Description)
	assert.Equal(t, "support@example.com", all[0].Email)
}

func TestUserRepository_GetByProjectID(t *testing.T) {
	const apiResponse = `{"users":[{"id":"6322872d9c7e445dbbb49c1f9ca28adc","username":"example-support","description":"Supporter account","email":"support@example.com"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects/6322872d9c7e445dbbb49c1f9ca28adc/users", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	all, err := repo.GetByProjectID("6322872d9c7e445dbbb49c1f9ca28adc")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "6322872d9c7e445dbbb49c1f9ca28adc", all[0].ID)
	assert.Equal(t, "example-support", all[0].Username)
	assert.Equal(t, "Supporter account", all[0].Description)
	assert.Equal(t, "support@example.com", all[0].Email)
}

func TestUserRepository_GetByID(t *testing.T) {
	const apiResponse = `{"user":{"id":"6322872d9c7e445dbbb49c1f9ca28adc","username":"example-support","description":"Supporter account","email":"support@example.com"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/users/6322872d9c7e445dbbb49c1f9ca28adc", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	user, err := repo.GetByID("6322872d9c7e445dbbb49c1f9ca28adc")
	require.NoError(t, err)

	assert.Equal(t, "6322872d9c7e445dbbb49c1f9ca28adc", user.ID)
	assert.Equal(t, "example-support", user.Username)
	assert.Equal(t, "Supporter account", user.Description)
	assert.Equal(t, "support@example.com", user.Email)
}

func TestUserRepository_Create(t *testing.T) {
	const expectedBody = `{"projectId":"6322872d9c7e445dbbb49c1f9ca28adc","username":"example-support","password":"banaan","description":"Supporter account","email":"support@example.com"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/users", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	err := repo.Create(CreateUserRequest{
		ProjectID:   "6322872d9c7e445dbbb49c1f9ca28adc",
		Username:    "example-support",
		Description: "Supporter account",
		Password:    "banaan",
		Email:       "support@example.com",
	})
	require.NoError(t, err)
}

func TestUserRepository_Update(t *testing.T) {
	const expectedBody = `{"user":{"id":"6322872d9c7e445dbbb49c1f9ca28adc","username":"example-support","description":"Supporter account","email":"support@example.com"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/users/6322872d9c7e445dbbb49c1f9ca28adc", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	err := repo.Update(User{
		ID:          "6322872d9c7e445dbbb49c1f9ca28adc",
		Username:    "example-support",
		Description: "Supporter account",
		Email:       "support@example.com",
	})
	require.NoError(t, err)
}

func TestUserRepository_ChangePassword(t *testing.T) {
	const expectedBody = `{"newPassword":"banaan"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/users/6322872d9c7e445dbbb49c1f9ca28adc", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	err := repo.ChangePassword("6322872d9c7e445dbbb49c1f9ca28adc", "banaan")
	require.NoError(t, err)
}

func TestUserRepository_Delete(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/users/6322872d9c7e445dbbb49c1f9ca28adc", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	err := repo.Delete("6322872d9c7e445dbbb49c1f9ca28adc")
	require.NoError(t, err)
}

func TestUserRepository_AddToProject(t *testing.T) {
	const expectedBody = `{"userId":"6322872d9c7e445dbbb49c1f9ca28adc"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects/6322872d9c7e445dbbb49c1f9ca28adc/users", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	err := repo.AddToProject("6322872d9c7e445dbbb49c1f9ca28adc", "6322872d9c7e445dbbb49c1f9ca28adc")
	require.NoError(t, err)
}

func TestUserRepository_RemoveFromProject(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/openstack/projects/6322872d9c7e445dbbb49c1f9ca28adc/users/6322872d9c7e445dbbb49c1f9ca28adc", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := UserRepository{Client: *client}

	err := repo.RemoveFromProject("6322872d9c7e445dbbb49c1f9ca28adc", "6322872d9c7e445dbbb49c1f9ca28adc")
	require.NoError(t, err)
}
