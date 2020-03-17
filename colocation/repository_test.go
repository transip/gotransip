package colocation

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/ipaddress"
	"github.com/transip/gotransip/v6/repository"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t               *testing.T
	expectedUrl     string
	expectedMethod  string
	statusCode      int
	expectedRequest string
	response        string
	skipRequestBody bool
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedUrl, req.URL.String()) // check if right expectedUrl is called

		if m.skipRequestBody == false && req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := ioutil.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequest, string(body))
		}

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequest expectedMethod is used
		rw.WriteHeader(m.statusCode)                    // respond with given status code

		if m.response != "" {
			rw.Write([]byte(m.response))
		}
	}))
}

func (m *mockServer) getClient() (*repository.Client, func()) {
	httpServer := m.getHTTPServer()
	config := gotransip.DemoClientConfiguration
	config.URL = httpServer.URL

	client, err := gotransip.NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return &client, tearDown
}

func TestRepository_GetAll(t *testing.T) {
	const apiResponse = `{ "colocations": [ { "name": "example2", "ipRanges": [ "2a01:7c8:c038:6::/64" ] } ] } `
	server := mockServer{t: t, expectedUrl: "/colocations", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example2", all[0].Name)
	require.Equal(t, 1, len(all[0].IpRanges))

	assert.Equal(t, "2a01:7c8:c038:6::/64", all[0].IpRanges[0].String())
}

func TestRepository_GetByName(t *testing.T) {
	const apiResponse = `{ "colocation": { "name": "example2", "ipRanges": [ "2a01:7c8:c038:6::/64" ] } } `
	server := mockServer{t: t, expectedUrl: "/colocations/example2", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	colo, err := repo.GetByName("example2")
	require.NoError(t, err)

	assert.Equal(t, "example2", colo.Name)
	require.Equal(t, 1, len(colo.IpRanges))

	assert.Equal(t, "2a01:7c8:c038:6::/64", colo.IpRanges[0].String())
}

func TestRepository_CreateRemoteHandsRequest(t *testing.T) {
	const expectedRequest = `{"remoteHands":{"coloName":"example2","contactName":"Herman Kaakdorst","phoneNumber":"+31 612345678","expectedDuration":15,"instructions":"Reboot server with label Loadbalancer0"}}`
	server := mockServer{t: t, expectedUrl: "/colocations/example2/remote-hands", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
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
	server := mockServer{t: t, expectedUrl: "/colocations/example2/ip-addresses", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	ips, err := repo.GetIPAddresses("example2")
	require.NoError(t, err)
	require.Equal(t, 2, len(ips))

	assert.EqualValues(t, "149.210.192.184", ips[0].Address.String())
	assert.EqualValues(t, "00000000000000000000ffffffffff00", ips[0].SubnetMask.String())
	assert.EqualValues(t, "149.210.192.1", ips[0].Gateway.String())
	assert.EqualValues(t, "195.8.195.8", ips[0].DnsResolvers[0].String())
	assert.EqualValues(t, "195.135.195.135", ips[0].DnsResolvers[1].String())
	assert.EqualValues(t, "example.com", ips[0].ReverseDns)

	assert.EqualValues(t, "2a01:7c8:aab5:5d5::1", ips[1].Address.String())
	assert.EqualValues(t, "ffffffffffff00000000000000000000", ips[1].SubnetMask.String())
	assert.EqualValues(t, "2a01:7c8:aab5::1", ips[1].Gateway.String())
	assert.EqualValues(t, "2a01:7c8:7000:195:0:8:195:8", ips[1].DnsResolvers[0].String())
	assert.EqualValues(t, "2a01:7c8:7000:195:0:135:195:135", ips[1].DnsResolvers[1].String())
	assert.EqualValues(t, "example.com", ips[1].ReverseDns)
}

func TestRepository_GetIPAddressByAddress(t *testing.T) {
	const apiResponse = `{ "ipAddress": { "address": "37.97.254.6", "subnetMask": "255.255.255.0", "gateway": "37.97.254.1", "dnsResolvers": [ "195.8.195.8", "195.135.195.135" ], "reverseDns": "example.com" } } `
	server := mockServer{t: t, expectedUrl: "/colocations/example2/ip-addresses/37.97.254.6", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("37.97.254.6")
	ip, err := repo.GetIPAddressByAddress("example2", address)
	require.NoError(t, err)

	assert.EqualValues(t, "37.97.254.6", ip.Address.String())
	assert.EqualValues(t, "00000000000000000000ffffffffff00", ip.SubnetMask.String())
	assert.EqualValues(t, "37.97.254.1", ip.Gateway.String())
	assert.EqualValues(t, "195.8.195.8", ip.DnsResolvers[0].String())
	assert.EqualValues(t, "195.135.195.135", ip.DnsResolvers[1].String())
	assert.EqualValues(t, "example.com", ip.ReverseDns)
}

func TestRepository_AddIPAddress(t *testing.T) {
	const expectedRequest = `{"ipAddress":"2a01:7c8:3:1337::6","reverseDns":"example.com"}`
	server := mockServer{t: t, expectedUrl: "/colocations/example2/ip-addresses", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01:7c8:3:1337::6")
	err := repo.AddIPAddress("example2", address, "example.com")
	require.NoError(t, err)
}

func TestRepository_AddIPAddressWithoutReverseDns(t *testing.T) {
	const expectedRequest = `{"ipAddress":"2a01:7c8:3:1337::6"}`
	server := mockServer{t: t, expectedUrl: "/colocations/example2/ip-addresses", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01:7c8:3:1337::6")
	err := repo.AddIPAddress("example2", address, "")
	require.NoError(t, err)
}

func TestRepository_UpdateReverseDNS(t *testing.T) {
	const expectedRequest = `{"ipAddress":{"address":"37.97.254.6","gateway":"37.97.254.1","reverseDns":"example.com","subnetMask":"255.0.0.0"}}`
	server := mockServer{t: t, expectedUrl: "/colocations/example2/ip-addresses/37.97.254.6", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	ip := net.ParseIP("37.97.254.6")
	address := ipaddress.IPAddress{
		Address:    ip,
		Gateway:    net.ParseIP("37.97.254.1"),
		ReverseDns: "example.com",
		SubnetMask: ipaddress.SubnetMask{IPMask: ip.DefaultMask()},
	}
	err := repo.UpdateReverseDNS("example2", address)
	require.NoError(t, err)
}

func TestRepository_RemoveIPAddress(t *testing.T) {
	server := mockServer{t: t, expectedUrl: "/colocations/example2/ip-addresses/2a01::1", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01::1")
	err := repo.RemoveIPAddress("example2", address)
	require.NoError(t, err)
}
