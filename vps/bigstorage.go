package vps

import (
	"fmt"
	"net/url"
	"time"

	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
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
	// Description that the big storage should have after ordering
	Description string `json:"description,omitempty"`
}

// BigStorageStatus is one of the following strings
// 'active', 'attaching', 'detaching'
type BigStorageStatus string

// Definition of all of the possible bigstorage backup statuses
const (
	// BigStorageStatusActive is the status field for an active BigStorage, ready to use
	BigStorageStatusActive BigStorageStatus = "active"
	// BigStorageStatusAttaching is the status field for a BigStorage that is being attached to a vps
	BigStorageStatusAttaching BigStorageStatus = "attaching"
	// BigStorageStatusDetaching is the status field for a BigStorage that is being detached from a vps
	BigStorageStatusDetaching BigStorageStatus = "detaching"
)

// BigStorage struct for a BigStorage
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
	Status BigStorageStatus `json:"status,omitempty"`
	// Serial of the big storage. This is a unique identifier that is visible by the vps it has been attached to. On
	// linux servers it is visible using udevadm info /dev/vdb where it will be the value of ID_SERIAL. A symlink will
	// also be created in /dev/disk-by-id/ containing the serial. This is useful if you want to map a disk inside a VPS
	// to a big storage.
	Serial string `json:"serial"`
	// Lock status of the big storage, when it is locked, it cannot be attached or detached.
	IsLocked bool `json:"isLocked"`
	// The availability zone the bigstorage is located in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// BigStorageBackup struct for a BigStorageBackup
type BigStorageBackup struct {
	// ID of the big storage
	ID int64 `json:"id,omitempty"`
	// Status of the big storage backup ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status BackupStatus `json:"status,omitempty"`
	// The backup disk size in kB
	DiskSize int64 `json:"diskSize"`
	// Date of the big storage backup
	DateTimeCreate rest.Time `json:"dateTimeCreate,omitempty"`
	// The name of the availability zone the backup is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// GetAll returns a list of your bigstorages
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) GetAll() ([]BigStorage, error) {
	var response bigStoragesWrapper
	restRequest := rest.Request{Endpoint: "/big-storages"}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorages, err
}

// GetSelection returns a limited list of bigstorages,
// specify how many and which page/chunk of your bigstorage you want to retrieve
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) GetSelection(page int, itemsPerPage int) ([]BigStorage, error) {
	var response bigStoragesWrapper
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/big-storages", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorages, err
}

// GetByName returns a specific BigStorage struct by name
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) GetByName(bigStorageName string) (BigStorage, error) {
	var response bigStorageWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorage, err
}

// Order allows you to order a new bigstorage
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) Order(order BigStorageOrder) error {
	restRequest := rest.Request{Endpoint: "/big-storages", Body: &order}

	return r.Client.Post(restRequest)
}

// OrderWithResponse allows you to order a new bigstorage and returns a response
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) OrderWithResponse(order BigStorageOrder) (rest.Response, error) {
	restRequest := rest.Request{Endpoint: "/big-storages", Body: &order}

	return r.Client.PostWithResponse(restRequest)
}

// Upgrade allows you to upgrade a BigStorage's size or/and to enable off-site backups
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) Upgrade(bigStorageName string, size int, offsiteBackups bool) error {
	requestBody := bigStorageUpgradeRequest{BigStorageName: bigStorageName, Size: size, OffsiteBackups: offsiteBackups}
	restRequest := rest.Request{Endpoint: "/big-storages", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Update allows you to alter the BigStorage in several ways outlined below:
//   - Changing the description of a Big Storage;
//   - One Big Storages can only be attached to one VPS at a time;
//   - One VPS can have a maximum of 10 bigstorages attached;
//   - Set the vpsName property to the VPS name to attach to for attaching Big Storage;
//   - Set the vpsName property to null to detach the Big Storage from the currently attached VPS.
//
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) Update(bigStorage BigStorage) error {
	requestBody := bigStorageWrapper{BigStorage: bigStorage}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorage.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// UpdateWithResponse returns a response
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) UpdateWithResponse(bigStorage BigStorage) (rest.Response, error) {
	requestBody := bigStorageWrapper{BigStorage: bigStorage}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorage.Name), Body: &requestBody}

	return r.Client.PutWithResponse(restRequest)
}

