package haip

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest/request"
)

type Repository repository.RestRepository

func (r *Repository) GetAll() ([]Haip, error) {
	var response HaipsResponse
	err := r.Client.Get(request.RestRequest{Endpoint: "/haips"}, &response)

	return response.Haips, err
}

func (r *Repository) GetByHaipName(haipName string) (Haip, error) {
	var response haipWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/haips/%s", haipName)}
	err := r.Client.Get(restRequest, &response)

	return response.Haip, err
}

func (r *Repository) OrderHaip(productName string, description string) error {
	requestBody := haipOrderWrapper{ProductName: productName, Description: description}
	err := r.Client.Post(request.RestRequest{Endpoint: "/haips", Body: requestBody})

	return err
}

func (r *Repository) UpdateHaip(haipName string, configurations Haip) error {
	requestBody := haipWrapper{Haip: configurations}
	err := r.Client.Patch(request.RestRequest{Endpoint: fmt.Sprintf("/haips/%s", haipName), Body: requestBody})
	return err
}

func (r *Repository) CancelHaip(haipName string, endTime gotransip.CancellationTime) error {
	var requestBody gotransip.CancellationRequest
	requestBody.EndTime = endTime
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/haips/%s", haipName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

func (r *Repository) GetAllCertificates(haipName string) ([]HaipCertificate, error)  {
	var response HaipCertificates
	err := r.Client.Get(request.RestRequest{Endpoint: "/haips"}, &response)

	return response.Haips, err
}
