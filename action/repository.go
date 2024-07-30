package action

import (
	"errors"
	"fmt"
	"strings"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Return error if action uuid was not found
var (
	ErrNoActionReturned = errors.New("no action uuid found for this action")
)

// Repository allows you to retrieve information about your running actions
type Repository repository.RestRepository

// GetActions allows you to gather all your actions
func (r *Repository) GetActions() ([]Action, error) {
	var response actionsWrapper
	restRequest := rest.Request{Endpoint: "/actions"}
	err := r.Client.Get(restRequest, &response)

	return response.Actions, err
}

// GetByID returns information about an action with a specific ID, also available for finished actions
func (r *Repository) GetByID(actionID string) (Action, error) {
	var response actionWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/actions/%s", actionID)}
	err := r.Client.Get(restRequest, &response)

	return response.Action, err
}

// GetChildActionsByParentID returns information about the child actions of a specific parent
func (r *Repository) GetChildActionsByParentID(actionID string) ([]Action, error) {
	var response actionsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/actions/children/%s", actionID)}
	err := r.Client.Get(restRequest, &response)

	return response.Actions, err
}

// ParseActionFromResponse parses the actionUuid from the content location header
func (r *Repository) ParseActionFromResponse(response rest.Response) (Action, error) {
	if response.ContentLocation == "" {
		return Action{}, ErrNoActionReturned
	}
	actionUUID := strings.Replace(response.ContentLocation, "/v6/actions/", "", 1)
	return r.GetByID(actionUUID)
}
