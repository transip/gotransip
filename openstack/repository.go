package openstack

import (
	"fmt"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// ProjectRepository is the repository for creating and modifying OpenStack projects
type ProjectRepository repository.RestRepository

// GetAll returns all OpenStack projects
func (r *ProjectRepository) GetAll() ([]Project, error) {
	var response projectsWrapper
	restRequest := rest.Request{Endpoint: "/openstack/projects"}
	err := r.Client.Get(restRequest, &response)

	return response.Projects, err
}

// GetByID returns information about an OpenStack project by ID
func (r *ProjectRepository) GetByID(projectID string) (Project, error) {
	var response projectWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/projects/%s", projectID)}
	err := r.Client.Get(restRequest, &response)

	return response.Project, err
}

// Create creates a new OpenStack project
func (r *ProjectRepository) Create(project Project) error {
	restRequest := rest.Request{Endpoint: "/openstack/projects", Body: project}

	return r.Client.Post(restRequest)
}

// Update allows for updating the project name and description
func (r *ProjectRepository) Update(project Project) error {
	requestBody := projectWrapper{Project: project}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/projects/%s", project.ID), Body: requestBody}

	return r.Client.Put(restRequest)
}

// Handover initiates a handover procedure to another TransIP account.
func (r *ProjectRepository) Handover(projectID string, targetCustomerName string) error {
	requestBody := handoverRequest{Action: "handover", TargetCustomerName: targetCustomerName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/projects/%s", projectID), Body: requestBody}

	return r.Client.Patch(restRequest)
}

// Cancel will cancel an OpenStack project, deleting all the resources it contains
func (r *ProjectRepository) Cancel(projectID string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/openstack/projects/%s", projectID)}

	return r.Client.Delete(restRequest)
}
