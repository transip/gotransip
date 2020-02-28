package request

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/url"
	"testing"
)

func TestRequestMarshalling(t *testing.T) {
	request := RestRequest{}

	body, err := request.GetBody()
	assert.NoError(t, err)
	assert.Equal(t, string(body), "null")
}

func TestHttpRequestForRestRequest(t *testing.T) {
	order := struct {
		AvailabilityZone string `json:"availabilityZone"`
		OperatingSystem  string `json:"operatingSystem"`
		ProductName      string `json:"productName"`
	}{AvailabilityZone: "ams", OperatingSystem: "ubuntu-18.04", ProductName: "vps-bladevps-x1",}

	values := url.Values{"test": []string{"1"}}

	request := RestRequest{Endpoint: "/vps", Parameters: values, Body: order}
	httpRequest, err := request.GetHTTPRequest("https://example.com", "POST")
	assert.NoError(t, err)
	assert.Equal(t, "POST", httpRequest.Method)
	assert.Equal(t, "test=1", httpRequest.URL.RawQuery)
	assert.Equal(t, "https://example.com/vps?test=1", httpRequest.URL.String())
	assert.Equal(t, "application/json", httpRequest.Header.Get("Content-Type"))
	assert.Equal(t, "application/json", httpRequest.Header.Get("Accept"))
	assert.Equal(t, int64(91), httpRequest.ContentLength)

	body, err := ioutil.ReadAll(httpRequest.Body)
	assert.Equal(t, "{\"availabilityZone\":\"ams\",\"operatingSystem\":\"ubuntu-18.04\",\"productName\":\"vps-bladevps-x1\"}", string(body))
}

func TestHttpRequestForEmptyGetRestRequest(t *testing.T) {
	request := RestRequest{Endpoint: "/domains"}
	httpRequest, err := request.GetHTTPRequest("https://example.com", "GET")
	assert.NoError(t, err)
	assert.Equal(t, "GET", httpRequest.Method)
	assert.Equal(t, "https://example.com/domains", httpRequest.URL.String())
	assert.Equal(t, "application/json", httpRequest.Header.Get("Content-Type"))
	assert.Equal(t, "application/json", httpRequest.Header.Get("Accept"))
	assert.Zero(t, httpRequest.ContentLength)
}

func TestBodyReader(t *testing.T) {
	order := struct {
		AvailabilityZone string `json:"availabilityZone"`
		OperatingSystem  string `json:"operatingSystem"`
		ProductName      string `json:"productName"`
	}{AvailabilityZone: "ams", OperatingSystem: "ubuntu-18.04", ProductName: "vps-bladevps-x1",}

	request := RestRequest{Endpoint: "/vps", Body: order}

	reader, err := request.GetBodyReader()
	require.NoError(t, err)

	body, err := ioutil.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, "{\"availabilityZone\":\"ams\",\"operatingSystem\":\"ubuntu-18.04\",\"productName\":\"vps-bladevps-x1\"}", string(body))
}

func TestEmptyReaderReturnsNil(t *testing.T) {
	request := RestRequest{Endpoint: "/domains"}

	reader, err := request.GetBodyReader()
	require.NoError(t, err)
	assert.Nil(t, reader)
}
