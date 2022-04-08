package email

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
)

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t               *testing.T
	expectedURL     string
	expectedMethod  string
	statusCode      int
	expectedRequest string
	response        string
	skipRequestBody bool
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedURL, req.URL.String()) // check if right expectedURL is called

		if m.skipRequestBody == false && req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := io.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequest, string(body))
		}

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequest expectedMethod is used
		rw.WriteHeader(m.statusCode)                    // respond with given status code

		if m.response != "" {
			_, err := rw.Write([]byte(m.response))
			require.NoError(m.t, err, "error when writing mock response")
		}
	}))
}

func (m *mockServer) getClient() (*repository.Client, func()) {
	httpServer := m.getHTTPServer()
	config := gotransip.DemoClientConfiguration
	config.URL = httpServer.URL
	client, err := gotransip.NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return &client, tearDown
}

func TestRepository_GetMailboxesByDomainName(t *testing.T) {
	const apiResponse = `{"mailboxes":[{"identifier":"test@example.com","localPart":"test","domain":"example.com","forwardTo":"test@example.com","availableDiskSpace":100,"usedDiskSpace":100,"status":"creating","isLocked":true,"imapServer":"imap.example.com","imapPort":123,"smtpServer":"smtp.example.com","smtpPort":123,"pop3Server":"pop3.example.com","pop3Port":123}]}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mailboxes", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetMailboxesByDomainName("example.com")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, "test@example.com", all[0].Identifier)
	assert.Equal(t, "test", all[0].LocalPart)
	assert.Equal(t, "example.com", all[0].Domain)
	assert.Equal(t, "test@example.com", all[0].ForwardTo)
	assert.Equal(t, 100, all[0].AvailableDiskSpace)
	assert.Equal(t, 100, all[0].UsedDiskSpace)
	assert.Equal(t, "creating", all[0].Status)
	assert.Equal(t, true, all[0].IsLocked)
	assert.Equal(t, "imap.example.com", all[0].ImapServer)
	assert.Equal(t, 123, all[0].ImapPort)
	assert.Equal(t, "smtp.example.com", all[0].SMTPServer)
	assert.Equal(t, 123, all[0].SMTPPort)
	assert.Equal(t, "pop3.example.com", all[0].Pop3Server)
	assert.Equal(t, 123, all[0].Pop3Port)
}

func TestRepository_GetMailboxByEmailAddress(t *testing.T) {
	const apiResponse = `{"mailbox":{"identifier":"test@example.com","localPart":"test","domain":"example.com","forwardTo":"test@example.com","availableDiskSpace":100,"usedDiskSpace":100,"status":"creating","isLocked":true,"imapServer":"imap.example.com","imapPort":123,"smtpServer":"smtp.example.com","smtpPort":123,"pop3Server":"pop3.example.com","pop3Port":123}}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mailboxes/test@example.com", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	mailbox, err := repo.GetMailboxByEmailAddress("test@example.com")
	require.NoError(t, err)

	assert.Equal(t, "test@example.com", mailbox.Identifier)
	assert.Equal(t, "test", mailbox.LocalPart)
	assert.Equal(t, "example.com", mailbox.Domain)
	assert.Equal(t, "test@example.com", mailbox.ForwardTo)
	assert.Equal(t, 100, mailbox.AvailableDiskSpace)
	assert.Equal(t, 100, mailbox.UsedDiskSpace)
	assert.Equal(t, "creating", mailbox.Status)
	assert.Equal(t, true, mailbox.IsLocked)
	assert.Equal(t, "imap.example.com", mailbox.ImapServer)
	assert.Equal(t, 123, mailbox.ImapPort)
	assert.Equal(t, "smtp.example.com", mailbox.SMTPServer)
	assert.Equal(t, 123, mailbox.SMTPPort)
	assert.Equal(t, "pop3.example.com", mailbox.Pop3Server)
	assert.Equal(t, 123, mailbox.Pop3Port)
}

func TestRepository_CreateMailbox(t *testing.T) {
	const expectedRequestBody = `{"localPart":"test","password":"Password123","maxDiskUsage":100,"forwardTo":"test@example.com"}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mailboxes", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	m := CreateMailboxRequest{
		LocalPart:    "test",
		Password:     "Password123",
		MaxDiskUsage: 100,
		ForwardTo:    "test@example.com",
	}

	err := repo.CreateMailbox("example.com", m)
	require.NoError(t, err)
}

func TestRepository_UpdateMailbox(t *testing.T) {
	const expectedRequestBody = `{"password":"Password123","maxDiskUsage":100,"forwardTo":"test@example.com"}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mailboxes/test@example.com", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	m := UpdateMailboxRequest{
		Password:     "Password123",
		MaxDiskUsage: 100,
		ForwardTo:    "test@example.com",
	}

	err := repo.UpdateMailbox("test@example.com", m)
	require.NoError(t, err)
}

