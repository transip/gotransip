package rest

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestResponseParsing(t *testing.T) {
	responseBody := []byte(`{"name": "test"}`)
	restResponse := Response{Body: responseBody, StatusCode: 200, Method: GetMethod}

	var responseObject struct {
		Name string `json:"name"`
	}

	err := restResponse.ParseResponse(&responseObject)
	assert.NoError(t, err)
	assert.Equal(t, "test", responseObject.Name)
}

func TestErrorResponse(t *testing.T) {
	error := RestError{Message: "this should be returned"}
	data, err := json.Marshal(error)
	assert.NoError(t, err)

	restResponse := Response{Body: data, StatusCode: 406, Method: GetMethod}

	err = restResponse.ParseResponse(nil)
	require.Error(t, err)
	assert.Equal(t, errors.New("this should be returned"), err)

	restResponse.Body = []byte{0x41}
	err = restResponse.ParseResponse(nil)
	require.Error(t, err)
	assert.Equal(t, errors.New("response error could not be decoded, response = A"), err)
}

func TestEmptyResponse(t *testing.T) {
	restResponse := Response{StatusCode: 201, Method: PostMethod}

	err := restResponse.ParseResponse(nil)
	require.NoError(t, err)
}

func TestEmptyErrorResponse(t *testing.T) {
	restResponse := Response{
		StatusCode: 500,
		Method:     PostMethod,
	}

	err := restResponse.ParseResponse(nil)
	require.Error(t, err)
	assert.Equal(t, errors.New("error response without body from api server status code '500'"), err)
}

func TestResponseDateParsing(t *testing.T) {
	responseBody := []byte(`{"date": "2020-01-02"}`)
	restResponse := Response{Body: responseBody, StatusCode: 200, Method: GetMethod}

	var responseObject struct {
		Date Date `json:"date"`
	}

	err := restResponse.ParseResponse(&responseObject)
	assert.NoError(t, err)
	assert.Equal(t, 2020, responseObject.Date.Year())
	assert.Equal(t, time.Month(1), responseObject.Date.Month())
	assert.Equal(t, 2, responseObject.Date.Day())
}

func TestResponseTimeParsing(t *testing.T) {
	responseBody := []byte(`{"cancellationDate": "2020-01-02 12:13:37"}`)
	restResponse := Response{Body: responseBody, StatusCode: 200, Method: GetMethod}

	var responseObject struct {
		Date Time `json:"cancellationDate"`
	}

	err := restResponse.ParseResponse(&responseObject)
	assert.NoError(t, err)
	assert.Equal(t, "2020-01-02 12:13:37 +0100 CET", responseObject.Date.String())
}

func TestResponseEmptyTimeParsing(t *testing.T) {
	responseBody := []byte(`{"cancellationDate": "", "cancellationDatetime": ""}`)
	restResponse := Response{Body: responseBody, StatusCode: 200, Method: GetMethod}

	var responseObject struct {
		DateTime Time `json:"cancellationDatetime"`
		Date     Date `json:"cancellationDate"`
	}

	err := restResponse.ParseResponse(&responseObject)
	assert.NoError(t, err)
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", responseObject.DateTime.String())
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", responseObject.Date.String())
}
