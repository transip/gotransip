package forwardingservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

const forwardingServiceJSON = `{
  "hostname": "example.com",
  "isBlocked": false,
  "isLocked": true,
  "forwardTo": "https://www.transip.nl/domeinnaam/doorsturen/",
  "forwardMethod": "iframe",
  "frameTitle": "This is a forwarding service",
  "frameIcon": "https://www.example.com/favicon.ico",
  "forwardEverything": true,
  "forwardSubdomains": false,
  "sslStatus": "active"
}`

func assertForwardingService(t *testing.T, fs ForwardingService) {
	assert.Equal(t, "example.com", fs.Hostname)
	assert.False(t, fs.IsBlocked)
	assert.True(t, fs.IsLocked)
	assert.Equal(t, "https://www.transip.nl/domeinnaam/doorsturen/", fs.ForwardTo)
	assert.Equal(t, ForwardMethodIframe, fs.ForwardMethod)
	assert.Equal(t, "This is a forwarding service", fs.FrameTitle)
	assert.Equal(t, "https://www.example.com/favicon.ico", fs.FrameIcon)
	assert.True(t, fs.ForwardEverything)
	assert.False(t, fs.ForwardSubdomains)
	assert.Equal(t, SSLStatusActive, fs.SSLStatus)
}

func TestRepository_GetAll(t *testing.T) {
	const apiResponse = `{ "forwardingServices": [ ` + forwardingServiceJSON + ` ] }`
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assertForwardingService(t, all[0])
}

func TestRepository_GetSelection(t *testing.T) {
	const apiResponse = `{ "forwardingServices": [ ` + forwardingServiceJSON + ` ] }`
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services?page=1&pageSize=25", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	selection, err := repo.GetSelection(1, 25)
	require.NoError(t, err)
	require.Equal(t, 1, len(selection))

	assertForwardingService(t, selection[0])
}

func TestRepository_GetByDomainName(t *testing.T) {
	const apiResponse = `{ "forwardingService": ` + forwardingServiceJSON + ` }`
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	fs, err := repo.GetByDomainName("example.com")
	require.NoError(t, err)

	assertForwardingService(t, fs)
}

func TestRepository_Update(t *testing.T) {
	const expectedRequest = `{"forwardingService":{"hostname":"example.com","forwardTo":"https://www.example.com","forwardMethod":"movedPermanently","frameTitle":"","frameIcon":"","forwardEverything":true,"forwardSubdomains":false,"sslStatus":"active"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Update(ForwardingService{
		Hostname:          "example.com",
		ForwardTo:         "https://www.example.com",
		ForwardMethod:     ForwardMethodMovedPermanently,
		ForwardEverything: true,
		SSLStatus:         SSLStatusActive,
	})
	require.NoError(t, err)
}

func TestRepository_GetSSLStatus(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com/ssl", ExpectedMethod: "GET", StatusCode: 200, Response: `{ "sslStatus": "pending" }`}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	status, err := repo.GetSSLStatus("example.com")
	require.NoError(t, err)

	assert.Equal(t, SSLStatusPending, status)
}

func TestRepository_EnableSSL(t *testing.T) {
	const expectedRequest = `{"action":"enable"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com/ssl", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	require.NoError(t, repo.EnableSSL("example.com"))
}

func TestRepository_DisableSSL(t *testing.T) {
	const expectedRequest = `{"action":"disable"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com/ssl", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	require.NoError(t, repo.DisableSSL("example.com"))
}

func TestRepository_GetEmailForwardTo(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com/email", ExpectedMethod: "GET", StatusCode: 200, Response: `{ "forwardTo": "john.doe@example.com" }`}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	forwardTo, err := repo.GetEmailForwardTo("example.com")
	require.NoError(t, err)

	assert.Equal(t, "john.doe@example.com", forwardTo)
}

func TestRepository_GetEmailForwardToWithoutForward(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com/email", ExpectedMethod: "GET", StatusCode: 200, Response: `{ "forwardTo": null }`}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	forwardTo, err := repo.GetEmailForwardTo("example.com")
	require.NoError(t, err)

	assert.Equal(t, "", forwardTo)
}

func TestRepository_SetEmailForwardTo(t *testing.T) {
	const expectedRequest = `{"forwardTo":"john.doe@example.com"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com/email", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	require.NoError(t, repo.SetEmailForwardTo("example.com", "john.doe@example.com"))
}

func TestRepository_RemoveEmailForwardTo(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/forwarding-services/example.com/email", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	require.NoError(t, repo.RemoveEmailForwardTo("example.com"))
}
