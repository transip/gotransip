package vps

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestBlockStorageRepository_GetBlockStorages(t *testing.T) {
	const apiResponse = `{ "blockStorages": [ { "name": "example-blockstorage", "description": "Block storage description", "diskSize": 2147483648, "offsiteBackups": true, "vpsName": "example-vps", "status": "active", "isLocked": false, "availabilityZone": "ams0", "serial": "e7e12b3c7c6602973ac7" } ] } `
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-blockstorage", all[0].Name)
	assert.Equal(t, "Block storage description", all[0].Description)
	assert.EqualValues(t, 2147483648, all[0].DiskSize)
	assert.Equal(t, true, all[0].OffsiteBackups)
	assert.Equal(t, "example-vps", all[0].VpsName)
	assert.EqualValues(t, "active", all[0].Status)
	assert.Equal(t, false, all[0].IsLocked)
	assert.Equal(t, "ams0", all[0].AvailabilityZone)
	assert.Equal(t, "e7e12b3c7c6602973ac7", all[0].Serial)
}

func TestBlockStorageRepository_GetSelection(t *testing.T) {
	const apiResponse = `{ "blockStorages": [ { "name": "example-blockstorage", "description": "Block storage description", "diskSize": 2147483648, "offsiteBackups": true, "vpsName": "example-vps", "status": "active", "isLocked": false, "availabilityZone": "ams0", "serial": "e7e12b3c7c6602973ac7"} ] } `
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages?page=1&pageSize=25", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	all, err := repo.GetSelection(1, 25)
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "example-blockstorage", all[0].Name)
	assert.Equal(t, "Block storage description", all[0].Description)
	assert.EqualValues(t, 2147483648, all[0].DiskSize)
	assert.Equal(t, true, all[0].OffsiteBackups)
	assert.Equal(t, "example-vps", all[0].VpsName)
	assert.EqualValues(t, "active", all[0].Status)
	assert.Equal(t, false, all[0].IsLocked)
	assert.Equal(t, "ams0", all[0].AvailabilityZone)
	assert.Equal(t, "e7e12b3c7c6602973ac7", all[0].Serial)
}

func TestBlockStorageRepository_GetBlockStorageByName(t *testing.T) {
	const apiResponse = `{ "blockStorage": { "name": "example-blockstorage", "description": "Block storage description", "diskSize": 2147483648, "offsiteBackups": true, "vpsName": "example-vps", "status": "active", "isLocked": false, "availabilityZone": "ams0", "serial": "e7e12b3c7c6602973ac7" } } `
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	blockstorage, err := repo.GetByName("example-blockstorage")
	require.NoError(t, err)
	assert.Equal(t, "example-blockstorage", blockstorage.Name)
	assert.Equal(t, "Block storage description", blockstorage.Description)
	assert.EqualValues(t, 2147483648, blockstorage.DiskSize)
	assert.Equal(t, true, blockstorage.OffsiteBackups)
	assert.Equal(t, "example-vps", blockstorage.VpsName)
	assert.EqualValues(t, "active", blockstorage.Status)
	assert.Equal(t, false, blockstorage.IsLocked)
	assert.Equal(t, "ams0", blockstorage.AvailabilityZone)
	assert.Equal(t, "e7e12b3c7c6602973ac7", blockstorage.Serial)
}

func TestBlockStorageRepository_OrderBlockStorage(t *testing.T) {
	const expectedRequest = `{"type":"fast-storage","size":8,"offsiteBackups":true,"availabilityZone":"ams0","vpsName":"example-vps","description":"test-description"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	order := BlockStorageOrder{Type: "fast-storage", Size: 8, OffsiteBackups: true, AvailabilityZone: "ams0", VpsName: "example-vps", Description: "test-description"}
	err := repo.Order(order)

	require.NoError(t, err)
}

func TestBlockStorageRepository_UpgradeBlockStorage(t *testing.T) {
	const expectedRequest = `{"blockStorageName":"example-blockstorage","size":8,"offsiteBackups":true}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	err := repo.Upgrade("example-blockstorage", 8, true)

	require.NoError(t, err)
}

