// Package kubernetes is an EXPERIMENTAL endpoint for controlling the resources of a kubernetes
// cluster. It is not yet stable and may receive breaking changes.
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
func (r *Repository) GetNodePools(clusterName string) ([]NodePool, error) {
	var response nodePoolsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools", clusterName)}
	err := r.Client.Get(restRequest, &response)

	return response.NodePools, err
}

// GetNodePool returns the NodePool for given nodePoolUUID
func (r *Repository) GetNodePool(clusterName, nodePoolUUID string) (NodePool, error) {
	var response nodePoolWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s", clusterName, nodePoolUUID)}
	err := r.Client.Get(restRequest, &response)

	return response.NodePool, err
}

// AddNodePool allows you to order a new node pool to a cluster
func (r *Repository) AddNodePool(nodePoolOrder NodePoolOrder) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools", nodePoolOrder.ClusterName), Body: &nodePoolOrder}

	return r.Client.Post(restRequest)
}

// UpdateNodePool allows you to update the description and desired node count of a node pool
func (r *Repository) UpdateNodePool(nodePool NodePool) error {
	requestBody := nodePoolWrapper{NodePool: nodePool}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s", nodePool.ClusterName, nodePool.UUID), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveNodePool will cancel the node pool, thus deleting it
func (r *Repository) RemoveNodePool(clusterName, nodePoolUUID string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s", clusterName, nodePoolUUID)}

	return r.Client.Delete(restRequest)
}

// GetNodes returns all nodes
func (r *Repository) GetNodes(clusterName string) ([]Node, error) {
	var response nodesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/nodes", clusterName)}
	err := r.Client.Get(restRequest, &response)

	return response.Nodes, err
}

// GetNodesByNodePoolUUID returns all nodes for a node pool
func (r *Repository) GetNodesByNodePoolUUID(clusterName, nodePoolUUID string) ([]Node, error) {
	var response nodesWrapper
	restRequest := rest.Request{
		Endpoint:   fmt.Sprintf("/kubernetes/clusters/%s/nodes", clusterName),
		Parameters: url.Values{"nodePoolUuid": []string{nodePoolUUID}},
	}
	err := r.Client.Get(restRequest, &response)

	return response.Nodes, err
}

// GetNode return a node
func (r *Repository) GetNode(clusterName, nodeUUID string) (Node, error) {
	var response nodeWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/nodes/%s", clusterName, nodeUUID)}
	err := r.Client.Get(restRequest, &response)

	return response.Node, err
}

// GetBlockStorageVolumes returns all block storage volumes
func (r *Repository) GetBlockStorageVolumes(clusterName string) ([]BlockStorage, error) {
	var response blockStoragesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/block-storages", clusterName)}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorages, err
}

// GetBlockStorageVolume returns a specific block storage volume
func (r *Repository) GetBlockStorageVolume(clusterName, name string) (BlockStorage, error) {
	var response blockStorageWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/block-storages/%s", clusterName, name)}
	err := r.Client.Get(restRequest, &response)

	return response.BlockStorage, err
}

// AddBlockStorageVolume creates a block storage volume
func (r *Repository) AddBlockStorageVolume(order BlockStorageOrder) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/block-storages", order.ClusterName), Body: &order}

	return r.Client.Post(restRequest)
}

// UpdateBlockStorageVolume allows you to update the name and attached node for a block storage volumes
func (r *Repository) UpdateBlockStorageVolume(volume BlockStorage) error {
	requestBody := blockStorageWrapper{BlockStorage: volume}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/block-storages/%s", volume.ClusterName, volume.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveBlockStorageVolume will remove a block storage volume
func (r *Repository) RemoveBlockStorageVolume(clusterName, name string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/block-storages/%s", clusterName, name)}

	return r.Client.Delete(restRequest)
}

// GetLoadBalancers returns all load balancers
func (r *Repository) GetLoadBalancers(clusterName string) ([]LoadBalancer, error) {
	var response lbsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/load-balancers", clusterName)}
	err := r.Client.Get(restRequest, &response)

	return response.LoadBalancers, err
}

// GetLoadBalancer returns a load balancer
func (r *Repository) GetLoadBalancer(clusterName, name string) (LoadBalancer, error) {
	var response lbWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/load-balancers/%s", clusterName, name)}
	err := r.Client.Get(restRequest, &response)

	return response.LoadBalancer, err
}

// CreateLoadBalancer creates a new load balancer
func (r *Repository) CreateLoadBalancer(clusterName, name string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/load-balancers", clusterName), Body: &lbOrder{Name: name}}

	return r.Client.Post(restRequest)
}

// UpdateLoadBalancer updates the entire state of the load balancer
func (r *Repository) UpdateLoadBalancer(clusterName, name string, config LoadBalancerConfig) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/load-balancers/%s", clusterName, name), Body: &lbcWrapper{Config: config}}

	return r.Client.Put(restRequest)
}

// RemoveLoadBalancer will remove a load balancer
func (r *Repository) RemoveLoadBalancer(clusterName, name string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/load-balancers/%s", clusterName, name)}

	return r.Client.Delete(restRequest)
}

// GetTaints will get all the taints on a NodePool
func (r *Repository) GetTaints(clusterName, nodePoolUUID string) ([]Taint, error) {
	var response taintWrapper
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s/taints", clusterName, nodePoolUUID),
	}

	err := r.Client.Get(restRequest, &response)
	return response.Taints, err
}

// SetTaints will set the taints on a NodePool
func (r *Repository) SetTaints(clusterName, nodePoolUUID string, taints []Taint) error {
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s/taints", clusterName, nodePoolUUID),
		Body:     &taintWrapper{Taints: taints},
	}

	return r.Client.Put(restRequest)
}

// GetLabels will get the labels on a NodePool
func (r *Repository) GetLabels(clusterName, nodePoolUUID string) ([]Label, error) {
	var response labelWrapper
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s/labels", clusterName, nodePoolUUID),
	}

	err := r.Client.Get(restRequest, &response)
	return response.Labels, err
}

// SetLabels will set the labels on a NodePool
func (r *Repository) SetLabels(clusterName, nodePoolUUID string, labels []Label) error {
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s/labels", clusterName, nodePoolUUID),
		Body:     &labelWrapper{Labels: labels},
	}

	return r.Client.Put(restRequest)
}
