package domain

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v5"
)

func TestBatchCheckAvailability(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/batchcheckavailability.xml")
	require.NoError(t, err)

	lst, err := BatchCheckAvailability(c, []string{"example.org", "example.com"})
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
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
	require.NoError(t, err)

	s, err := CheckAvailability(c, "example.org")
	require.NoError(t, err)
	assert.Equal(t, StatusNotFree, s)
}

func TestGetWhois(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getwhois.xml")
	require.NoError(t, err)

	s, err := GetWhois(c, "example.org")
	require.NoError(t, err)
	assert.Equal(t, "Domain Name: EXAMPLE.ORG", s)
}

func TestGetDomainNames(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdomainnames.xml")
	require.NoError(t, err)

	lst, err := GetDomainNames(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []string{}, lst)
	assert.Equal(t, "example.org", lst[0])
	assert.Equal(t, "example.com", lst[1])
}

func TestGetInfo(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getinfo.xml")
	require.NoError(t, err)

	d, err := GetInfo(c, "example.org")
	require.NoError(t, err)
	assert.IsType(t, Domain{}, d)
	assert.Equal(t, "Z290cmFuc2lwLXRlc3RpbmcK", d.AuthorizationCode)
	assert.Equal(t, "TransIP API", d.Branding.BannerLine1)
	assert.Equal(t, "gotransip", d.Branding.BannerLine2)
	assert.Equal(t, "https://www.transip.nl/api/", d.Branding.BannerLine3)
	assert.Equal(t, "TransIP", d.Branding.CompanyName)
	assert.Equal(t, "http://www.transip.nl/", d.Branding.CompanyURL)
	assert.Equal(t, "support@transip.nl", d.Branding.SupportEmail)
	assert.Equal(t, "http://www.transip.nl/tou", d.Branding.TermsOfUsageURL)
	require.Equal(t, 2, len(d.Contacts))
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
	require.Equal(t, 2, len(d.DNSEntries))
	assert.Equal(t, "1.2.3.4", d.DNSEntries[0].Content)
	assert.Equal(t, "@", d.DNSEntries[0].Name)
	assert.Equal(t, int64(86400), d.DNSEntries[0].TTL)
	assert.Equal(t, DNSEntryTypeA, d.DNSEntries[0].Type)
	assert.Equal(t, "www", d.DNSEntries[1].Name)
	assert.Equal(t, true, d.IsLocked)
	assert.Equal(t, "example.org", d.Name)
	require.Equal(t, 2, len(d.Nameservers))
	assert.Equal(t, "ns0.transip.net", d.Nameservers[0].Hostname)
	assert.Equal(t, net.ParseIP("1.2.3.4"), d.Nameservers[0].IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), d.Nameservers[0].IPv6Address)
	assert.Equal(t, "ns1.transip.nl", d.Nameservers[1].Hostname)
	x, _ := time.Parse("2006-01-02", "2017-12-28")
	assert.Equal(t, x, d.RegistrationDate.Time)
	x, _ = time.Parse("2006-01-02", "2018-12-28")
	assert.Equal(t, x, d.RenewalDate.Time)
}

func TestBatchGetInfo(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/batchgetinfo.xml")
	require.NoError(t, err)

	lst, err := BatchGetInfo(c, []string{"example.org"})
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Domain{}, lst)
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
	require.Equal(t, 2, len(d.Contacts))
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
	require.Equal(t, 2, len(d.DNSEntries))
	assert.Equal(t, "1.2.3.4", d.DNSEntries[0].Content)
	assert.Equal(t, "@", d.DNSEntries[0].Name)
	assert.Equal(t, int64(86400), d.DNSEntries[0].TTL)
	assert.Equal(t, DNSEntryTypeA, d.DNSEntries[0].Type)
	assert.Equal(t, "www", d.DNSEntries[1].Name)
	assert.Equal(t, true, d.IsLocked)
	assert.Equal(t, "example.org", d.Name)
	require.Equal(t, 2, len(d.Nameservers))
	assert.IsType(t, []Nameserver{}, d.Nameservers)
	assert.Equal(t, "ns0.transip.net", d.Nameservers[0].Hostname)
	assert.Equal(t, net.ParseIP("1.2.3.4"), d.Nameservers[0].IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), d.Nameservers[0].IPv6Address)
	assert.Equal(t, "ns1.transip.nl", d.Nameservers[1].Hostname)
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
	require.NoError(t, err)

	a, err := GetAuthCode(c, "example.org")
	require.NoError(t, err)
	assert.Equal(t, "Z290cmFuc2lwLXRlc3RpbmcK", a)
}

func TestGetIsLocked(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getislocked.xml")
	require.NoError(t, err)

	l, err := GetIsLocked(c, "example.org")
	require.NoError(t, err)
	assert.Equal(t, true, l)
}

