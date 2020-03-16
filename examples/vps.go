package main

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/vps"
	"log"
)

func main() {
	// Create a new client with a default client config, using the demo token
	clientConfig := gotransip.ClientConfiguration{DemoMode: true}
	client, err := gotransip.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	vpsRepo := vps.Repository{Client: client}
	log.Println("Getting a list of vpses")
	vpss, err := vpsRepo.GetAll()
	if err != nil {
		panic(err)
	}

	// Simple loop to print vpses with their ip addresses
	// Check out the vps structs to learn more about which data you can use
	// For more info about the vps api, see: https://api.transip.nl/rest/docs.html#vps-vps-get
	for _, v := range vpss {
		ips, err := vpsRepo.GetIPAddresses(v.Name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Vps '%s' with ips:\n", v.Name)
		for _, ip := range ips {
			fmt.Printf("- %s\n", ip.Address)
		}
	}
}
