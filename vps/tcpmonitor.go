package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net"
)

// TcpMonitorRepository allows you to manage all tcp monitor and tcp monitor contact api actions
// like listing, getting information, adding, updating, deleting tcp monitors
// updating, creating, deleting contacts
type TcpMonitorRepository repository.RestRepository

// TcpMonitor struct for TcpMonitor
type TcpMonitor struct {
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
	// TcpMonitorContact that will be notified for this monitor
	Contacts []TcpMonitorContact `json:"contacts"`
	// The hours when the TCP monitoring is ignored (no notifications are sent out)
	IgnoreTimes []IgnoreTime `json:"ignoreTimes"`
}

// TcpMonitorContact struct for TcpMonitorContact
type TcpMonitorContact struct {
	// Monitoring contact id
	Id int64 `json:"id"`
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
	// Id number of the contact
	Id int64 `json:"id,omitempty"`
	// Name of the contact
	Name string `json:"name"`
	// Telephone number of the contact
	Telephone string `json:"telephone"`
	// Email address of the contact
	Email string `json:"email"`
}

// GetTCPMonitors returns an overview of all existing monitors attached to a VPS
func (r *TcpMonitorRepository) GetTCPMonitors(vpsName string) ([]TcpMonitor, error) {
	var response tcpMonitorsWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.TcpMonitors, err
}

// CreateTCPMonitor allows you to create a tcp monitor and specify which ports you would like to monitor
// to get a better grip on which fields exist and which can be changes have a look at the TcpMonitor struct
// or see the documentation: https://api.transip.nl/rest/docs.html#vps-tcp-monitors-post
func (r *TcpMonitorRepository) CreateTCPMonitor(vpsName string, tcpMonitor TcpMonitor) error {
	requestBody := tcpMonitorWrapper{TcpMonitor: tcpMonitor}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateTCPMonitor allows you to update your monitor settings for a given tcp monitored ip
func (r *TcpMonitorRepository) UpdateTCPMonitor(vpsName string, tcpMonitor TcpMonitor) error {
	requestBody := tcpMonitorWrapper{TcpMonitor: tcpMonitor}
	restRequest := rest.RestRequest{
		Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors/%s", vpsName, tcpMonitor.IPAddress.String()),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// RemoveTCPMonitor allows you to remove a tcp monitor for a specific ip address on a specifc VPS
func (r *TcpMonitorRepository) RemoveTCPMonitor(vpsName string, ip net.IP) error {
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors/%s", vpsName, ip.String())}

	return r.Client.Delete(restRequest)
}

// GetContacts returns a list of all your monitoring contacts
func (r *TcpMonitorRepository) GetContacts() ([]MonitoringContact, error) {
	var response contactsWrapper
	restRequest := rest.RestRequest{Endpoint: "/monitoring-contacts"}
	err := r.Client.Get(restRequest, &response)

	return response.Contacts, err
}

// CreateContact allows you to add a new contact which could be used by the tcp monitoring
func (r *TcpMonitorRepository) CreateContact(contact MonitoringContact) error {
	restRequest := rest.RestRequest{Endpoint: "/monitoring-contacts", Body: &contact}

	return r.Client.Post(restRequest)
}

// UpdateContact updates the specified contact
func (r *TcpMonitorRepository) UpdateContact(contact MonitoringContact) error {
	requestBody := contactWrapper{Contact: contact}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/monitoring-contacts/%d", contact.Id), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveContact allows you to delete a specific contact by id
func (r *TcpMonitorRepository) RemoveContact(contactId int64) error {
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/monitoring-contacts/%d", contactId)}

	return r.Client.Delete(restRequest)
}
