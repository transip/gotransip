package vps

import (
	"fmt"
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
	"time"
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

	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "GET", statusCode: 200, response: apiResponse}
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
	server := mockServer{t: t, expectedUrl: "/vps/example-vps", expectedMethod: "GET", statusCode: 200, response: apiResponse}
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
	server := mockServer{t: t, expectedUrl: "/vps", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
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
	const apiResponse = `{"usage":{"cpu":[{"percentage":3.11,"date":1574783109}]}} `
	const expectedRequest = `{"types":"cpu","dateTimeStart":1500538995,"dateTimeEnd":1500542619}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/usage", expectedMethod: "GET", statusCode: 200, expectedRequest: expectedRequest, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	usageData, err := repo.GetUsageDataByVps("example-vps", []UsageType{UsageTypeCpu}, UsagePeriod{TimeStart: 1500538995, TimeEnd: 1500542619})
	require.NoError(t, err)

	require.Equal(t, 1, len(usageData.Cpu))
	assert.EqualValues(t, 3.11, usageData.Cpu[0].Percentage)
	assert.EqualValues(t, 1574783109, usageData.Cpu[0].Date)
}

func TestRepository_GetAllUsageDataByVps(t *testing.T) {
	const apiResponse = `{ "usage": { "cpu": [ { "percentage": 3.11, "date": 1574783109 } ], "disk": [ { "iopsRead": 0.27, "iopsWrite": 0.13, "date": 1574783109 } ], "network": [ { "mbitOut": 100.2, "mbitIn": 249.93, "date": 1574783109 } ] } } `
	const expectedRequest = `{"types":"cpu,disk,network","dateTimeStart":1500538995,"dateTimeEnd":1500542619}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/usage", expectedMethod: "GET", statusCode: 200, expectedRequest: expectedRequest, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	usageData, err := repo.GetAllUsageDataByVps("example-vps", UsagePeriod{TimeStart: 1500538995, TimeEnd: 1500542619})
	require.NoError(t, err)

	require.Equal(t, 1, len(usageData.Cpu))
	require.Equal(t, 1, len(usageData.Disk))
	require.Equal(t, 1, len(usageData.Network))

	assert.EqualValues(t, 3.11, usageData.Cpu[0].Percentage)
	assert.EqualValues(t, 1574783109, usageData.Cpu[0].Date)

	assert.EqualValues(t, 0.27, usageData.Disk[0].IopsRead)
	assert.EqualValues(t, 0.13, usageData.Disk[0].IopsWrite)
	assert.EqualValues(t, 1574783109, usageData.Disk[0].Date)

	assert.EqualValues(t, 100.2, usageData.Network[0].MbitOut)
	assert.EqualValues(t, 249.93, usageData.Network[0].MbitIn)
	assert.EqualValues(t, 1574783109, usageData.Network[0].Date)
}

func TestRepository_GetAllUsageDataByVps24Hours(t *testing.T) {
	const apiResponse = `{ "usage": { "cpu": [ { "percentage": 3.11, "date": 1574783109 } ], "disk": [ { "iopsRead": 0.27, "iopsWrite": 0.13, "date": 1574783109 } ], "network": [ { "mbitOut": 100.2, "mbitIn": 249.93, "date": 1574783109 } ] } } `
	expectedRequest := fmt.Sprintf(`{"types":"cpu,disk,network","dateTimeStart":%d,"dateTimeEnd":%d}`, time.Now().Unix()-24*3600, time.Now().Unix())
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/usage", expectedMethod: "GET", statusCode: 200, expectedRequest: expectedRequest, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	usageData, err := repo.GetAllUsageDataByVps24Hours("example-vps")
	require.NoError(t, err)

	require.Equal(t, 1, len(usageData.Cpu))
	require.Equal(t, 1, len(usageData.Disk))
	require.Equal(t, 1, len(usageData.Network))

	assert.EqualValues(t, 3.11, usageData.Cpu[0].Percentage)
	assert.EqualValues(t, 1574783109, usageData.Cpu[0].Date)

	assert.EqualValues(t, 0.27, usageData.Disk[0].IopsRead)
	assert.EqualValues(t, 0.13, usageData.Disk[0].IopsWrite)
	assert.EqualValues(t, 1574783109, usageData.Disk[0].Date)

	assert.EqualValues(t, 100.2, usageData.Network[0].MbitOut)
	assert.EqualValues(t, 249.93, usageData.Network[0].MbitIn)
	assert.EqualValues(t, 1574783109, usageData.Network[0].Date)
}

