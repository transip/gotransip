package email

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository for creating and modifying mailboxes, mail forwards and mail lists
type Repository repository.RestRepository

// GetMailboxesByDomainName returns all mailboxes by domain name
func (r *Repository) GetMailboxesByDomainName(domainName string) ([]Mailbox, error) {
	var mailboxResponse mailboxesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mailboxes", domainName)}
	err := r.Client.Get(restRequest, &mailboxResponse)

	return mailboxResponse.Mailboxes, err
}

// GetMailboxByEmailAddress returns a mailbox
func (r *Repository) GetMailboxByEmailAddress(emailAddress string) (Mailbox, error) {
	_, err := mail.ParseAddress(emailAddress)

	if err != nil {
		return Mailbox{}, err
	}

	components := strings.Split(emailAddress, "@")
	domainName := components[1]

	var mailboxResponse mailboxWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mailboxes/%s", domainName, emailAddress)}
	err = r.Client.Get(restRequest, &mailboxResponse)

	return mailboxResponse.Mailbox, err
}

// CreateMailbox creates a new mailbox
func (r *Repository) CreateMailbox(domainName string, createRequest CreateMailboxRequest) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mailboxes", domainName), Body: createRequest}

	return r.Client.Post(restRequest)
}

// UpdateMailbox updates a mailbox
func (r *Repository) UpdateMailbox(emailAddress string, updateRequest UpdateMailboxRequest) error {
	_, err := mail.ParseAddress(emailAddress)

	if err != nil {
		return err
	}

	components := strings.Split(emailAddress, "@")
	domainName := components[1]

	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mailboxes/%s", domainName, emailAddress), Body: updateRequest}

	return r.Client.Put(restRequest)
}

// DeleteMailbox deletes a mailbox
func (r *Repository) DeleteMailbox(emailAddress string) error {
	_, err := mail.ParseAddress(emailAddress)

	if err != nil {
		return err
	}

	components := strings.Split(emailAddress, "@")
	domainName := components[1]

	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mailboxes/%s", domainName, emailAddress)}

	return r.Client.Delete(restRequest)
}

// GetMailforwardsByDomainName returns all mail forwards by domain name
func (r *Repository) GetMailforwardsByDomainName(domainName string) ([]Mailforward, error) {
	var response mailforwardsWrappper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-forwards", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Mailforwards, err
}

// GetMailforwardByDomainNameAndID returns a mailbox
func (r *Repository) GetMailforwardByDomainNameAndID(domainName string, mailforwardID int) (Mailforward, error) {
	var response mailforwardWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-forwards/%d", domainName, mailforwardID)}
	err := r.Client.Get(restRequest, &response)

	return response.Mailforward, err
}

// CreateMailforward creates a mail forward
func (r *Repository) CreateMailforward(domainName string, createRequest CreateMailforwardRequest) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-forwards", domainName), Body: createRequest}

	return r.Client.Post(restRequest)
}

// UpdateMailforward updates a mail forward
func (r *Repository) UpdateMailforward(domainName string, forwardID int, updateRequest UpdateMailforwardRequest) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-forwards/%d", domainName, forwardID), Body: updateRequest}

	return r.Client.Put(restRequest)
}

// DeleteMailforward deletes a mail forward
func (r *Repository) DeleteMailforward(domainName string, forwardID int) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-forwards/%d", domainName, forwardID)}

	return r.Client.Delete(restRequest)
}

// GetMaillistsByDomainName returns all maillists by domain name
func (r *Repository) GetMaillistsByDomainName(domainName string) ([]Maillist, error) {
	var response maillistsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-lists", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Maillists, err
}

// GetMaillistByDomainNameAndID returns a mail list by domain name and ID
func (r *Repository) GetMaillistByDomainNameAndID(domainName string, maillistID int) (Maillist, error) {
	var response maillistWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-lists/%d", domainName, maillistID)}
	err := r.Client.Get(restRequest, &response)

	return response.Maillist, err
}

// CreateMaillist creates a mail list
func (r *Repository) CreateMaillist(domainName string, createRequest CreateMaillistRequest) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-lists", domainName), Body: createRequest}

	return r.Client.Post(restRequest)
}

// UpdateMaillist updates a mail list
func (r *Repository) UpdateMaillist(domainName string, maillistID int, updateRequest UpdateMaillistRequest) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-lists/%d", domainName, maillistID), Body: updateRequest}

	return r.Client.Put(restRequest)
}

// DeleteMaillist deletes a mail list
func (r *Repository) DeleteMaillist(domainName string, maillistID int) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/email/%s/mail-lists/%d", domainName, maillistID)}

	return r.Client.Delete(restRequest)
}
