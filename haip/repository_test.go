package haip

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
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
	expectedURL     string
	expectedMethod  string
	statusCode      int
	expectedRequest string
	response        string
	skipRequestBody bool
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedURL, req.URL.String()) // check if right expectedURL is called

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
			_, err := rw.Write([]byte(m.response))
			require.NoError(m.t, err, "error when writing mock response")
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
	const apiResponse = `{ "haips": [ { "name": "example-haip", "description": "frontend cluster", "status": "active", "isLoadBalancingEnabled": true, "loadBalancingMode": "cookie", "stickyCookieName": "PHPSESSID", "healthCheckInterval": 3000, "httpHealthCheckPath": "/status.php", "httpHealthCheckPort": 443, "httpHealthCheckSsl": true, "ipv4Address": "37.97.254.7", "ipv6Address": "2a01:7c8:3:1337::1", "ipSetup": "ipv6to4", "ptrRecord": "frontend.example.com", "ipAddresses": [ "10.3.37.1", "10.3.38.1" ] } ] } `
	server := mockServer{t: t, expectedURL: "/haips", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-haip", all[0].Name)
	assert.Equal(t, "frontend cluster", all[0].Description)
	assert.EqualValues(t, "active", all[0].Status)
	assert.Equal(t, true, all[0].IsLoadBalancingEnabled)
	assert.EqualValues(t, "cookie", all[0].LoadBalancingMode)
	assert.Equal(t, "PHPSESSID", all[0].StickyCookieName)
	assert.EqualValues(t, 3000, all[0].HealthCheckInterval)
	assert.Equal(t, "/status.php", all[0].HTTPHealthCheckPath)
	assert.Equal(t, 443, all[0].HTTPHealthCheckPort)
	assert.Equal(t, true, all[0].HTTPHealthCheckSsl)
	assert.Equal(t, "37.97.254.7", all[0].IPv4Address.String())
	assert.Equal(t, "2a01:7c8:3:1337::1", all[0].IPv6Address.String())
	assert.EqualValues(t, "ipv6to4", all[0].IPSetup)
	assert.Equal(t, "frontend.example.com", all[0].PtrRecord)

	require.Equal(t, 2, len(all[0].IPAddresses))
	assert.Equal(t, "10.3.37.1", all[0].IPAddresses[0].String())
	assert.Equal(t, "10.3.38.1", all[0].IPAddresses[1].String())
}

