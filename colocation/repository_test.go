package colocation

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
	"github.com/transip/gotransip/v6/ipaddress"
)

func TestRepository_GetAll(t *testing.T) {
	const apiResponse = `{ "colocations": [ { "name": "example2", "ipRanges": [ "2a01:7c8:c038:6::/64" ] } ] } `
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example2", all[0].Name)
	require.Equal(t, 1, len(all[0].IPRanges))

	assert.Equal(t, "2a01:7c8:c038:6::/64", all[0].IPRanges[0].String())
}

func TestRepository_GetByName(t *testing.T) {
	const apiResponse = `{ "colocation": { "name": "example2", "ipRanges": [ "2a01:7c8:c038:6::/64" ] } } `
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	colo, err := repo.GetByName("example2")
	require.NoError(t, err)

	assert.Equal(t, "example2", colo.Name)
	require.Equal(t, 1, len(colo.IPRanges))

	assert.Equal(t, "2a01:7c8:c038:6::/64", colo.IPRanges[0].String())
}

func TestRepository_CreateRemoteHandsRequest(t *testing.T) {
	const expectedRequest = `{"remoteHands":{"coloName":"example2","contactName":"Herman Kaakdorst","phoneNumber":"+31 612345678","expectedDuration":15,"instructions":"Reboot server with label Loadbalancer0"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2/remote-hands", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	remoteHands := RemoteHandsRequest{
		ColoName:         "example2",
		ContactName:      "Herman Kaakdorst",
		PhoneNumber:      "+31 612345678",
		ExpectedDuration: 15,
		Instructions:     "Reboot server with label Loadbalancer0",
	}

	err := repo.CreateRemoteHandsRequest(remoteHands)
	require.NoError(t, err)
}

func TestRepository_GetIPAddresses(t *testing.T) {
	const apiResponse = `{ "ipAddresses" : [ { "dnsResolvers" : [ "195.8.195.8", "195.135.195.135" ], "subnetMask" : "255.255.255.0", "reverseDns" : "example.com", "address" : "149.210.192.184", "gateway" : "149.210.192.1" }, { "address" : "2a01:7c8:aab5:5d5::1", "gateway" : "2a01:7c8:aab5::1", "dnsResolvers" : [ "2a01:7c8:7000:195::8:195:8", "2a01:7c8:7000:195::135:195:135" ], "subnetMask" : "/48", "reverseDns" : "example.com" } ] }`
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2/ip-addresses", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	ips, err := repo.GetIPAddresses("example2")
	require.NoError(t, err)
	require.Equal(t, 2, len(ips))

	assert.EqualValues(t, "149.210.192.184", ips[0].Address.String())
	assert.EqualValues(t, "00000000000000000000ffffffffff00", ips[0].SubnetMask.String())
	assert.EqualValues(t, "149.210.192.1", ips[0].Gateway.String())
	assert.EqualValues(t, "195.8.195.8", ips[0].DNSResolvers[0].String())
	assert.EqualValues(t, "195.135.195.135", ips[0].DNSResolvers[1].String())
	assert.EqualValues(t, "example.com", ips[0].ReverseDNS)

	assert.EqualValues(t, "2a01:7c8:aab5:5d5::1", ips[1].Address.String())
	assert.EqualValues(t, "ffffffffffff00000000000000000000", ips[1].SubnetMask.String())
	assert.EqualValues(t, "2a01:7c8:aab5::1", ips[1].Gateway.String())
	assert.EqualValues(t, "2a01:7c8:7000:195:0:8:195:8", ips[1].DNSResolvers[0].String())
	assert.EqualValues(t, "2a01:7c8:7000:195:0:135:195:135", ips[1].DNSResolvers[1].String())
	assert.EqualValues(t, "example.com", ips[1].ReverseDNS)
}

func TestRepository_GetIPAddressByAddress(t *testing.T) {
	const apiResponse = `{ "ipAddress": { "address": "37.97.254.6", "subnetMask": "255.255.255.0", "gateway": "37.97.254.1", "dnsResolvers": [ "195.8.195.8", "195.135.195.135" ], "reverseDns": "example.com" } } `
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2/ip-addresses/37.97.254.6", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("37.97.254.6")
	ip, err := repo.GetIPAddressByAddress("example2", address)
	require.NoError(t, err)

	assert.EqualValues(t, "37.97.254.6", ip.Address.String())
	assert.EqualValues(t, "00000000000000000000ffffffffff00", ip.SubnetMask.String())
	assert.EqualValues(t, "37.97.254.1", ip.Gateway.String())
	assert.EqualValues(t, "195.8.195.8", ip.DNSResolvers[0].String())
	assert.EqualValues(t, "195.135.195.135", ip.DNSResolvers[1].String())
	assert.EqualValues(t, "example.com", ip.ReverseDNS)
}

func TestRepository_AddIPAddress(t *testing.T) {
	const expectedRequest = `{"ipAddress":"2a01:7c8:3:1337::6","reverseDns":"example.com"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2/ip-addresses", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01:7c8:3:1337::6")
	err := repo.AddIPAddress("example2", address, "example.com")
	require.NoError(t, err)
}

func TestRepository_AddIPAddressWithoutReverseDns(t *testing.T) {
	const expectedRequest = `{"ipAddress":"2a01:7c8:3:1337::6"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2/ip-addresses", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01:7c8:3:1337::6")
	err := repo.AddIPAddress("example2", address, "")
	require.NoError(t, err)
}

func TestRepository_UpdateReverseDNS(t *testing.T) {
	const expectedRequest = `{"ipAddress":{"address":"37.97.254.6","gateway":"37.97.254.1","reverseDns":"example.com","subnetMask":"255.0.0.0"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2/ip-addresses/37.97.254.6", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	ip := net.ParseIP("37.97.254.6")
	address := ipaddress.IPAddress{
		Address:    ip,
		Gateway:    net.ParseIP("37.97.254.1"),
		ReverseDNS: "example.com",
		SubnetMask: ipaddress.SubnetMask{IPMask: ip.DefaultMask()},
	}
	err := repo.UpdateReverseDNS("example2", address)
	require.NoError(t, err)
}

func TestRepository_RemoveIPAddress(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/colocations/example2/ip-addresses/2a01::1", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01::1")
	err := repo.RemoveIPAddress("example2", address)
	require.NoError(t, err)
}
