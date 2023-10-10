package vps

import (
	"fmt"
	"net/url"
	"time"

	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// BlockStorageRepository allows you to manage all api actions on a blockstorage
// getting information, ordering, upgrading, attaching/detaching it to a vps
type BlockStorageRepository repository.RestRepository

// BlockStorageOrder struct which is used to construct a new order request for a blockstorage
type BlockStorageOrder struct {
	// The type of the block storage. It can be big-storage or fast-storage.
	Type string `json:"type"`
	// The size of the block storage in KB.
	// Big storages: The minimum size is 2 TiB and storage can be extended with up to maximum of 40 TiB. Make sure to
	// use a multiple of 2 TiB. Note that 2 TiB equals 2147483648 KiB.
	// Fast storages: The minimum size is 10 GiB and storage can be extended with up to maximum of 10000 GiB. Make sure
	// to use a multiple of 10 GiB. Note that 10 GiB equals 10485760 KiB.
	Size int `json:"size"`
	// Whether to order offsite backups, omit this to use current value
	OffsiteBackups bool `json:"offsiteBackups"`
	// The name of the availabilityZone where the BlockStorage should be created. This parameter can not be used in conjunction with vpsName
	// If a vpsName is provided as well as an availabilityZone, the zone of the vps is leading
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The name of the VPS to attach the block storage to
	VpsName string `json:"vpsName"`
	// Description that the block storage should have after ordering
	Description string `json:"description,omitempty"`
}

// BlockStorageStatus is one of the following strings
// 'active', 'attaching', 'detaching'
type BlockStorageStatus string

// Definition of all the possible blockstorage backup statuses
const (
	// BlockStorageStatusActive is the status field for an active BlockStorage, ready to use
	BlockStorageStatusActive BlockStorageStatus = "active"
	// BlockStorageStatusAttaching is the status field for a BlockStorage that is being attached to a vps
	BlockStorageStatusAttaching BlockStorageStatus = "attaching"
	// BlockStorageStatusDetaching is the status field for a BlockStorage that is being detached from a vps
	BlockStorageStatusDetaching BlockStorageStatus = "detaching"
)

// BlockStorage struct for a BlockStorage
type BlockStorage struct {
	// Name of the block storage
	Name string `json:"name,omitempty"`
	// Name that can be set by customer
	Description string `json:"description"`
	// Disk size of the block storage in kB
	DiskSize int64 `json:"diskSize,omitempty"`
	// Whether a blockstorage has backups
	OffsiteBackups bool `json:"offsiteBackups"`
	// The VPS that the block storage is attached to
	VpsName string `json:"vpsName"`
	// Status of the block storage can be 'active', 'attaching' or 'detachting'
	Status BlockStorageStatus `json:"status,omitempty"`
	// Serial of the block storage. This is a unique identifier that is visible by the vps it has been attached to. On
	// linux servers it is visible using udevadm info /dev/vdb where it will be the value of ID_SERIAL. A symlink will
	// also be created in /dev/disk-by-id/ containing the serial. This is useful if you want to map a disk inside a VPS
	// to a block storage.
	Serial string `json:"serial"`
	// Lock status of the block storage, when it is locked, it cannot be attached or detached.
	IsLocked bool `json:"isLocked"`
	// The availability zone the blockstorage is located in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The type of the block storage. It can be big-storage or fast-storage.
	ProductType string `json:"productType"`
}

// BlockStorageBackup struct for a BlockStorageBackup
type BlockStorageBackup struct {
	// ID of the block storage
	ID int64 `json:"id,omitempty"`
	// Status of the block storage backup ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status BackupStatus `json:"status,omitempty"`
	// The backup disk size in kB
	DiskSize int64 `json:"diskSize"`
	// Date of the block storage backup
	DateTimeCreate rest.Time `json:"dateTimeCreate,omitempty"`
	// The name of the availability zone the backup is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// GetAll returns a list of your blockstorages
func (r *BlockStorageRepository) GetAll() ([]BlockStorage, error) {
	var response blockStoragesWrapper
	restRequest := rest.Request{Endpoint: "/block-storages"}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorages, err
}

// GetSelection returns a limited list of blockstorages,
// specify how many and which page/chunk of your blockstorage you want to retrieve
func (r *BlockStorageRepository) GetSelection(page int, itemsPerPage int) ([]BlockStorage, error) {
	var response blockStoragesWrapper
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/block-storages", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorages, err
}

// GetByName returns a specific BlockStorage struct by name
func (r *BlockStorageRepository) GetByName(blockStorageName string) (BlockStorage, error) {
	var response blockStorageWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s", blockStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorage, err
}

// Order allows you to order a new blockstorage
func (r *BlockStorageRepository) Order(order BlockStorageOrder) error {
	restRequest := rest.Request{Endpoint: "/block-storages", Body: &order}

	return r.Client.Post(restRequest)
}

// OrderWithResponse allows you to order a new blockstorage and returns a response
func (r *BlockStorageRepository) OrderWithResponse(order BlockStorageOrder) (rest.Response, error) {
	restRequest := rest.Request{Endpoint: "/block-storages", Body: &order}

	return r.Client.PostWithResponse(restRequest)
}

// Upgrade allows you to upgrade a BlockStorage's size or/and to enable off-site backups
func (r *BlockStorageRepository) Upgrade(blockStorageName string, size int, offsiteBackups bool) error {
	requestBody := blockStorageUpgradeRequest{BlockStorageName: blockStorageName, Size: size, OffsiteBackups: offsiteBackups}
	restRequest := rest.Request{Endpoint: "/block-storages", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Update allows you to alter the BlockStorage in several ways outlined below:
//   - Changing the description of a Block Storage;
//   - One Block Storages can only be attached to one VPS at a time;
//   - One VPS can have a maximum of 10 blockstorages attached;
//   - Set the vpsName property to the VPS name to attach to for attaching Block Storage;
//   - Set the vpsName property to null to detach the Block Storage from the currently attached VPS.
func (r *BlockStorageRepository) Update(blockStorage BlockStorage) error {
	requestBody := blockStorageWrapper{BlockStorage: blockStorage}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s", blockStorage.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// UpdateWithResponse returns a response
func (r *BlockStorageRepository) UpdateWithResponse(blockStorage BlockStorage) (rest.Response, error) {
	requestBody := blockStorageWrapper{BlockStorage: blockStorage}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s", blockStorage.Name), Body: &requestBody}

	return r.Client.PutWithResponse(restRequest)
}

// DetachFromVps allows you to detach a blockstorage from the vps it is attached to
func (r *BlockStorageRepository) DetachFromVps(blockStorage BlockStorage) error {
	blockStorage.VpsName = ""

	return r.Update(blockStorage)
}

// AttachToVps allows you to attach a given VPS by name to a BlockStorage
func (r *BlockStorageRepository) AttachToVps(vpsName string, blockStorage BlockStorage) error {
	blockStorage.VpsName = vpsName

	return r.Update(blockStorage)
}

// Cancel cancels a blockstorage for the specified endTime.
// You can set the endTime to end or immediately, this has the following implications:
//
//   - end: The Block Storage will be terminated from the end date of the agreement as can be found in the applicable quote;
//   - immediately: The Block Storage will be terminated immediately.
func (r *BlockStorageRepository) Cancel(blockStorageName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s", blockStorageName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetBackups returns a list of backups for a specific blockstorage
func (r *BlockStorageRepository) GetBackups(blockStorageName string) ([]BlockStorageBackup, error) {
	var response blockStorageBackupsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s/backups", blockStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorageBackups, err
}

// RevertBackup allows you to revert a blockstorage by blockstorage name and backupID
// if you want to revert a backup to a different block storage you can use the RevertBackupToOtherBlockStorage method
func (r *BlockStorageRepository) RevertBackup(blockStorageName string, backupID int64) error {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s/backups/%d", blockStorageName, backupID), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// RevertBackupWithResponse allows you to revert a blockstorage by blockstorage name and backupID and returns a response
// if you want to revert a backup to a different block storage you can use the RevertBackupToOtherBlockStorage method
func (r *BlockStorageRepository) RevertBackupWithResponse(blockStorageName string, backupID int64) (rest.Response, error) {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s/backups/%d", blockStorageName, backupID), Body: &requestBody}

	return r.Client.PatchWithResponse(restRequest)
}

// RevertBackupToOtherBlockStorage allows you to revert a backup to a different block storage
func (r *BlockStorageRepository) RevertBackupToOtherBlockStorage(blockStorageName string, backupID int64, destinationBlockStorageName string) error {
	requestBody := blockStorageRestoreBackupsWrapper{Action: "revert", DestinationBlockStorageName: destinationBlockStorageName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s/backups/%d", blockStorageName, backupID), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// RevertBackupToOtherBlockStorageWithResponse allows you to revert a backup to a different block storage and returns a response
func (r *BlockStorageRepository) RevertBackupToOtherBlockStorageWithResponse(
	blockStorageName string,
	backupID int64,
	destinationBlockStorageName string,
) (rest.Response, error) {
	requestBody := blockStorageRestoreBackupsWrapper{Action: "revert", DestinationBlockStorageName: destinationBlockStorageName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s/backups/%d", blockStorageName, backupID), Body: &requestBody}

	return r.Client.PatchWithResponse(restRequest)
}

// GetUsage allows you to query your blockstorage usage within a certain period
func (r *BlockStorageRepository) GetUsage(blockStorageName string, period UsagePeriod) ([]UsageDataDisk, error) {
	var response usageDataDiskWrapper
	parameters := url.Values{
		"dateTimeStart": []string{fmt.Sprintf("%d", period.TimeStart)},
		"dateTimeEnd":   []string{fmt.Sprintf("%d", period.TimeEnd)},
	}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/block-storages/%s/usage", blockStorageName), Parameters: parameters}

	err := r.Client.Get(restRequest, &response)

	return response.Usage, err
}

// GetUsageLast24Hours allows you to get usage statistics for a given blockstorage within the last 24 hours
func (r *BlockStorageRepository) GetUsageLast24Hours(blockStorageName string) ([]UsageDataDisk, error) {
	// always define a period body, this way we don't have to depend on the empty body logic on the api server
	period := UsagePeriod{TimeStart: time.Now().Add(-24 * time.Hour).Unix(), TimeEnd: time.Now().Unix()}

	return r.GetUsage(blockStorageName, period)
}
