package vps

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestTcpMonitorRepository_GetTCPMonitors(t *testing.T) {
	const apiResponse = `{ "tcpMonitors": [ { "ipAddress": "10.3.37.1", "label": "HTTP", "ports": [ 80, 443 ], "interval": 6, "allowedTimeouts": 1, "contacts": [ { "id": 1, "enableEmail": true, "enableSMS": false } ], "ignoreTimes": [ { "timeFrom": "18:00", "timeTo": "08:30" } ] } ] }`

	server := mockServer{t: t, expectedURL: "/vps/example-vps/tcp-monitors", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	all, err := repo.GetTCPMonitors("example-vps")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "10.3.37.1", all[0].IPAddress.String())
	assert.Equal(t, "HTTP", all[0].Label)
	assert.Equal(t, []int{80, 443}, all[0].Ports)
	assert.Equal(t, 6, all[0].Interval)
	assert.Equal(t, 1, all[0].AllowedTimeouts)

	require.Equal(t, 1, len(all[0].Contacts))
	assert.EqualValues(t, 1, all[0].Contacts[0].ID)
	assert.Equal(t, true, all[0].Contacts[0].EnableEmail)
	assert.Equal(t, false, all[0].Contacts[0].EnableSMS)

	require.Equal(t, 1, len(all[0].IgnoreTimes))
	assert.Equal(t, "18:00", all[0].IgnoreTimes[0].TimeFrom)
	assert.Equal(t, "08:30", all[0].IgnoreTimes[0].TimeTo)

}

func TestTcpMonitorRepository_CreateTCPMonitor(t *testing.T) {
	const expectedRequest = `{"tcpMonitor":{"ipAddress":"10.3.37.1","label":"HTTP","ports":[80,443],"interval":6,"allowedTimeouts":1,"contacts":[{"id":1,"enableEmail":true,"enableSMS":false}],"ignoreTimes":[{"timeFrom":"18:00","timeTo":"08:30"}]}}`
	server := mockServer{t: t, expectedURL: "/vps/example-vps/tcp-monitors", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	tcpMonitor := TCPMonitor{
		IPAddress:       net.ParseIP("10.3.37.1"),
		Label:           "HTTP",
		Ports:           []int{80, 443},
		Interval:        6,
		AllowedTimeouts: 1,
		Contacts: []TCPMonitorContact{{
			ID:          1,
			EnableEmail: true,
			EnableSMS:   false,
		}},
		IgnoreTimes: []IgnoreTime{{
			TimeFrom: "18:00",
			TimeTo:   "08:30",
		}},
	}

	err := repo.CreateTCPMonitor("example-vps", tcpMonitor)
	require.NoError(t, err)
}

func TestTcpMonitorRepository_UpdateTCPMonitor(t *testing.T) {
	const expectedRequest = `{"tcpMonitor":{"ipAddress":"10.3.37.1","label":"HTTP","ports":[80,443],"interval":6,"allowedTimeouts":1,"contacts":[{"id":1,"enableEmail":true,"enableSMS":false}],"ignoreTimes":[{"timeFrom":"18:00","timeTo":"08:30"}]}}`
	server := mockServer{t: t, expectedURL: "/vps/example-vps/tcp-monitors/10.3.37.1", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	tcpMonitor := TCPMonitor{
		IPAddress:       net.ParseIP("10.3.37.1"),
		Label:           "HTTP",
		Ports:           []int{80, 443},
		Interval:        6,
		AllowedTimeouts: 1,
		Contacts: []TCPMonitorContact{{
			ID:          1,
			EnableEmail: true,
			EnableSMS:   false,
		}},
		IgnoreTimes: []IgnoreTime{{
			TimeFrom: "18:00",
			TimeTo:   "08:30",
		}},
	}

	err := repo.UpdateTCPMonitor("example-vps", tcpMonitor)
	require.NoError(t, err)
}

func TestTcpMonitorRepository_RemoveTCPMonitor(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedURL: "/vps/example-vps/tcp-monitors/10.3.37.1", expectedMethod: "DELETE", statusCode: 204, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	ip := net.ParseIP("10.3.37.1")
	err := repo.RemoveTCPMonitor("example-vps", ip)
	require.NoError(t, err)
}

func TestTcpMonitorRepository_GetContacts(t *testing.T) {
	const apiResponse = `{ "contacts": [ { "id": 1, "name": "John Wick", "telephone": "+31612345678", "email": "j.wick@example.com" } ] }`
	server := mockServer{t: t, expectedURL: "/monitoring-contacts", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	all, err := repo.GetContacts()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 1, all[0].ID)
	assert.Equal(t, "John Wick", all[0].Name)
	assert.Equal(t, "+31612345678", all[0].Telephone)
	assert.Equal(t, "j.wick@example.com", all[0].Email)
}

func TestTcpMonitorRepository_CreateContact(t *testing.T) {
	const expectedRequest = `{"name":"John Wick","telephone":"+31612345678","email":"j.wick@example.com"}`
	server := mockServer{t: t, expectedURL: "/monitoring-contacts", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	contact := MonitoringContact{
		Name:      "John Wick",
		Telephone: "+31612345678",
		Email:     "j.wick@example.com",
	}
	err := repo.CreateContact(contact)
	require.NoError(t, err)
}

func TestTcpMonitorRepository_UpdateContact(t *testing.T) {
	const expectedRequest = `{"contact":{"id":1,"name":"John Wick","telephone":"+31612345678","email":"j.wick@example.com"}}`
	server := mockServer{t: t, expectedURL: "/monitoring-contacts/1", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	contact := MonitoringContact{
		ID:        1,
		Name:      "John Wick",
		Telephone: "+31612345678",
		Email:     "j.wick@example.com",
	}

	err := repo.UpdateContact(contact)
	require.NoError(t, err)
}

func TestTcpMonitorRepository_DeleteContact(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/monitoring-contacts/1", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := TCPMonitorRepository{Client: *client}

	err := repo.RemoveContact(1)
	require.NoError(t, err)
}
