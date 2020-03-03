package domain

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	error404Response   = `{ "error": "Domain with name 'example2.com' not found" }`
	domainsAPIResponse = `{ "domains": [
    {
      "name": "example.com",
      "authCode": "kJqfuOXNOYQKqh/jO4bYSn54YDqgAt1ksCe+ZG4Ud",
      "isTransferLocked": false,
      "registrationDate": "2016-01-01",
      "renewalDate": "2020-01-01",
      "isWhitelabel": false,
      "cancellationDate": "2020-01-01 12:00:00",
      "cancellationStatus": "signed",
      "isDnsOnly": false,
      "tags": [ "customTag", "anotherTag" ]
    }
  ] }`
	domainAPIResponse = `{ "domain": {
    "name": "example.com",
    "authCode": "kJqfuOXNOYQKqh/jO4bYSn54YDqgAt1ksCe+ZG4Ud",
    "isTransferLocked": false,
    "registrationDate": "2016-01-01",
    "renewalDate": "2020-01-01",
    "isWhitelabel": false,
    "cancellationDate": "2020-01-01 12:00:00",
    "cancellationStatus": "signed",
    "isDnsOnly": false,
    "tags": [ "customTag", "anotherTag" ]
  } } `
	brandingAPIResponse = `{
		"branding": {
		"companyName": "Example B.V.",
		"supportEmail": "admin@example.com",
		"companyUrl": "www.example.com",
		"termsOfUsageUrl": "www.example.com/tou",
		"bannerLine1": "Example B.V.",
		"bannerLine2": "Example",
		"bannerLine3": "http://www.example.com/products"
	} }`
	contactsAPIResponse = `{ "contacts": [ {
      "type": "registrant",
      "firstName": "John",
      "lastName": "Doe",
      "companyName": "Example B.V.",
      "companyKvk": "83057825",
      "companyType": "BV",
      "street": "Easy street",
      "number": "12",
      "postalCode": "1337 XD",
      "city": "Leiden",
      "phoneNumber": "+31 715241919",
      "faxNumber": "+31 715241919",
      "email": "example@example.com",
      "country": "nl"
    } ] }`
)

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t                   *testing.T
	expectedUrl         string
	expectedMethod      string
	statusCode          int
	expectedRequestBody string
	response            string
	skipRequestBody     bool
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedUrl, req.URL.String()) // check if right expectedUrl is called

		if m.skipRequestBody == false && req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := ioutil.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequestBody, string(body))
		}

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequestBody expectedMethod is used
		rw.WriteHeader(m.statusCode)                    // respond with given status code

		if m.response != "" {
			rw.Write([]byte(m.response))
		}
	}))
}

func (m *mockServer) getClient() (*repository.Client, func()) {
	httpServer := m.getHTTPServer()
	config := gotransip.ClientConfiguration{DemoMode: true, URL: httpServer.URL}
	client, err := gotransip.NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return &client, tearDown
}

func TestRepository_GetAll(t *testing.T) {
	server := mockServer{t: t, expectedMethod: "GET", expectedUrl: "/domains", statusCode: 200, response: domainsAPIResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))
	assert.Equal(t, "example.com", all[0].Name)
	assert.Equal(t, "kJqfuOXNOYQKqh/jO4bYSn54YDqgAt1ksCe+ZG4Ud", all[0].AuthCode)
	assert.Equal(t, false, all[0].IsTransferLocked)
	assert.Equal(t, "2016-01-01 00:00:00 +0100 CET", all[0].RegistrationDate.String())
	assert.Equal(t, "2020-01-01 00:00:00 +0100 CET", all[0].RenewalDate.String())
	assert.Equal(t, false, all[0].IsWhitelabel)
	assert.Equal(t, "2020-01-01 12:00:00 +0100 CET", all[0].CancellationDate.String())
	assert.Equal(t, "signed", all[0].CancellationStatus)
	assert.Equal(t, false, all[0].IsDnsOnly)
	assert.Equal(t, []string{"customTag", "anotherTag"}, all[0].Tags)
}

