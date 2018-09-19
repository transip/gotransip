package webhosting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
)

func TestGetAvailablePackages(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailablepackages.xml")
	assert.NoError(t, err)

	lst, err := GetAvailablePackages(c)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
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
	assert.NoError(t, err)

	lst, err := GetAvailableUpgrades(c, "example.org")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
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
	assert.NoError(t, err)

	h, err := GetInfo(c, "example.org")
	assert.NoError(t, err)
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
	assert.Equal(t, 2, len(h.Database))
	assert.Equal(t, 100, h.Database[0].MaxDiskUsage)
	assert.Equal(t, "example_org_db", h.Database[0].Name)
	assert.Equal(t, "foobar", h.Database[0].Username)
	assert.Equal(t, "example.org", h.DomainName)
	assert.Equal(t, 2, len(h.MailBoxes))
	assert.Equal(t, "info@example.org", h.MailBoxes[0].Address)
	assert.Equal(t, true, h.MailBoxes[0].HasVacationReply)
	assert.Equal(t, 1, h.MailBoxes[0].MaxDiskUsage)
	assert.Equal(t, SpamCheckStrengthAverage, h.MailBoxes[0].SpamCheckerStrength)
	assert.Equal(t, "I'm on holiday", h.MailBoxes[0].VacationReplyMessage)
	assert.Equal(t, "Out of office", h.MailBoxes[0].VacationReplySubject)
	assert.Equal(t, 2, len(h.MailForward))
	assert.Equal(t, "foobar@example.org", h.MailForward[0].Name)
	assert.Equal(t, "info@example.org", h.MailForward[0].TargetAddress)
	assert.Equal(t, "barfoo@example.org", h.MailForward[1].Name)
	assert.Equal(t, "foobar@example.org", h.MailForward[1].TargetAddress)
	assert.Equal(t, 2, len(h.SubDomains))
	assert.IsType(t, []SubDomain{}, h.SubDomains)
	assert.Equal(t, "demo.example.org", h.SubDomains[0].Name)
	assert.Equal(t, "beta.example.org", h.SubDomains[1].Name)
}

func TestGetWebhostingDomainNames(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getwebhostingdomainnames.xml")
	assert.NoError(t, err)

	lst, err := GetWebhostingDomainNames(c)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.Equal(t, "example.org", lst[0])
	assert.Equal(t, "example.com", lst[1])
}
