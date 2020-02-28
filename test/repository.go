package test

import (
	"errors"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest/request"
)

// Repository can be used to test the api,
// do a simple ping and retrieve pong from the server
type Repository repository.RestRepository

// Test will execute a test and respond with an error if the test failed
func (r *Repository) Test() error {
	var testResponse ApiTest
	restRequest := request.RestRequest{Endpoint: "/api-test"}

	err := r.Client.Get(restRequest, &testResponse)
	if err != nil {
		return err
	}

	if testResponse.Response != "pong" {
		return errors.New("Test api response doesn't equal 'pong'")
	}

	return nil
}
