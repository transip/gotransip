package webhosting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v5"
)

func TestGetAvailablePackages(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailablepackages.xml")
	require.NoError(t, err)

	lst, err := GetAvailablePackages(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Package{}, lst)
	assert.Equal(t, "New Webhosting S", lst[0].Description)
	assert.Equal(t, "webhosting-small", lst[0].Name)
	assert.Equal(t, 3.0, lst[0].Price)
	assert.Equal(t, 5.0, lst[0].RenewalPrice)
	assert.Equal(t, "webhosting-large", lst[1].Name)
}

func TestGetAvailableUpgrades(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailableupgrades.xml")
	require.NoError(t, err)

	lst, err := GetAvailableUpgrades(c, "example.org")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Package{}, lst)
	assert.Equal(t, "New Webhosting XL", lst[0].Description)
	assert.Equal(t, "webhosting-extra-large", lst[0].Name)
	assert.Equal(t, 10.0, lst[0].Price)
	assert.Equal(t, 15.0, lst[0].RenewalPrice)
	assert.Equal(t, "webhosting-large", lst[1].Name)
}

func TestGetInfo(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getinfo.xml")
	require.NoError(t, err)

	h, err := GetInfo(c, "example.org")
	require.NoError(t, err)
	assert.IsType(t, Host{}, h)
	assert.Equal(t, 2, len(h.CronJobs))
	assert.Equal(t, "3", h.CronJobs[0].DayTrigger)
	assert.Equal(t, "info@example.org", h.CronJobs[0].Email)
	assert.Equal(t, "2", h.CronJobs[0].HourTrigger)
	assert.Equal(t, "1", h.CronJobs[0].MinuteTrigger)
	assert.Equal(t, "4", h.CronJobs[0].MonthTrigger)
	assert.Equal(t, "test", h.CronJobs[0].Name)
	assert.Equal(t, "http://example.org/?foobar", h.CronJobs[0].URL)
	assert.Equal(t, "*", h.CronJobs[0].WeekdayTrigger)
	require.Equal(t, 2, len(h.Database))
	assert.Equal(t, 100, h.Database[0].MaxDiskUsage)
	assert.Equal(t, "example_org_db", h.Database[0].Name)
	assert.Equal(t, "foobar", h.Database[0].Username)
	assert.Equal(t, "example_org_test", h.Database[1].Name)
	assert.Equal(t, "example.org", h.DomainName)
	require.Equal(t, 2, len(h.MailBoxes))
	assert.Equal(t, "info@example.org", h.MailBoxes[0].Address)
	assert.Equal(t, true, h.MailBoxes[0].HasVacationReply)
	assert.Equal(t, 1, h.MailBoxes[0].MaxDiskUsage)
	assert.Equal(t, SpamCheckStrengthAverage, h.MailBoxes[0].SpamCheckerStrength)
	assert.Equal(t, "I'm on holiday", h.MailBoxes[0].VacationReplyMessage)
	assert.Equal(t, "Out of office", h.MailBoxes[0].VacationReplySubject)
	assert.Equal(t, "support@example.org", h.MailBoxes[1].Address)
	require.Equal(t, 2, len(h.MailForward))
	assert.Equal(t, "foobar@example.org", h.MailForward[0].Name)
	assert.Equal(t, "info@example.org", h.MailForward[0].TargetAddress)
	assert.Equal(t, "barfoo@example.org", h.MailForward[1].Name)
	assert.Equal(t, "foobar@example.org", h.MailForward[1].TargetAddress)
	require.Equal(t, 2, len(h.SubDomains))
	assert.IsType(t, []SubDomain{}, h.SubDomains)
	assert.Equal(t, "demo.example.org", h.SubDomains[0].Name)
	assert.Equal(t, "beta.example.org", h.SubDomains[1].Name)
}

