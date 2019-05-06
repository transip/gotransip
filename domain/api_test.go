package domain

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
)

func TestBatchCheckAvailability(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/batchcheckavailability.xml")
	assert.NoError(t, err)

	lst, err := BatchCheckAvailability(c, []string{"example.org", "example.com"})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []CheckResult{}, lst)
	assert.Equal(t, []Action{ActionRegister}, lst[0].Actions)
	assert.Equal(t, "example.com", lst[0].DomainName)
	assert.Equal(t, StatusFree, lst[0].Status)
	assert.Equal(t, "example.org", lst[1].DomainName)
}

func TestCheckAvailability(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/checkavailability.xml")
	assert.NoError(t, err)

	s, err := CheckAvailability(c, "example.org")
	assert.NoError(t, err)
	assert.Equal(t, StatusNotFree, s)
}

func TestGetWhois(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getwhois.xml")
	assert.NoError(t, err)

	s, err := GetWhois(c, "example.org")
	assert.NoError(t, err)
	assert.Equal(t, "Domain Name: EXAMPLE.ORG", s)
}

func TestGetDomainNames(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdomainnames.xml")
	assert.NoError(t, err)

	lst, err := GetDomainNames(c)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []string{}, lst)
	assert.Equal(t, "example.org", lst[0])
	assert.Equal(t, "example.com", lst[1])
}

func TestGetInfo(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getinfo.xml")
	assert.NoError(t, err)

	d, err := GetInfo(c, "example.org")
	assert.NoError(t, err)
	assert.IsType(t, Domain{}, d)
	assert.Equal(t, "Z290cmFuc2lwLXRlc3RpbmcK", d.AuthorizationCode)
	assert.Equal(t, "TransIP API", d.Branding.BannerLine1)
	assert.Equal(t, "gotransip", d.Branding.BannerLine2)
	assert.Equal(t, "https://www.transip.nl/api/", d.Branding.BannerLine3)
	assert.Equal(t, "TransIP", d.Branding.CompanyName)
	assert.Equal(t, "http://www.transip.nl/", d.Branding.CompanyURL)
	assert.Equal(t, "support@transip.nl", d.Branding.SupportEmail)
	assert.Equal(t, "http://www.transip.nl/tou", d.Branding.TermsOfUsageURL)
	assert.Equal(t, 2, len(d.Contacts))
	assert.Equal(t, "Leiden", d.Contacts[0].City)
	assert.Equal(t, "1234", d.Contacts[0].CompanyKvk)
	assert.Equal(t, "TransIP B.V.", d.Contacts[0].CompanyName)
	assert.Equal(t, "BV", d.Contacts[0].CompanyType)
	assert.Equal(t, "nl", d.Contacts[0].Country)
	assert.Equal(t, "support@transip.nl", d.Contacts[0].Email)
	assert.Equal(t, "+31 715241918", d.Contacts[0].FaxNumber)
	assert.Equal(t, "foo", d.Contacts[0].FirstName)
	assert.Equal(t, "baz", d.Contacts[0].LastName)
	assert.Equal(t, "bar", d.Contacts[0].MiddleName)
	assert.Equal(t, "9B", d.Contacts[0].Number)
	assert.Equal(t, "+31 715241919", d.Contacts[0].PhoneNumber)
	assert.Equal(t, "2316 XB", d.Contacts[0].PostalCode)
	assert.Equal(t, "Schipholweg", d.Contacts[0].Street)
	assert.Equal(t, "registrant", d.Contacts[0].Type)
	assert.Equal(t, "administrative", d.Contacts[1].Type)
	assert.IsType(t, []DNSEntry{}, d.DNSEntries)
	assert.Equal(t, 2, len(d.DNSEntries))
	assert.Equal(t, "1.2.3.4", d.DNSEntries[0].Content)
	assert.Equal(t, "@", d.DNSEntries[0].Name)
	assert.Equal(t, int64(86400), d.DNSEntries[0].TTL)
	assert.Equal(t, DNSEntryTypeA, d.DNSEntries[0].Type)
	assert.Equal(t, "www", d.DNSEntries[1].Name)
	assert.Equal(t, true, d.IsLocked)
	assert.Equal(t, "example.org", d.Name)
	assert.Equal(t, 2, len(d.Nameservers))
	assert.Equal(t, "ns0.transip.net", d.Nameservers[0].Hostname)
	assert.Equal(t, net.ParseIP("1.2.3.4"), d.Nameservers[0].IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), d.Nameservers[0].IPv6Address)
	x, _ := time.Parse("2006-01-02", "2017-12-28")
	assert.Equal(t, x, d.RegistrationDate.Time)
	x, _ = time.Parse("2006-01-02", "2018-12-28")
	assert.Equal(t, x, d.RenewalDate.Time)
}

