package kubernetes

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
}

// NodeStatus is one of the following strings
// 'active'
type NodeStatus string

// Definition of all of the possible vps statuses
const (
	// NodeStatusActive is the status for an active node ready for workload
	NodeStatusActive NodeStatus = "active"
)
