package main

import (
	"fmt"
	"log"

	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/action"
	"github.com/transip/gotransip/v6/vps"
)

func main() {
	// Create a new client with the default demo client config, using the demo token
	client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)
	if err != nil {
		panic(err)
	}

	vpsName := "example-vps"
	snapshotName := "example-snapshot"

	vpsRepo := vps.Repository{Client: client}
	actionRepo := action.Repository{Client: client}
	log.Println("Execution snapshot revert")
	response, err := vpsRepo.RevertSnapshotWithResponse(vpsName, snapshotName)
	if err != nil {
		panic(err)
	}
	action, err := actionRepo.ParseActionFromResponse(response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", action)
}
