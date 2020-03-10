package vps

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"io/ioutil"
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
	config := gotransip.ClientConfiguration{DemoMode: true, URL: httpServer.URL}
	client, err := gotransip.NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return &client, tearDown
}

func TestRepository_GetAll(t *testing.T) {
	const apiResponse = `{ "vpss": [ { "name": "example-vps", "description": "example VPS", "productName": "vps-bladevps-x1", "operatingSystem": "ubuntu-18.04", "diskSize": 157286400, "memorySize": 4194304, "cpus": 2, "status": "running", "ipAddress": "37.97.254.6", "macAddress": "52:54:00:3b:52:65", "currentSnapshots": 1, "maxSnapshots": 10, "isLocked": false, "isBlocked": false, "isCustomerLocked": false, "availabilityZone": "ams0", "tags": [ "customTag", "anotherTag" ] } ] }`

	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-vps", all[0].Name)
	assert.Equal(t, "example VPS", all[0].Description)
	assert.Equal(t, "vps-bladevps-x1", all[0].ProductName)
	assert.Equal(t, "ubuntu-18.04", all[0].OperatingSystem)
	assert.EqualValues(t, 157286400, all[0].DiskSize)
	assert.EqualValues(t, 4194304, all[0].MemorySize)
	assert.EqualValues(t, 2, all[0].Cpus)
	assert.Equal(t, "running", all[0].Status)
	assert.Equal(t, "37.97.254.6", all[0].IpAddress)
	assert.Equal(t, "52:54:00:3b:52:65", all[0].MacAddress)
	assert.EqualValues(t, 1, all[0].CurrentSnapshots)
	assert.EqualValues(t, 10, all[0].MaxSnapshots)
	assert.Equal(t, false, all[0].IsLocked)
	assert.Equal(t, false, all[0].IsBlocked)
	assert.Equal(t, false, all[0].IsCustomerLocked)
	assert.Equal(t, "ams0", all[0].AvailabilityZone)
	assert.Equal(t, []string{"customTag", "anotherTag"}, all[0].Tags)
}

func TestRepository_GetByName(t *testing.T) {
	const apiResponse = `{ "vps": { "name": "example-vps", "description": "example VPS", "productName": "vps-bladevps-x1", "operatingSystem": "ubuntu-18.04", "diskSize": 157286400, "memorySize": 4194304, "cpus": 2, "status": "running", "ipAddress": "37.97.254.6", "macAddress": "52:54:00:3b:52:65", "currentSnapshots": 1, "maxSnapshots": 10, "isLocked": false, "isBlocked": false, "isCustomerLocked": false, "availabilityZone": "ams0", "tags": [ "customTag", "anotherTag" ] } }`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	vps, err := repo.GetByName("example-vps")
	require.NoError(t, err)

	assert.Equal(t, "example-vps", vps.Name)
	assert.Equal(t, "example VPS", vps.Description)
	assert.Equal(t, "vps-bladevps-x1", vps.ProductName)
	assert.Equal(t, "ubuntu-18.04", vps.OperatingSystem)
	assert.EqualValues(t, 157286400, vps.DiskSize)
	assert.EqualValues(t, 4194304, vps.MemorySize)
	assert.EqualValues(t, 2, vps.Cpus)
	assert.Equal(t, "running", vps.Status)
	assert.Equal(t, "37.97.254.6", vps.IpAddress)
	assert.Equal(t, "52:54:00:3b:52:65", vps.MacAddress)
	assert.EqualValues(t, 1, vps.CurrentSnapshots)
	assert.EqualValues(t, 10, vps.MaxSnapshots)
	assert.Equal(t, false, vps.IsLocked)
	assert.Equal(t, false, vps.IsBlocked)
	assert.Equal(t, false, vps.IsCustomerLocked)
	assert.Equal(t, "ams0", vps.AvailabilityZone)
	assert.Equal(t, []string{"customTag", "anotherTag"}, vps.Tags)
}

func TestRepository_Order(t *testing.T) {
	const expectedRequestBody = `{"productName":"vps-bladevps-x8","operatingSystem":"ubuntu-18.04","availabilityZone":"ams0","addons":["vpsAddon-1-extra-cpu-core"],"hostname":"server01.example.com","description":"example vps description","base64InstallText":"testtext123"}`
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	order := VpsOrder{
		ProductName:       "vps-bladevps-x8",
		OperatingSystem:   "ubuntu-18.04",
		AvailabilityZone:  "ams0",
		Addons:            []string{"vpsAddon-1-extra-cpu-core"},
		Hostname:          "server01.example.com",
		Description:       "example vps description",
		Base64InstallText: "testtext123",
	}

	err := repo.Order(order)
	require.NoError(t, err)
}

func TestRepository_OrderMultiple(t *testing.T) {
	const expectedRequestBody = `{"vpss":[{"productName":"vps-bladevps-x8","operatingSystem":"ubuntu-18.04","availabilityZone":"ams0","addons":["vpsAddon-1-extra-cpu-core"],"hostname":"server01.example.com","description":"webserver01","base64InstallText":"testtext123"},{"productName":"vps-bladevps-x8","operatingSystem":"ubuntu-18.04","availabilityZone":"ams0","hostname":"server01.example.com","description":"backupserver01"}]}`
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	orders := []VpsOrder{{
		ProductName:       "vps-bladevps-x8",
		OperatingSystem:   "ubuntu-18.04",
		AvailabilityZone:  "ams0",
		Addons:            []string{"vpsAddon-1-extra-cpu-core"},
		Hostname:          "server01.example.com",
		Description:       "webserver01",
		Base64InstallText: "testtext123",
	}, {
		ProductName:      "vps-bladevps-x8",
		OperatingSystem:  "ubuntu-18.04",
		AvailabilityZone: "ams0",
		Hostname:         "server01.example.com",
		Description:      "backupserver01",
	}}

	err := repo.OrderMultiple(orders)
	require.NoError(t, err)
}

