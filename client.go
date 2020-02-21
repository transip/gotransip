package gotransip

import (
	"context"
	"errors"
	"fmt"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/jwt"
	"github.com/transip/gotransip/v6/rest"
	"github.com/transip/gotransip/v6/rest/request"
	"github.com/transip/gotransip/v6/rest/response"
	"io/ioutil"
	"net/http"
	"os"
)

// Client manages communication with the TransIP API
// In most cases there should be only one, shared, Client.
type Client struct {
	// Client configuration file, allows you to:
	// - setting a custom useragent
	// - enable test mode
	// - use the demo token
	// - enable debugging
	config ClientConfiguration
	// provides you the possibility to specify timeouts
	context context.Context
	// authenticator wraps all authentication logic
	authenticator authenticator.Authenticator
}

type Service struct {
	Client *Client
}

// NewClient creates a new API client.
// optionally you could put a custom http.Client in the configuration struct
// to allow for advanced features such as caching.
func NewClient(config ClientConfiguration) (Client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}
	var privateKeyBody []byte
	var token jwt.Token

	// check account name
	if len(config.AccountName) == 0 && len(config.Token) == 0 {
		return Client{}, errors.New("AccountName is required")
	}

	// check if token or private key is set
	if len(config.Token) == 0 && len(config.PrivateKeyPath) == 0 {
		return Client{}, errors.New("PrivateKeyPath, Token or PrivateKeyBody is required")
	}

	// if PrivateKeyPath was set, this will override any given PrivateKeyBody
	if len(config.PrivateKeyPath) > 0 {
		// try to open private key and read contents
		if _, err := os.Stat(config.PrivateKeyPath); err != nil {
			return Client{}, fmt.Errorf("Could not open private key: %s", err.Error())
		}

		// read private key so we can pass the body to the soapClient
		var err error
		privateKeyBody, err = ioutil.ReadFile(config.PrivateKeyPath)
		if err != nil {
			return Client{}, err
		}
	}

	if len(config.Token) > 0 {
		var err error
		token, err = jwt.New(config.Token)

		if err != nil {
			return Client{}, err
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

	return Client{
		authenticator: authenticator.Authenticator{PrivateKeyBody: privateKeyBody, Token: token, HTTPClient: config.HTTPClient},
		config:        config,
	}, nil
}

// This method is used by all rest client methods, thus: 'get','post','put','delete'
// It uses the authenticator to get a token, either statically provided by the user or requested from the authentication server
// Then decodes the json response to a supplied interface
func (c *Client) call(method rest.RestMethod, request request.RestRequest, result interface{}) error {
	token, err := c.authenticator.GetToken()
	if err != nil {
		return fmt.Errorf("could not get token from authenticator %s", err.Error())
	}
	httpRequest, err := request.GetHttpRequest(c.config.URL, method.Method)
	if err != nil {
		return fmt.Errorf("error during request creation %s", err.Error())
	}

	httpRequest.Header.Add("Authorization", token.GetAuthenticationHeaderValue())
	client := c.config.HTTPClient
	fmt.Println(httpRequest.URL)
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
func (c *Client) ChangeBasePath(path string) {
	c.config.URL = path
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *Client) GetConfig() ClientConfiguration {
	return c.config
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *Client) GetAuthenticator() authenticator.Authenticator {
	return c.authenticator
}

// This method that executes a http Get request
func (c *Client) Get(request request.RestRequest, response interface{}) error {
	return c.call(rest.GetRestMethod, request, response)
}
