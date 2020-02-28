package rest

import "github.com/transip/gotransip/v6/util"

// RestMethod is a struct which is used by the client to present a HTTP method of choice
// and the expected return status codes on which you can check to see if the response is correct
// thus not an error
type RestMethod struct {
	// Method is where a HTTP method like "GET" would go
	Method              string
	// ExpectedStatusCodes are the expected status codes with which
	// you can check if the response status code is correct
	ExpectedStatusCodes []int
}

var (
	// GetRestMethod is a wrapper with expected status codes around the HTTP "GET" method
	GetRestMethod    = RestMethod{Method: "GET", ExpectedStatusCodes: []int{200, 200}}
	// PostRestMethod is a wrapper with expected status codes around the HTTP "GET" method
	PostRestMethod   = RestMethod{Method: "POST", ExpectedStatusCodes: []int{200, 201}}
	// PutRestMethod is a wrapper with expected status codes around the HTTP "GET" method
	PutRestMethod    = RestMethod{Method: "PUT", ExpectedStatusCodes: []int{204}}
	// DeleteRestMethod is a wrapper with expected status codes around the HTTP "GET" method
	DeleteRestMethod = RestMethod{Method: "DELETE", ExpectedStatusCodes: []int{204}}
)

// StatusCodeIsCorrect returns true when the status code
// method used by the rest client to check if the given status code is correct
func (r *RestMethod) StatusCodeIsCorrect(statusCode int) bool {
	return util.Contains(r.ExpectedStatusCodes, statusCode)
}
