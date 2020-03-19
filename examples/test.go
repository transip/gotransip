package main

import (
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/test"
	"log"
)

func main() {
	// Create a new client with the default demo client config, using the demo token
	client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)
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
