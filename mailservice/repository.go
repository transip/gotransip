package mailservice

import (
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

type Repository repository.RestRepository

// GetInformation allows you to gather detailed information
// regarding mail service usage and credentials using this API call
func (r *Repository) GetInformation() (MailServiceInformation, error) {
	var response mailServiceInformationWrapper
	restRequest := rest.RestRequest{Endpoint: "/mail-service"}
	err := r.Client.Get(restRequest, &response)

	return response.MailServiceInformation, err

}

// RegeneratePassword allows you to regenerate your transip mail service password
func (r *Repository) RegeneratePassword() error {
	restRequest := rest.RestRequest{Endpoint: "/mail-service"}

	return r.Client.Patch(restRequest)
}

// AddDNSEntriesDomains allows you to add default DNS records to you domains
// In order to reduce spam score, several DNS records should be added to your domains
func (r *Repository) AddDNSEntriesDomains(domainNames []string) error {
	var requestBody domainNamesWrapper
	requestBody.DomainNames = domainNames
	restRequest := rest.RestRequest{Endpoint: "/mail-service", Body:requestBody}

	return r.Client.Post(restRequest)
}