func TestRepository_GetByDomainName(t *testing.T) {
	server := mockServer{t: t, expectedMethod: "GET", expectedUrl: "/domains/example.com", statusCode: 200, response: domainAPIResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	domain, err := repo.GetByDomainName("example.com")
	require.NoError(t, err)
	assert.Equal(t, "example.com", domain.Name)
	assert.Equal(t, "kJqfuOXNOYQKqh/jO4bYSn54YDqgAt1ksCe+ZG4Ud", domain.AuthCode)
	assert.Equal(t, false, domain.IsTransferLocked)
	assert.Equal(t, "2016-01-01 00:00:00 +0100 CET", domain.RegistrationDate.String())
	assert.Equal(t, "2020-01-01 00:00:00 +0100 CET", domain.RenewalDate.String())
	assert.Equal(t, false, domain.IsWhitelabel)
	assert.Equal(t, "2020-01-01 12:00:00 +0100 CET", domain.CancellationDate.String())
	assert.Equal(t, "signed", domain.CancellationStatus)
	assert.Equal(t, false, domain.IsDnsOnly)
	assert.Equal(t, []string{"customTag", "anotherTag"}, domain.Tags)
}

func TestRepository_GetByDomainNameError(t *testing.T) {
	domainName := "example2.com"
	server := mockServer{t: t, expectedMethod: "GET", expectedUrl: fmt.Sprintf("/domains/%s", domainName), statusCode: 404, response: error404Response}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	domain, err := repo.GetByDomainName(domainName)
	require.Error(t, err)
	require.Empty(t, domain.Name)
	assert.Equal(t, errors.New("Domain with name 'example2.com' not found"), err)
}

func TestRepository_Register(t *testing.T) {
	expectedRequest := `{"domainName":"example.com"}`
	server := mockServer{t: t, expectedMethod: "POST", expectedUrl: "/domains", statusCode: 201, expectedRequestBody: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	register := Register{DomainName: "example.com"}
	err := repo.Register(register)
	require.NoError(t, err)
}

func TestRepository_RegisterError(t *testing.T) {
	errorResponse := `{"error":"The domain 'example.com' is not free and thus cannot be registered"}`
	server := mockServer{t: t, expectedMethod: "POST", expectedUrl: "/domains", statusCode: 406, skipRequestBody: true, response: errorResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	register := Register{DomainName: "example.com"}
	err := repo.Register(register)
	require.Error(t, err)
	assert.Error(t, errors.New("The domain 'example.com' is not free and thus cannot be registered"), err)
}

func TestRepository_Transfer(t *testing.T) {
	expectedRequest := `{"domainName":"example.com","authCode":"test123"}`
	server := mockServer{t: t, expectedMethod: "POST", expectedUrl: "/domains", statusCode: 201, expectedRequestBody: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	transfer := Transfer{DomainName: "example.com", AuthCode: "test123"}

	err := repo.Transfer(transfer)
	require.NoError(t, err)
}

func TestRepository_TransferError(t *testing.T) {
	errorResponse := `{"error":"The domain 'example.com' is not registered and thus cannot be transferred"}`
	server := mockServer{t: t, expectedMethod: "POST", expectedUrl: "/domains", statusCode: 409, skipRequestBody: true, response: errorResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	transfer := Transfer{DomainName: "example.com", AuthCode: "test123"}
	err := repo.Transfer(transfer)
	require.Error(t, err)
	assert.Error(t, errors.New("The domain 'example.com' is not registered and thus cannot be transferred"), err)
}

func TestRepository_Update(t *testing.T) {
	expectedRequest := `{"domain":{"tags":["test123","test1234"],"cancellationDate":"0001-01-01T00:00:00Z","isTransferLocked":false,"isWhitelabel":false,"name":"example.com","registrationDate":"0001-01-01T00:00:00Z","renewalDate":"0001-01-01T00:00:00Z"}}`
	server := mockServer{t: t, expectedMethod: "PUT", expectedUrl: "/domains/example.com", statusCode: 204, expectedRequestBody: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	domain := Domain{Tags: []string{"test123", "test1234"}, IsTransferLocked: false, IsWhitelabel: false, Name: "example.com"}

	err := repo.Update(domain)
	require.NoError(t, err)
}

func TestRepository_CancelEnd(t *testing.T) {
	expectedRequest := `{"endTime":"end"}`
	server := mockServer{t: t, expectedMethod: "DELETE", expectedUrl: "/domains/example.com", statusCode: 204, expectedRequestBody: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Cancel("example.com", gotransip.CancellationTimeEnd)
	require.NoError(t, err)
}

func TestRepository_Cancel(t *testing.T) {
	expectedRequest := `{"endTime":"immediately"}`
	server := mockServer{t: t, expectedMethod: "DELETE", expectedUrl: "/domains/example.com", statusCode: 204, expectedRequestBody: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Cancel("example.com", gotransip.CancellationTimeImmediately)
	require.NoError(t, err)
}

func TestRepository_GetDomainBranding(t *testing.T) {
	domainName := "example.com"
	server := mockServer{t: t, expectedMethod: "GET", expectedUrl: fmt.Sprintf("/domains/%s/branding", domainName), statusCode: 200, response: brandingAPIResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	branding, err := repo.GetDomainBranding(domainName)
	require.NoError(t, err)

	assert.Equal(t, "Example B.V.", branding.CompanyName)
	assert.Equal(t, "admin@example.com", branding.SupportEmail)
	assert.Equal(t, "www.example.com", branding.CompanyUrl)
	assert.Equal(t, "www.example.com/tou", branding.TermsOfUsageUrl)
	assert.Equal(t, "Example B.V.", branding.BannerLine1)
	assert.Equal(t, "Example", branding.BannerLine2)
	assert.Equal(t, "http://www.example.com/products", branding.BannerLine3)
}

func TestRepository_UpdateDomainBranding(t *testing.T) {
	expectedRequest := `{"branding":{"bannerLine1":"Example B.V.","bannerLine2":"admin@example.com","bannerLine3":"www.example.com","companyName":"www.example.com/tou","companyUrl":"Example B.V.","supportEmail":"Example","termsOfUsageUrl":"http://www.example.com/products"}}`
	server := mockServer{t: t, expectedMethod: "PUT", expectedUrl: "/domains/example.com/branding", statusCode: 204, expectedRequestBody: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	branding := Branding{
		BannerLine1:     "Example B.V.",
		BannerLine2:     "admin@example.com",
		BannerLine3:     "www.example.com",
		CompanyName:     "www.example.com/tou",
		CompanyUrl:      "Example B.V.",
		SupportEmail:    "Example",
		TermsOfUsageUrl: "http://www.example.com/products",
	}

	err := repo.UpdateDomainBranding("example.com", branding)
	require.NoError(t, err)
}

func TestRepository_GetContacts(t *testing.T) {
	domainName := "example.com"
	server := mockServer{t: t, expectedMethod: "GET", expectedUrl: fmt.Sprintf("/domains/%s/contacts", domainName), statusCode: 200, response: contactsAPIResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	contacts, err := repo.GetContacts(domainName)
	require.NoError(t, err)
	require.Equal(t, 1, len(contacts))

	assert.Equal(t, "registrant", contacts[0].Type)
	assert.Equal(t, "John", contacts[0].FirstName)
	assert.Equal(t, "Doe", contacts[0].LastName)
	assert.Equal(t, "Example B.V.", contacts[0].CompanyName)
	assert.Equal(t, "83057825", contacts[0].CompanyKvk)
	assert.Equal(t, "BV", contacts[0].CompanyType)
	assert.Equal(t, "Easy street", contacts[0].Street)
	assert.Equal(t, "12", contacts[0].Number)
	assert.Equal(t, "1337 XD", contacts[0].PostalCode)
	assert.Equal(t, "Leiden", contacts[0].City)
	assert.Equal(t, "+31 715241919", contacts[0].PhoneNumber)
	assert.Equal(t, "+31 715241919", contacts[0].FaxNumber)
	assert.Equal(t, "example@example.com", contacts[0].Email)
	assert.Equal(t, "nl", contacts[0].Country)
}

func TestRepository_UpdateContacts(t *testing.T) {
	expectedRequest := `{"contacts":[{"city":"Leiden","companyKvk":"83057825","companyName":"Example B.V.","companyType":"BV","country":"nl","email":"example@example.com","faxNumber":"+31 715241919","firstName":"John","lastName":"Doe","number":"12","phoneNumber":"+31 715241919","postalCode":"1337 XD","street":"Easy street","type":"registrant"}]}`
	server := mockServer{t: t, expectedMethod: "PUT", expectedUrl: "/domains/example.com/contacts", statusCode: 204, expectedRequestBody: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	contacts := []WhoisContact{
		{
			Type:        "registrant",
			FirstName:   "John",
			LastName:    "Doe",
			CompanyName: "Example B.V.",
			CompanyKvk:  "83057825",
			CompanyType: "BV",
			Street:      "Easy street",
			Number:      "12",
			PostalCode:  "1337 XD",
			City:        "Leiden",
			PhoneNumber: "+31 715241919",
			FaxNumber:   "+31 715241919",
			Email:       "example@example.com",
			Country:     "nl",
		},
	}

	err := repo.UpdateContacts("example.com", contacts)
	require.NoError(t, err)
}
