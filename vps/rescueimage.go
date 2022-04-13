package vps

import (
	"fmt"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// RescueImageRepository show the rescue images available to boot from. A rescue image
// can be used to recover a VPS in case it doesn't normally boot
type RescueImageRepository repository.RestRepository

// RescueImage contains the information known about a rescue image
type RescueImage struct {
	Name string `json:"name"`
}

// The RescueImageLinux and RescueImageBSD constants exists for convenience.
// There may be more or less images available for a VPS in the future.
const (
	RescueImageLinux = "RescueLinux"
	RescueImageBSD   = "RescueBSD"
)

// rescueImageWrapper struct contains RescueImages in it,
// this is solely used for unmarshalling
type rescueImageWrapper struct {
	RescueImages []RescueImage `json:"rescueImages"`
}

// bootRescueImageRequest struct contains the data needed to
// perform a BootRescueImage call. It exists soley for marshalling
type bootRescueImageRequest struct {
	Name string `json:"name"`
}

// GetAll returns all the rescue images available for a vps
func (r *RescueImageRepository) GetAll(vpsName string) ([]RescueImage, error) {
	var response rescueImageWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/rescue-images", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.RescueImages, err
}

// BootRescueImage boots a rescue image for the provided vps
func (r *RescueImageRepository) BootRescueImage(vpsName string, imageName string) error {
	requestBody := bootRescueImageRequest{imageName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/rescue-images", vpsName), Body: requestBody}
	return r.Client.Patch(restRequest)
}
