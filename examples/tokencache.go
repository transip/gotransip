package main

import (
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/test"
	"log"
)

func main() {
	// specify a filesystem token cache with a location in tmp
	// in this case the new requested token will be stored at: /tmp/gotransip_token_cache
	cache, err := authenticator.NewFileTokenCache("/tmp/gotransip_token_cache")
	if err != nil {
		panic(err.Error())
	}
	// Create a new client and provide the private key and the tokencache
	config := gotransip.ClientConfiguration{
		AccountName:    "accountname",
		PrivateKeyPath: "/path/to/private.key",
		TokenCache:     cache,
	}
	client, err := gotransip.NewClient(config)
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
