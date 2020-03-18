/*
	Package gotransip implements a client for the TransIP Rest API.
	This package is a complete implementation for communicating with the TransIP RestAPI.
	It covers resource calls available in the TransIP RestAPI Docs and it allows your
	project(s) to connect to the TransIP RestAPI easily. Using this library you can order,
	update and remove products from your TransIP account.

	As of version 6.0 this package is no longer compatible with TransIP SOAP API because the
	library is now organized around REST. The SOAP API library versions 5.* are now deprecated
	and will no longer receive future updates.

	Authentication

	If you want to tinker out with the api first without setting up your authentication,
	we defined a static DemoClientConfiguration.
	Which can be used to create a new client:
		client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)

	Create a new client using a token:
		client, err := gotransip.NewClient(gotransip.ClientConfiguration{
			Token:      "this_is_where_your_token_goes",
		})

	As tokens have a limited expiry time you can also request new tokens using the private key
	acquired from your transip control panel:
		client, err := gotransip.NewClient(gotransip.ClientConfiguration{
			AccountName:    "accountName",
			PrivateKeyPath: "/path/to/api/private.key",
		})

	We also implemented a PrivateKeyReader option, for users that want to store their key elsewhere,
	not on a filesystem but on X datastore:
		file, err := os.Open("/path/to/api/private.key")
		if err != nil {
			panic(err.Error())
		}
		client, err := gotransip.NewClient(gotransip.ClientConfiguration{
			AccountName:      "accountName",
			PrivateKeyReader: file,
		})

	Repositories

	All resource calls as can be seen on https://api.transip.nl/rest/docs.html
	have been grouped in the following repositories, these are subpackages under the gotransip package:
		availabilityzone.Repository
		colocation.Repository
		domain.Repository
		haip.Repository
		invoice.Repository
		ipaddress.Repository
		mailservice.Repository
		product.Repository
		test.Repository
		traffic.Repository
		vps.BigstorageRepository
		vps.PrivateNetworkRepository
		vps.Repository

	Such a repository can be initialised with a client as follows:
		domainRepo := domain.Repository{Client: client}

	Each repository has a bunch methods you can use to call get/modify/update resources in that specific subpackage.
	For example, here we get a list of domains from a transip account:
		log.Println("Getting a list of domains")
		domains, err := domainRepo.GetAll()

*/
package gotransip