func TestBatchGetInfo(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/batchgetinfo.xml")
	assert.NoError(t, err)

	lst, err := BatchGetInfo(c, []string{"example.org"})
	assert.NoError(t, err)
	assert.IsType(t, []Domain{}, lst)
	assert.Equal(t, 2, len(lst))
	d := lst[0]
	// tests copy/pasted from TestGetInfo
	assert.Equal(t, "Z290cmFuc2lwLXRlc3RpbmcK", d.AuthorizationCode)
	assert.Equal(t, "TransIP API", d.Branding.BannerLine1)
	assert.Equal(t, "gotransip", d.Branding.BannerLine2)
	assert.Equal(t, "https://www.transip.nl/api/", d.Branding.BannerLine3)
	assert.Equal(t, "TransIP", d.Branding.CompanyName)
	assert.Equal(t, "http://www.transip.nl/", d.Branding.CompanyURL)
	assert.Equal(t, "support@transip.nl", d.Branding.SupportEmail)
	assert.Equal(t, "http://www.transip.nl/tou", d.Branding.TermsOfUsageURL)
	assert.Equal(t, 2, len(d.Contacts))
	assert.Equal(t, "Leiden", d.Contacts[0].City)
	assert.Equal(t, "1234", d.Contacts[0].CompanyKvk)
	assert.Equal(t, "TransIP B.V.", d.Contacts[0].CompanyName)
	assert.Equal(t, "BV", d.Contacts[0].CompanyType)
	assert.Equal(t, "nl", d.Contacts[0].Country)
	assert.Equal(t, "support@transip.nl", d.Contacts[0].Email)
	assert.Equal(t, "+31 715241918", d.Contacts[0].FaxNumber)
	assert.Equal(t, "foo", d.Contacts[0].FirstName)
	assert.Equal(t, "baz", d.Contacts[0].LastName)
	assert.Equal(t, "bar", d.Contacts[0].MiddleName)
	assert.Equal(t, "9B", d.Contacts[0].Number)
	assert.Equal(t, "+31 715241919", d.Contacts[0].PhoneNumber)
	assert.Equal(t, "2316 XB", d.Contacts[0].PostalCode)
	assert.Equal(t, "Schipholweg", d.Contacts[0].Street)
	assert.Equal(t, "registrant", d.Contacts[0].Type)
	assert.Equal(t, "administrative", d.Contacts[1].Type)
	assert.IsType(t, []DNSEntry{}, d.DNSEntries)
	assert.Equal(t, 2, len(d.DNSEntries))
	assert.Equal(t, "1.2.3.4", d.DNSEntries[0].Content)
	assert.Equal(t, "@", d.DNSEntries[0].Name)
	assert.Equal(t, int64(86400), d.DNSEntries[0].TTL)
	assert.Equal(t, DNSEntryTypeA, d.DNSEntries[0].Type)
	assert.Equal(t, "www", d.DNSEntries[1].Name)
	assert.Equal(t, true, d.IsLocked)
	assert.Equal(t, "example.org", d.Name)
	assert.Equal(t, 2, len(d.Nameservers))
	assert.IsType(t, []Nameserver{}, d.Nameservers)
	assert.Equal(t, "ns0.transip.net", d.Nameservers[0].Hostname)
	assert.Equal(t, net.ParseIP("1.2.3.4"), d.Nameservers[0].IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), d.Nameservers[0].IPv6Address)
	x, _ := time.Parse("2006-01-02", "2017-12-28")
	assert.Equal(t, x, d.RegistrationDate.Time)
	x, _ = time.Parse("2006-01-02", "2018-12-28")
	assert.Equal(t, x, d.RenewalDate.Time)
	// /tests copy/pasted from TestGetInfo
	assert.Equal(t, "example.com", lst[1].Name)
}