func TestBlockStorageRepository_UpdateBlockStorage(t *testing.T) {
	const expectedRequest = `{"blockStorage":{"name":"example-blockstorage","description":"Block storage description","diskSize":2147483648,"offsiteBackups":true,"vpsName":"example-vps","status":"active","serial":"e7e12b3c7c6602973ac7","isLocked":false,"availabilityZone":"ams0","productType":"fast-storage"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	blockStorage := BlockStorage{
		Name:             "example-blockstorage",
		Description:      "Block storage description",
		DiskSize:         2147483648,
		OffsiteBackups:   true,
		VpsName:          "example-vps",
		Status:           BlockStorageStatusActive,
		IsLocked:         false,
		AvailabilityZone: "ams0",
		Serial:           "e7e12b3c7c6602973ac7",
		ProductType:      "fast-storage",
	}
	err := repo.Update(blockStorage)

	require.NoError(t, err)
}

func TestBlockStorageRepository_DetachVpsFromBlockStorage(t *testing.T) {
	const expectedRequest = `{"blockStorage":{"name":"example-blockstorage","description":"Block storage description","diskSize":2147483648,"offsiteBackups":true,"vpsName":"example-vps","status":"active","serial":"e7e12b3c7c6602973ac7","isLocked":false,"availabilityZone":"ams0","productType":"fast-storage"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	blockStorage := BlockStorage{
		Name:             "example-blockstorage",
		Description:      "Block storage description",
		DiskSize:         2147483648,
		OffsiteBackups:   true,
		Status:           "active",
		IsLocked:         false,
		AvailabilityZone: "ams0",
		Serial:           "e7e12b3c7c6602973ac7",
		ProductType:      "fast-storage",
	}
	err := repo.AttachToVps("example-vps", blockStorage)
	require.NoError(t, err)
}

func TestBlockStorageRepository_AttachVpsToBlockStorage(t *testing.T) {
	const expectedRequest = `{"blockStorage":{"name":"example-blockstorage","description":"Block storage description","diskSize":2147483648,"offsiteBackups":true,"vpsName":"","status":"active","serial":"e7e12b3c7c6602973ac7","isLocked":false,"availabilityZone":"ams0","productType":"fast-storage"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	blockStorage := BlockStorage{
		Name:             "example-blockstorage",
		Description:      "Block storage description",
		DiskSize:         2147483648,
		OffsiteBackups:   true,
		VpsName:          "example-vps",
		Status:           "active",
		IsLocked:         false,
		AvailabilityZone: "ams0",
		Serial:           "e7e12b3c7c6602973ac7",
		ProductType:      "fast-storage",
	}
	err := repo.DetachFromVps(blockStorage)
	require.NoError(t, err)
}

func TestBlockStorageRepository_CancelBlockStorage(t *testing.T) {
	const expectedRequest = `{"endTime":"end"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage", ExpectedMethod: "DELETE", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	err := repo.Cancel("example-blockstorage", gotransip.CancellationTimeEnd)
	require.NoError(t, err)
}

func TestBlockStorageRepository_GetBlockStorageBackups(t *testing.T) {
	const apiResponse = `{ "backups": [ { "id": 1583, "status": "active", "diskSize": 4294967296, "dateTimeCreate": "2019-12-31 09:13:55", "availabilityZone": "ams0" } ] }`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage/backups", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	all, err := repo.GetBackups("example-blockstorage")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 1583, all[0].ID)
	assert.EqualValues(t, "active", all[0].Status)
	assert.EqualValues(t, 4294967296, all[0].DiskSize)
	assert.Equal(t, "2019-12-31 09:13:55", all[0].DateTimeCreate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "ams0", all[0].AvailabilityZone)
}

func TestBlockStorageRepository_RevertBlockStorageBackup(t *testing.T) {
	const expectedRequest = `{"action":"revert"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage/backups/123", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	err := repo.RevertBackup("example-blockstorage", 123)
	require.NoError(t, err)
}

func TestBlockStorageRepository_RevertBackupToOtherBlockStorage(t *testing.T) {
	const expectedRequest = `{"action":"revert","destinationBlockStorageName":"example-blockStorage2"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/block-storages/example-blockstorage/backups/123", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	err := repo.RevertBackupToOtherBlockStorage("example-blockstorage", 123, "example-blockStorage2")
	require.NoError(t, err)
}

func TestBlockStorageRepository_GetBlockStorageUsage(t *testing.T) {
	const apiResponse = `{ "usage": [ { "iopsRead": 0.27, "iopsWrite": 0.13, "date": 1574783109 } ] }`

	parameters := url.Values{
		"dateTimeStart": []string{"1500538995"},
		"dateTimeEnd":   []string{"1500542619"},
	}

	expectedURL := "/block-storages/example-blockstorage/usage?" + parameters.Encode()
	server := testutil.MockServer{T: t, ExpectedURL: expectedURL, ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	usageData, err := repo.GetUsage("example-blockstorage", UsagePeriod{TimeStart: 1500538995, TimeEnd: 1500542619})
	require.NoError(t, err)

	require.Equal(t, 1, len(usageData))
	assert.EqualValues(t, 0.27, usageData[0].IopsRead)
	assert.EqualValues(t, 0.13, usageData[0].IopsWrite)
	assert.EqualValues(t, 1574783109, usageData[0].Date)
}

func TestBlockStorageRepository_GetBlockStorageUsageLast24Hours(t *testing.T) {
	const apiResponse = `{ "usage": [ { "iopsRead": 0.27, "iopsWrite": 0.13, "date": 1574783109 } ] }`

	parameters := url.Values{
		"dateTimeStart": []string{fmt.Sprintf("%d", time.Now().Add(-24*time.Hour).Unix())},
		"dateTimeEnd":   []string{fmt.Sprintf("%d", time.Now().Unix())},
	}

	expectedURL := "/block-storages/example-blockstorage/usage?" + parameters.Encode()
	server := testutil.MockServer{T: t, ExpectedURL: expectedURL, ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := BlockStorageRepository{Client: *client}

	usageData, err := repo.GetUsageLast24Hours("example-blockstorage")
	require.NoError(t, err)

	require.Equal(t, 1, len(usageData))
	assert.EqualValues(t, 0.27, usageData[0].IopsRead)
	assert.EqualValues(t, 0.13, usageData[0].IopsWrite)
	assert.EqualValues(t, 1574783109, usageData[0].Date)
}
