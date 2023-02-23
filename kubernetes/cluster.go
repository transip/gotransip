package kubernetes

// clustersWrapper struct contains a list of Clusters in it,
// this is solely used for unmarshalling/marshalling
type clustersWrapper struct {
	Clusters []Cluster `json:"clusters"`
}

// clusterWrapper struct contains a cluster in it,
// this is solely used for unmarshalling/marshalling
type clusterWrapper struct {
	Cluster Cluster `json:"cluster"`
}

// handoverRequest is used to request a handover, this is solely used for marshalling
type handoverRequest struct {
	Action             string `json:"action"`
	TargetCustomerName string `json:"targetCustomerName"`
}

// Cluster struct for a Kubernetes cluster
type Cluster struct {
	// The unique Cluster name
	Name string `json:"name"`
	// The name that can be set by customer
	Description string `json:"description"`
	// Whether or not another process is already doing stuff with this cluster
	IsLocked bool `json:"isLocked,omitempty"`
	// If the cluster is administratively blocked
	IsBlocked bool `json:"isBlocked,omitempty"`
}

// ClusterOrder struct can be used to order a new cluster
type ClusterOrder struct {
	// The description of the cluster
	Description string `json:"description,omitempty"`
}