func TestRepository_GetSelection(t *testing.T) {
	const apiResponse = `{ "haips": [ { "name": "example-haip", "description": "frontend cluster", "status": "active", "isLoadBalancingEnabled": true, "loadBalancingMode": "cookie", "stickyCookieName": "PHPSESSID", "healthCheckInterval": 3000, "httpHealthCheckPath": "/status.php", "httpHealthCheckPort": 443, "httpHealthCheckSsl": true, "ipv4Address": "37.97.254.7", "ipv6Address": "2a01:7c8:3:1337::1", "ipSetup": "ipv6to4", "ptrRecord": "frontend.example.com", "ipAddresses": [ "10.3.37.1", "10.3.38.1" ] } ] } `
	server := mockServer{t: t, expectedURL: "/haips?page=1&pageSize=25", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetSelection(1, 25)
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-haip", all[0].Name)
	assert.Equal(t, "frontend cluster", all[0].Description)
	assert.EqualValues(t, "active", all[0].Status)
	assert.Equal(t, true, all[0].IsLoadBalancingEnabled)
	assert.EqualValues(t, "cookie", all[0].LoadBalancingMode)
	assert.Equal(t, "PHPSESSID", all[0].StickyCookieName)
	assert.EqualValues(t, 3000, all[0].HealthCheckInterval)
	assert.Equal(t, "/status.php", all[0].HTTPHealthCheckPath)
	assert.Equal(t, 443, all[0].HTTPHealthCheckPort)
	assert.Equal(t, true, all[0].HTTPHealthCheckSsl)
	assert.Equal(t, "37.97.254.7", all[0].IPv4Address.String())
	assert.Equal(t, "2a01:7c8:3:1337::1", all[0].IPv6Address.String())
	assert.EqualValues(t, "ipv6to4", all[0].IPSetup)
	assert.Equal(t, "frontend.example.com", all[0].PtrRecord)

	require.Equal(t, 2, len(all[0].IPAddresses))
	assert.Equal(t, "10.3.37.1", all[0].IPAddresses[0].String())
	assert.Equal(t, "10.3.38.1", all[0].IPAddresses[1].String())
}

func TestRepository_GetByName(t *testing.T) {
	const apiResponse = `{ "haip": { "name": "example-haip", "description": "frontend cluster", "status": "active", "isLoadBalancingEnabled": true, "loadBalancingMode": "cookie", "stickyCookieName": "PHPSESSID", "healthCheckInterval": 3000, "httpHealthCheckPath": "/status.php", "httpHealthCheckPort": 443, "httpHealthCheckSsl": true, "ipv4Address": "37.97.254.7", "ipv6Address": "2a01:7c8:3:1337::1", "ipSetup": "ipv6to4", "ptrRecord": "frontend.example.com", "ipAddresses": [ "10.3.37.1", "10.3.38.1" ] } }`
	server := mockServer{t: t, expectedURL: "/haips/example-haip", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	haip, err := repo.GetByName("example-haip")
	require.NoError(t, err)

	assert.Equal(t, "example-haip", haip.Name)
	assert.Equal(t, "frontend cluster", haip.Description)
	assert.EqualValues(t, "active", haip.Status)
	assert.Equal(t, true, haip.IsLoadBalancingEnabled)
	assert.EqualValues(t, "cookie", haip.LoadBalancingMode)
	assert.Equal(t, "PHPSESSID", haip.StickyCookieName)
	assert.EqualValues(t, 3000, haip.HealthCheckInterval)
	assert.Equal(t, "/status.php", haip.HTTPHealthCheckPath)
	assert.Equal(t, 443, haip.HTTPHealthCheckPort)
	assert.Equal(t, true, haip.HTTPHealthCheckSsl)
	assert.Equal(t, "37.97.254.7", haip.IPv4Address.String())
	assert.Equal(t, "2a01:7c8:3:1337::1", haip.IPv6Address.String())
	assert.EqualValues(t, "ipv6to4", haip.IPSetup)
	assert.Equal(t, "frontend.example.com", haip.PtrRecord)

	require.Equal(t, 2, len(haip.IPAddresses))
	assert.Equal(t, "10.3.37.1", haip.IPAddresses[0].String())
	assert.Equal(t, "10.3.38.1", haip.IPAddresses[1].String())
}

func TestRepository_Order(t *testing.T) {
	const expectedRequestBody = `{"productName":"haip-pro-contract","description":"myhaip01"}`
	server := mockServer{t: t, expectedURL: "/haips", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Order("haip-pro-contract", "myhaip01")
	require.NoError(t, err)
}

func TestRepository_Update(t *testing.T) {
	const expectedRequestBody = `{"haip":{"name":"example-haip","description":"frontend cluster","status":"active","isLoadBalancingEnabled":true,"loadBalancingMode":"cookie","stickyCookieName":"PHPSESSID","healthCheckInterval":3000,"httpHealthCheckPath":"/status.php","httpHealthCheckPort":443,"httpHealthCheckSsl":true,"ipv4Address":"37.97.254.7","ipv6Address":"2a01:7c8:3:1337::1","ipSetup":"ipv6to4","ptrRecord":"frontend.example.com","ipAddresses":["10.3.37.1","10.3.38.1"]}}`
	server := mockServer{t: t, expectedURL: "/haips/example-haip", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	haip := Haip{
		Name:                   "example-haip",
		Description:            "frontend cluster",
		Status:                 "active",
		IsLoadBalancingEnabled: true,
		LoadBalancingMode:      "cookie",
		StickyCookieName:       "PHPSESSID",
		HealthCheckInterval:    3000,
		HTTPHealthCheckPath:    "/status.php",
		HTTPHealthCheckPort:    443,
		HTTPHealthCheckSsl:     true,
		IPv4Address:            net.ParseIP("37.97.254.7"),
		IPv6Address:            net.ParseIP("2a01:7c8:3:1337::1"),
		IPSetup:                "ipv6to4",
		PtrRecord:              "frontend.example.com",
		IPAddresses:            []net.IP{net.ParseIP("10.3.37.1"), net.ParseIP("10.3.38.1")},
	}

	err := repo.Update(haip)
	require.NoError(t, err)
}

func TestRepository_Cancel(t *testing.T) {
	const expectedRequestBody = `{"endTime":"immediately"}`
	server := mockServer{t: t, expectedURL: "/haips/example-haip", expectedMethod: "DELETE", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Cancel("example-haip", gotransip.CancellationTimeImmediately)
	require.NoError(t, err)
}

func TestRepository_GetAllCertificates(t *testing.T) {
	const apiResponse = `{ "certificates": [ { "id": 25478, "commonName": "example.com", "expirationDate": "2019-11-23" } ] }`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/certificates", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAllCertificates("example-haip")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 25478, all[0].ID)
	assert.Equal(t, "example.com", all[0].CommonName)
	assert.Equal(t, "2019-11-23", all[0].ExpirationDate)
}

func TestRepository_AddCertificate(t *testing.T) {
	const expectedRequestBody = `{"sslCertificateId":1337}`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/certificates", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.AddCertificate("example-haip", 1337)
	require.NoError(t, err)
}

func TestRepository_AddLetsEncryptCertificate(t *testing.T) {
	const expectedRequestBody = `{"commonName":"foobar.example.com"}`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/certificates", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.AddLetsEncryptCertificate("example-haip", "foobar.example.com")
	require.NoError(t, err)
}

func TestRepository_DetachCertificate(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/haips/example-haip/certificates/1337", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DetachCertificate("example-haip", 1337)
	require.NoError(t, err)
}

func TestRepository_GetAttachedIPAddresses(t *testing.T) {
	const apiResponse = `{ "ipAddresses": [ "149.13.3.7", "149.31.33.7" ] }`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/ip-addresses", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAttachedIPAddresses("example-haip")
	require.NoError(t, err)
	require.Equal(t, 2, len(all))

	assert.Equal(t, "149.13.3.7", all[0].String())
	assert.Equal(t, "149.31.33.7", all[1].String())
}

func TestRepository_SetAttachedIPAddresses(t *testing.T) {
	const expectedRequestBody = `{"ipAddresses":["10.3.37.1","10.3.37.2"]}`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/ip-addresses", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	ips := []net.IP{net.ParseIP("10.3.37.1"), net.ParseIP("10.3.37.2")}

	err := repo.SetAttachedIPAddresses("example-haip", ips)
	require.NoError(t, err)
}

func TestRepository_DetachIPAddresses(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/haips/example-haip/ip-addresses", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DetachIPAddresses("example-haip")
	require.NoError(t, err)
}

func TestRepository_GetPortConfigurations(t *testing.T) {
	const apiResponse = `{ "portConfigurations": [ { "id": 9865, "name": "Website Traffic", "sourcePort": 80, "targetPort": 80, "mode": "http", "endpointSslMode": "off" } ] } `
	server := mockServer{t: t, expectedURL: "/haips/example-haip/port-configurations", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetPortConfigurations("example-haip")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 9865, all[0].ID)
	assert.Equal(t, "Website Traffic", all[0].Name)
	assert.Equal(t, 80, all[0].SourcePort)
	assert.Equal(t, 80, all[0].TargetPort)
	assert.EqualValues(t, "http", all[0].Mode)
	assert.Equal(t, "off", all[0].EndpointSslMode)
}

func TestRepository_GetPortConfiguration(t *testing.T) {
	const apiResponse = `{ "portConfiguration": { "id": 9865, "name": "Website Traffic", "sourcePort": 80, "targetPort": 80, "mode": "http", "endpointSslMode": "off" } } `
	server := mockServer{t: t, expectedURL: "/haips/example-haip/port-configurations/9865", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	configuration, err := repo.GetPortConfiguration("example-haip", 9865)
	require.NoError(t, err)

	assert.EqualValues(t, 9865, configuration.ID)
	assert.Equal(t, "Website Traffic", configuration.Name)
	assert.Equal(t, 80, configuration.SourcePort)
	assert.Equal(t, 80, configuration.TargetPort)
	assert.EqualValues(t, "http", configuration.Mode)
	assert.Equal(t, "off", configuration.EndpointSslMode)
}

func TestRepository_AddPortConfiguration(t *testing.T) {
	const expectedRequestBody = `{"name":"Website Traffic","sourcePort":443,"targetPort":443,"mode":"https","endpointSslMode":"on"}`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/port-configurations", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	configuration := PortConfiguration{
		Name:            "Website Traffic",
		SourcePort:      443,
		TargetPort:      443,
		Mode:            "https",
		EndpointSslMode: "on",
	}
	err := repo.AddPortConfiguration("example-haip", configuration)
	require.NoError(t, err)
}

func TestRepository_UpdatePortConfiguration(t *testing.T) {
	const expectedRequestBody = `{"portConfiguration":{"id":9865,"name":"Website Traffic","sourcePort":443,"targetPort":443,"mode":"https","endpointSslMode":"on"}}`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/port-configurations/9865", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	configuration := PortConfiguration{
		ID:              9865,
		Name:            "Website Traffic",
		SourcePort:      443,
		TargetPort:      443,
		Mode:            "https",
		EndpointSslMode: "on",
	}
	err := repo.UpdatePortConfiguration("example-haip", configuration)
	require.NoError(t, err)
}

func TestRepository_RemovePortConfiguration(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/haips/example-haip/port-configurations/1337", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RemovePortConfiguration("example-haip", 1337)
	require.NoError(t, err)
}

func TestRepository_GetStatusReport(t *testing.T) {
	const apiResponse = `{ "statusReport": [ { "ipAddress": "136.10.14.1", "port": 80, "ipVersion": 4, "loadBalancerName": "lb0", "loadBalancerIp": "136.144.151.255", "state": "up", "lastChange": "2019-09-29 16:51:18" } ] }`
	server := mockServer{t: t, expectedURL: "/haips/example-haip/status-reports", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetStatusReport("example-haip")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "136.10.14.1", all[0].IPAddress.String())
	assert.Equal(t, 80, all[0].Port)
	assert.Equal(t, 4, all[0].IPVersion)
	assert.Equal(t, "lb0", all[0].LoadBalancerName)
	assert.Equal(t, "136.144.151.255", all[0].LoadBalancerIP.String())
	assert.Equal(t, "up", all[0].State)
	assert.Equal(t, "2019-09-29 16:51:18", all[0].LastChange.Format("2006-01-02 15:04:05"))
}
