package domain

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest/request"
)

// Repository can be used to get a list of your domains
// order new ones and changing specific domain properties
type Repository repository.RestRepository

// GetAll returns all domains listed in your account
func (r *Repository) GetAll() ([]Domain, error) {
	var response domainsResponse
	err := r.Client.Get(request.RestRequest{Endpoint: "/domains"}, &response)

	return response.Domains, err
}

// GetByDomainName returns an object for specific domain name]
// requires a domainName, for example: 'example.com'
func (r *Repository) GetByDomainName(domainName string) (Domain, error) {
	var response domainWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/domains/%s", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Domain, err
}

// Register allows you to registers a new domain
// You can set the contacts, nameservers and DNS entries immediately, but it’s not mandatory for registration
func (r *Repository) Register(domainRegister Register) error {
	restRequest := request.RestRequest{Endpoint: "/domains", Body: &domainRegister}

	return r.Client.Post(restRequest)
}

// Transfer allows you to transfer a domain to TransIP using its transfer key
// (or ‘EPP code’) by specifying it in the authCode parameter
func (r *Repository) Transfer(domainTransfer Transfer) error {
	restRequest := request.RestRequest{Endpoint: "/domains", Body: &domainTransfer}

	return r.Client.Post(restRequest)
}

// Update an existing domain.
// To apply or release a lock, change the IsTransferLocked property.
// To change tags, update the tags property
func (r *Repository) Update(domain Domain) error {
	requestBody := domainWrapper{Domain: domain}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/domains/%s", domain.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Cancel cancels the specified domain
// Depending on the time you want to cancel the domain,
// specify gotransip.CancellationTimeEnd or gotransip.CancellationTimeImmediately for the endTime attribute
func (r *Repository) Cancel(domainName string, endTime gotransip.CancellationTime) error {
	var requestBody gotransip.CancellationRequest
	requestBody.EndTime = endTime
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/domains/%s", domainName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetDomainBranding returns a Branding struct for the given domain
// Branding can be altered using the method below
func (r *Repository) GetDomainBranding(domainName string) (Branding, error) {
	var response domainBrandingWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/branding", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Branding, err
}

// UpdateDomainBranding allows you to change the branding information on a domain
func (r *Repository) UpdateDomainBranding(domainName string, branding Branding) error {
	requestBody := domainBrandingWrapper{Branding: branding}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/branding", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetContacts returns a list of contacts for the given domain name
func (r *Repository) GetContacts(domainName string) ([]WhoisContact, error) {
	var response contactsWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/contacts", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Contacts, err
}

// UpdateContacts allows you to replace the whois contacts currently on a domain
func (r *Repository) UpdateContacts(domainName string, contacts []WhoisContact) error {
	requestBody := contactsWrapper{Contacts: contacts}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/contacts", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}