func TestGetAllTLDInfos(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getalltldinfos.xml")
	require.NoError(t, err)

	lst, err := GetAllTLDInfos(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []TLD{}, lst)
	// tests copy/pasted from TestGetTldInfo
	tld := lst[0]
	assert.Equal(t, int64(2), tld.CancelTimeFrame)
	require.Equal(t, 2, len(tld.Capabilities))
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
	require.NoError(t, err)

	tld, err := GetTldInfo(c, ".org")
	require.NoError(t, err)
	assert.IsType(t, TLD{}, tld)
	assert.Equal(t, int64(2), tld.CancelTimeFrame)
	require.Equal(t, 2, len(tld.Capabilities))
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
	require.NoError(t, err)

	act, err := GetCurrentDomainAction(c, "example.org")
	require.NoError(t, err)
	assert.IsType(t, ActionResult{}, act)
	assert.Equal(t, true, act.HasFailed)
	assert.Equal(t, "test message", act.Message)
	assert.Equal(t, "test", act.Name)
}

func TestRequestAuthCode(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/requestauthcode.xml")
	require.NoError(t, err)

	code, err := RequestAuthCode(c, "example.org")
	require.NoError(t, err)
	assert.IsType(t, "", code)
	assert.Equal(t, "are0AeThe7er1Uyoo1aifowoMilohnae", code)
}

func TestGetDNSSecEntries(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdnssecentries.xml")
	require.NoError(t, err)

	dns, err := GetDNSSecEntries(c, "example.org")
	require.NoError(t, err)
	require.Equal(t, 2, len(dns))
	assert.IsType(t, DNSSecEntries{}, dns)
	assert.Equal(t, 1337, dns[0].KeyTag)
	assert.Equal(t, DNSSecFlagKSK, dns[0].Flags)
	assert.Equal(t, DNSSecAlgorithmRSASHA512, dns[0].Algorithm)
	assert.Equal(t, "emFpcmFvcHUzdm9vNGFpczlxdWVpN211dTVvc2g1QWVHaG9pY2kxaWVnaDZmYWk5ZWVtb2hwOWdhaU5nb29DZWlxdWFoZjNhaWdhaDJyYWhRdTFhaGxvaHNlaTh3ZW9zaDRhZXBoYWhzb29raTZFaWNoOGFpdnU5Y2llcGhvYWwK", dns[0].PublicKey)
	assert.Equal(t, 12, dns[1].KeyTag)
	assert.Equal(t, DNSSecFlagZSK, dns[1].Flags)
	assert.Equal(t, DNSSecAlgorithmECDSAP384SHA384, dns[1].Algorithm)
	assert.Equal(t, "dWl4YWl4MHBoZWVtN3lhcGhhaWIwYWhsYWVqMW9odzB1YThYaTFoYUJhaHBvOWhhZXNhaDJBaGQ4b2s4VGhvU2hhaWozc2hhaDluYWljYWljN2lvaG83aW9YZWRvb2w0YWhXYWl0bzNYZWlQaGFlNWVpZ2VpcGVlZzdhZXhpZTAK", dns[1].PublicKey)
}

func TestGetDefaultDNSEntries(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdefaultdnsentries.xml")
	require.NoError(t, err)

	entries, err := GetDefaultDNSEntries(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(entries))
	assert.IsType(t, DNSEntries{}, entries)
	assert.Equal(t, "1.2.3.4", entries[0].Content)
	assert.Equal(t, int64(86400), entries[0].TTL)
	assert.Equal(t, "@", entries[0].Name)
	assert.Equal(t, DNSEntryTypeA, entries[0].Type)
	assert.Equal(t, "fe80::1", entries[1].Content)
}

func TestGetDefaultDNSEntriesByDomainName(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdefaultdnsentriesbydomainname.xml")
	require.NoError(t, err)

	entries, err := GetDefaultDNSEntriesByDomainName(c, "example.org")
	require.NoError(t, err)
	require.Equal(t, 2, len(entries))
	assert.IsType(t, DNSEntries{}, entries)
	assert.Equal(t, "1.2.3.4", entries[0].Content)
	assert.Equal(t, int64(86400), entries[0].TTL)
	assert.Equal(t, "@", entries[0].Name)
	assert.Equal(t, DNSEntryTypeA, entries[0].Type)
	assert.Equal(t, "fe80::1", entries[1].Content)
}

