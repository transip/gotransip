package domain

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository can be used to get a list of your domains
// order new ones and changing specific domain properties
type Repository repository.RestRepository

// GetAll returns all domains listed in your account
func (r *Repository) GetAll() ([]Domain, error) {
	var response domainsResponse
	err := r.Client.Get(rest.RestRequest{Endpoint: "/domains"}, &response)

	return response.Domains, err
}

// GetByDomainName returns an object for specific domain name]
// requires a domainName, for example: 'example.com'
func (r *Repository) GetByDomainName(domainName string) (Domain, error) {
	var response domainWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Domain, err
}

// Register allows you to registers a new domain
// You can set the contacts, nameservers and DNS entries immediately, but it’s not mandatory for registration
func (r *Repository) Register(domainRegister Register) error {
	restRequest := rest.RestRequest{Endpoint: "/domains", Body: &domainRegister}

	return r.Client.Post(restRequest)
}

// Transfer allows you to transfer a domain to TransIP using its transfer key
// (or ‘EPP code’) by specifying it in the authCode parameter
func (r *Repository) Transfer(domainTransfer Transfer) error {
	restRequest := rest.RestRequest{Endpoint: "/domains", Body: &domainTransfer}

	return r.Client.Post(restRequest)
}

// Update an existing domain.
// To apply or release a lock, change the IsTransferLocked property.
// To change tags, update the tags property
func (r *Repository) Update(domain Domain) error {
	requestBody := domainWrapper{Domain: domain}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s", domain.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Cancel cancels the specified domain
// Depending on the time you want to cancel the domain,
// specify gotransip.CancellationTimeEnd or gotransip.CancellationTimeImmediately for the endTime attribute
func (r *Repository) Cancel(domainName string, endTime gotransip.CancellationTime) error {
	var requestBody gotransip.CancellationRequest
	requestBody.EndTime = endTime
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s", domainName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetDomainBranding returns a Branding struct for the given domain
// Branding can be altered using the method below
func (r *Repository) GetDomainBranding(domainName string) (Branding, error) {
	var response domainBrandingWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/branding", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Branding, err
}

// UpdateDomainBranding allows you to change the branding information on a domain
func (r *Repository) UpdateDomainBranding(domainName string, branding Branding) error {
	requestBody := domainBrandingWrapper{Branding: branding}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/branding", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetContacts returns a list of contacts for the given domain name
func (r *Repository) GetContacts(domainName string) ([]WhoisContact, error) {
	var response contactsWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/contacts", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Contacts, err
}

// UpdateContacts allows you to replace the whois contacts currently on a domain
func (r *Repository) UpdateContacts(domainName string, contacts []WhoisContact) error {
	requestBody := contactsWrapper{Contacts: contacts}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/contacts", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetDnsEntries returns a list of all DNS entries for a domain by domainName
func (r *Repository) GetDnsEntries(domainName string) ([]DnsEntry, error) {
	var response dnsEntriesWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.DnsEntries, err
}

// AddDnsEntry allows you to add a single dns entry to a domain
func (r *Repository) AddDnsEntry(domainName string, dnsEntry DnsEntry) error {
	requestBody := dnsEntryWrapper{DnsEntry: dnsEntry}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Update the content of a single DNS entry,
// the dns entry is identified by the 'Name', 'Expire' and 'Type' properties of the DnsEntry struct
func (r *Repository) UpdateDnsEntry(domainName string, dnsEntry DnsEntry) error {
	requestBody := dnsEntryWrapper{DnsEntry: dnsEntry}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// ReplaceDnsEntries will wipe the entire zone replacing it with the given dns entries
func (r *Repository) ReplaceDnsEntries(domainName string, dnsEntries []DnsEntry) error {
	requestBody := dnsEntriesWrapper{DnsEntries: dnsEntries}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveDnsEntry allows you to remove a single DNS entry from a domain
func (r *Repository) RemoveDnsEntry(domainName string, dnsEntry DnsEntry) error {
	requestBody := dnsEntryWrapper{DnsEntry: dnsEntry}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetDnsEntries returns a list of all DNS entries for a domain by domainName
func (r *Repository) GetDnsSecEntries(domainName string) ([]DnsSecEntry, error) {
	var response dnsSecEntriesWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/dnssec", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.DnsSecEntries, err
}

// ReplaceDnsSecEntries allows you to replace all DNSSEC entries with the ones that are provided
func (r *Repository) ReplaceDnsSecEntries(domainName string, dnsSecEntries []DnsSecEntry) error {
	requestBody := dnsSecEntriesWrapper{DnsSecEntries: dnsSecEntries}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/dnssec", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetNameservers will list all nameservers currently set for a domain.
func (r *Repository) GetNameservers(domainName string) ([]Nameserver, error) {
	var response nameserversWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/nameservers", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Nameservers, err
}

// UpdateNameservers allows you to change the nameservers for a domain
func (r *Repository) UpdateNameservers(domainName string, nameservers []Nameserver) error {
	requestBody := nameserversWrapper{Nameservers: nameservers}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/nameservers", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Domain actions are kept track of by TransIP. Domain actions include, for example, changing nameservers
func (r *Repository) GetDomainAction(domainName string) (Action, error) {
	var response actionWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/actions", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Action, err
}

// Domain actions can fail due to wrong information, this method allows you to retry an action
func (r *Repository) RetryDomainAction(domainName string, authCode string, dnsEntries []DnsEntry, nameservers []Nameserver, contacts []WhoisContact) error {
	var requestBody retryActionWrapper
	requestBody.AuthCode = authCode
	requestBody.DnsEntries = dnsEntries
	requestBody.Nameservers = nameservers
	requestBody.Contacts = contacts
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/actions", domainName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// With this method you are able to cancel a domain action while it is still pending or being processed
func (r *Repository) CancelDomainAction(domainName string) error {
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/actions", domainName)}

	return r.Client.Delete(restRequest)
}

// GetSSLCertificates allows you to get a list of SSL certificates for a specific domain
func (r *Repository) GetSSLCertificates(domainName string) ([]SslCertificate, error) {
	var response certificatesWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/ssl", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Certificates, err
}

// GetSSLCertificateById allows you to get a single SSL certificate by id.
func (r *Repository) GetSSLCertificateById(domainName string, certificateId int64) (SslCertificate, error) {
	var response certificateWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/ssl/%d", domainName, certificateId)}
	err := r.Client.Get(restRequest, &response)

	return response.Certificate, err
}

// This method will return the WHOIS information for a domain name as a string
func (r *Repository) GetWHOIS(domainName string) (string, error) {
	var response whoisWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domains/%s/whois", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Whois, err
}

// OrderWhitelabel allows you to order a whitelabel account
// Note that you do not need to order a whitelabel account for every registered domain name
func (r *Repository) OrderWhitelabel() error {
	restRequest := rest.RestRequest{Endpoint: "/whitelabel"}

	return r.Client.Post(restRequest)
}

// GetAvailability method allows you to check the availability for a domain name
func (r *Repository) GetAvailability(domainName string) (Availability, error) {
	var response availabilityWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/domain-availability/%s", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Availability, err
}

// GetAvailability method allows you to check the availability for a domain name
func (r *Repository) GetAvailabilityForMultipleDomains(domainNames []string) ([]Availability, error) {
	var response availabilityListWrapper
	var requestBody multipleAvailabilityRequest
	requestBody.DomainNames = domainNames

	restRequest := rest.RestRequest{Endpoint: "/domain-availability", Body: requestBody}
	err := r.Client.Get(restRequest, &response)

	return response.AvailabilityList, err
}

// GetTLDs will return a list of all available TLDs currently offered by TransIP
func (r *Repository) GetTLDs() ([]Tld, error) {
	var response tldsWrapper
	restRequest := rest.RestRequest{Endpoint: "/tlds"}
	err := r.Client.Get(restRequest, &response)

	return response.Tlds, err
}

// GetTLDByTLD returns information about a specific TLD
// General details such as price, renewal price and minimum registration length are outlined
func (r *Repository) GetTLDByTLD(tld string) (Tld, error) {
	var response tldWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/tlds/%s", tld)}
	err := r.Client.Get(restRequest, &response)

	return response.Tld, err
}
