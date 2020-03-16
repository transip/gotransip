package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"time"
)

// BigStorageRepository allows you to manage all api actions on a bigstorage
// getting information, ordering, upgrading, attaching/detaching it to a vps
type BigStorageRepository repository.RestRepository

// BigStorageOrder struct which is used to construct a new order request for a bigstorage
type BigStorageOrder struct {
	// The size of the big storage in TB's, use a multitude of 2. The maximum size is 40.
	Size int `json:"size"`
	// Whether to order offsite backups, omit this to use current value
	OffsiteBackups bool `json:"offsiteBackups"`
	// The name of the availabilityZone where the BigStorage should be created. This parameter can not be used in conjunction with vpsName
	// If a vpsName is provided as well as an availabilityZone, the zone of the vps is leading
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The name of the VPS to attach the big storage to
	VpsName string `json:"vpsName"`
}

// BigStorage struct for BigStorage
type BigStorage struct {
	// Name of the big storage
	Name string `json:"name,omitempty"`
	// Name that can be set by customer
	Description string `json:"description"`
	// Disk size of the big storage in kB
	DiskSize int64 `json:"diskSize,omitempty"`
	// Whether a bigstorage has backups
	OffsiteBackups bool `json:"offsiteBackups"`
	// The VPS that the big storage is attached to
	VpsName string `json:"vpsName"`
	// Status of the big storage can be 'active', 'attaching' or 'detachting'
	Status string `json:"status,omitempty"`
	// Lock status of the big storage, when it is locked, it cannot be attached or detached.
	IsLocked bool `json:"isLocked"`
	// The availability zone the bigstorage is located in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// BigStorageBackup struct for BigStorageBackup
type BigStorageBackup struct {
	// Id of the big storage
	Id int64 `json:"id,omitempty"`
	// Status of the big storage backup ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status string `json:"status,omitempty"`
	// The backup disk size in kB
	DiskSize int64 `json:"diskSize"`
	// Date of the big storage backup
	DateTimeCreate rest.Time `json:"dateTimeCreate,omitempty"`
	// The name of the availability zone the backup is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// GetBigStorages returns a list of your bigstorages
func (r *BigStorageRepository) GetBigStorages() ([]BigStorage, error) {
	var response bigStoragesWrapper
	restRequest := rest.RestRequest{Endpoint: "/big-storages"}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorages, err
}

// GetBigStorageByName returns a specific BigStorage struct by name
func (r *BigStorageRepository) GetBigStorageByName(bigStorageName string) (BigStorage, error) {
	var response bigStorageWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorage, err
}

// OrderBigStorage allows you to order a new bigstorage
func (r *BigStorageRepository) OrderBigStorage(order BigStorageOrder) error {
	restRequest := rest.RestRequest{Endpoint: "/big-storages", Body: &order}

	return r.Client.Post(restRequest)
}

// UpgradeBigStorage allows you to upgrade a BigStorage's size or/and to enable off-site backups
func (r *BigStorageRepository) UpgradeBigStorage(bigStorageName string, size int, offsiteBackups bool) error {
	requestBody := bigStorageUpgradeRequest{BigStorageName: bigStorageName, Size: size, OffsiteBackups: offsiteBackups}
	restRequest := rest.RestRequest{Endpoint: "/big-storages", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateBigStorage allows you to alter the BigStorage in several ways outlined below:
// - Changing the description of a Big Storage;
// - One Big Storages can only be attached to one VPS at a time;
// - One VPS can have a maximum of 10 bigstorages attached;
// - Set the vpsName property to the VPS name to attach to for attaching Big Storage;
// - Set the vpsName property to null to detach the Big Storage from the currently attached VPS.
func (r *BigStorageRepository) UpdateBigStorage(bigStorage BigStorage) error {
	requestBody := bigStorageWrapper{BigStorage: bigStorage}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorage.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// DetachVpsFromBigStorage allows you to detach a bigstorage from the vps it is attached to
func (r *BigStorageRepository) DetachVpsFromBigStorage(bigStorage BigStorage) error {
	bigStorage.VpsName = ""

	return r.UpdateBigStorage(bigStorage)
}

// AttachVpsToBigStorage allows you to attach a given VPS by name to a BigStorage
func (r *BigStorageRepository) AttachVpsToBigStorage(vpsName string, bigStorage BigStorage) error {
	bigStorage.VpsName = vpsName

	return r.UpdateBigStorage(bigStorage)
}

// CancelBigStorage cancels a bigstorage for the specified endTime.
// You can set the endTime to end or immediately, this has the following implications:
// - end: The Big Storage will be terminated from the end date of the agreement as can be found in the applicable quote;
// - immediately: The Big Storage will be terminated immediately.
func (r *BigStorageRepository) CancelBigStorage(bigStorageName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorageName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetBigStorageBackups returns a list of backups for a specific bigstorage
func (r *BigStorageRepository) GetBigStorageBackups(bigStorageName string) ([]BigStorageBackup, error) {
	var response bigStorageBackupsWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s/backups", bigStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorageBackups, err
}

// RevertBigStorageBackup allows you to revert a bigstorage by bigstorage name and backupId
func (r *BigStorageRepository) RevertBigStorageBackup(bigStorageName string, backupId int64) error {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s/backups/%d", bigStorageName, backupId), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// GetBigStorageUsage allows you to query your bigstorage usage within a certain period
func (r *BigStorageRepository) GetBigStorageUsage(bigStorageName string, period UsagePeriod) ([]UsageDataDisk, error) {
	var response usageDataDiskWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s/usage", bigStorageName), Body: &period}

	err := r.Client.Get(restRequest, &response)

	return response.Usage, err
}

// GetBigStorageUsageLast24Hours allows you to get usage statistics for a given bigstorage within the last 24 hours
func (r *BigStorageRepository) GetBigStorageUsageLast24Hours(bigStorageName string) ([]UsageDataDisk, error) {
	// always define a period body, this way we don't have to depend on the empty body logic on the api server
	period := UsagePeriod{TimeStart: time.Now().Unix() - 24*3600, TimeEnd: time.Now().Unix()}

	return r.GetBigStorageUsage(bigStorageName, period)
}
