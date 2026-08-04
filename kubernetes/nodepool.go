package kubernetes

type NodePoolAutoscaling struct {
	// The minimum amount of nodes allowed in the NodePool
	MinNodeCount int
	// The maximum amount of nodes allowed in the NodePool
	MaxNodeCount int
	// Whether autoscaling is enabled on the NodePool
	Enabled bool
}

// NodePool struct for a kubernetes cluster node pool
type NodePool struct {
	// The unique identifier for the node pool
	UUID string
	// The name of the cluster the node pool belongs to
	ClusterName string
	// The name that can be set by customer
	Description string
	// Amount of desired nodes in the node pool
	DesiredNodeCount int
	// Specification for nodes in this node pool
	NodeSpec string
	// Availability zone of the node pool
	AvailabilityZone string
	// The NodePoolAutoScaling settings
	Autoscaling *NodePoolAutoscaling
	// Nodes in this node pool
	Nodes []Node
}

type nodePoolResponse struct {
	UUID               string `json:"uuid"`
	ClusterName        string `json:"clusterName"`
	Description        string `json:"description"`
	DesiredNodeCount   int    `json:"desiredNodeCount"`
	NodeSpec           string `json:"nodeSpec"`
	AvailabilityZone   string `json:"availabilityZone"`
	MinNodeCount       int    `json:"minNodeCount,omitempty"`
	MaxNodeCount       int    `json:"maxNodeCount,omitempty"`
	AutoscalingEnabled bool   `json:"autoscalingEnabled,omitempty"`
	Nodes              []Node `json:"nodes,omitempty"`
}

type nodePoolUpdateRequest struct {
	UUID             string `json:"uuid"`
	ClusterName      string `json:"clusterName"`
	Description      string `json:"description"`
	DesiredNodeCount int    `json:"desiredNodeCount"`
	NodeSpec         string `json:"nodeSpec"`
	AvailabilityZone string `json:"availabilityZone"`
	Nodes            []Node `json:"nodes,omitempty"`
}

type nodePoolAutoscalingUpdateRequest struct {
	UUID               string `json:"uuid"`
	ClusterName        string `json:"clusterName"`
	Description        string `json:"description"`
	DesiredNodeCount   int    `json:"desiredNodeCount"`
	NodeSpec           string `json:"nodeSpec"`
	AvailabilityZone   string `json:"availabilityZone"`
	MinNodeCount       int    `json:"minNodeCount"`
	MaxNodeCount       int    `json:"maxNodeCount"`
	AutoscalingEnabled bool   `json:"autoscalingEnabled"`
	Nodes              []Node `json:"nodes,omitempty"`
}

// NodePoolOrder struct can be used to order a new node pool to a cluster
type NodePoolOrder struct {
	// The name of the cluster the node pool should be ordered for
	ClusterName string `json:"-"`
	// The description of the node pool
	Description string `json:"description,omitempty"`
	// Amount of desired nodes in the node pool
	DesiredNodeCount int `json:"desiredNodeCount"`
	// Specification for nodes in this node pool
	NodeSpec string `json:"nodeSpec"`
	// Availability zone of the node pool
	AvailabilityZone string `json:"availabilityZone"`
}

// Taint truct for a taints on a NodePool
type Taint struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Effect     string `json:"effect"`
	Modifiable bool   `json:"modifiable"`
}

type taintWrapper struct {
	Taints []Taint `json:"taints"`
}

// Label Struct for a labels on a NodePool
type Label struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Modifiable bool   `json:"modifiable"`
}

type labelWrapper struct {
	Labels []Label `json:"labels"`
}

func (npr nodePoolResponse) intoNodePool() NodePool {
	return NodePool{
		UUID:             npr.UUID,
		ClusterName:      npr.ClusterName,
		Description:      npr.Description,
		DesiredNodeCount: npr.DesiredNodeCount,
		NodeSpec:         npr.NodeSpec,
		AvailabilityZone: npr.AvailabilityZone,
		Autoscaling: &NodePoolAutoscaling{
			MinNodeCount: npr.MinNodeCount,
			MaxNodeCount: npr.MaxNodeCount,
			Enabled:      npr.AutoscalingEnabled,
		},
		Nodes: npr.Nodes,
	}
}
