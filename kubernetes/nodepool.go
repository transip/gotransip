package kubernetes

// nodePoolsWrapper struct contains a list of NodePools in it,
// this is solely used for unmarshalling/marshalling
type nodePoolsWrapper struct {
	NodePools []NodePool `json:"nodePools"`
}

// nodePoolWrapper struct contains a NodePool in it,
// this is solely used for unmarshalling/marshalling
type nodePoolWrapper struct {
	NodePool NodePool `json:"nodePool"`
}

// NodePool struct for a kubernetes cluster node pool
type NodePool struct {
	// The unique identifier for the node pool
	UUID string `json:"uuid"`
	// The name that can be set by customer
	Description string `json:"description"`
	// Amount of desired nodes in the node pool
	DesiredNodeCount int `json:"desiredNodeCount"`
	// Specification for nodes in this node pool
	NodeSpec string `json:"nodeSpec"`
	// Nodes in this node pool
	Nodes []Node `json:"nodes,omitempty"`
}

// Node struct is a single node in a kubernetes cluster node pool
type Node struct {
	// The unique identifier for the node
	UUID string `json:"uuid"`
	// The VPS status, either 'created', 'installing', 'running', 'stopped' or 'paused'
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

// NodePoolOrder struct can be used to order a new node pool to a cluster
type NodePoolOrder struct {
	// The description of the node pool
	Description string `json:"description,omitempty"`
	// Amount of desired nodes in the node pool
	DesiredNodeCount int `json:"desiredNodeCount"`
	// Specification for nodes in this node pool
	NodeSpec string `json:"nodeSpec"`
}
