package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	contentType = "application/json"
)

type RestRequest struct {
	Endpoint   string
	Parameters url.Values
	Body       interface{}
}

// GetBody returns the request object as a json
func (r *RestRequest) GetBody() ([]byte, error) {
	return json.Marshal(r.Body)
}

// Get body io.Reader
func (r *RestRequest) GetBodyReader() (io.Reader, error) {
	if r.Body == nil {
		return nil, nil
	}

	// try to get the marshalled body
	body, err := r.GetBody()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}

// GetHttpRequest generates and returns a http.Request object
// It does this with the RestRequest struct and the basePath and method,
// that are provided by the client itself
func (r *RestRequest) GetHttpRequest(basePath string, method string) (*http.Request, error) {
	url := basePath + r.Endpoint
	bodyReader, err := r.GetBodyReader()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	// set json headers, because our this library sends and expects that
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)

	// set the custom parameters on the rawquery
	request.URL.RawQuery = r.Parameters.Encode()

	return request, nil
}
