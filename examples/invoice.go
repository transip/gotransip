package main

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/invoice"
	"log"
)

func main() {
	// Create a new client with a default client config, using the demo token
	clientConfig := gotransip.ClientConfiguration{DemoMode: true}
	client, err := gotransip.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	invoiceRepo := invoice.Repository{Client: client}
	log.Println("Getting a list of invoices")
	invoices, err := invoiceRepo.GetAll()
	if err != nil {
		panic(err)
	}

	// Simple loop to print invoices with total amount including vat
	// Check out the inv structs to learn more about which data you can use
	// For more info about the invoices on the api, see: https://api.transip.nl/rest/docs.html#account-invoices
	for _, inv := range invoices {
		fmt.Printf("Invoice '%s' with total amount = '%d' \n", inv.InvoiceNumber, inv.TotalAmountInclVat)
	}
}
