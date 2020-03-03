package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/transip/gotransip/v6/rest"
	"time"
)

// RestException is used to unpack every error returned by the api contains
type RestException struct {
	// Message contains the error from the api as string
	Message string `json:"error"`
}

// RestResponse will contain a body (which can be empty), status code and the RestMethod
// this struct will be used to decode a response from the api server
type RestResponse struct {
	Body       []byte
	StatusCode int
	Method     rest.RestMethod
}

// Time is defined because the transip api server does not return a rfc 3339 time string
// and golang requires this, so we need to do manual time parsing, by defining our own time struct
// encapsulating time.Time
type Time struct {
	// Time item containing the actual parsed time object
	time.Time
}

// Date is defined because the transip api server returns date strings, not parsed by golang by default
// so we need to do manual time parsing, by defining our own date struct
// encapsulating time.Time
type Date struct {
	// Time item containing the actual parsed time object
	time.Time
}

// UnmarshalJSON parses datetime strings returned by the transip api
func (tt *Time) UnmarshalJSON(input []byte) error {
	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		return err
	}
	newTime, err := time.ParseInLocation("\"2006-01-02 15:04:05\"", string(input), loc)
	if err != nil {
		return err
	}

	tt.Time = newTime
	return nil
}

// UnmarshalJSON parses date strings returned by the transip api
func (td *Date) UnmarshalJSON(input []byte) error {
	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		return err
	}
	newTime, err := time.ParseInLocation("\"2006-01-02\"", string(input), loc)
	if err != nil {
		return err
	}

	td.Time = newTime
	return nil
}

// ParseResponse will convert a RestResponse struct to the given interface
// on error it will pass this back
// when the rest response has no body it will return without filling the dest variable
// todo: look into specific types of errors
func (r *RestResponse) ParseResponse(dest interface{}) error {
	// do response error checking
	if !r.Method.StatusCodeIsCorrect(r.StatusCode) {
		// there is no response content so we also don't need to parse it
		if len(r.Body) == 0 {
			return fmt.Errorf("error response without body from api server status code '%d'", r.StatusCode)
		}

		var errorResponse RestException
		err := json.Unmarshal(r.Body, &errorResponse)
		if err != nil {
			// todo: look into error wrapping, nested errors
			// todo: add http status code
			return fmt.Errorf("response error could not be decoded, response = %s", string(r.Body))
		}

		return errors.New(errorResponse.Message)
	}

	if len(r.Body) == 0 {
		return nil
	}

	return json.Unmarshal(r.Body, &dest)
}
