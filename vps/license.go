package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// LicenseRepository allows you to manage all vps license api actions
// like listing, getting license keys, ordering, updating, deleting licenses
type LicenseRepository repository.RestRepository

// LicenseType is one of the following strings
// 'addon', 'operating-system'
type LicenseType string

// Definition of all of the possible license types
const (
	// Addon licenses can be purchased individually
	LicenseTypeAddon LicenseType = "addon"
	// Operating system licenses cannot be directly purchased, or cancelled,
	// they are attached to your VPS the moment you install an operating system that requires a license.
	// Operating systems such as Plesk, DirectAdmin, cPanel and etc need a valid license.
	// An operating system license can only be upgraded or downgraded
	LicenseTypeOperatingSystem LicenseType = "operating-system"
)

// Licenses struct contains Active, Available and Cancellable License structs in it
type Licenses struct {
	// A list of licenses active on your VPS
	Active []License `json:"active"`
	// A list of available licenses that you can order for your VPS
	Available []LicenseProduct `json:"available"`
	// A list of licenses active on your VPS that you can cancel
	Cancellable []License `json:"cancellable"`
}

// License struct for a vps license
type License struct {
	// The License id
	ID int64 `json:"id"`
	// License name
	Name string `json:"name"`
	// Price in cents
	Price int `json:"price"`
	// Recurring price in cents
	RecurringPrice int `json:"recurringPrice"`
	// License type: 'operating-system', 'addon'
	Type LicenseType `json:"type"`
	// Quantity already purchased
	Quantity int `json:"quantity"`
	// Maximum quantity you are allowed to purchase
	MaxQuantity int `json:"maxQuantity"`
	// License keys belonging to this License
	Keys []LicenseKey `json:"keys"`
}

// LicenseProduct struct for a orderable license
type LicenseProduct struct {
	// License name
	Name string `json:"name"`
	// Price in cents
	Price int `json:"price"`
	// Recurring price in cents
	RecurringPrice int `json:"recurringPrice"`
	// License type: 'operating-system', 'addon'
	Type LicenseType `json:"type"`
	// Maximum quantity you are allowed to purchase
	MaxQuantity int `json:"maxQuantity"`
}

// licensesWrapper struct contains Licenses in it,
// this is solely used for unmarshalling
type licensesWrapper struct {
	Licenses Licenses `json:"licenses"`
}

// licenseReplaceRequest struct contains a License Name in it,
// this is solely used for marshalling
type licenseReplaceRequest struct {
	NewLicenseName string `json:"newLicenseName"`
}

// LicenseKey struct contains a license key and name for a specific License
type LicenseKey struct {
	// License name
	Name string `json:"name"`
	// License key
	Key string `json:"key"`
}

// The LicenseOrder struct is used for ordering a new license for a VPS
type LicenseOrder struct {
	// Name of the license that you want to order
	LicenseName string `json:"licenseName"`
	// Quantity of this license that you want to order
	Quantity int `json:"quantity"`
}

// ReplaceLicenseRequest this struct is used for replacing a 'operating-system' type license with a new license
type ReplaceLicenseRequest struct {
	// The License id
	LicenseID int64
	// NewLicenseName is the name of the license with which you want to replace the current 'operating-system' license
	NewLicenseName string
}

// GetAll returns a struct with 'cancellable', 'available' and 'active' licenses in it for the given VPS
//
// Operating system licenses cannot be directly purchased, or cancelled,
// they are attached to your VPS the moment you install an operating system that requires a license.
// Operating systems such as Plesk, DirectAdmin, cPanel and etc need a valid license.
// An operating system license can only be upgraded or downgraded by using the Update an operating system license API call.
//
// Addon licenses can be purchased individually through the Order an addon license API call.
func (r *LicenseRepository) GetAll(vpsName string) (Licenses, error) {
	var response licensesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/licenses", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Licenses, err
}

// Order allows you to order a new license for a given Vps.
// In order to purchase an addon license for your VPS, use this API call.
// The licenses that can be ordered can be requested using the get licenses api call
func (r *LicenseRepository) Order(vpsName string, order LicenseOrder) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/licenses", vpsName), Body: &order}

	return r.Client.Post(restRequest)
}

// Replace allows you to switch between operating system licenses
//
// Provide your desired license name in the licenseName parameter for either to upgrade or downgrade.
// Only operating system licenses can be passed through this API call.
func (r *LicenseRepository) Replace(vpsName string, request ReplaceLicenseRequest) error {
	requestBody := licenseReplaceRequest{NewLicenseName: request.NewLicenseName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/licenses/%d", vpsName, request.LicenseID), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Cancel allows you to cancel a license for a given vps by its id and the VPS name.
// Operating system licenses cannot be cancelled.
func (r *LicenseRepository) Cancel(vpsName string, licenseID int64) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/licenses/%d", vpsName, licenseID)}

	return r.Client.Delete(restRequest)
}
