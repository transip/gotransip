package main

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/email"
	"log"
	"strings"
)

func main() {
	// Create a new client with the default demo client config, using the demo token
	client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)
	if err != nil {
		panic(err)
	}

	emailRepo := email.Repository{Client: client}

	//transipdemo.net
	log.Println("Getting a list of email boxes")
	mailboxes, err := emailRepo.GetMailboxesByDomainName("transipdemo.net")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 50))
	for _, mailbox := range mailboxes {
		alias := mailbox.LocalPart
		if len(alias) == 0 {
			alias = "*"
		}
		fmt.Printf("Mailbox Identifier: %s, Available Disk Space: %d Status is %s \n", mailbox.Identifier, mailbox.AvailableDiskSpace, mailbox.Status)
	}
	fmt.Println(strings.Repeat("-", 50))

	log.Println("Getting a list of email forwards")
	forwards, err := emailRepo.GetMailforwardsByDomainName("transipdemo.net")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 50))
	for _, forward := range forwards {
		alias := forward.LocalPart
		if len(alias) == 0 {
			alias = "*"
		}
		fmt.Printf("Forward ID %d Forwards to: '%s' whene receiving mail on %s@%s. Status is %s \n", forward.ID, forward.ForwardTo, alias, forward.Domain, forward.Status)
	}

	fmt.Println(strings.Repeat("-", 50))

	log.Println("Getting a list of email lists")
	maillists, err := emailRepo.GetMaillistsByDomainName("transipdemo.net")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 50))
	for _, list := range maillists {
		fmt.Printf("List ID %d Name: '%s'. Email Address: %s \n", list.ID, list.Name, list.EmailAddress)
	}

	fmt.Println(strings.Repeat("-", 50))

	log.Println("Getting a list of email addons")
	addons, err := emailRepo.GetAddonsByDomainName("transipdemo.net")
	if err != nil {
		panic(err)
	}

	// Simple loop to print your addons
	// For more info about the email api, see: https://api.transip.nl/rest/docs.html#email
	fmt.Println(strings.Repeat("-", 50))
	for _, addon := range addons {
		fmt.Printf("Addon ID %d with disk space: %d, additional mailboxes: '%d' \n", addon.ID, addon.DiskSpace, addon.Mailboxes)
	}
	fmt.Println(strings.Repeat("-", 50))

	log.Println("Getting a list of mail packages")
	mailpackages, err := emailRepo.GetMailpackages()
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 50))
	for _, pkg := range mailpackages {
		fmt.Printf("Package for domain %s and status %s", pkg.Domain, pkg.Status)
	}
	fmt.Println(strings.Repeat("-", 50))
}
