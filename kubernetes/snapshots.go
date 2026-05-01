package kubernetes

import "github.com/transip/gotransip/v6/rest"

type blockStorageSnapshotsWrapper struct {
	BlockStorageSnapshots []BlockStorageSnapshot `json:"snapshots"`
}

type blockStorageSnapshotWrapper struct {
	BlockStorageSnapshot BlockStorageSnapshot `json:"snapshot"`
}

type BlockStorageSnapshot struct {
	// The unique identifier for the volume
	UUID string `json:"uuid"`
	// User configurable unique identifier (max 64 chars)
	Name string `json:"name"`
	// CreationDate is the date when the snapshot was created
	CreationDate rest.Date `json:"creationDate"`
	// ClusterName is the name of the cluster the blockstorage belongs to
	ClusterName string `json:"clusterName"`
	// The volume's size in gibibytes
	SizeInGiB int `json:"sizeInGib"`
	// Status of the volume 'attached', 'attaching', 'available', 'creating',
	// 'deleting' or 'detaching'
	Status BlockStorageSnapshotStatus `json:"status,omitempty"`
	// blockStorageName references the source of this snapshot
	BlockStorageName string `json:"blockstorageName"`
}

type BlockStorageSnapshotStatus string

const (
	BlockStorageSnapshotReady     BlockStorageSnapshotStatus = "ready"
	BlockStorageSnapshotCreating  BlockStorageSnapshotStatus = "creating"
	BlockStorageSnapshotReverting BlockStorageSnapshotStatus = "reverting"
	BlockStorageSnapshotDeleting  BlockStorageSnapshotStatus = "deleting"
)

type BlockStorageSnapshotOrder struct {
	ClusterName      string `json:"clusterName"`
	BlockStorageName string `json:"blockStorageName"`
	Name             string `json:"name"`
}
