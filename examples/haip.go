package main

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/haip"
	"log"
)

func main() {
	// Create a new client with a default client config, using the demo token
	clientConfig := gotransip.ClientConfiguration{DemoMode: true}
	client, err := gotransip.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	haipRepo := haip.Repository{Client: client}
	log.Println("Getting a list of haips")
	haips, err := haipRepo.GetAll()
	if err != nil {
		panic(err)
	}

	// Simple loop to print haips with their ip addresses
	// Check out the Haip structs and haip.Repository to learn more about which data you can use
	// For more info about the haip api, see: https://api.transip.nl/rest/docs.html#ha-ip
	for _, v := range haips {
		ips, err := haipRepo.GetAttachedIPAddresses(v.Name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("HA-IP '%s' with attached IPs:\n", v.Name)
		for _, ip := range ips {
			fmt.Printf("- %s\n", ip)
		}
	}
}
