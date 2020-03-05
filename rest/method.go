package rest

// RestMethod is a struct which is used by the client to present a HTTP method of choice
// and the expected return status codes on which you can check to see if the response is correct
// thus not an error
type RestMethod struct {
	// Method is where a HTTP method like "GET" would go
	Method string
	// ExpectedStatusCodes are the expected status codes with which
	// you can check if the response status code is correct
	ExpectedStatusCodes []int
}

var (
	// GetRestMethod is a wrapper with expected status codes around the HTTP "GET" method
	GetRestMethod = RestMethod{Method: "GET", ExpectedStatusCodes: []int{200, 200}}
	// PostRestMethod is a wrapper with expected status codes around the HTTP "POST" method
	PostRestMethod = RestMethod{Method: "POST", ExpectedStatusCodes: []int{200, 201}}
	// PutRestMethod is a wrapper with expected status codes around the HTTP "PUT" method
	PutRestMethod = RestMethod{Method: "PUT", ExpectedStatusCodes: []int{204}}
	// PutRestMethod is a wrapper with expected status codes around the HTTP "PATCH" method
	PatchRestMethod = RestMethod{Method: "PATCH", ExpectedStatusCodes: []int{204}}
	// DeleteRestMethod is a wrapper with expected status codes around the HTTP "DELETE" method
	DeleteRestMethod = RestMethod{Method: "DELETE", ExpectedStatusCodes: []int{204}}
)

// StatusCodeIsCorrect returns true when the status code
// method used by the rest client to check if the given status code is correct
func (r *RestMethod) StatusCodeIsCorrect(statusCode int) bool {
	return contains(r.ExpectedStatusCodes, statusCode)
}

// Contains is a case insenstive match, finding needle in a haystack
func contains(haystack []int, needle int) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}
