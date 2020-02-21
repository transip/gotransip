package test

import (
	"errors"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/rest/request"
)

type ApiTestRepository gotransip.Service

func (r *ApiTestRepository) Test() error {
	var testResponse ApiTest
	request := request.RestRequest{Endpoint: "/api-test"}

	err := r.Client.Get(request, &testResponse)
	if err != nil {
		return err
	}

	if testResponse.Response != "pong" {
		return errors.New("Test api response doesn't equal 'pong'")
	}

	return nil
}
