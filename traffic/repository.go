package traffic

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest/request"
)

type Repository repository.RestRepository

// GetTrafficPool returns all the traffic of your VPSes combined, overusage will also be billed based on this information
func (r *Repository) GetTrafficPool() (TrafficInformation, error) {
	var response trafficWrapper
	restRequest := request.RestRequest{Endpoint: "/traffic"}
	err := r.Client.Get(restRequest, &response)

	return response.TrafficInformation, err
}

// GetTrafficInformationForVps allows you to get specific traffic information for a given VPS
func (r *Repository) GetTrafficInformationForVps(vpsName string) (TrafficInformation, error) {
	var response trafficWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/traffic/%s", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.TrafficInformation, err
}
