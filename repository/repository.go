package repository

import (
	"github.com/transip/gotransip/v6/rest"
)

// Client interface, this is the client interface as far as other packages should care about
type Client interface {
	// Executes a GET rest request and returns the response into the destination struct
	Get(request rest.RestRequest, dest interface{}) error
	// Executes a PUT request, not expecting any response from the api server
	Put(request rest.RestRequest) error
	// Executes a POST request, not expecting any response from the api server
	Post(request rest.RestRequest) error
	// Executes a DELETE request, not expecting any response from the api server
	Delete(request rest.RestRequest) error
	// Executes a PATCH request, not expecting any response from the api server
	Patch(restRequest rest.RestRequest) error
}

// RestRepository is the struct which is going to be used by all other repositories in the gotransip package
type RestRepository struct {
	// we have a client that follows the Client interface
	Client Client
}
