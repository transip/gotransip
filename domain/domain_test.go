package domain

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v5"
)

func getFixture(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("could not read fixture from file %s: %s", filename, err.Error())
	}

	fixt := strings.TrimSpace(string(data))

	return fixt, err
}

func TestDomainEncoding(t *testing.T) {
	// start with a simple domain without any complex fields
	domain := Domain{
		Name:              "example.org",
		AuthorizationCode: "abcxyz",
		IsLocked:          true,
	}

	// read fixture files and compare XML and encoded parameters
	fixtArgs, err := getFixture("testdata/encoding/domain_simple.xml")
	require.NoError(t, err)
	assert.Equal(t, fixtArgs, domain.EncodeArgs("domain"))

	fixtPrm, err := getFixture("testdata/encoding/domain_simple.prm")
	require.NoError(t, err)
	prm := gotransip.TestParamsContainer{}
	domain.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)

	// dress up domain some more
	domain.Nameservers = Nameservers{
		{
			Hostname:    "ns0.transip.net",
			IPv4Address: net.IP{195, 135, 195, 195},
			IPv6Address: net.ParseIP("2a01:7c8:dddd:195::195"),
		},
	}
	domain.Contacts = WhoisContacts{
		{
			Type:        "registrant",
			FirstName:   "foo",
			MiddleName:  "bar",
			LastName:    "baz",
			CompanyName: "TransIP BV",
			CompanyKvk:  "1234",
			CompanyType: "BV",
			Street:      "Schipholweg",
			Number:      "9B",
			PostalCode:  "2316XB",
			City:        "Leiden",
			PhoneNumber: "+31 715241919",
			FaxNumber:   "+31 715241918",
			Email:       "support@transip.nl",
			Country:     "nl",
		},
	}
	domain.DNSEntries = DNSEntries{
		{
			Name:    "www",
			Content: "1.2.3.4",
			TTL:     3600,
			Type:    DNSEntryTypeA,
		},
	}
	domain.Branding = Branding{
		CompanyName:     "TransIP BV",
		SupportEmail:    "support@transip.nl",
		CompanyURL:      "https://transip.nl/",
		TermsOfUsageURL: "https://transip.nl/tou",
		BannerLine1:     "TransIP 1",
		BannerLine2:     "TransIP 2",
		BannerLine3:     "TransIP 3",
	}

	// read fixture files and compare XML and encoded parameters
	fixtArgs, err = getFixture("testdata/encoding/domain_full.xml")
	require.NoError(t, err)
	assert.Equal(t, fixtArgs, domain.EncodeArgs("domain"))

	fixtPrm, err = getFixture("testdata/encoding/domain_full.prm")
	require.NoError(t, err)
	prm = gotransip.TestParamsContainer{}
	domain.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)

	// try again with prefixed parameters
	fixtPrm, err = getFixture("testdata/encoding/domain_full_prefixed.prm")
	require.NoError(t, err)
	prm = gotransip.TestParamsContainer{}
	domain.EncodeParams(&prm, "5[domain]")
	assert.Equal(t, fixtPrm, prm.Prm)
}

func TestDnsEntriesEncoding(t *testing.T) {
	entries := DNSEntries{
		{
			Name:    "www",
			Content: "1.2.3.4",
			TTL:     3600,
			Type:    DNSEntryTypeA,
		},
		{
			Name:    "@",
			Content: "ns0.transip.net",
			TTL:     86400,
			Type:    DNSEntryTypeNS,
		},
	}

	// test EncodeArgs
	fixtArgs, err := getFixture("testdata/encoding/dnsentries.xml")
	require.NoError(t, err)
	assert.Equal(t, fixtArgs, entries.EncodeArgs("dnsEntries"))

	// test EncodeParams
	fixtPrm, err := getFixture("testdata/encoding/dnsentries.prm")
	require.NoError(t, err)
	prm := gotransip.TestParamsContainer{}
	entries.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)
}

func TestNameserversEncoding(t *testing.T) {
	nameservers := Nameservers{
		{
			Hostname:    "ns0.transip.net",
			IPv4Address: net.IP{195, 135, 195, 195},
			IPv6Address: net.ParseIP("2a01:7c8:dddd:195::195"),
		},
		{
			Hostname:    "ns1.transip.nl",
			IPv4Address: net.IP{195, 8, 195, 195},
		},
		{
			Hostname:    "ns2.transip.eu",
			IPv6Address: net.ParseIP("2a01:7c8:f:c1f::195"),
		},
	}

	// test EncodeArgs
	fixtArgs, err := getFixture("testdata/encoding/nameservers.xml")
	require.NoError(t, err)
	assert.Equal(t, fixtArgs, nameservers.EncodeArgs("nameServers"))

	// test EncodeParams
	fixtPrm, err := getFixture("testdata/encoding/nameservers.prm")
	require.NoError(t, err)
	prm := gotransip.TestParamsContainer{}
	nameservers.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)
}

