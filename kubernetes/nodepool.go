package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

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
	// The name of the cluster the node pool should be ordered for
	ClusterName string `json:"clusterName"`
	// The name that can be set by customer
	Description string `json:"description"`
	// Amount of desired nodes in the node pool
	DesiredNodeCount int `json:"desiredNodeCount"`
	// Specification for nodes in this node pool
	NodeSpec string `json:"nodeSpec"`
	// Labels to be set on new nodes in the node pool
	Labels map[string]string `json:"labels"`
	// Taints to be set on new nodes in the node pool
	Taints []v1.Taint `json:"taints"`
	// Nodes in this node pool
	Nodes []Node `json:"nodes,omitempty"`
}

// NodePoolOrder struct can be used to order a new node pool to a cluster
type NodePoolOrder struct {
	// The name of the cluster the node pool should be ordered for
	ClusterName string `json:"clusterName"`
	// The description of the node pool
	Description string `json:"description,omitempty"`
	// Amount of desired nodes in the node pool
	DesiredNodeCount int `json:"desiredNodeCount"`
	// Specification for nodes in this node pool
	NodeSpec string `json:"nodeSpec"`
	// Labels to be set on new nodes in the node pool
	Labels map[string]string `json:"labels"`
	// Taints to be set on new nodes in the node pool
	Taints []v1.Taint `json:"taints"`
}
