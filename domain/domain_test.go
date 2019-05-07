package domain

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
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
	assert.NoError(t, err)
	assert.Equal(t, fixtArgs, domain.EncodeArgs("domain"))

	fixtPrm, err := getFixture("testdata/encoding/domain_simple.prm")
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	assert.Equal(t, fixtArgs, domain.EncodeArgs("domain"))

	fixtPrm, err = getFixture("testdata/encoding/domain_full.prm")
	assert.NoError(t, err)
	prm = gotransip.TestParamsContainer{}
	domain.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)

	// try again with prefixed parameters
	fixtPrm, err = getFixture("testdata/encoding/domain_full_prefixed.prm")
	assert.NoError(t, err)
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

	fixtArgs := `<dnsEntries SOAP-ENC:arrayType="ns1:DnsEntry[2]" xsi:type="ns1:ArrayOfDnsEntry">
	<item xsi:type="ns1:DnsEntry">
		<name xsi:type="xsd:string">www</name>
		<expire xsi:type="xsd:int">3600</expire>
		<type xsi:type="xsd:string">A</type>
		<content xsi:type="xsd:string">1.2.3.4</content>
	</item>
	<item xsi:type="ns1:DnsEntry">
		<name xsi:type="xsd:string">@</name>
		<expire xsi:type="xsd:int">86400</expire>
		<type xsi:type="xsd:string">NS</type>
		<content xsi:type="xsd:string">ns0.transip.net</content>
	</item>
</dnsEntries>`
	assert.Equal(t, fixtArgs, entries.EncodeArgs("dnsEntries"))

	prm := gotransip.TestParamsContainer{}
	entries.EncodeParams(&prm, "")
	assert.Equal(t, "00[0][name]=www&150[0][expire]=3600&350[0][type]=A&500[0][content]=1.2.3.4&740[1][name]=@&890[1][expire]=86400&1100[1][type]=NS&1270[1][content]=ns0.transip.net", prm.Prm)
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

	fixtArgs := `<nameServers SOAP-ENC:arrayType="ns1:Nameserver[3]" xsi:type="ns1:ArrayOfNameserver">
	<item xsi:type="ns1:Nameserver">
		<hostname xsi:type="xsd:string">ns0.transip.net</hostname>
		<ipv4 xsi:type="xsd:string">195.135.195.195</ipv4>
		<ipv6 xsi:type="xsd:string">2a01:7c8:dddd:195::195</ipv6>
	</item>
	<item xsi:type="ns1:Nameserver">
		<hostname xsi:type="xsd:string">ns1.transip.nl</hostname>
		<ipv4 xsi:type="xsd:string">195.8.195.195</ipv4>
		<ipv6 xsi:type="xsd:string"></ipv6>
	</item>
	<item xsi:type="ns1:Nameserver">
		<hostname xsi:type="xsd:string">ns2.transip.eu</hostname>
		<ipv4 xsi:type="xsd:string"></ipv4>
		<ipv6 xsi:type="xsd:string">2a01:7c8:f:c1f::195</ipv6>
	</item>
</nameServers>`
	assert.Equal(t, fixtArgs, nameservers.EncodeArgs("nameServers"))
	prm := gotransip.TestParamsContainer{}
	nameservers.EncodeParams(&prm, "")
	assert.Equal(t, "00[0][hostname]=ns0.transip.net&310[0][ipv4]=195.135.195.195&600[0][ipv6]=2a01:7c8:dddd:195::195&960[1][hostname]=ns1.transip.nl&1280[1][ipv4]=195.8.195.195&1560[1][ipv6]=&1710[2][hostname]=ns2.transip.eu&2040[2][ipv4]=&2190[2][ipv6]=2a01:7c8:f:c1f::195", prm.Prm)
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

	fixtArgs := `<branding xsi:type="ns1:DomainBranding">
		<companyName xsi:type="xsd:string">TransIP BV</companyName>
		<supportEmail xsi:type="xsd:string">support@transip.nl</supportEmail>
		<companyUrl xsi:type="xsd:string">https://transip.nl/</companyUrl>
		<termsOfUsageUrl xsi:type="xsd:string">https://transip.nl/tou</termsOfUsageUrl>
		<bannerLine1 xsi:type="xsd:string">TransIP 1</bannerLine1>
		<bannerLine2 xsi:type="xsd:string">TransIP 2</bannerLine2>
		<bannerLine3 xsi:type="xsd:string">TransIP 3</bannerLine3>
	</branding>`
	assert.Equal(t, fixtArgs, brand.EncodeArgs("branding"))
	prm := gotransip.TestParamsContainer{}
	brand.EncodeParams(&prm, "")
	assert.Equal(t, "00[companyName]=TransIP BV&260[supportEmail]=support@transip.nl&630[companyUrl]=https://transip.nl/&990[termsOfUsageUrl]=https://transip.nl/tou&1430[bannerLine1]=TransIP 1&1710[bannerLine2]=TransIP 2&1990[bannerLine3]=TransIP 3", prm.Prm)
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

	fixtArgs := `<contacts SOAP-ENC:arrayType="ns1:WhoisContact[1]" xsi:type="ns1:ArrayOfWhoisContact">
	<item xsi:type="ns1:WhoisContact">
		<type xsi:type="xsd:string">registrant</type>
		<firstName xsi:type="xsd:string">foo</firstName>
		<middleName xsi:type="xsd:string">bar</middleName>
		<lastName xsi:type="xsd:string">baz</lastName>
		<companyName xsi:type="xsd:string">TransIP BV</companyName>
		<companyKvk xsi:type="xsd:string">1234</companyKvk>
		<companyType xsi:type="xsd:string">BV</companyType>
		<street xsi:type="xsd:string">Schipholweg</street>
		<number xsi:type="xsd:string">9B</number>
		<postalCode xsi:type="xsd:string">2316XB</postalCode>
		<city xsi:type="xsd:string">Leiden</city>
		<phoneNumber xsi:type="xsd:string">+31 715241919</phoneNumber>
		<faxNumber xsi:type="xsd:string">+31 715241918</faxNumber>
		<email xsi:type="xsd:string">support@transip.nl</email>
		<country xsi:type="xsd:string">nl</country>
	</item>
</contacts>`
	assert.Equal(t, fixtArgs, whois.EncodeArgs("contacts"))
	prm := gotransip.TestParamsContainer{}
	whois.EncodeParams(&prm, "")
	assert.Equal(t, "00[0][type]=registrant&220[0][firstName]=foo&440[0][middleName]=bar&670[0][lastName]=baz&880[0][companyName]=TransIP BV&1190[0][companyKvk]=1234&1440[0][companyType]=BV&1680[0][street]=Schipholweg&1960[0][number]=9B&2150[0][postalCode]=2316XB&2420[0][city]=Leiden&2630[0][phoneNumber]=+31 715241919&2980[0][faxNumber]=+31 715241918&3310[0][email]=support@transip.nl&3650[0][country]=nl", prm.Prm)
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
	assert.NoError(t, err)
	assert.Equal(t, fixtArgs, entries.EncodeArgs("dnsSecEntries"))

	// test EncodeParams
	fixtPrm, err := getFixture("testdata/encoding/dnssecentries.prm")
	assert.NoError(t, err)
	prm := gotransip.TestParamsContainer{}
	entries.EncodeParams(&prm, "")
	assert.Equal(t, fixtPrm, prm.Prm)
}
