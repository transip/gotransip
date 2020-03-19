package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net/url"
)

// PrivateNetworkRepository allows you to manage all private network api actions
// like listing, ordering, canceling, getting information, updating description, attaching and detaching vpses
type PrivateNetworkRepository repository.RestRepository

// PrivateNetwork struct for a PrivateNetwork
type PrivateNetwork struct {
	// The unique private network name
	Name string `json:"name"`
	// The custom name that can be set by customer
	Description string `json:"description"`
	// If the Private Network is administratively blocked
	IsBlocked bool `json:"isBlocked"`
	// When locked, another process is already working with this private network
	IsLocked bool `json:"isLocked"`
	// The VPSes in this private network
	VpsNames []string `json:"vpsNames,omitempty"`
}

// GetAll returns a list of all your private networks
func (r *PrivateNetworkRepository) GetAll() ([]PrivateNetwork, error) {
	var response privateNetworksWrapper
	restRequest := rest.Request{Endpoint: "/private-networks"}
	err := r.Client.Get(restRequest, &response)

	return response.PrivateNetworks, err
}

// GetSelection returns a limited list of private networks,
// specify how many and which page/chunk of private networks you want to retrieve
func (r *PrivateNetworkRepository) GetSelection(page int, itemsPerPage int) ([]PrivateNetwork, error) {
	var response privateNetworksWrapper
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/private-networks", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.PrivateNetworks, err
}

// GetByName allows you to get a specific PrivateNetwork by name
func (r *PrivateNetworkRepository) GetByName(privateNetworkName string) (PrivateNetwork, error) {
	var response privateNetworkWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName)}
	err := r.Client.Get(restRequest, &response)

	return response.PrivateNetwork, err
}

// Order allows you to order new private network with a given description
func (r *PrivateNetworkRepository) Order(description string) error {
	requestBody := privateNetworkOrderRequest{Description: description}
	restRequest := rest.Request{Endpoint: "/private-networks", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Update allows you to update the private network.
// You can change the description by changing the Description field
// on the PrivateNetwork struct Updating it using this function.
func (r *PrivateNetworkRepository) Update(privateNetwork PrivateNetwork) error {
	requestBody := privateNetworkWrapper{PrivateNetwork: privateNetwork}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetwork.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// AttachVps allows you to attach a VPS to a PrivateNetwork
func (r *PrivateNetworkRepository) AttachVps(vpsName string, privateNetworkName string) error {
	requestBody := privateNetworkActionwrapper{Action: "attachvps", VpsName: vpsName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// DetachVps allows you to detach a VPS from a PrivateNetwork
func (r *PrivateNetworkRepository) DetachVps(vpsName string, privateNetworkName string) error {
	requestBody := privateNetworkActionwrapper{Action: "detachvps", VpsName: vpsName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Cancel allows you to cancel a private network
func (r *PrivateNetworkRepository) Cancel(privateNetworkName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}