func TestRepository_GetVNCData(t *testing.T) {
	const apiResponse = `{ "vncData": { "host": "vncproxy.transip.nl", "path": "websockify?token=testtokentje", "url": "https://vncproxy.transip.nl/websockify?token=testtokentje", "token": "testtokentje", "password": "esisteinpassw0rd" } }`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/vnc-data", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	vncData, err := repo.GetVNCData("example-vps")
	require.NoError(t, err)
	assert.Equal(t, "vncproxy.transip.nl", vncData.Host)
	assert.Equal(t, "websockify?token=testtokentje", vncData.Path)
	assert.Equal(t, "https://vncproxy.transip.nl/websockify?token=testtokentje", vncData.Url)
	assert.Equal(t, "esisteinpassw0rd", vncData.Password)
	assert.Equal(t, "testtokentje", vncData.Token)

}

func TestRepository_RegenerateVNCToken(t *testing.T) {
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/vnc-data", expectedMethod: "PATCH", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RegenerateVNCToken("example-vps")
	require.NoError(t, err)
}

func TestRepository_GetAddons(t *testing.T) {
	const apiResponse = `{ "addons": { "active": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ], "cancellable": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ], "available": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ] } }`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/addons", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	allAddons, err := repo.GetAddons("example-vps")
	require.NoError(t, err)
	require.Equal(t, 1, len(allAddons.Active))
	require.Equal(t, 1, len(allAddons.Cancellable))
	require.Equal(t, 1, len(allAddons.Available))

	assert.Equal(t, "example-product-name", allAddons.Active[0].Name)
	assert.Equal(t, "This is an example product", allAddons.Active[0].Description)
	assert.Equal(t, 499, allAddons.Active[0].Price)
	assert.Equal(t, 799, allAddons.Active[0].RecurringPrice)

	assert.Equal(t, "example-product-name", allAddons.Cancellable[0].Name)
	assert.Equal(t, "This is an example product", allAddons.Cancellable[0].Description)
	assert.Equal(t, 499, allAddons.Cancellable[0].Price)
	assert.Equal(t, 799, allAddons.Cancellable[0].RecurringPrice)

	assert.Equal(t, "example-product-name", allAddons.Available[0].Name)
	assert.Equal(t, "This is an example product", allAddons.Available[0].Description)
	assert.Equal(t, 499, allAddons.Available[0].Price)
	assert.Equal(t, 799, allAddons.Available[0].RecurringPrice)
}

func TestRepository_OrderAddons(t *testing.T) {
	const expectedRequest = `{"addons":["vps-addon-1-extra-ip-address"]}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/addons", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.OrderAddons("example-vps", []string{"vps-addon-1-extra-ip-address"})
	require.NoError(t, err)
}

func TestRepository_CancelAddon(t *testing.T) {
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/addons/einaddon", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.CancelAddon("example-vps", "einaddon")
	require.NoError(t, err)
}

func TestRepository_GetUpgrades(t *testing.T) {
	const apiResponse = `{ "upgrades": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ] } `
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/upgrades", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	allUpgrades, err := repo.GetUpgrades("example-vps")
	require.NoError(t, err)
	require.Equal(t, 1, len(allUpgrades))

	assert.Equal(t, "example-product-name", allUpgrades[0].Name)
	assert.Equal(t, "This is an example product", allUpgrades[0].Description)
	assert.Equal(t, 499, allUpgrades[0].Price)
	assert.Equal(t, 799, allUpgrades[0].RecurringPrice)
}

func TestRepository_Upgrade(t *testing.T) {
	const expectedRequest = `{"productName":"vps-bladevps-pro-x16"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/upgrades", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Upgrade("example-vps", "vps-bladevps-pro-x16")
	require.NoError(t, err)
}

