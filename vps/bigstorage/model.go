package bigstorage

// Order struct which is used to construct a new order request for a bigstorage
type Order struct {
	// The size of the big storage in TB's, use a multitude of 2. The maximum size is 40.
	Size float32 `json:"size"`
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
