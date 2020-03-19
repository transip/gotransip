package test

// APITest struct contains a response from the server
type APITest struct {
	// response string of the api-test endpoint
	Response string `json:"ping"`
}