func TestGetAuthCode(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getauthcode.xml")
	assert.NoError(t, err)

	a, err := GetAuthCode(c, "example.org")
	assert.NoError(t, err)
	assert.Equal(t, "Z290cmFuc2lwLXRlc3RpbmcK", a)
}

func TestGetIsLocked(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getislocked.xml")
	assert.NoError(t, err)

	l, err := GetIsLocked(c, "example.org")
	assert.NoError(t, err)
	assert.Equal(t, true, l)
}

func TestGetAllTLDInfos(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getalltldinfos.xml")
	assert.NoError(t, err)

	lst, err := GetAllTLDInfos(c)
	assert.NoError(t, err)
	assert.IsType(t, []TLD{}, lst)
	assert.Equal(t, 2, len(lst))
	// tests copy/pasted from TestGetTldInfo
	tld := lst[0]
	assert.Equal(t, int64(2), tld.CancelTimeFrame)
	assert.Equal(t, 2, len(tld.Capabilities))
	assert.IsType(t, []Capability{}, tld.Capabilities)
	assert.Equal(t, CapabilityCanRegister, tld.Capabilities[0])
	assert.Equal(t, CapabilityCanSetContacts, tld.Capabilities[1])
	assert.Equal(t, ".org", tld.Name)
	assert.Equal(t, float64(3.49), tld.Price)
	assert.Equal(t, int64(12), tld.RegistrationPeriodLength)
	assert.Equal(t, float64(7.49), tld.RenewalPrice)
	// /tests copy/pasted from TestGetTldInfo
	assert.Equal(t, ".com", lst[1].Name)
}

func TestGetTldInfo(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/gettldinfo.xml")
	assert.NoError(t, err)

	tld, err := GetTldInfo(c, ".org")
	assert.NoError(t, err)
	assert.IsType(t, TLD{}, tld)
	assert.Equal(t, int64(2), tld.CancelTimeFrame)
	assert.Equal(t, 2, len(tld.Capabilities))
	assert.IsType(t, []Capability{}, tld.Capabilities)
	assert.Equal(t, CapabilityCanRegister, tld.Capabilities[0])
	assert.Equal(t, CapabilityCanSetContacts, tld.Capabilities[1])
	assert.Equal(t, ".org", tld.Name)
	assert.Equal(t, float64(3.49), tld.Price)
	assert.Equal(t, int64(12), tld.RegistrationPeriodLength)
	assert.Equal(t, float64(7.49), tld.RenewalPrice)
}

func TestGetCurrentDomainAction(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getcurrentdomainaction.xml")
	assert.NoError(t, err)

	act, err := GetCurrentDomainAction(c, "example.org")
	assert.NoError(t, err)
	assert.IsType(t, ActionResult{}, act)
	assert.Equal(t, true, act.HasFailed)
	assert.Equal(t, "test message", act.Message)
	assert.Equal(t, "test", act.Name)
}

func TestRequestAuthCode(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/requestauthcode.xml")
	assert.NoError(t, err)

	code, err := RequestAuthCode(c, "example.org")
	assert.NoError(t, err)
	assert.IsType(t, "", code)
	assert.Equal(t, "are0AeThe7er1Uyoo1aifowoMilohnae", code)
}
