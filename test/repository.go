package test

import (
	"errors"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository can be used to test the api,
// do a simple ping and retrieve pong from the server
type Repository repository.RestRepository

var (
	// ErrWrongResponse will be thrown when the api's response does not match what we expect
	ErrWrongResponse = errors.New("Test api response doesn't equal 'pong'")
)

// Test will execute an api test and respond with an error if the test failed
func (r *Repository) Test() error {
	var testResponse APITest
	restRequest := rest.Request{Endpoint: "/api-test"}

	if err := r.Client.Get(restRequest, &testResponse); err != nil {
		return err
	}

	if testResponse.Response != "pong" {
		return ErrWrongResponse
	}

	return nil
}
