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
	// Version of kubernetes this cluster is running
	Version string `json:"version"`
	// URL to connect to with kubectl
	Endpoint string `json:"endpoint"`
	// Whether or not another process is already doing stuff with this cluster
	IsLocked bool `json:"isLocked"`
	// If the cluster is administratively blocked
	IsBlocked bool `json:"isBlocked"`
}

// ClusterOrder struct can be used to order a new cluster
type ClusterOrder struct {
	// VPS ProductName for the WorkerNodes in the initial NodePool
	NodeSpec string `json:"nodeSpec"`
	// The desired amount of nodes in the initial NodePool
	NodeCount int `json:"desiredNodeCount"`
	// Availability Zone the WorkerNodes of the initial pool will spawn
	AvailabilityZone string `json:"availabilityZone"`
	// The description of the cluster
	Description string `json:"description"`
	// The specific version the Cluster should run on
	Version string `json:"kubernetesVersion"`
}
