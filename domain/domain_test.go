package domain

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
)

func TestDomainEncoding(t *testing.T) {
	domain := Domain{
		Name: "example.org",
		Nameservers: Nameservers{
			{
				Hostname:    "ns0.transip.net",
				IPv4Address: net.IP{195, 135, 195, 195},
				IPv6Address: net.ParseIP("2a01:7c8:dddd:195::195"),
			},
		},
		Contacts: WhoisContacts{
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
		},
		DNSEntries: DNSEntries{
			{
				Name:    "www",
				Content: "1.2.3.4",
				TTL:     3600,
				Type:    DNSEntryTypeA,
			},
		},
		Branding: Branding{
			CompanyName:     "TransIP BV",
			SupportEmail:    "support@transip.nl",
			CompanyURL:      "https://transip.nl/",
			TermsOfUsageURL: "https://transip.nl/tou",
			BannerLine1:     "TransIP 1",
			BannerLine2:     "TransIP 2",
			BannerLine3:     "TransIP 3",
		},
		AuthorizationCode: "abcxyz",
		IsLocked:          true,
	}

	fixtArgs := `<domain xsi:type="ns1:Domain">
	<name xsi:type="xsd:string">example.org</name>
	<authCode xsi:type="xsd:string">abcxyz</authCode>
	<isLocked xsi:type="xsd:boolean">true</isLocked>
	<registrationDate xsi:type="xsd:string">0001-01-01</registrationDate>
	<renewalDate xsi:type="xsd:string">0001-01-01</renewalDate>
<nameservers SOAP-ENC:arrayType="ns1:Nameserver[1]" xsi:type="ns1:ArrayOfNameserver">
	<item xsi:type="ns1:Nameserver">
		<hostname xsi:type="xsd:string">ns0.transip.net</hostname>
		<ipv4 xsi:type="xsd:string">195.135.195.195</ipv4>
		<ipv6 xsi:type="xsd:string">2a01:7c8:dddd:195::195</ipv6>
	</item>
</nameservers>
<contacts SOAP-ENC:arrayType="ns1:WhoisContact[1]" xsi:type="ns1:ArrayOfWhoisContact">
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
</contacts>
<dnsEntries SOAP-ENC:arrayType="ns1:DnsEntry[1]" xsi:type="ns1:ArrayOfDnsEntry">
	<item xsi:type="ns1:DnsEntry">
		<name xsi:type="xsd:string">www</name>
		<expire xsi:type="xsd:int">3600</expire>
		<type xsi:type="xsd:string">A</type>
		<content xsi:type="xsd:string">1.2.3.4</content>
	</item>
</dnsEntries>
<branding xsi:type="ns1:DomainBranding">
		<companyName xsi:type="xsd:string">TransIP BV</companyName>
		<supportEmail xsi:type="xsd:string">support@transip.nl</supportEmail>
		<companyUrl xsi:type="xsd:string">https://transip.nl/</companyUrl>
		<termsOfUsageUrl xsi:type="xsd:string">https://transip.nl/tou</termsOfUsageUrl>
		<bannerLine1 xsi:type="xsd:string">TransIP 1</bannerLine1>
		<bannerLine2 xsi:type="xsd:string">TransIP 2</bannerLine2>
		<bannerLine3 xsi:type="xsd:string">TransIP 3</bannerLine3>
	</branding>
</domain>`
	assert.Equal(t, fixtArgs, domain.EncodeArgs("domain"))

	prm := gotransip.TestParamsContainer{}
	domain.EncodeParams(&prm)
	assert.Equal(t, "00[name]=example.org&200[authCode]=abcxyz&410[isLocked]=1&570[registrationDate]=0001-01-01&900[renewalDate]=0001-01-01&1180[nameservers][0][hostname]=ns0.transip.net&1650[nameservers][0][ipv4]=195.135.195.195&2080[nameservers][0][ipv6]=2a01:7c8:dddd:195::195&2580[contacts][0][type]=registrant&2930[contacts][0][firstName]=foo&3260[contacts][0][middleName]=bar&3600[contacts][0][lastName]=baz&3920[contacts][0][companyName]=TransIP BV&4340[contacts][0][companyKvk]=1234&4690[contacts][0][companyType]=BV&5030[contacts][0][street]=Schipholweg&5410[contacts][0][number]=9B&5700[contacts][0][postalCode]=2316XB&6070[contacts][0][city]=Leiden&6380[contacts][0][phoneNumber]=+31 715241919&6830[contacts][0][faxNumber]=+31 715241918&7260[contacts][0][email]=support@transip.nl&7700[contacts][0][country]=nl&8000[dnsEntries][0][name]=www&8300[dnsEntries][0][expire]=3600&8630[dnsEntries][0][type]=A&8910[dnsEntries][0][content]=1.2.3.4&9280[branding][companyName]=TransIP BV&9670[branding][supportEmail]=support@transip.nl&10150[branding][companyUrl]=https://transip.nl/&10630[branding][termsOfUsageUrl]=https://transip.nl/tou&11190[branding][bannerLine1]=TransIP 1&11580[branding][bannerLine2]=TransIP 2&11970[branding][bannerLine3]=TransIP 3", prm.Prm)
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
	entries.EncodeParams(&prm)
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
	nameservers.EncodeParams(&prm)
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
	brand.EncodeParams(&prm)
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
	whois.EncodeParams(&prm)
	assert.Equal(t, "00[0][type]=registrant&220[0][firstName]=foo&440[0][middleName]=bar&670[0][lastName]=baz&880[0][companyName]=TransIP BV&1190[0][companyKvk]=1234&1440[0][companyType]=BV&1680[0][street]=Schipholweg&1960[0][number]=9B&2150[0][postalCode]=2316XB&2420[0][city]=Leiden&2630[0][phoneNumber]=+31 715241919&2980[0][faxNumber]=+31 715241918&3310[0][email]=support@transip.nl&3650[0][country]=nl", prm.Prm)
}
