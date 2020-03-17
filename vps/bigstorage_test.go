package vps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"testing"
	"time"
)

func TestBigStorageRepository_GetBigStorages(t *testing.T) {
	const apiResponse = `{ "bigStorages": [ { "name": "example-bigstorage", "description": "Big storage description", "diskSize": 2147483648, "offsiteBackups": true, "vpsName": "example-vps", "status": "active", "isLocked": false, "availabilityZone": "ams0" } ] } `
	server := mockServer{t: t, expectedUrl: "/big-storages", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	all, err := repo.GetBigStorages()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-bigstorage", all[0].Name)
	assert.Equal(t, "Big storage description", all[0].Description)
	assert.EqualValues(t, 2147483648, all[0].DiskSize)
	assert.Equal(t, true, all[0].OffsiteBackups)
	assert.Equal(t, "example-vps", all[0].VpsName)
	assert.EqualValues(t, "active", all[0].Status)
	assert.Equal(t, false, all[0].IsLocked)
	assert.Equal(t, "ams0", all[0].AvailabilityZone)
}

func TestBigStorageRepository_GetBigStorageByName(t *testing.T) {
	const apiResponse = `{ "bigStorage": { "name": "example-bigstorage", "description": "Big storage description", "diskSize": 2147483648, "offsiteBackups": true, "vpsName": "example-vps", "status": "active", "isLocked": false, "availabilityZone": "ams0" } } `
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	bigstorage, err := repo.GetBigStorageByName("example-bigstorage")
	require.NoError(t, err)
	assert.Equal(t, "example-bigstorage", bigstorage.Name)
	assert.Equal(t, "Big storage description", bigstorage.Description)
	assert.EqualValues(t, 2147483648, bigstorage.DiskSize)
	assert.Equal(t, true, bigstorage.OffsiteBackups)
	assert.Equal(t, "example-vps", bigstorage.VpsName)
	assert.EqualValues(t, "active", bigstorage.Status)
	assert.Equal(t, false, bigstorage.IsLocked)
	assert.Equal(t, "ams0", bigstorage.AvailabilityZone)
}

func TestBigStorageRepository_OrderBigStorage(t *testing.T) {
	const expectedRequest = `{"size":8,"offsiteBackups":true,"availabilityZone":"ams0","vpsName":"example-vps"}`
	server := mockServer{t: t, expectedUrl: "/big-storages", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	order := BigStorageOrder{Size: 8, OffsiteBackups: true, AvailabilityZone: "ams0", VpsName: "example-vps"}
	err := repo.OrderBigStorage(order)

	require.NoError(t, err)
}

func TestBigStorageRepository_UpgradeBigStorage(t *testing.T) {
	const expectedRequest = `{"bigStorageName":"example-bigstorage","size":8,"offsiteBackups":true}`
	server := mockServer{t: t, expectedUrl: "/big-storages", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	err := repo.UpgradeBigStorage("example-bigstorage", 8, true)

	require.NoError(t, err)
}

func TestBigStorageRepository_UpdateBigStorage(t *testing.T) {
	const expectedRequest = `{"bigStorage":{"name":"example-bigstorage","description":"Big storage description","diskSize":2147483648,"offsiteBackups":true,"vpsName":"example-vps","status":"active","isLocked":false,"availabilityZone":"ams0"}}`
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	bigStorage := BigStorage{
		Name:             "example-bigstorage",
		Description:      "Big storage description",
		DiskSize:         2147483648,
		OffsiteBackups:   true,
		VpsName:          "example-vps",
		Status:           BigStorageStatusActive,
		IsLocked:         false,
		AvailabilityZone: "ams0",
	}
	err := repo.UpdateBigStorage(bigStorage)

	require.NoError(t, err)
}

func TestBigStorageRepository_DetachVpsFromBigStorage(t *testing.T) {
	const expectedRequest = `{"bigStorage":{"name":"example-bigstorage","description":"Big storage description","diskSize":2147483648,"offsiteBackups":true,"vpsName":"example-vps","status":"active","isLocked":false,"availabilityZone":"ams0"}}`
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	bigStorage := BigStorage{
		Name:             "example-bigstorage",
		Description:      "Big storage description",
		DiskSize:         2147483648,
		OffsiteBackups:   true,
		Status:           "active",
		IsLocked:         false,
		AvailabilityZone: "ams0",
	}
	err := repo.AttachVpsToBigStorage("example-vps", bigStorage)
	require.NoError(t, err)
}

func TestBigStorageRepository_AttachVpsToBigStorage(t *testing.T) {
	const expectedRequest = `{"bigStorage":{"name":"example-bigstorage","description":"Big storage description","diskSize":2147483648,"offsiteBackups":true,"vpsName":"","status":"active","isLocked":false,"availabilityZone":"ams0"}}`
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	bigStorage := BigStorage{
		Name:             "example-bigstorage",
		Description:      "Big storage description",
		DiskSize:         2147483648,
		OffsiteBackups:   true,
		VpsName:          "example-vps",
		Status:           "active",
		IsLocked:         false,
		AvailabilityZone: "ams0",
	}
	err := repo.DetachVpsFromBigStorage(bigStorage)
	require.NoError(t, err)
}

func TestBigStorageRepository_CancelBigStorage(t *testing.T) {
	const expectedRequest = `{"endTime":"end"}`
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage", expectedMethod: "DELETE", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	err := repo.CancelBigStorage("example-bigstorage", gotransip.CancellationTimeEnd)
	require.NoError(t, err)
}

func TestBigStorageRepository_GetBigStorageBackups(t *testing.T) {
	const apiResponse = `{ "backups": [ { "id": 1583, "status": "active", "diskSize": 4294967296, "dateTimeCreate": "2019-12-31 09:13:55", "availabilityZone": "ams0" } ] }`
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage/backups", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	all, err := repo.GetBigStorageBackups("example-bigstorage")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 1583, all[0].Id)
	assert.EqualValues(t, "active", all[0].Status)
	assert.EqualValues(t, 4294967296, all[0].DiskSize)
	assert.Equal(t, "2019-12-31 09:13:55", all[0].DateTimeCreate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "ams0", all[0].AvailabilityZone)
}

func TestBigStorageRepository_RevertBigStorageBackup(t *testing.T) {
	const expectedRequest = `{"action":"revert"}`
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage/backups/123", expectedMethod: "PATCH", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	err := repo.RevertBigStorageBackup("example-bigstorage", 123)
	require.NoError(t, err)
}

func TestBigStorageRepository_GetBigStorageUsage(t *testing.T) {
	const apiResponse = `{ "usage": [ { "iopsRead": 0.27, "iopsWrite": 0.13, "date": 1574783109 } ] }`
	const expectedRequest = `{"dateTimeStart":1500538995,"dateTimeEnd":1500542619}`
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage/usage", expectedMethod: "GET", statusCode: 200, expectedRequest: expectedRequest, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	usageData, err := repo.GetBigStorageUsage("example-bigstorage", UsagePeriod{TimeStart: 1500538995, TimeEnd: 1500542619})
	require.NoError(t, err)

	require.Equal(t, 1, len(usageData))
	assert.EqualValues(t, 0.27, usageData[0].IopsRead)
	assert.EqualValues(t, 0.13, usageData[0].IopsWrite)
	assert.EqualValues(t, 1574783109, usageData[0].Date)
}

func TestBigStorageRepository_GetBigStorageUsageLast24Hours(t *testing.T) {
	const apiResponse = `{ "usage": [ { "iopsRead": 0.27, "iopsWrite": 0.13, "date": 1574783109 } ] }`
	expectedRequest := fmt.Sprintf(`{"dateTimeStart":%d,"dateTimeEnd":%d}`, time.Now().Unix()-24*3600, time.Now().Unix())
	server := mockServer{t: t, expectedUrl: "/big-storages/example-bigstorage/usage", expectedMethod: "GET", statusCode: 200, response: apiResponse, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := BigStorageRepository{Client: *client}

	usageData, err := repo.GetBigStorageUsageLast24Hours("example-bigstorage")
	require.NoError(t, err)

	require.Equal(t, 1, len(usageData))
	assert.EqualValues(t, 0.27, usageData[0].IopsRead)
	assert.EqualValues(t, 0.13, usageData[0].IopsWrite)
	assert.EqualValues(t, 1574783109, usageData[0].Date)
}