func TestGetDefaultNameservers(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdefaultnameservers.xml")
	require.NoError(t, err)

	ns, err := GetDefaultNameservers(c)
	require.NoError(t, err)
	require.Equal(t, 3, len(ns))
	assert.IsType(t, Nameservers{}, ns)
	assert.Equal(t, "ns0.transip.net", ns[0].Hostname)
	assert.Equal(t, net.ParseIP("195.135.195.195"), ns[0].IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8:dddd:195::195"), ns[0].IPv6Address)
	assert.Equal(t, "ns1.transip.nl", ns[1].Hostname)
	assert.Equal(t, "ns2.transip.eu", ns[2].Hostname)
}

func TestGetDefaultNameserversByDomainName(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdefaultnameserversbydomainname.xml")
	require.NoError(t, err)

	ns, err := GetDefaultNameserversByDomainName(c, "example.org")
	require.NoError(t, err)
	require.Equal(t, 3, len(ns))
	assert.IsType(t, Nameservers{}, ns)
	assert.Equal(t, "ns0.transip.net", ns[0].Hostname)
	assert.Equal(t, net.ParseIP("195.135.195.195"), ns[0].IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8:dddd:195::195"), ns[0].IPv6Address)
	assert.Equal(t, "ns1.transip.nl", ns[1].Hostname)
	assert.Equal(t, "ns2.transip.eu", ns[2].Hostname)
}

func TestCancel(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/cancel.xml")
	require.NoError(t, err)

	err = Cancel(c, "example.org", gotransip.CancellationTimeImmediately)
	require.NoError(t, err)
}

func TestCancelDomainAction(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/canceldomainaction.xml")
	require.NoError(t, err)

	err = CancelDomainAction(c, Domain{Name: "example.org"})
	require.NoError(t, err)
}

func TestCanEditDNSSec(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/caneditdnssec.xml")
	require.NoError(t, err)

	can, err := CanEditDNSSec(c, "example.org")
	require.NoError(t, err)
	assert.Equal(t, true, can)
}

func TestHandover(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/handover.xml")
	require.NoError(t, err)

	err = Handover(c, "example.org", "not-my-account")
	require.NoError(t, err)
}

func TestHandoverWithAuthCode(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/handoverwithauthcode.xml")
	require.NoError(t, err)

	err = HandoverWithAuthCode(c, "example.org", "s3cr3t")
	require.NoError(t, err)
}

func TestRegister(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/register.xml")
	require.NoError(t, err)

	prop, err := Register(c, Domain{Name: "example.org"})
	require.NoError(t, err)
	assert.Equal(t, "1234-5678", prop)
}

func TestRemoveAllDNSSecEntries(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/removealldnssecentries.xml")
	require.NoError(t, err)

	err = RemoveAllDNSSecEntries(c, "example.org")
	require.NoError(t, err)
}

func TestRetryCurrentDomainActionWithNewData(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/retrycurrentdomainactionwithnewdata.xml")
	require.NoError(t, err)

	err = RetryCurrentDomainActionWithNewData(c, Domain{Name: "example.org"})
	require.NoError(t, err)
}

func TestRetryTransferWithDifferentAuthCode(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/retrytransferwithdifferentauthcode.xml")
	require.NoError(t, err)

	err = RetryTransferWithDifferentAuthCode(c, Domain{Name: "example.org"}, "s3cr3t")
	require.NoError(t, err)
}

func TestSetContacts(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setcontacts.xml")
	require.NoError(t, err)

	err = SetContacts(c, "example.org", []WhoisContact{{FirstName: "John", LastName: "Doe"}})
	require.NoError(t, err)
}

func TestSetDNSEntries(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setdnsentries.xml")
	require.NoError(t, err)

	err = SetDNSEntries(c, "example.org", []DNSEntry{
		{Type: DNSEntryTypeA, Name: "www"},
	})
	require.NoError(t, err)
}

func TestSetDNSSecEntries(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setdnssecentries.xml")
	require.NoError(t, err)

	err = SetDNSSecEntries(c, "example.org", []DNSSecEntry{
		{PublicKey: "s3cr3t"},
	})
	require.NoError(t, err)
}

func TestSetLock(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setlock.xml")
	require.NoError(t, err)

	err = SetLock(c, "example.org")
	require.NoError(t, err)
}

func TestSetNameservers(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setnameservers.xml")
	require.NoError(t, err)

	err = SetNameservers(c, "example.org", []Nameserver{
		{Hostname: "ns1.transip.nl"},
	})
	require.NoError(t, err)
}

func TestSetOwner(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setowner.xml")
	require.NoError(t, err)

	err = SetOwner(c, "example.org", WhoisContact{FirstName: "John", LastName: "Doe"})
	require.NoError(t, err)
}

func TestTransferWithoutOwnerChange(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/transferwithoutownerchange.xml")
	require.NoError(t, err)

	prop, err := TransferWithoutOwnerChange(c, Domain{Name: "example.org"}, "s3cr3t")
	require.NoError(t, err)
	assert.Equal(t, "5678-9012", prop)
}

func TestTransferWithOwnerChange(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/transferwithownerchange.xml")
	require.NoError(t, err)

	prop, err := TransferWithOwnerChange(c, Domain{Name: "example.org"}, "s3cr3t")
	require.NoError(t, err)
	assert.Equal(t, "3456-7890", prop)
}

func TestUnsetLock(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/unsetlock.xml")
	require.NoError(t, err)

	err = UnsetLock(c, "example.org")
	require.NoError(t, err)
}
