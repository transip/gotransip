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
	// ClusterName is the name of the cluster the blockstorage belongs to
	ClusterName string `json:"clusterName"`
	// The volume's size in gibibytes
	SizeInGiB int `json:"sizeInGib"`
	// Type of storage
	Type BlockStorageType `json:"type"`
	// AvailabilityZone where this volume is located
	AvailabilityZone string `json:"availabilityZone"`
	// Status of the volume 'attached', 'attaching', 'available', 'creating',
	// 'deleting' or 'detaching'
	Status BlockStorageStatus `json:"status,omitempty"`
	// UUID of node this volume is attached to
	NodeUUID string `json:"nodeUuid"`
	// The serial of the disk. This is a unique identifier that is visible by the node it has been attached to.
	Serial string `json:"serial"`
}

// BlockStorageStatus is one of the following statuses
// 'attached', 'attaching', 'available', 'creating', 'deleting' or 'detaching'
type BlockStorageStatus string

// Definition of all of the possible block storage statuses
const (
	// BlockStorageStatusAttached is the status for a volume that currently is attached to a node
	BlockStorageStatusAttached BlockStorageStatus = "attached"
	// BlockStorageStatusAttaching is the status for a volume that is being attached to a node
	BlockStorageStatusAttaching BlockStorageStatus = "attaching"
	// BlockStorageStatusAvailable is the status for a volume that is available
	BlockStorageStatusAvailable BlockStorageStatus = "available"
	// BlockStorageStatusCreating is the status for a volume that is being created
	BlockStorageStatusCreating BlockStorageStatus = "creating"
	// BlockStorageStatusDeleting is the status for a volume that is being deleted
	BlockStorageStatusDeleting BlockStorageStatus = "deleting"
	// BlockStorageStatusDetaching is the status for a volume that is being detached from a node
	BlockStorageStatusDetaching BlockStorageStatus = "detaching"
)

// BlockStorageType is one of the following types
// 'hdd'
type BlockStorageType string

// Description of all the possible block storage types
const (
	// BlockStorageTypeHDD reflects a block storage volume of type HDD
	BlockStorageTypeHDD BlockStorageType = "hdd"
)

// BlockStorageOrder struct can be used to order a new block storage volume
type BlockStorageOrder struct {
	// user adjustable unique identifier for volume (max 64 chars), when none is given, the uuid will be used.
	Name string `json:"name"`
	// ClusterName name of the cluster the block storage will be available for
	ClusterName string `json:"-"`
	// amount of storage in gibibytes
	SizeInGiB int `json:"sizeInGib"`
	// type of storage
	Type BlockStorageType `json:"type"`
	// location of the volume
	AvailabilityZone string `json:"availabilityZone"`
}
