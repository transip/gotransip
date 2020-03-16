package main

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/domain"
	"log"
	"strings"
)

func main() {
	// Create a new client with a default client config, using the demo token
	clientConfig := gotransip.ClientConfiguration{DemoMode: true}
	client, err := gotransip.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	domainRepo := domain.Repository{Client: client}
	log.Println("Getting a list of domains")
	domains, err := domainRepo.GetAll()
	if err != nil {
		panic(err)
	}

	// Simple loop to print your listed domains
	// For more info about the domains api, see: https://api.transip.nl/rest/docs.html#account-invoices
	fmt.Println(strings.Repeat("-", 50))
	for _, domain := range domains {
		fmt.Printf("Domain '%s' with tags: '%s' \n", domain.Name, domain.Tags)
	}
	fmt.Println(strings.Repeat("-", 50))

	defaultDnsEntry := domain.DnsEntry{
		Name:    "localhost",
		Expire:  86400,
		Type:    "A",
		Content: "127.0.0.1",
	}
	// Add a default dns entry to all of my domains
	for _, domain := range domains {
		log.Printf("Added dnsEntry '%v' to domain '%s'\n", defaultDnsEntry, domain.Name)
		domainRepo.AddDnsEntry(domain.Name, defaultDnsEntry)
	}
}
