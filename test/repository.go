package test

import (
	"errors"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository can be used to test the api,
// do a simple ping and retrieve pong from the server
type Repository repository.RestRepository

// Test will execute an api test and respond with an error if the test failed
func (r *Repository) Test() error {
	var testResponse ApiTest
	restRequest := rest.RestRequest{Endpoint: "/api-test"}

	if err := r.Client.Get(restRequest, &testResponse); err != nil {
		return err
	}

	if testResponse.Response != "pong" {
		return errors.New("Test api response doesn't equal 'pong'")
	}

	return nil
}
