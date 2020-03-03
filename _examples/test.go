package main

import (
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/test"
	"log"
)

func main() {
	// Create a new client with a default client config, using the demo token
	clientConfig := gotransip.ClientConfiguration{DemoMode: true}
	client, err := gotransip.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	testRepo := test.Repository{Client: client}
	log.Println("Executing test call to the api server")
	if err := testRepo.Test(); err != nil {
		panic(err)
	}
	log.Println("Test successful")
}