func TestRepository_Clone(t *testing.T) {
	const expectedRequest = `{"vpsName":"example-vps"}`
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Clone("example-vps")
	require.NoError(t, err)
}

func TestRepository_CloneToAvailabilityZone(t *testing.T) {
	const expectedRequest = `{"vpsName":"example-vps","availabilityZone":"ams0"}`
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.CloneToAvailabilityZone("example-vps", "ams0")
	require.NoError(t, err)
}

func TestRepository_Update(t *testing.T) {
	const expectedRequest = `{"vps":{"name":"example-vps","description":"example VPS","productName":"vps-bladevps-x1","operatingSystem":"ubuntu-18.04","diskSize":157286400,"memorySize":4194304,"cpus":2,"status":"running","ipAddress":"37.97.254.6","macAddress":"52:54:00:3b:52:65","currentSnapshots":1,"maxSnapshots":10,"availabilityZone":"ams0","tags":["customTag","anotherTag"]}}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	vpsToUpdate := Vps{
		Name:             "example-vps",
		Description:      "example VPS",
		ProductName:      "vps-bladevps-x1",
		OperatingSystem:  "ubuntu-18.04",
		DiskSize:         157286400,
		MemorySize:       4194304,
		Cpus:             2,
		Status:           "running",
		IpAddress:        "37.97.254.6",
		MacAddress:       "52:54:00:3b:52:65",
		CurrentSnapshots: 1,
		MaxSnapshots:     10,
		IsLocked:         false,
		IsBlocked:        false,
		IsCustomerLocked: false,
		AvailabilityZone: "ams0",
		Tags:             []string{"customTag", "anotherTag"},
	}

	err := repo.Update(vpsToUpdate)

	require.NoError(t, err)
}

func TestRepository_Start(t *testing.T) {
	const expectedRequest = `{"action":"start"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Start("example-vps")
	require.NoError(t, err)
}

func TestRepository_Stop(t *testing.T) {
	const expectedRequest = `{"action":"stop"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Stop("example-vps")
	require.NoError(t, err)
}

func TestRepository_Reset(t *testing.T) {
	const expectedRequest = `{"action":"reset"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Reset("example-vps")
	require.NoError(t, err)
}

func TestRepository_Handover(t *testing.T) {
	const expectedRequest = `{"action":"handover","targetCustomerName":"bobexample"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Handover("example-vps", "bobexample")
	require.NoError(t, err)
}

func TestRepository_Cancel(t *testing.T) {
	const expectedRequest = `{"endTime":"end"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "DELETE", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Cancel("example-vps", gotransip.CancellationTimeEnd)
	require.NoError(t, err)
}

func TestRepository_GetUsageData(t *testing.T) {
	const apiResponse = `{ "usage": { "cpu": [ { "percentage": 3.11, "date": 1574783109 } ], "disk": [ { "iopsRead": 0.27, "iopsWrite": 0.13, "date": 1574783109 } ], "network": [ { "mbitOut": 100.2, "mbitIn": 249.93, "date": 1574783109 } ] } } `
	const expectedRequest = `{"types":"cpu,disk,network","dateTimeStart":1500538995,"dateTimeEnd":1500542619}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/usage", expectedMethod: "GET", statusCode: 200, expectedRequest: expectedRequest, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	types := []VpsUsageType{VpsUsageTypeCpu, VpsUsageTypeDisk, VpsUsageTypeNetwork}
	_, err := repo.GetUsage("example-vps", types, UsagePeriod{TimeStart: 1500538995, TimeEnd: 1500542619})
	require.NoError(t, err)
}

func TestRepository_GetVNCData(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_RegenerateVNCToken(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetAddons(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_OrderAddons(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_CancelAddon(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetUpgrades(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_Upgrade(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetOperatingSystems(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_InstallOperatingSystem(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetIPAddresses(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetIPAddressByAddress(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_AddIPv6Address(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_UpdateReverseDNS(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_RemoveIPv6Address(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetSnapshots(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetSnapshotByName(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_CreateSnapshot(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_RevertSnapshot(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_DeleteSnapshot(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetBackups(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_RevertBackup(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_ConvertBackupToSnapshot(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetFirewall(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_UpdateFirewall(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetPrivateNetworks(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetPrivateNetworkByName(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_OrderPrivateNetwork(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_UpdatePrivateNetwork(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_AttachVpsToPrivateNetwork(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_DetachVpsFromPrivateNetwork(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_CancelPrivateNetwork(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetBigStorages(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetBigStorageByName(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_OrderBigStorage(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_UpgradeBigStorage(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_UpdateBigStorage(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_DetachVpsFromBigStorage(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_AttachVpsToBigStorage(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_CancelBigStorage(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetBigStorageBackups(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_RevertBigStorageBackup(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetBigStorageUsage(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetBigStorageUsageLast24Hours(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetTCPMonitors(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_CreateTCPMonitor(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_UpdateTCPMonitor(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_DeleteTCPMonitor(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_GetContacts(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_CreateContact(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_UpdateContact(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}

func TestRepository_DeleteContact(t *testing.T) {
	const apiResponse = ""
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse,}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 0, len(all))
}
