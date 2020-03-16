package traffic

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository allows you to get information about your usage in your traffic pool
// you can retrieve this information globally or per vps
type Repository repository.RestRepository

// GetTrafficPool returns all the traffic of your VPSes combined, overusage will also be billed based on this information
func (r *Repository) GetTrafficPool() (TrafficInformation, error) {
	var response trafficWrapper
	restRequest := rest.RestRequest{Endpoint: "/traffic"}
	err := r.Client.Get(restRequest, &response)

	return response.TrafficInformation, err
}

// GetTrafficInformationForVps allows you to get specific traffic information for a given VPS
func (r *Repository) GetTrafficInformationForVps(vpsName string) (TrafficInformation, error) {
	var response trafficWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/traffic/%s", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.TrafficInformation, err
}
