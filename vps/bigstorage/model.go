package bigstorage

// BigStorage struct for BigStorage
type BigStorage struct {
	// The availability zone the bigstorage is located in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// Name that can be set by customer
	Description string `json:"description"`
	// Disk size of the big storage in kB
	DiskSize float32 `json:"diskSize,omitempty"`
	// Lock status of the big storage, when it is locked, it cannot be attached or detached.
	IsLocked bool `json:"isLocked,omitempty"`
	// Name of the big storage
	Name string `json:"name,omitempty"`
	// Whether a bigstorage has backups
	OffsiteBackups bool `json:"offsiteBackups,omitempty"`
	// Status of the big storage can be 'active', 'attaching' or 'detachting'
	Status string `json:"status,omitempty"`
	// The VPS that the big storage is attached to
	VpsName string `json:"vpsName"`
}

// BigStorageBackup struct for BigStorageBackup
type BigStorageBackup struct {
	// The name of the availability zone the backup is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// Date of the big storage backup
	DateTimeCreate string `json:"dateTimeCreate,omitempty"`
	// The backup disk size in kB
	DiskSize float32 `json:"diskSize"`
	// Id of the big storage
	Id float32 `json:"id,omitempty"`
	// Status of the big storage backup ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status string `json:"status,omitempty"`
}

// BigStorageBackups struct for BigStorageBackups
type BigStorageBackups struct {
	// array of BigStorageBackup
	BigStorageBackups []BigStorageBackup `json:"bigStorageBackups"`
}

// BigStorages struct for BigStorages
type BigStorages struct {
	// array of BigStorage
	BigStorages []BigStorage `json:"bigStorages"`
}