func TestRepository_GetOperatingSystems(t *testing.T) {
	const apiResponse = `{ "operatingSystems": [ { "name": "ubuntu-18.04", "description": "Ubuntu 18.04 LTS", "isPreinstallableImage": false, "version": "18.04 LTS", "price": 1250 } ] }`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/operating-systems", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	oses, err := repo.GetOperatingSystems("example-vps")
	require.NoError(t, err)
	require.Equal(t, 1, len(oses))

	assert.Equal(t, "ubuntu-18.04", oses[0].Name)
	assert.Equal(t, "Ubuntu 18.04 LTS", oses[0].Description)
	assert.Equal(t, false, oses[0].IsPreinstallableImage)
	assert.Equal(t, "18.04 LTS", oses[0].Version)
	assert.Equal(t, 1250, oses[0].Price)

}

func TestRepository_InstallOperatingSystemOptionalFields(t *testing.T) {
	const expectedRequest = `{"operatingSystemName":"ubuntu-18.04"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/operating-systems", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.InstallOperatingSystem("example-vps", "ubuntu-18.04", "", "")
	require.NoError(t, err)
}

func TestRepository_InstallOperatingSystem(t *testing.T) {
	const expectedRequest = `{"operatingSystemName":"ubuntu-18.04","hostname":"test","base64InstallText":"ZGFzaXN0YmFzZTY0"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/operating-systems", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.InstallOperatingSystem("example-vps", "ubuntu-18.04", "test", "ZGFzaXN0YmFzZTY0")
	require.NoError(t, err)
}

func TestRepository_GetIPAddresses(t *testing.T) {
	const apiResponse = `{ "ipAddresses" : [ { "dnsResolvers" : [ "195.8.195.8", "195.135.195.135" ], "subnetMask" : "255.255.255.0", "reverseDns" : "example.com", "address" : "149.210.192.184", "gateway" : "149.210.192.1" }, { "address" : "2a01:7c8:aab5:5d5::1", "gateway" : "2a01:7c8:aab5::1", "dnsResolvers" : [ "2a01:7c8:7000:195::8:195:8", "2a01:7c8:7000:195::135:195:135" ], "subnetMask" : "/48", "reverseDns" : "example.com" } ] }`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/ip-addresses", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	ips, err := repo.GetIPAddresses("example-vps")
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
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/ip-addresses/37.97.254.6", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("37.97.254.6")
	ip, err := repo.GetIPAddressByAddress("example-vps", address)
	require.NoError(t, err)

	assert.EqualValues(t, "37.97.254.6", ip.Address.String())
	assert.EqualValues(t, "00000000000000000000ffffffffff00", ip.SubnetMask.String())
	assert.EqualValues(t, "37.97.254.1", ip.Gateway.String())
	assert.EqualValues(t, "195.8.195.8", ip.DnsResolvers[0].String())
	assert.EqualValues(t, "195.135.195.135", ip.DnsResolvers[1].String())
	assert.EqualValues(t, "example.com", ip.ReverseDns)
}

func TestRepository_AddIPv6Address(t *testing.T) {
	const expectedRequest = `{"ipAddress":"2a01:7c8:3:1337::6"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/ip-addresses", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01:7c8:3:1337::6")
	err := repo.AddIPv6Address("example-vps", address)
	require.NoError(t, err)
}

func TestRepository_UpdateReverseDNS(t *testing.T) {
	const expectedRequest = `{"ipAddress":{"address":"37.97.254.6","gateway":"37.97.254.1","reverseDns":"example.com","subnetMask":"255.0.0.0"}}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/ip-addresses/37.97.254.6", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
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
	err := repo.UpdateReverseDNS("example-vps", address)
	require.NoError(t, err)
}

func TestRepository_RemoveIPv6Address(t *testing.T) {
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/ip-addresses/2a01::1", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	address := net.ParseIP("2a01::1")
	err := repo.RemoveIPv6Address("example-vps", address)
	require.NoError(t, err)
}

func TestRepository_GetSnapshots(t *testing.T) {
	const apiResponse = `{ "snapshots": [ { "name": "1572607577", "description": "before upgrade", "diskSize": 314572800, "status": "creating", "dateTimeCreate": "2019-07-14 12:21:11", "operatingSystem": "ubuntu-18.04" } ] } `
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/snapshots", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetSnapshots("example-vps")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "1572607577", all[0].Name)
	assert.Equal(t, "before upgrade", all[0].Description)
	assert.EqualValues(t, 314572800, all[0].DiskSize)
	assert.Equal(t, "creating", all[0].Status)
	assert.Equal(t, "2019-07-14 12:21:11", all[0].DateTimeCreate)
	assert.Equal(t, "ubuntu-18.04", all[0].OperatingSystem)

}

func TestRepository_GetSnapshotByName(t *testing.T) {
	const apiResponse = `{ "snapshot": { "name": "1572607577", "description": "before upgrade", "diskSize": 314572800, "status": "creating", "dateTimeCreate": "2019-07-14 12:21:11", "operatingSystem": "ubuntu-18.04" } }`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/snapshots/1572607577", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	snap, err := repo.GetSnapshotByName("example-vps", "1572607577")
	require.NoError(t, err)

	assert.Equal(t, "1572607577", snap.Name)
	assert.Equal(t, "before upgrade", snap.Description)
	assert.EqualValues(t, 314572800, snap.DiskSize)
	assert.Equal(t, "creating", snap.Status)
	assert.Equal(t, "2019-07-14 12:21:11", snap.DateTimeCreate)
	assert.Equal(t, "ubuntu-18.04", snap.OperatingSystem)
}

func TestRepository_CreateSnapshot(t *testing.T) {
	const expectedRequest = `{"description":"BeforeItsAllBroken","shouldStartVps":true}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/snapshots", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.CreateSnapshot("example-vps", "BeforeItsAllBroken", true)
	require.NoError(t, err)
}

func TestRepository_RevertSnapshot(t *testing.T) {
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/snapshots/1337", expectedMethod: "PATCH", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RevertSnapshot("example-vps", "1337")
	require.NoError(t, err)
}

func TestRepository_RevertSnapshotToOtherVps(t *testing.T) {
	const expectedRequest = `{"destinationVpsName":"example-vps2"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/snapshots/1337", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RevertSnapshotToOtherVps("example-vps", "1337", "example-vps2")
	require.NoError(t, err)
}

func TestRepository_RemoveSnapshot(t *testing.T) {
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/snapshots/1337", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RemoveSnapshot("example-vps", "1337")
	require.NoError(t, err)
}

func TestRepository_GetBackups(t *testing.T) {
	const apiResponse = `{ "backups": [ { "id": 712332, "status": "active", "dateTimeCreate": "2019-11-29 22:11:20", "diskSize": 157286400, "operatingSystem": "Ubuntu 19.10", "availabilityZone": "ams0" } ] }`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/backups", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetBackups("example-vps")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 712332, all[0].Id)
	assert.EqualValues(t, "active", all[0].Status)
	assert.Equal(t, "2019-11-29 22:11:20", all[0].DateTimeCreate.Format("2006-01-02 15:04:05"))
	assert.EqualValues(t, 157286400, all[0].DiskSize)
	assert.Equal(t, "Ubuntu 19.10", all[0].OperatingSystem)
	assert.Equal(t, "ams0", all[0].AvailabilityZone)
}

func TestRepository_RevertBackup(t *testing.T) {
	const expectedRequest = `{"action":"revert"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/backups/1337", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RevertBackup("example-vps", 1337)
	require.NoError(t, err)
}

func TestRepository_ConvertBackupToSnapshot(t *testing.T) {
	const expectedRequest = `{"action":"convert","description":"BeforeItsAllBroken"}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/backups/1337", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.ConvertBackupToSnapshot("example-vps", 1337, "BeforeItsAllBroken")
	require.NoError(t, err)
}
