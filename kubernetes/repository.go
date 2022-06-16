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

// OrderCluster allows you to order a new cluster
func (r *Repository) OrderCluster(clusterOrder ClusterOrder) error {
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

// CancelCluster will cancel the cluster, thus deleting it
func (r *Repository) CancelCluster(clusterName string) error {
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

// GetNodePools returns all node pools for a cluster
func (r *Repository) GetNodePools(clusterName string) ([]NodePool, error) {
	var response nodePoolsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools", clusterName)}
	err := r.Client.Get(restRequest, &response)

	return response.NodePools, err
}

// GetNodePool returns the NodePool for given clusterName and nodePoolUUID
func (r *Repository) GetNodePool(clusterName, nodePoolUUID string) (NodePool, error) {
	var response nodePoolWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s", clusterName, nodePoolUUID)}
	err := r.Client.Get(restRequest, &response)

	return response.NodePool, err
}

// OrderNodePool allows you to order a new node pool to a cluster
func (r *Repository) OrderNodePool(clusterName string, nodePoolOrder NodePoolOrder) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools", clusterName), Body: &nodePoolOrder}

	return r.Client.Post(restRequest)
}

// UpdateNodePool allows you to update the description and desired node count of a node pool
func (r *Repository) UpdateNodePool(clusterName string, nodePool NodePool) error {
	requestBody := nodePoolWrapper{NodePool: nodePool}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s", clusterName, nodePool.UUID), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// CancelNodePool will cancel the node pool of a cluster, thus deleting it
func (r *Repository) CancelNodePool(clusterName, nodePoolUUID string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s/node-pools/%s", clusterName, nodePoolUUID)}

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
