package vps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestPrivateNetworkRepository_GetPrivateNetworks(t *testing.T) {
	const apiResponse = `{ "privateNetworks": [ { "name": "example-privatenetwork", "description": "FilesharingNetwork", "isBlocked": false, "isLocked": false, "vpsNames": [ "example-vps", "example-vps2" ] } ] } `
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-privatenetwork", all[0].Name)
	assert.Equal(t, "FilesharingNetwork", all[0].Description)
	assert.Equal(t, false, all[0].IsBlocked)
	assert.Equal(t, false, all[0].IsLocked)

	assert.Equal(t, []string{"example-vps", "example-vps2"}, all[0].VpsNames)
}

func TestPrivateNetworkRepository_GetSelection(t *testing.T) {
	const apiResponse = `{ "privateNetworks": [ { "name": "example-privatenetwork", "description": "FilesharingNetwork", "isBlocked": false, "isLocked": false, "vpsNames": [ "example-vps", "example-vps2" ] } ] } `
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks?page=1&pageSize=25", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	all, err := repo.GetSelection(1, 25)
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-privatenetwork", all[0].Name)
	assert.Equal(t, "FilesharingNetwork", all[0].Description)
	assert.Equal(t, false, all[0].IsBlocked)
	assert.Equal(t, false, all[0].IsLocked)

	assert.Equal(t, []string{"example-vps", "example-vps2"}, all[0].VpsNames)
}

func TestPrivateNetworkRepository_GetPrivateNetworkByName(t *testing.T) {
	const apiResponse = `{ "privateNetwork": { "name": "example-privatenetwork", "description": "FilesharingNetwork", "isBlocked": false, "isLocked": false, "vpsNames": [ "example-vps", "example-vps2" ] } } `
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks/example-privatenetwork", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	privateNetwork, err := repo.GetByName("example-privatenetwork")
	require.NoError(t, err)
	assert.Equal(t, "example-privatenetwork", privateNetwork.Name)
	assert.Equal(t, "FilesharingNetwork", privateNetwork.Description)
	assert.Equal(t, false, privateNetwork.IsBlocked)
	assert.Equal(t, false, privateNetwork.IsLocked)

	assert.Equal(t, []string{"example-vps", "example-vps2"}, privateNetwork.VpsNames)
}

func TestPrivateNetworkRepository_OrderPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"description":"test123"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	err := repo.Order("test123")
	require.NoError(t, err)
}

func TestPrivateNetworkRepository_UpdatePrivateNetwork(t *testing.T) {
	const expectedRequest = `{"privateNetwork":{"name":"example-privatenetwork","description":"einnetwork","isBlocked":false,"isLocked":false,"vpsNames":["example-vps","example-vps2"]}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks/example-privatenetwork", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	privateNetwork := PrivateNetwork{
		Name:        "example-privatenetwork",
		Description: "einnetwork",
		IsBlocked:   false,
		IsLocked:    false,
		VpsNames:    []string{"example-vps", "example-vps2"},
	}

	err := repo.Update(privateNetwork)
	require.NoError(t, err)
}

func TestPrivateNetworkRepository_AttachVpsToPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"action":"attachvps","vpsName":"example-vps"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks/example-privatenetwork", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	err := repo.AttachVps("example-vps", "example-privatenetwork")
	require.NoError(t, err)
}

func TestPrivateNetworkRepository_DetachVpsFromPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"action":"detachvps","vpsName":"example-vps"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks/example-privatenetwork", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	err := repo.DetachVps("example-vps", "example-privatenetwork")
	require.NoError(t, err)
}

func TestPrivateNetworkRepository_CancelPrivateNetwork(t *testing.T) {
	const expectedRequest = `{"endTime":"end"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/private-networks/example-privatenetwork", ExpectedMethod: "DELETE", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := PrivateNetworkRepository{Client: *client}

	err := repo.Cancel("example-privatenetwork", gotransip.CancellationTimeEnd)
	require.NoError(t, err)
}
