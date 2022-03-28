package openstack

import (
	"fmt"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// UserRepository lists and manages OpenStack users. It can also add and remove users
// from projects
type UserRepository repository.RestRepository

// GetAll list all OpenStack users
func (r *UserRepository) GetAll() ([]User, error) {
	var response usersWrapper
	restRequest := rest.Request{Endpoint: "/openstack/users"}
	err := r.Client.Get(restRequest, &response)

	return response.Users, err
}

// GetByProjectID gets all the users assigned to a project
func (r *UserRepository) GetByProjectID(projectID string) ([]User, error) {
	var response usersWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/projects/%s/users", projectID)}
	err := r.Client.Get(restRequest, &response)

	return response.Users, err
}

// AddToProject adds a user to an OpenStack project
func (r *UserRepository) AddToProject(userID string, projectID string) error {
	type addToProjectRequest struct {
		UserID string `json:"userId"`
	}
	request := addToProjectRequest{UserID: userID}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/projects/%s/users", projectID), Body: request}
	return r.Client.Post(restRequest)
}

// RemoveFromProject removes a user from an OpenStack project
func (r *UserRepository) RemoveFromProject(userID string, projectID string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/projects/%s/users/%s", projectID, userID)}
	return r.Client.Delete(restRequest)
}

// GetByID fetches information about an user by their ID
func (r *UserRepository) GetByID(userID string) (User, error) {
	var response userWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/users/%s", userID)}
	err := r.Client.Get(restRequest, &response)

	return response.User, err
}

// Create creates a new user and grants it access to the provided projectID
func (r *UserRepository) Create(request CreateUserRequest) error {
	restRequest := rest.Request{Endpoint: "/openstack/users", Body: request}

	return r.Client.Post(restRequest)
}

// Update can be used to update the description and email of user. To change the password of
// as user see ChangePassword
func (r *UserRepository) Update(user User) error {
	requestBody := userWrapper{User: user}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/users/%s", user.ID), Body: requestBody}

	return r.Client.Put(restRequest)
}

// ChangePassword changes the password of an existing user
func (r *UserRepository) ChangePassword(userID string, newPassword string) error {
	requestBody := changePasswordRequest{NewPassword: newPassword}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/users/%s", userID), Body: requestBody}

	return r.Client.Patch(restRequest)
}

// Delete removes an OpenStack user entirely
func (r *UserRepository) Delete(userID string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/users/%s", userID)}

	return r.Client.Delete(restRequest)
}
