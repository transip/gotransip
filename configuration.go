package gotransip

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

const (
	libraryVersion = "6.0.0"
	basePath       = "https://api.transip.nl/v6"
	userAgent      = "go-client-gotransip/" + libraryVersion
)

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// APIMode specifies in which mode the API is used. Currently this is only
// supports either readonly or readwrite
type APIMode string

var (
	// APIModeReadOnly specifies that no changes can be made from API calls
	APIModeReadOnly APIMode = "readonly"
	// APIModeReadWrite specifies that changes can be made from API calls
	APIModeReadWrite APIMode = "readwrite"
)

// ClientConfiguration stores the configuration of the API client
type ClientConfiguration struct {
	AccountName    string
	URL            string
	PrivateKeyPath string
	Token          string
	DemoMode       bool
	TestMode       bool
	HTTPClient     *http.Client
	Mode           APIMode
}
