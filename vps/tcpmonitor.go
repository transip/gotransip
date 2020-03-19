package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net"
)

// TCPMonitorRepository allows you to manage all tcp monitor and tcp monitor contact api actions
// like listing, getting information, adding, updating, deleting tcp monitors
// updating, creating, deleting contacts
type TCPMonitorRepository repository.RestRepository

// TCPMonitor struct for TCPMonitor
type TCPMonitor struct {
	// IP Address that is monitored
	IPAddress net.IP `json:"ipAddress"`
	// Title of the monitor
	Label string `json:"label"`
	// Ports that are monitored
	Ports []int `json:"ports"`
	// Checking interval in minutes (numbers 1-6)
	Interval int `json:"interval"`
	// Allowed time outs (numbers 1-5)
	AllowedTimeouts int `json:"allowedTimeouts"`
	// TCPMonitorContact that will be notified for this monitor
	Contacts []TCPMonitorContact `json:"contacts"`
	// The hours when the TCP monitoring is ignored (no notifications are sent out)
	IgnoreTimes []IgnoreTime `json:"ignoreTimes"`
}

// TCPMonitorContact struct for TCPMonitorContact
type TCPMonitorContact struct {
	// Monitoring contact id
	ID int64 `json:"id"`
	// Send emails to contact
	EnableEmail bool `json:"enableEmail"`
	// Send SMS text messages to contact
	EnableSMS bool `json:"enableSMS"`
}

// IgnoreTime struct for IgnoreTime
type IgnoreTime struct {
	// Start from (24 hour format)
	TimeFrom string `json:"timeFrom"`
	// End at (24 hour format)
	TimeTo string `json:"timeTo"`
}

// MonitoringContact struct for MonitoringContact
type MonitoringContact struct {
	// ID number of the contact
	ID int64 `json:"id,omitempty"`
	// Name of the contact
	Name string `json:"name"`
	// Telephone number of the contact
	Telephone string `json:"telephone"`
	// Email address of the contact
	Email string `json:"email"`
}

// GetTCPMonitors returns an overview of all existing monitors attached to a VPS
func (r *TCPMonitorRepository) GetTCPMonitors(vpsName string) ([]TCPMonitor, error) {
	var response tcpMonitorsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.TCPMonitors, err
}

// CreateTCPMonitor allows you to create a tcp monitor and specify which ports you would like to monitor
// to get a better grip on which fields exist and which can be changes have a look at the TCPMonitor struct
// or see the documentation: https://api.transip.nl/rest/docs.html#vps-tcp-monitors-post
func (r *TCPMonitorRepository) CreateTCPMonitor(vpsName string, tcpMonitor TCPMonitor) error {
	requestBody := tcpMonitorWrapper{TCPMonitor: tcpMonitor}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateTCPMonitor allows you to update your monitor settings for a given tcp monitored ip
func (r *TCPMonitorRepository) UpdateTCPMonitor(vpsName string, tcpMonitor TCPMonitor) error {
	requestBody := tcpMonitorWrapper{TCPMonitor: tcpMonitor}
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors/%s", vpsName, tcpMonitor.IPAddress.String()),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// RemoveTCPMonitor allows you to remove a tcp monitor for a specific ip address on a specifc VPS
func (r *TCPMonitorRepository) RemoveTCPMonitor(vpsName string, ip net.IP) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors/%s", vpsName, ip.String())}

	return r.Client.Delete(restRequest)
}

// GetContacts returns a list of all your monitoring contacts
func (r *TCPMonitorRepository) GetContacts() ([]MonitoringContact, error) {
	var response contactsWrapper
	restRequest := rest.Request{Endpoint: "/monitoring-contacts"}
	err := r.Client.Get(restRequest, &response)

	return response.Contacts, err
}

// CreateContact allows you to add a new contact which could be used by the tcp monitoring
func (r *TCPMonitorRepository) CreateContact(contact MonitoringContact) error {
	restRequest := rest.Request{Endpoint: "/monitoring-contacts", Body: &contact}

	return r.Client.Post(restRequest)
}

// UpdateContact updates the specified contact
func (r *TCPMonitorRepository) UpdateContact(contact MonitoringContact) error {
	requestBody := contactWrapper{Contact: contact}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/monitoring-contacts/%d", contact.ID), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveContact allows you to delete a specific contact by id
func (r *TCPMonitorRepository) RemoveContact(contactID int64) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/monitoring-contacts/%d", contactID)}

	return r.Client.Delete(restRequest)
}
