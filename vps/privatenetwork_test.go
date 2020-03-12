package vps

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"testing"
)

func TestRepository_GetPrivateNetworks(t *testing.T) {
	const apiResponse = `{ "privateNetworks": [ { "name": "example-privatenetwork", "description": "FilesharingNetwork", "isBlocked": false, "isLocked": false, "vpsNames": [ "example-vps", "example-vps2" ] } ] } `
	server := mockServer{t: t, expectedUrl: "/private-networks", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetPrivateNetworks()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t,"example-privatenetwork", all[0].Name)
	assert.Equal(t,"FilesharingNetwork", all[0].Description)
	assert.Equal(t,false, all[0].IsBlocked)
	assert.Equal(t,false, all[0].IsLocked)

	assert.Equal(t, []string{"example-vps", "example-vps2"}, all[0].VpsNames)

}

func TestRepository_GetPrivateNetworkByName(t *testing.T) {
	const apiResponse = `{ "privateNetwork": { "name": "example-privatenetwork", "description": "FilesharingNetwork", "isBlocked": false, "isLocked": false, "vpsNames": [ "example-vps", "example-vps2" ] } } `
	server := mockServer{t: t, expectedUrl: "/private-networks/example-privatenetwork", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	privateNetwork, err := repo.GetPrivateNetworkByName("example-privatenetwork")
	require.NoError(t, err)
	assert.Equal(t,"example-privatenetwork", privateNetwork.Name)
	assert.Equal(t,"FilesharingNetwork", privateNetwork.Description)
	assert.Equal(t,false, privateNetwork.IsBlocked)
	assert.Equal(t,false, privateNetwork.IsLocked)

	assert.Equal(t, []string{"example-vps", "example-vps2"}, privateNetwork.VpsNames)
}

func TestRepository_OrderPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"description":"test123"}`
	server := mockServer{t: t, expectedUrl: "/private-networks", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.OrderPrivateNetwork("test123")
	require.NoError(t, err)
}

func TestRepository_UpdatePrivateNetwork(t *testing.T) {
	const expectedRequest = `{"privateNetwork":{"name":"example-privatenetwork","description":"einnetwork","isBlocked":false,"isLocked":false,"vpsNames":["example-vps","example-vps2"]}}`
	server := mockServer{t: t, expectedUrl: "/private-networks/example-privatenetwork", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	privateNetwork := PrivateNetwork{
		Name:        "example-privatenetwork",
		Description: "einnetwork",
		IsBlocked:   false,
		IsLocked:    false,
		VpsNames:    []string{"example-vps", "example-vps2"},
	}
	
	err := repo.UpdatePrivateNetwork(privateNetwork)
	require.NoError(t, err)
}

func TestRepository_AttachVpsToPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"action":"attachvps","vpsName":"example-vps"}`
	server := mockServer{t: t, expectedUrl: "/private-networks/example-privatenetwork", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.AttachVpsToPrivateNetwork("example-vps", "example-privatenetwork")
	require.NoError(t, err)
}

func TestRepository_DetachVpsFromPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"action":"detachvps","vpsName":"example-vps"}`
	server := mockServer{t: t, expectedUrl: "/private-networks/example-privatenetwork", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DetachVpsFromPrivateNetwork("example-vps", "example-privatenetwork")
	require.NoError(t, err)
}

func TestRepository_CancelPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"endTime":"end"}`
	server := mockServer{t: t, expectedUrl: "/private-networks/example-privatenetwork", expectedMethod: "DELETE", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.CancelPrivateNetwork("example-privatenetwork", gotransip.CancellationTimeEnd)
	require.NoError(t, err)
}