// DetachFromVps allows you to detach a bigstorage from the vps it is attached to
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) DetachFromVps(bigStorage BigStorage) error {
	bigStorage.VpsName = ""

	return r.Update(bigStorage)
}

// AttachToVps allows you to attach a given VPS by name to a BigStorage
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) AttachToVps(vpsName string, bigStorage BigStorage) error {
	bigStorage.VpsName = vpsName

	return r.Update(bigStorage)
}

// Cancel cancels a bigstorage for the specified endTime.
// You can set the endTime to end or immediately, this has the following implications:
//
//   - end: The Big Storage will be terminated from the end date of the agreement as can be found in the applicable quote;
//   - immediately: The Big Storage will be terminated immediately.
//
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) Cancel(bigStorageName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorageName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetBackups returns a list of backups for a specific bigstorage
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) GetBackups(bigStorageName string) ([]BigStorageBackup, error) {
	var response bigStorageBackupsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s/backups", bigStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorageBackups, err
}

// RevertBackup allows you to revert a bigstorage by bigstorage name and backupID
// if you want to revert a backup to a different big storage you can use the RevertBackupToOtherBigStorage method
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) RevertBackup(bigStorageName string, backupID int64) error {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s/backups/%d", bigStorageName, backupID), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// RevertBackupWithResponse allows you to revert a bigstorage by bigstorage name and backupID and returns a response
// if you want to revert a backup to a different big storage you can use the RevertBackupToOtherBigStorage method
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) RevertBackupWithResponse(bigStorageName string, backupID int64) (rest.Response, error) {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s/backups/%d", bigStorageName, backupID), Body: &requestBody}

	return r.Client.PatchWithResponse(restRequest)
}

// RevertBackupToOtherBigStorage allows you to revert a backup to a different big storage
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) RevertBackupToOtherBigStorage(bigStorageName string, backupID int64, destinationBigStorageName string) error {
	requestBody := bigStorageRestoreBackupsWrapper{Action: "revert", DestinationBigStorageName: destinationBigStorageName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s/backups/%d", bigStorageName, backupID), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// RevertBackupToOtherBigStorageWithResponse allows you to revert a backup to a different big storage and returns a response
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) RevertBackupToOtherBigStorageWithResponse(
	bigStorageName string,
	backupID int64,
	destinationBigStorageName string,
) (rest.Response, error) {
	requestBody := bigStorageRestoreBackupsWrapper{Action: "revert", DestinationBigStorageName: destinationBigStorageName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s/backups/%d", bigStorageName, backupID), Body: &requestBody}

	return r.Client.PatchWithResponse(restRequest)
}

// GetUsage allows you to query your bigstorage usage within a certain period
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) GetUsage(bigStorageName string, period UsagePeriod) ([]UsageDataDisk, error) {
	var response usageDataDiskWrapper
	parameters := url.Values{
		"dateTimeStart": []string{fmt.Sprintf("%d", period.TimeStart)},
		"dateTimeEnd":   []string{fmt.Sprintf("%d", period.TimeEnd)},
	}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/big-storages/%s/usage", bigStorageName), Parameters: parameters}

	err := r.Client.Get(restRequest, &response)

	return response.Usage, err
}

// GetUsageLast24Hours allows you to get usage statistics for a given bigstorage within the last 24 hours
// Deprecated: Use block storage resource instead
func (r *BigStorageRepository) GetUsageLast24Hours(bigStorageName string) ([]UsageDataDisk, error) {
	// always define a period body, this way we don't have to depend on the empty body logic on the api server
	period := UsagePeriod{TimeStart: time.Now().Add(-24 * time.Hour).Unix(), TimeEnd: time.Now().Unix()}

	return r.GetUsage(bigStorageName, period)
}
