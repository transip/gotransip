package response

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/rest"
	"testing"
)

func TestResponseParsing(t *testing.T) {
	responseBody := []byte(`{"name": "test"}`)
	restResponse := RestResponse{
		Body:       responseBody,
		StatusCode: 200,
		Method:     rest.GetRestMethod,
	}

	var responseObject struct {
		Name string `json:"name"`
	}

	err := restResponse.ParseResponse(&responseObject)
	assert.NoError(t, err)
	assert.Equal(t, "test", responseObject.Name)
}

func TestErrorResponse(t *testing.T) {
	error := RestException{Message: "this should be returned"}
	data, err := json.Marshal(error)
	assert.NoError(t, err)

	restResponse := RestResponse{
		Body:       data,
		StatusCode: 406,
		Method:     rest.GetRestMethod,
	}

	err = restResponse.ParseResponse(nil)
	require.Error(t, err)
	assert.Equal(t, errors.New("this should be returned"), err)

	restResponse.Body = []byte{0x41}
	err = restResponse.ParseResponse(nil)
	require.Error(t, err)
	assert.Equal(t, errors.New("response error could not be decoded, response = A"), err)
}

func TestEmptyResponse(t *testing.T) {
	restResponse := RestResponse{
		StatusCode: 201,
		Method:     rest.PostRestMethod,
	}

	err := restResponse.ParseResponse(nil)
	require.NoError(t, err)
}

func TestEmptyErrorResponse(t *testing.T) {
	restResponse := RestResponse{
		StatusCode: 500,
		Method:     rest.PostRestMethod,
	}

	err := restResponse.ParseResponse(nil)
	require.Error(t, err)
	assert.Equal(t, errors.New("error response without body from api server status code '500'"), err)
}
