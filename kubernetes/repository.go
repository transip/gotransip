package kubernetes

import (
	"encoding/base64"
	"fmt"

	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository is the kubernetes repository
// this repository allows you to manage all Kubernetes services for your TransIP account
type Repository repository.RestRepository

// GetAll returns a list of all your VPSs
func (r *Repository) GetClusters() ([]Cluster, error) {
	var response clustersWrapper
	restRequest := rest.Request{Endpoint: "/kubernetes/clusters"}
	err := r.Client.Get(restRequest, &response)

	return response.Clusters, err
}

// GetByName returns information on a specific cluster by name
func (r *Repository) GetClusterByName(clusterName string) (Cluster, error) {
	var response clusterWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", clusterName)}
	err := r.Client.Get(restRequest, &response)

	return response.Cluster, err
}

// Order allows you to order a new cluster
func (r *Repository) OrderCluster(clusterOrder ClusterOrder) error {
	restRequest := rest.Request{Endpoint: "/kubernetes/clusters", Body: &clusterOrder}

	return r.Client.Post(restRequest)
}

func (r *Repository) UpdateCluster(cluster Cluster) error {
	requestBody := clusterWrapper{Cluster: cluster}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", cluster.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Handover will handover a cluster to another TransIP Account. This call will initiate the handover process.
// The actual handover will be done when the target customer accepts the handover
func (r *Repository) HandoverCluster(clusterName string, targetCustomerName string) error {
	requestBody := handoverRequest{Action: "handover", TargetCustomerName: targetCustomerName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", clusterName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Cancel will cancel the cluster, thus deleting it
func (r *Repository) CancelCluster(clusterName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/kubernetes/clusters/%s", clusterName)}

	return r.Client.Delete(restRequest)
}

func (r *Repository) GetKubeConfig(clusterName string) (string, error) {
	var response struct {
		Config struct {
			YAML string `json:"yaml"`
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
