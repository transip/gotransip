package forwardingservice

import (
	"fmt"
	"net/url"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository can be used to list, modify and configure the forwarding services in your account
type Repository repository.RestRepository

// GetAll returns an array of all forwarding services in your account
func (r *Repository) GetAll() ([]ForwardingService, error) {
	var response forwardingServicesWrapper
	err := r.Client.Get(rest.Request{Endpoint: "/forwarding-services"}, &response)

	return response.ForwardingServices, err
}

// GetSelection returns a limited list of forwarding services,
// specify how many and which page/chunk of forwarding services you want to retrieve
func (r *Repository) GetSelection(page int, itemsPerPage int) ([]ForwardingService, error) {
	var response forwardingServicesWrapper
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/forwarding-services", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.ForwardingServices, err
}

// GetByDomainName returns the forwarding service of a domain name
func (r *Repository) GetByDomainName(domainName string) (ForwardingService, error) {
	var response forwardingServiceWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/forwarding-services/%s", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.ForwardingService, err
}

// Update allows you to update the settings of a forwarding service,
// ForwardTo and ForwardMethod are required. The Hostname selects the forwarding service,
// IsBlocked and SSLStatus are read only. Use EnableSSL and DisableSSL to change SSL.
func (r *Repository) Update(forwardingService ForwardingService) error {
	requestBody := forwardingServiceWrapper{ForwardingService: forwardingService}
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/forwarding-services/%s", forwardingService.Hostname),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// GetSSLStatus returns the SSL status of the forwarding service of a domain name
func (r *Repository) GetSSLStatus(domainName string) (SSLStatus, error) {
	var response sslStatusWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/forwarding-services/%s/ssl", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.SSLStatus, err
}

// EnableSSL enables SSL for the forwarding service of a domain name,
// after which it can be reached over HTTPS. The SSL status changes when the certificate has been issued.
func (r *Repository) EnableSSL(domainName string) error {
	return r.patchSSL(domainName, sslActionEnable)
}

// DisableSSL disables SSL for the forwarding service of a domain name,
// after which it can only be reached over HTTP
func (r *Repository) DisableSSL(domainName string) error {
	return r.patchSSL(domainName, sslActionDisable)
}

func (r *Repository) patchSSL(domainName string, action sslAction) error {
	requestBody := sslRequest{Action: action}
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/forwarding-services/%s/ssl", domainName),
		Body:     &requestBody,
	}

	return r.Client.Patch(restRequest)
}

// GetEmailForwardTo returns the email address that email for the domain name is forwarded to,
// it returns an empty string when email is not forwarded
func (r *Repository) GetEmailForwardTo(domainName string) (string, error) {
	var response emailForwardToWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/forwarding-services/%s/email", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.ForwardTo, err
}

// SetEmailForwardTo forwards email for the domain name to the given email address
func (r *Repository) SetEmailForwardTo(domainName string, forwardTo string) error {
	requestBody := emailForwardToWrapper{ForwardTo: forwardTo}
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/forwarding-services/%s/email", domainName),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// RemoveEmailForwardTo stops forwarding email for the domain name
func (r *Repository) RemoveEmailForwardTo(domainName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/forwarding-services/%s/email", domainName)}

	return r.Client.Delete(restRequest)
}
