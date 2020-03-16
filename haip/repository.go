package haip

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

type Repository repository.RestRepository

func (r *Repository) GetAll() ([]Haip, error) {
	var response haipsWrapper
	err := r.Client.Get(rest.RestRequest{Endpoint: "/haips"}, &response)

	return response.Haips, err
}

func (r *Repository) GetByHaipName(haipName string) (Haip, error) {
	var response haipWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/haips/%s", haipName)}
	err := r.Client.Get(restRequest, &response)

	return response.Haip, err
}

func (r *Repository) OrderHaip(productName string, description string) error {
	requestBody := haipOrderWrapper{ProductName: productName, Description: description}

	return r.Client.Post(rest.RestRequest{Endpoint: "/haips", Body: requestBody})
}

func (r *Repository) UpdateHaip(haipName string, configurations Haip) error {
	requestBody := haipWrapper{Haip: configurations}

	return r.Client.Patch(rest.RestRequest{Endpoint: fmt.Sprintf("/haips/%s", haipName), Body: requestBody})
}

func (r *Repository) CancelHaip(haipName string, endTime gotransip.CancellationTime) error {
	var requestBody gotransip.CancellationRequest
	requestBody.EndTime = endTime
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/haips/%s", haipName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

func (r *Repository) GetAllCertificates(haipName string) ([]HaipCertificate, error)  {
	var response certificatesWrapper
	err := r.Client.Get(rest.RestRequest{Endpoint: fmt.Sprintf("/haips/%s/certificates", haipName)}, &response)

	return response.Certificates, err
}
