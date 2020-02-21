package rest

type RestMethod struct {
	Method              string
	ExpectedStatusCodes []int
}

var (
	GetRestMethod    = RestMethod{Method: "GET", ExpectedStatusCodes: []int{200, 200}}
	PostRestMethod   = RestMethod{Method: "POST", ExpectedStatusCodes: []int{200, 201}}
	PutRestMethod    = RestMethod{Method: "PUT", ExpectedStatusCodes: []int{204}}
	DeleteRestMethod = RestMethod{Method: "DELETE", ExpectedStatusCodes: []int{204}}
)

// method used by the rest client to check if the given status code is correct
func (r *RestMethod) StatusCodeIsCorrect(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
