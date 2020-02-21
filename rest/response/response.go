package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/transip/gotransip/v6/rest"
)

// Every error returned by the api contains an error key with the message
type RestException struct {
	Message string `json:"error"`
}

// A typical rest response will contain a body, status code and the method
// which the response corresponds to
type RestResponse struct {
	Body       []byte
	StatusCode int
	Method     rest.RestMethod
}

func (r *RestResponse) ParseResponse(result interface{}) error {
	// do response error checking
	if !r.Method.StatusCodeIsCorrect(r.StatusCode) {
		// there is no response content so we also don't need to parse it
		if len(r.Body) == 0 {
			return fmt.Errorf("error response without body from api server status code '%d'", r.StatusCode)
		}

		var errorResponse RestException
		err := json.Unmarshal(r.Body, &errorResponse)
		if err != nil {
			return fmt.Errorf("response error could not be decoded, response = %s", string(r.Body))
		}

		return errors.New(errorResponse.Message)
	}

	if len(r.Body) == 0 {
		return nil
	}

	return json.Unmarshal(r.Body, &result)
}
