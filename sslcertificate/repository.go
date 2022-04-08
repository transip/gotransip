package sslcertificate

import (
	"fmt"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository for ordering and viewing SSL certificates
type Repository repository.RestRepository

// GetAll returns all SSL certificates in account
func (r *Repository) GetAll() ([]SSLCertificate, error) {
	var response sslcertificatesWrapper
	restRequest := rest.Request{Endpoint: "/ssl-certificates"}
	err := r.Client.Get(restRequest, &response)

	return response.Sslcertificates, err
}

// GetByID returns a SSL certificate by ID
func (r *Repository) GetByID(id int) (SSLCertificate, error) {
	var response wrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/ssl-certificates/%d", id)}
	err := r.Client.Get(restRequest, &response)

	return response.Sslcertificate, err
}

// GetDetails returns details for SSL certificate
func (r *Repository) GetDetails(id int) (Details, error) {
	var response detailsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/ssl-certificates/%d/details", id)}
	err := r.Client.Get(restRequest, &response)

	return response.Details, err
}

// Order a new SSL certificate
func (r *Repository) Order(orderRequest OrderSSLCertificateRequest) error {
	restRequest := rest.Request{Endpoint: "/ssl-certificates", Body: orderRequest}

	return r.Client.Post(restRequest)
}
