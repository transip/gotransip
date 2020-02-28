package repository

import (
	"github.com/transip/gotransip/v6/rest/request"
)

// Client interface, this is the client interface as far as other packages should care about
type Client interface {
	Get(request request.RestRequest, response interface{}) error
	Put(request request.RestRequest) error
	Post(request request.RestRequest) error
	Delete(request request.RestRequest) error
}

// RestRepository is the struct which is going to be used by all other repositories in the gotransip package
type RestRepository struct {
	// we have a client that adheres to the client interface
	Client Client
}
