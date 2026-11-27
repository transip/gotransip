package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/transip/gotransip/v6/rest"
)

// verificationsItemsPerPageForGetAll is the page size used by GetVerifications to retrieve all verifications,
// the api paginates this resource by default
const verificationsItemsPerPageForGetAll = 100

// VerificationStatus is the status of a domain verification.
// Possible statuses are: 'cancelled', 'failed', 'pending' and 'success'
type VerificationStatus string

// define all possible verification statuses
const (
	// VerificationStatusCancelled is the status for a verification that was cancelled
	VerificationStatusCancelled VerificationStatus = "cancelled"
	// VerificationStatusFailed is the status for a verification that has failed
	VerificationStatusFailed VerificationStatus = "failed"
	// VerificationStatusPending is the status for a verification that is waiting for a response
	VerificationStatusPending VerificationStatus = "pending"
	// VerificationStatusSuccess is the status for a verification that has succeeded
	VerificationStatusSuccess VerificationStatus = "success"
)

// verificationsWrapper struct contains a Verification list in it,
// this is solely used for unmarshalling/marshalling
type verificationsWrapper struct {
	Verifications []Verification `json:"verifications"`
}

// verificationWrapper struct contains a Verification in it,
// this is solely used for unmarshalling/marshalling
type verificationWrapper struct {
	Verification Verification `json:"verification"`
}

// Verification struct for a domain verification, used to verify actions on a domain, for example a change of registrant.
// The Contact and Domain are only filled when they are part of the response of the api,
// use GetVerificationByID to retrieve those.
type Verification struct {
	// The unique id of this verification
	ID int64 `json:"id"`
	// The name of the domain this verification belongs to
	DomainName string `json:"domainName"`
	// The status of this verification, 'cancelled', 'failed', 'pending' or 'success'
	Status VerificationStatus `json:"status"`
	// The creation date of this verification as a unix timestamp
	CreatedAt int64 `json:"createdAt"`
	// The last update date of this verification as a unix timestamp, 0 when it was never updated
	UpdatedAt int64 `json:"updatedAt"`
	// The date the response was received as a unix timestamp, 0 when no response was received yet
	ResponseReceivedAt int64 `json:"responseReceivedAt"`
	// The contact that was submitted for this verification
	Contact VerificationContact `json:"contact"`
	// The domain this verification belongs to
	Domain VerificationDomain `json:"domain"`
}

// VerificationContact struct for the contact that was submitted for a verification.
// This is read only and not a WhoisContact because it contains a middle name.
type VerificationContact struct {
	// The type of this contact, 'registrant', 'administrative' or 'technical'
	Type string `json:"type"`
	// The first name of this contact
	FirstName string `json:"firstName"`
	// The middle name of this contact
	MiddleName string `json:"middleName"`
	// The last name of this contact
	LastName string `json:"lastName"`
	// The company name of this contact, in case of a company
	CompanyName string `json:"companyName"`
	// The kvk number of this contact, in case of a company
	CompanyKvk string `json:"companyKvk"`
	// The company type of this contact, in case of a company
	CompanyType string `json:"companyType"`
	// The street of the address of this contact
	Street string `json:"street"`
	// The number part of the address of this contact
	Number string `json:"number"`
	// The postal code part of the address of this contact
	PostalCode string `json:"postalCode"`
	// The city part of the address of this contact
	City string `json:"city"`
	// The phone number of this contact
	PhoneNumber string `json:"phoneNumber"`
	// The fax number of this contact
	FaxNumber string `json:"faxNumber"`
	// The email address of this contact
	Email string `json:"email"`
	// The country of this contact, one of the ISO 3166-1 2 letter country codes
	Country string `json:"country"`
}

// UnmarshalJSON parses a JSON object, or the empty JSON array the api returns when there is no contact
func (c *VerificationContact) UnmarshalJSON(input []byte) error {
	if isEmptyJSONArray(input) {
		*c = VerificationContact{}
		return nil
	}

	type plainContact VerificationContact
	return json.Unmarshal(input, (*plainContact)(c))
}

// VerificationDomain struct for the domain a verification belongs to
type VerificationDomain struct {
	// The name of this domain
	Name string `json:"name"`
	// The registration status of this domain
	RegistrationStatus string `json:"registrationStatus"`
}

// UnmarshalJSON parses a JSON object, or the empty JSON array the api returns when there is no domain
func (d *VerificationDomain) UnmarshalJSON(input []byte) error {
	if isEmptyJSONArray(input) {
		*d = VerificationDomain{}
		return nil
	}

	type plainDomain VerificationDomain
	return json.Unmarshal(input, (*plainDomain)(d))
}

// GetVerifications returns all verifications in your account, optionally filtered on a domain name and statuses.
// Use an empty domainName and no statuses to not filter. The Contact and Domain of the verifications are
// not filled, use GetVerificationByID to retrieve those.
func (r *Repository) GetVerifications(domainName string, statuses []VerificationStatus) ([]Verification, error) {
	var verifications []Verification

	for page := 1; ; page++ {
		selection, err := r.GetVerificationsSelection(page, verificationsItemsPerPageForGetAll, domainName, statuses)
		if err != nil {
			return nil, err
		}

		verifications = append(verifications, selection...)
		if len(selection) < verificationsItemsPerPageForGetAll {
			return verifications, nil
		}
	}
}

// GetVerificationsSelection returns a limited list of verifications, optionally filtered on a domain name and statuses.
// Specify how many and which page/chunk of verifications you want to retrieve.
func (r *Repository) GetVerificationsSelection(page int, itemsPerPage int, domainName string, statuses []VerificationStatus) ([]Verification, error) {
	var response verificationsWrapper
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	if domainName != "" {
		params.Set("domainName", domainName)
	}

	if len(statuses) > 0 {
		// the api expects the statuses as a comma separated list
		statusStrings := make([]string, len(statuses))
		for i, status := range statuses {
			statusStrings[i] = string(status)
		}
		params.Set("statuses", strings.Join(statusStrings, ","))
	}

	restRequest := rest.Request{Endpoint: "/domain-verifications", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	// a list item is a summary, the contact and domain are only part of GetVerificationByID
	for i := range response.Verifications {
		response.Verifications[i].Contact = VerificationContact{}
		response.Verifications[i].Domain = VerificationDomain{}
	}

	return response.Verifications, err
}

// GetVerificationByID returns a single verification, including its contact and domain
func (r *Repository) GetVerificationByID(id int64) (Verification, error) {
	var response verificationWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domain-verifications/%d", id)}
	err := r.Client.Get(restRequest, &response)

	return response.Verification, err
}

// isEmptyJSONArray returns whether the input is a JSON array without items.
// The api is written in PHP, which encodes an empty object as an empty array.
func isEmptyJSONArray(input []byte) bool {
	trimmed := bytes.TrimSpace(input)
	if len(trimmed) == 0 || trimmed[0] != '[' {
		return false
	}

	var items []interface{}
	return json.Unmarshal(trimmed, &items) == nil && len(items) == 0
}
