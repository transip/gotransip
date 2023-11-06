package main

import (
	"fmt"
	"log"

	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/kubernetes"
)

func main() {
	client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)
	if err != nil {
		panic(err)
	}

	k8sRepo := kubernetes.Repository{Client: client}

	// List all clusters with nodepools
	log.Println("List all clusters and the acccompanying nodepools")
	clusters, err := k8sRepo.GetClusters()
	if err != nil {
		panic(err)
	}

	for _, c := range clusters {
		fmt.Printf("Cluster: %s\n", c.Name)
		nodepools, err := k8sRepo.GetNodePools(c.Name)
		if err != nil {
			continue
		}
		for _, n := range nodepools {
			fmt.Printf(
				"\tNodePool: [description=%s, zone=%s, spec=%s, nodeCount=%s]\n",
				n.Description, n.AvailabilityZone, n.NodeSpec, n.DesiredNodeCount,
			)
		}

	}

	log.Println("Create new cluster and wait for it to be available")
	clusters, err = k8sRepo.GetClusters()
	if err != nil {
		panic(err)
	}
}