func TestBrandingEncoding(t *testing.T) {
	brand := Branding{
		CompanyName:     "TransIP BV",
		SupportEmail:    "support@transip.nl",
		CompanyURL:      "https://transip.nl/",
		TermsOfUsageURL: "https://transip.nl/tou",
		BannerLine1:     "TransIP 1",
		BannerLine2:     "TransIP 2",
		BannerLine3:     "TransIP 3",
	}

	// test EncodeArgs
	fixtArgs, err := getFixture("testdata/encoding/branding.xml")
	require.NoError(t, err)
	assert.Equal(t, fixtArgs, brand.EncodeArgs("branding"))

	// test EncodeParams
	fixtPrm, err := getFixture("testdata/encoding/branding.prm")
	require.NoError(t, err)
	prm := gotransip.TestParamsContainer{}
	brand.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)
}

func TestWhoisContactsEncoding(t *testing.T) {
	whois := WhoisContacts{
		{
			Type:        "registrant",
			FirstName:   "foo",
			MiddleName:  "bar",
			LastName:    "baz",
			CompanyName: "TransIP BV",
			CompanyKvk:  "1234",
			CompanyType: "BV",
			Street:      "Schipholweg",
			Number:      "9B",
			PostalCode:  "2316XB",
			City:        "Leiden",
			PhoneNumber: "+31 715241919",
			FaxNumber:   "+31 715241918",
			Email:       "support@transip.nl",
			Country:     "nl",
		},
	}

	// test EncodeArgs
	fixtArgs, err := getFixture("testdata/encoding/whoiscontacts.xml")
	require.NoError(t, err)
	assert.Equal(t, fixtArgs, whois.EncodeArgs("contacts"))

	// test EncodeParams
	fixtPrm, err := getFixture("testdata/encoding/whoiscontacts.prm")
	require.NoError(t, err)
	prm := gotransip.TestParamsContainer{}
	whois.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)
}

func TestDNSSecAlgorithms(t *testing.T) {
	assert.Equal(t, 3, int(DNSSecAlgorithmDSA))
	assert.Equal(t, 5, int(DNSSecAlgorithmRSASHA1))
	assert.Equal(t, 6, int(DNSSecAlgorithmDSANSEC3SHA1))
	assert.Equal(t, 7, int(DNSSecAlgorithmRSASHA1NSEC3SHA1))
	assert.Equal(t, 8, int(DNSSecAlgorithmRSASHA256))
	assert.Equal(t, 10, int(DNSSecAlgorithmRSASHA512))
	assert.Equal(t, 12, int(DNSSecAlgorithmECCGOST))
	assert.Equal(t, 13, int(DNSSecAlgorithmECDSAP256SHA256))
	assert.Equal(t, 14, int(DNSSecAlgorithmECDSAP384SHA384))
	assert.Equal(t, 15, int(DNSSecAlgorithmED25519))
	assert.Equal(t, 16, int(DNSSecAlgorithmED448))
}

func TestDNSSecFlags(t *testing.T) {
	assert.Equal(t, 0, int(DNSSecFlagNone))
	assert.Equal(t, 256, int(DNSSecFlagZSK))
	assert.Equal(t, 257, int(DNSSecFlagKSK))
}

func TestDNSSecEntriesEncoding(t *testing.T) {
	entries := DNSSecEntries{
		{
			KeyTag:    1337,
			Flags:     DNSSecFlagKSK,
			Algorithm: DNSSecAlgorithmRSASHA512,
			PublicKey: "emFpcmFvcHUzdm9...cGhvYWwK",
		},
		{
			KeyTag:    12,
			Flags:     DNSSecFlagZSK,
			Algorithm: DNSSecAlgorithmECDSAP384SHA384,
			PublicKey: "dWl4YWl4MHBoZWV...ZXhpZTAK",
		},
	}

	fixtArgs, err := getFixture("testdata/encoding/dnssecentries.xml")
	require.NoError(t, err)
	assert.Equal(t, fixtArgs, entries.EncodeArgs("dnsSecEntries"))

	// test EncodeParams
	fixtPrm, err := getFixture("testdata/encoding/dnssecentries.prm")
	require.NoError(t, err)
	prm := gotransip.TestParamsContainer{}
	entries.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)

	// test EncodeParams with prefix
	fixtPrm, err = getFixture("testdata/encoding/dnssecentries_prefixed.prm")
	require.NoError(t, err)
	prm = gotransip.TestParamsContainer{}
	entries.EncodeParams(&prm, "test")
	assert.Equal(t, fixtPrm, prm.Prm)

	// test EncodeParams on empty set
	prm = gotransip.TestParamsContainer{}
	DNSSecEntries{}.EncodeParams(&prm, "")
	assert.Equal(t, "", prm.Prm)
}