func TestRepository_DeleteMailbox(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/email/example.com/mailboxes/test@example.com", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DeleteMailbox("test@example.com")
	require.NoError(t, err)
}

func TestRepository_GetMailforwardsByDomainName(t *testing.T) {
	const apiResponse = `{"forwards":[{"id":1,"localPart":"test","domain":"example.com","forwardTo":"test@example.com","status":"created"}]}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-forwards", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetMailforwardsByDomainName("example.com")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, 1, all[0].ID)
	assert.Equal(t, "test", all[0].LocalPart)
	assert.Equal(t, "example.com", all[0].Domain)
	assert.Equal(t, "test@example.com", all[0].ForwardTo)
	assert.Equal(t, "created", all[0].Status)
}

func TestRepository_GetMailforwardByDomainNameAndID(t *testing.T) {
	const apiResponse = `{"forward":{"id":1,"localPart":"test","domain":"example.com","forwardTo":"test@example.com","status":"created"}}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-forwards/1", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	mailforward, err := repo.GetMailforwardByDomainNameAndID("example.com", 1)
	require.NoError(t, err)

	assert.Equal(t, 1, mailforward.ID)
	assert.Equal(t, "test", mailforward.LocalPart)
	assert.Equal(t, "example.com", mailforward.Domain)
	assert.Equal(t, "test@example.com", mailforward.ForwardTo)
	assert.Equal(t, "created", mailforward.Status)
}

func TestRepository_CreateMailforward(t *testing.T) {
	const expectedRequestBody = `{"forwardTo":"test@example.com","localPart":"test"}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-forwards", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	m := CreateMailforwardRequest{
		ForwardTo: "test@example.com",
		LocalPart: "test",
	}

	err := repo.CreateMailforward("example.com", m)
	require.NoError(t, err)
}

func TestRepository_UpdateMailforward(t *testing.T) {
	const expectedRequestBody = `{"forwardTo":"test@example.com","localPart":"test"}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-forwards/1", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	m := UpdateMailforwardRequest{
		ForwardTo: "test@example.com",
		LocalPart: "test",
	}

	err := repo.UpdateMailforward("example.com", 1, m)
	require.NoError(t, err)
}

func TestRepository_DeleteMailforward(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-forwards/1", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DeleteMailforward("example.com", 1)
	require.NoError(t, err)
}

func TestRepository_GetMaillistsByDomainName(t *testing.T) {
	const apiResponse = `{"lists":[{"id":1,"name":"test","emailAddress":"test@example.com"}]}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-lists", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetMaillistsByDomainName("example.com")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, 1, all[0].ID)
	assert.Equal(t, "test", all[0].Name)
	assert.Equal(t, "test@example.com", all[0].EmailAddress)
}

func TestRepository_GetMaillistByDomainNameAndID(t *testing.T) {
	const apiResponse = `{"list":{"id":1,"name":"test","emailAddress":"test@example.com","entries":["test1@example.com","test2@example.com"]}}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-lists/1", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	maillist, err := repo.GetMaillistByDomainNameAndID("example.com", 1)
	require.NoError(t, err)

	assert.Equal(t, 1, maillist.ID)
	assert.Equal(t, "test", maillist.Name)
	assert.Equal(t, "test@example.com", maillist.EmailAddress)
}

func TestRepository_CreateMaillist(t *testing.T) {
	const expectedRequestBody = `{"emailAddress":"test@example.com","entries":["test1@example.com","test2@example.com"],"name":"test"}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-lists", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	m := CreateMaillistRequest{
		Name:         "test",
		EmailAddress: "test@example.com",
		Entries:      []string{"test1@example.com", "test2@example.com"},
	}

	err := repo.CreateMaillist("example.com", m)
	require.NoError(t, err)
}

func TestRepository_UpdateMaillist(t *testing.T) {
	const expectedRequestBody = `{"emailAddress":"test@example.com","entries":["test1@example.com","test2@example.com"]}`
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-lists/1", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	m := UpdateMaillistRequest{
		EmailAddress: "test@example.com",
		Entries:      []string{"test1@example.com", "test2@example.com"},
	}

	err := repo.UpdateMaillist("example.com", 1, m)
	require.NoError(t, err)
}

func TestRepository_DeleteMaillist(t *testing.T) {
	server := mockServer{t: t, expectedURL: "/email/example.com/mail-lists/1", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DeleteMaillist("example.com", 1)
	require.NoError(t, err)
}
