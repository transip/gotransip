package kubernetes

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository is the kubernetes repository
// this repository allows you to manage all Kubernetes services for your TransIP account
type Repository repository.RestRepository

// GetClusters returns a list of all your VPSs
func (r *Repository) GetClusters() ([]Cluster, error) {
	var response clustersWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/clusters"}
	err := r.Client.Get(restRequest, &response)

	return response.Clusters, err
}

// GetClusterByName returns information on a specific cluster by name
func (r *Repository) GetClusterByName(clusterName string) (Cluster, error) {
	var response clusterWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", clusterName)}
	err := r.Client.Get(restRequest, &response)

	return response.Cluster, err
}

// CreateCluster allows you to order a new cluster
func (r *Repository) CreateCluster(clusterOrder ClusterOrder) error {
	restRequest := rest.Request{Endpoint: "/kubernetes/clusters", Body: &clusterOrder}

	return r.Client.Post(restRequest)
}

// UpdateCluster allows you to updated the description of a cluster
func (r *Repository) UpdateCluster(cluster Cluster) error {
	requestBody := clusterWrapper{Cluster: cluster}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", cluster.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// HandoverCluster will handover a cluster to another TransIP Account. This call will initiate the handover process.
// The actual handover will be done when the target customer accepts the handover
func (r *Repository) HandoverCluster(clusterName string, targetCustomerName string) error {
	requestBody := handoverRequest{Action: "handover", TargetCustomerName: targetCustomerName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", clusterName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// RemoveCluster will cancel the cluster, thus deleting it
func (r *Repository) RemoveCluster(clusterName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", clusterName)}

	return r.Client.Delete(restRequest)
}

// GetKubeConfig returns the Config YAML with admin credentials for given cluster
func (r *Repository) GetKubeConfig(clusterName string) (string, error) {
	var response struct {
		Config struct {
			YAML string `json:"encodedYaml"`
		} `json:"kubeConfig"`
	}

	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/kubeconfig", clusterName)}
	err := r.Client.Get(restRequest, &response)
	if err != nil {
		return "", err
	}

	yaml, err := base64.URLEncoding.DecodeString(response.Config.YAML)
	return string(yaml), err
}

// GetNodePools returns all node pools
func (r *Repository) GetNodePools() ([]NodePool, error) {
	var response nodePoolsWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/node-pools"}
	err := r.Client.Get(restRequest, &response)

	return response.NodePools, err
}

// GetNodePoolsByClusterName returns all node pools for a given clusterName
func (r *Repository) GetNodePoolsByClusterName(clusterName string) ([]NodePool, error) {
	var response nodePoolsWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/node-pools", Parameters: url.Values{"clusterName": []string{clusterName}}}
	err := r.Client.Get(restRequest, &response)

	return response.NodePools, err
}

// GetNodePool returns the NodePool for given nodePoolUUID
func (r *Repository) GetNodePool(nodePoolUUID string) (NodePool, error) {
	var response nodePoolWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/node-pools/%s", nodePoolUUID)}
	err := r.Client.Get(restRequest, &response)

	return response.NodePool, err
}

// AddNodePool allows you to order a new node pool to a cluster
func (r *Repository) AddNodePool(nodePoolOrder NodePoolOrder) error {
	restRequest := rest.Request{Endpoint: "/kubernetes/node-pools", Body: &nodePoolOrder}

	return r.Client.Post(restRequest)
}

// UpdateNodePool allows you to update the description and desired node count of a node pool
func (r *Repository) UpdateNodePool(nodePool NodePool) error {
	requestBody := nodePoolWrapper{NodePool: nodePool}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/node-pools/%s", nodePool.UUID), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveNodePool will cancel the node pool, thus deleting it
func (r *Repository) RemoveNodePool(nodePoolUUID string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/node-pools/%s", nodePoolUUID)}

	return r.Client.Delete(restRequest)
}

// GetNodes returns all nodes
func (r *Repository) GetNodes() ([]Node, error) {
	var response nodesWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/nodes"}
	err := r.Client.Get(restRequest, &response)

	return response.Nodes, err
}

// GetNodesByClusterName returns all nodes for a cluster
func (r *Repository) GetNodesByClusterName(clusterName string) ([]Node, error) {
	var response nodesWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/nodes", Parameters: url.Values{"clusterName": []string{clusterName}}}
	err := r.Client.Get(restRequest, &response)

	return response.Nodes, err
}

// GetNodesByNodePoolUUID returns all nodes for a node pool
func (r *Repository) GetNodesByNodePoolUUID(nodePoolUUID string) ([]Node, error) {
	var response nodesWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/nodes", Parameters: url.Values{"nodePoolUuid": []string{nodePoolUUID}}}
	err := r.Client.Get(restRequest, &response)

	return response.Nodes, err
}

// GetNode return a node
func (r *Repository) GetNode(nodeUUID string) (Node, error) {
	var response nodeWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/nodes/%s", nodeUUID)}
	err := r.Client.Get(restRequest, &response)

	return response.Node, err
}

// GetBlockStorageVolumes returns all block storage volumes
func (r *Repository) GetBlockStorageVolumes() ([]BlockStorage, error) {
	var response blockStoragesWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/block-storages"}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorages, err
}

// GetBlockStorageVolume returns a specific block storage volume
func (r *Repository) GetBlockStorageVolume(name string) (BlockStorage, error) {
	var response blockStorageWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/block-storages/%s", name)}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorage, err
}

// AddBlockStorageVolume creates a block storage volume
func (r *Repository) AddBlockStorageVolume(order BlockStorageOrder) error {
	restRequest := rest.Request{Endpoint: "/kubernetes/block-storages", Body: &order}

	return r.Client.Post(restRequest)
}

// UpdateBlockStorageVolume allows you to update the name and attached node for a block storage volumes
func (r *Repository) UpdateBlockStorageVolume(volume BlockStorage) error {
	requestBody := blockStorageWrapper{BlockStorage: volume}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/block-storages/%s", volume.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveBlockStorageVolume will remove a block storage volume
func (r *Repository) RemoveBlockStorageVolume(name string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/block-storages/%s", name)}

	return r.Client.Delete(restRequest)
}

// GetLoadBalancers returns all load balancers
func (r *Repository) GetLoadBalancers() ([]LoadBalancer, error) {
	var response lbsWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/load-balancers"}
	err := r.Client.Get(restRequest, &response)

	return response.LoadBalancers, err
}

// GetLoadBalancer returns a load balancer
func (r *Repository) GetLoadBalancer(name string) (LoadBalancer, error) {
	var response lbWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/load-balancers/%s", name)}
	err := r.Client.Get(restRequest, &response)

	return response.LoadBalancer, err
}

// CreateLoadBalancer creates a new load balancer
func (r *Repository) CreateLoadBalancer(name string) error {
	restRequest := rest.Request{Endpoint: "/kubernetes/load-balancers", Body: &lbOrder{Name: name}}

	return r.Client.Post(restRequest)
}

// UpdateLoadBalancer updates the entire state of the load balancer
func (r *Repository) UpdateLoadBalancer(name string, config LoadBalancerConfig) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/load-balancers/%s", name), Body: &lbcWrapper{Config: config}}

	return r.Client.Put(restRequest)
}

// RemoveLoadBalancer will remove a load balancer
func (r *Repository) RemoveLoadBalancer(name string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/load-balancers/%s", name)}

	return r.Client.Delete(restRequest)
}
