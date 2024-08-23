package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestRepository_GetMailboxesByDomainName(t *testing.T) {
	const apiResponse = `{"mailboxes":[{"identifier":"test@example.com","localPart":"test","domain":"example.com","forwardTo":"test@example.com","availableDiskSpace":100,"usedDiskSpace":100,"status":"creating","isLocked":true,"imapServer":"imap.example.com","imapPort":123,"smtpServer":"smtp.example.com","smtpPort":123,"pop3Server":"pop3.example.com","pop3Port":123}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mailboxes", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mailboxes/test@example.com", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mailboxes", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mailboxes/test@example.com", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mailboxes/test@example.com", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DeleteMailbox("test@example.com")
	require.NoError(t, err)
}

func TestRepository_GetMailforwardsByDomainName(t *testing.T) {
	const apiResponse = `{"forwards":[{"id":1,"localPart":"test","domain":"example.com","forwardTo":"test@example.com","status":"created"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-forwards", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-forwards/1", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-forwards", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-forwards/1", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-forwards/1", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DeleteMailforward("example.com", 1)
	require.NoError(t, err)
}

func TestRepository_GetMaillistsByDomainName(t *testing.T) {
	const apiResponse = `{"lists":[{"id":1,"name":"test","emailAddress":"test@example.com"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-lists", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-lists/1", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-lists", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-lists/1", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
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
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-lists/1", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.DeleteMaillist("example.com", 1)
	require.NoError(t, err)
}

func TestRepository_LinkAddon(t *testing.T) {
	const expectedRequestBody = `{"action":"linkmailbox","addonId":7,"mailbox":"test@example.com"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-addons", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.LinkMailaddon(7, "test@example.com")
	require.NoError(t, err)
}

func TestRepository_LinkAddonInvalidMailbox(t *testing.T) {
	const expectedRequestBody = `{"action":"linkmailbox","addonId":7,"mailbox":"test@example.com"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-addons", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.LinkMailaddon(7, "testexample.com")
	require.Error(t, err)
	expectedErrorMessage := "invalid mailbox"
	assert.EqualErrorf(t, err, expectedErrorMessage, "Error should be: %v, got: %v", expectedErrorMessage, err)
}

func TestRepository_UnlinkAddon(t *testing.T) {
	const expectedRequestBody = `{"action":"unlinkmailbox","addonId":7,"mailbox":"test@example.com"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-addons", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.UnlinkMailaddon(7, "test@example.com")
	require.NoError(t, err)
}

func TestRepository_UnlinkAddonInvalidMailbox(t *testing.T) {
	const expectedRequestBody = `{"action":"unlinkmailbox","addonId":7,"mailbox":"test@example.com"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-addons", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.UnlinkMailaddon(7, "testexample.com")
	require.Error(t, err)
	expectedErrorMessage := "invalid mailbox"
	assert.EqualErrorf(t, err, expectedErrorMessage, "Error should be: %v, got: %v", expectedErrorMessage, err)
}

func TestRepository_GetMailAddonsByDomainName(t *testing.T) {
	const apiResponse = `{"addons": [{"id": 282154,"diskSpace": 1024,"mailboxes": 5,"linkedMailBox": "test@example.com","canBeLinked": false}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email/example.com/mail-addons", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAddonsByDomainName("example.com")
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, 282154, all[0].ID)
	assert.Equal(t, "test@example.com", all[0].LinkedMailBox)
	assert.Equal(t, false, all[0].CanBeLinked)
	assert.Equal(t, 5, all[0].Mailboxes)
	assert.Equal(t, 1024, all[0].DiskSpace)
}

func TestRepository_GetMailpackages(t *testing.T) {
	const apiResponse = `{"packages": [{"domain": "example.com", "status": "creating"},{"domain": "example2.com", "status": "created"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/email", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetMailpackages()
	require.NoError(t, err)
	require.Equal(t, 2, len(all))

	assert.Equal(t, "example.com", all[0].Domain)
	assert.Equal(t, "creating", all[0].Status)

	assert.Equal(t, "example2.com", all[1].Domain)
	assert.Equal(t, "created", all[1].Status)
}
