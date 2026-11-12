package domain

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/internal/testutil"
	"github.com/transip/gotransip/v6/rest"
)

const (
	verification404Response = `{ "error": "Verification with id '12345' not found" }`

	// the api does not return the contact and domain of a verification in a list
	verificationsAPIResponse = `{ "verifications": [ {
    "id": 33,
    "domainName": "example.com",
    "status": "pending",
    "createdAt": 1786111426,
    "updatedAt": 0,
    "responseReceivedAt": 0
  } ] }`
	verificationAPIResponse = `{
  "verification": {
    "id": 33,
    "domainName": "example.com",
    "status": "success",
    "createdAt": 1786111426,
    "updatedAt": 1786111500,
    "responseReceivedAt": 1786111499,
    "contact": {
      "type": "registrant",
      "firstName": "first",
      "middleName": "middle",
      "lastName": "last",
      "companyName": "company",
      "companyKvk": "",
      "companyType": "",
      "street": "street",
      "number": "1",
      "postalCode": "1234 AB",
      "city": "city",
      "phoneNumber": "+31 000000000",
      "faxNumber": "",
      "email": "firstlast@example.com",
      "country": "nl"
    },
    "domain": { "name": "example.com", "registrationStatus": "registered", "isBlocked": false }
  }
}`
)

func TestRepository_GetVerificationsSelectionLeavesOutContactAndDomain(t *testing.T) {
	response := `{ "verifications": [ {
    "id": 33,
    "domainName": "example.com",
    "status": "pending",
    "contact": { "type": "registrant", "firstName": "first" },
    "domain": { "name": "example.com" }
  } ] }`
	server := testutil.MockServer{T: t, ExpectedMethod: "GET", ExpectedURL: "/domain-verifications?page=1&pageSize=10", StatusCode: 200, Response: response}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	verifications, err := repo.GetVerificationsSelection(1, 10, "", nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(verifications))
	assert.Equal(t, VerificationContact{}, verifications[0].Contact)
	assert.Equal(t, VerificationDomain{}, verifications[0].Domain)
}

func TestRepository_GetVerificationsSelection(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedMethod: "GET", ExpectedURL: "/domain-verifications?page=2&pageSize=10", StatusCode: 200, Response: verificationsAPIResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	verifications, err := repo.GetVerificationsSelection(2, 10, "", nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(verifications))
	assert.Equal(t, int64(33), verifications[0].ID)
	assert.Equal(t, "example.com", verifications[0].DomainName)
	assert.Equal(t, VerificationStatusPending, verifications[0].Status)
	assert.Equal(t, int64(1786111426), verifications[0].CreatedAt)
	assert.Equal(t, int64(0), verifications[0].UpdatedAt)
	assert.Equal(t, int64(0), verifications[0].ResponseReceivedAt)
	// the list response has no contact and domain, so they stay empty
	assert.Equal(t, VerificationContact{}, verifications[0].Contact)
	assert.Equal(t, VerificationDomain{}, verifications[0].Domain)
}

func TestRepository_GetVerificationsSelectionWithFilters(t *testing.T) {
	expectedURL := "/domain-verifications?domainName=example.com&page=1&pageSize=25&statuses=pending%2Cfailed"
	server := testutil.MockServer{T: t, ExpectedMethod: "GET", ExpectedURL: expectedURL, StatusCode: 200, Response: verificationsAPIResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	verifications, err := repo.GetVerificationsSelection(1, 25, "example.com", []VerificationStatus{VerificationStatusPending, VerificationStatusFailed})
	require.NoError(t, err)
	assert.Equal(t, 1, len(verifications))
}

func TestRepository_GetVerificationsRetrievesAllPages(t *testing.T) {
	var requestedURLs []string
	httpServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestedURLs = append(requestedURLs, req.URL.String())

		// the first page is full, the second page is the last one
		amount := 100
		if req.URL.Query().Get("page") != "1" {
			amount = 5
		}

		items := make([]string, amount)
		for i := range items {
			items[i] = `{"id":` + strconv.Itoa(i+1) + `,"domainName":"example.com","status":"pending"}`
		}

		_, err := rw.Write([]byte(fmt.Sprintf(`{"verifications":[%s]}`, strings.Join(items, ","))))
		require.NoError(t, err)
	}))
	defer httpServer.Close()

	config := gotransip.DemoClientConfiguration
	config.URL = httpServer.URL
	client, err := gotransip.NewClient(config)
	require.NoError(t, err)
	repo := Repository{Client: client}

	verifications, err := repo.GetVerifications("example.com", nil)
	require.NoError(t, err)
	assert.Equal(t, 105, len(verifications))
	assert.Equal(t, []string{
		"/domain-verifications?domainName=example.com&page=1&pageSize=100",
		"/domain-verifications?domainName=example.com&page=2&pageSize=100",
	}, requestedURLs)
}

func TestRepository_GetVerificationByID(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedMethod: "GET", ExpectedURL: "/domain-verifications/33", StatusCode: 200, Response: verificationAPIResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	verification, err := repo.GetVerificationByID(33)
	require.NoError(t, err)
	assert.Equal(t, int64(33), verification.ID)
	assert.Equal(t, VerificationStatusSuccess, verification.Status)
	assert.Equal(t, int64(1786111500), verification.UpdatedAt)
	assert.Equal(t, int64(1786111499), verification.ResponseReceivedAt)

	assert.Equal(t, "registrant", verification.Contact.Type)
	assert.Equal(t, "first", verification.Contact.FirstName)
	assert.Equal(t, "middle", verification.Contact.MiddleName)
	assert.Equal(t, "last", verification.Contact.LastName)
	assert.Equal(t, "company", verification.Contact.CompanyName)
	assert.Equal(t, "nl", verification.Contact.Country)

	assert.Equal(t, VerificationDomain{Name: "example.com", RegistrationStatus: "registered", IsBlocked: false}, verification.Domain)
}

func TestRepository_GetVerificationByIDError(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedMethod: "GET", ExpectedURL: "/domain-verifications/12345", StatusCode: 404, Response: verification404Response}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	verification, err := repo.GetVerificationByID(12345)
	if assert.Errorf(t, err, "getverificationbyid server response error not returned") {
		require.Empty(t, verification.ID)
		assert.Equal(t, &rest.Error{Message: "Verification with id '12345' not found", StatusCode: 404}, err)
	}
}
