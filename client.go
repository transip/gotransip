package gotransip

import (
	"context"
	"errors"
	"fmt"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/jwt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"github.com/transip/gotransip/v6/rest/request"
	"github.com/transip/gotransip/v6/rest/response"
	"io/ioutil"
	"net/http"
)

// client manages communication with the TransIP API
// In most cases there should be only one, shared, client.
type client struct {
	// client configuration file, allows you to:
	// - setting a custom useragent
	// - enable test mode
	// - use the demo token
	// - enable debugging
	config ClientConfiguration
	// provides you the possibility to specify timeouts
	context context.Context
	// authenticator wraps all authentication logic
	// - checking if the token is not expired yet
	// - creating an authentication request
	// - requesting and setting a new token
	authenticator authenticator.Authenticator
}

// NewClient creates a new API client.
// optionally you could put a custom http.client in the configuration struct
// to allow for advanced features such as caching.
func NewClient(config ClientConfiguration) (repository.Client, error) {
	return newClient(config)
}

// newClient method is used internally for testing,
// the NewClient method is exported as it follows the repository.Client interface
// which is so that we don't have to bind to this specific implementation
// todo: change to pointer
func newClient(config ClientConfiguration) (*client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}
	var privateKeyBody []byte
	var token jwt.Token

	// copy demo token from authenticator on demo mode
	if config.DemoMode {
		config.Token = authenticator.DemoToken
	}

	// check account name
	if len(config.AccountName) == 0 && len(config.Token) == 0 {
		return &client{}, errors.New("AccountName is required")
	}

	// check if token or private key is set
	if len(config.Token) == 0 && config.PrivateKeyReader == nil {
		return &client{}, errors.New("PrivateKeyReader, token or PrivateKeyReader is required")
	}

	if config.PrivateKeyReader != nil {
		var err error
		privateKeyBody, err = ioutil.ReadAll(config.PrivateKeyReader)

		if err != nil {
			return &client{}, fmt.Errorf("error while reading private key: %s", err.Error())
		}
	}

	if len(config.Token) > 0 {
		var err error
		token, err = jwt.New(config.Token)

		if err != nil {
			return &client{}, err
		}
	}

	// default to APIMode read/write
	if len(config.Mode) == 0 {
		config.Mode = APIModeReadWrite
	}

	// set basePath by default
	if len(config.URL) == 0 {
		config.URL = basePath
	}

	return &client{
		authenticator: &authenticator.TransipAuthenticator{PrivateKeyBody: privateKeyBody, Token: token, HTTPClient: config.HTTPClient},
		config:        config,
	}, nil
}

// This method is used by all rest client methods, thus: 'get','post','put','delete'
// It uses the authenticator to get a token, either statically provided by the user or requested from the authentication server
// Then decodes the json response to a supplied interface
func (c *client) call(method rest.RestMethod, request request.RestRequest, result interface{}) error {
	token, err := c.authenticator.GetToken()
	if err != nil {
		return fmt.Errorf("could not get token from authenticator %s", err.Error())
	}
	httpRequest, err := request.GetHTTPRequest(c.config.URL, method.Method)
	if err != nil {
		return fmt.Errorf("error during request creation %s", err.Error())
	}

	httpRequest.Header.Add("Authorization", token.GetAuthenticationHeaderValue())
	client := c.config.HTTPClient
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("request error:\n%s", err.Error())
	}

	defer httpResponse.Body.Close()

	// read entire httpResponse body
	b, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	restResponse := response.RestResponse{
		Body:       b,
		StatusCode: httpResponse.StatusCode,
		Method:     method,
	}

	return restResponse.ParseResponse(&result)
}

// ChangeBasePath changes base path to allow switching to mocks
func (c *client) ChangeBasePath(path string) {
	c.config.URL = path
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *client) GetConfig() ClientConfiguration {
	return c.config
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *client) GetAuthenticator() authenticator.Authenticator {
	return c.authenticator
}

// This method that executes a http Get request
func (c *client) Get(request request.RestRequest, responseObject interface{}) error {
	return c.call(rest.GetRestMethod, request, responseObject)
}

// This method that executes a http Post request
// It expects no response, that is why it does not return one
func (c *client) Post(request request.RestRequest) error {
	return c.call(rest.PostRestMethod, request, nil)
}

// This method that executes a http Put request
// It expects no response, that is why it does not return one
func (c *client) Put(request request.RestRequest) error {
	return c.call(rest.PutRestMethod, request, nil)
}

// This method that executes a http Delete request
// It expects no response, that is why it does not return one
func (c *client) Delete(request request.RestRequest) error {
	return c.call(rest.DeleteRestMethod, request, nil)
}
