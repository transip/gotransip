package main

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/vps"
	"log"
	"time"
)

func main() {
	// Create a new client with the default demo client config, using the demo token
	client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)
	if err != nil {
		panic(err)
	}

	vpsRepo := vps.Repository{Client: client}

	// Create a usage period with which you can query Vps usage for the last 10 minutes
	log.Println("Getting a list of cpu usage for the last 10 minutes")
	period := vps.UsagePeriod{
		TimeStart: time.Now().Unix() - 600,
		TimeEnd:   time.Now().Unix(),
	}

	// Query with our usage period for the Vps 'transipdemo-vps4'
	usage, err := vpsRepo.GetUsage("transipdemo-vps4", []vps.UsageType{vps.UsageTypeCPU}, period)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", usage)
}
