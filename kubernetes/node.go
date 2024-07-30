package kubernetes

import "net"

// nodesWrapper struct contains Nodes in it,
// this is solely used for unmarshalling/marshalling
type nodesWrapper struct {
	Nodes []Node `json:"nodes"`
}

// nodeWrapper struct contains a Node in it,
// this is solely used for unmarshalling/marshalling
type nodeWrapper struct {
	Node Node `json:"node"`
}

// actionWrapper struct contains an action in it,
// this is solely used for marshalling
type actionWrapper struct {
	Action string `json:"action"`
}

// Node struct is a single node in a kubernetes cluster node pool
type Node struct {
	// The unique identifier for the node
	UUID string `json:"uuid"`
	// The unique identifier for the node pool this node belongs to
	NodePoolUUID string `json:"nodePoolUUID"`
	// The name of the cluster this node belongs to
	ClusterName string `json:"clusterName"`
	// The node's status
	Status NodeStatus `json:"status,omitempty"`
	// The node's IP addresses
	IPAddresses []NodeAddress `json:"ipAddresses"`
}

// NodeStatus is one of the following strings
// 'active', 'creating', 'deleting'
type NodeStatus string

// Definition of all of the possible node statuses
const (
	// NodeStatusActive is the status for an active node ready for workload
	NodeStatusActive NodeStatus = "active"
	// NodeStatusCreating is the status for a node that is being provisioned
	NodeStatusCreating NodeStatus = "creating"
	// NodeStatusDeleting is the status for a node that is being shutdown for removal
	NodeStatusDeleting NodeStatus = "deleting"
)

// NodeAddress defines the structure of 1 single node address
type NodeAddress struct {
	Address net.IP          `json:"address"`
	Netmask net.IP          `json:"subnetMask"`
	Type    NodeAddressType `json:"type"`
}

// NodeAddressType is one of the following strings
// 'external', 'internal'
type NodeAddressType string

// Definition of all of the possible node address types
const (
	// NodeAddressTypeExternal is an external node address
	NodeAddressTypeExternal NodeAddressType = "external"
	// NodeAddressTypeInternal is an internal node address
	NodeAddressTypeInternal NodeAddressType = "internal"
)