func TestGetWebhostingDomainNames(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getwebhostingdomainnames.xml")
	require.NoError(t, err)

	lst, err := GetWebhostingDomainNames(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.Equal(t, "example.org", lst[0])
	assert.Equal(t, "example.com", lst[1])
}

func TestCancel(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/cancel.xml")
	require.NoError(t, err)

	err = Cancel(c, "example.org", gotransip.CancellationTimeImmediately)
	require.NoError(t, err)
}

func TestCreateCronjob(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/createcronjob.xml")
	require.NoError(t, err)

	err = CreateCronjob(c, "example.org", CronJob{Name: "refresh-twitter"})
	require.NoError(t, err)
}

func TestCreateDatabase(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/createdatabase.xml")
	require.NoError(t, err)

	err = CreateDatabase(c, "example.org", Database{Name: "example_org_db"})
	require.NoError(t, err)
}

func TestCreateMailBox(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/createmailbox.xml")
	require.NoError(t, err)

	err = CreateMailBox(c, "example.org", MailBox{Address: "info@example.org"})
	require.NoError(t, err)
}

func TestCreateMailForward(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/createmailforward.xml")
	require.NoError(t, err)

	err = CreateMailForward(c, "example.org", MailForward{
		Name:          "info@example.org",
		TargetAddress: "devnull@example.org",
	})
	require.NoError(t, err)
}

func TestCreateSubdomain(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/createsubdomain.xml")
	require.NoError(t, err)

	err = CreateSubdomain(c, "example.org", SubDomain{Name: "ftp.example.org"})
	require.NoError(t, err)
}

func TestDeleteCronjob(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/deletecronjob.xml")
	require.NoError(t, err)

	err = DeleteCronjob(c, "example.org", CronJob{Name: "refresh-twitter"})
	require.NoError(t, err)
}

func TestDeleteDatabase(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/deletedatabase.xml")
	require.NoError(t, err)

	err = DeleteDatabase(c, "example.org", Database{Name: "example_org_db"})
	require.NoError(t, err)
}

func TestDeleteMailBox(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/deletemailbox.xml")
	require.NoError(t, err)

	err = DeleteMailBox(c, "example.org", MailBox{Address: "info@example.org"})
	require.NoError(t, err)
}

func TestDeleteMailForward(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/deletemailforward.xml")
	require.NoError(t, err)

	err = DeleteMailForward(c, "example.org", MailForward{Name: "info@example.org"})
	require.NoError(t, err)
}

func TestDeleteSubdomain(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/deletesubdomain.xml")
	require.NoError(t, err)

	err = DeleteSubdomain(c, "example.org", SubDomain{Name: "ftp.example.org"})
	require.NoError(t, err)
}

func TestModifyDatabase(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/modifydatabase.xml")
	require.NoError(t, err)

	err = ModifyDatabase(c, "example.org", Database{
		Name:         "example_org_db",
		MaxDiskUsage: 1000,
	})
	require.NoError(t, err)
}

func TestModifyMailBox(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/modifymailbox.xml")
	require.NoError(t, err)

	err = ModifyMailBox(c, "example.org", MailBox{
		Address:      "info@example.org",
		MaxDiskUsage: 1000,
	})
	require.NoError(t, err)
}

func TestModifyMailForward(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/modifymailforward.xml")
	require.NoError(t, err)

	err = ModifyMailForward(c, "example.org", MailForward{
		Name: "info@example.org", TargetAddress: "support@example.org",
	})
	require.NoError(t, err)
}

func TestOrder(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/order.xml")
	require.NoError(t, err)

	err = Order(c, "example.org", "webhosting-small")
	require.NoError(t, err)
}

func TestSetDatabasePassword(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setdatabasepassword.xml")
	require.NoError(t, err)

	err = SetDatabasePassword(c, "example.org", Database{Name: "example_org_db"}, "s3cr3t")
	require.NoError(t, err)
}

func TestSetFtpPassword(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setftppassword.xml")
	require.NoError(t, err)

	err = SetFtpPassword(c, "example.org", "s3cr3t")
	require.NoError(t, err)
}

func TestSetMailBoxPassword(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setmailboxpassword.xml")
	require.NoError(t, err)

	err = SetMailBoxPassword(c, "example.org", MailBox{Address: "info@example.org"}, "s3cr3t")
	require.NoError(t, err)
}

func TestUpgrade(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/upgrade.xml")
	require.NoError(t, err)

	err = Upgrade(c, "example.org", "webhosting-large")
	require.NoError(t, err)
}
