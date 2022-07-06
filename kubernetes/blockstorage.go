package kubernetes

// blockStoragesWrapper struct contains BlockStorages in it,
// this is solely used for unmarshalling/marshalling
type blockStoragesWrapper struct {
	BlockStorages []BlockStorage `json:"volumes"`
}

// blockStorageWrapper struct contains a BlockStorage in it,
// this is solely used for unmarshalling/marshalling
type blockStorageWrapper struct {
	BlockStorage BlockStorage `json:"volume"`
}

// BlockStorage struct is a single block storage volume that can be used in a kubernetes cluster
type BlockStorage struct {
	// The unique identifier for the volume
	UUID string `json:"uuid"`
	// User configurable unique identifier (max 64 chars)
	Name string `json:"name"`
	// The volume's size in gibibytes
	SizeInGiB int `json:"sizeInGib"`
	// Type of storage
	Type string `json:"type"`
	// AvailabilityZone where this volume is located
	AvailabilityZone string `json:"availabilityZone"`
	// Status of the volume ‘available’, ‘attached’
	Status BlockStorageStatus `json:"status,omitempty"`
	// UUID of node this volume is attached to
	NodeUUID string `json:"nodeUuid"`
	// The serial of the disk. This is a unique identifier that is visible by the node it has been attached to.
	Serial string `json:"serial"`
}

// BlockStorageStatus is one of the following strings
// 'available', 'attached'
type BlockStorageStatus string

// Definition of all of the possible block storage statuses
const (
	// BlockStorageStatusAvailable is the status for a volume that is available
	BlockStorageStatusAvailable BlockStorageStatus = "available"
	// BlockStorageStatusAttached is the status for a volume that currently attached to a node
	BlockStorageStatusAttached BlockStorageStatus = "attached"
)

// BlockStorageOrder struct can be used to order a new block storage volume
type BlockStorageOrder struct {
	// user adjustable unique identifier for volume (max 64 chars), when none is given, the uuid will be used.
	Name string `json:"name"`
	// amount of storage in gibibytes
	SizeInGiB int `json:"sizeInGib"`
	// type of storage
	Type string `json:"type"`
	// location of the volume
	AvailabilityZone string `json:"availabilityZone"`
}
