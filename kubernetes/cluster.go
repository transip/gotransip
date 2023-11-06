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

// resetRequest is used to request a cluster reset, this is solely used for marshalling
type resetRequest struct {
	Action       string `json:"action"`
	Confirmation string `json:"confirmation"`
}

// upgradeRequest is used to request a cluster upgrade, this is solely used for marshalling

type upgradeRequest struct {
	Action  string `json:"action"`
	Version string `json:"version"`
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
	// The description of the cluster
	Description string `json:"description,omitempty"`
}
