package dns

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
)

func TestKeyAlgorithms(t *testing.T) {
	assert.Equal(t, 3, int(KeyAlgorithmDSA))
	assert.Equal(t, 5, int(KeyAlgorithmRSASHA1))
	assert.Equal(t, 6, int(KeyAlgorithmDSANSEC3SHA1))
	assert.Equal(t, 7, int(KeyAlgorithmRSASHA1NSEC3SHA1))
	assert.Equal(t, 8, int(KeyAlgorithmRSASHA256))
	assert.Equal(t, 10, int(KeyAlgorithmRSASHA512))
	assert.Equal(t, 12, int(KeyAlgorithmECCGOST))
	assert.Equal(t, 13, int(KeyAlgorithmECDSAP256SHA256))
	assert.Equal(t, 14, int(KeyAlgorithmECDSAP384SHA384))
	assert.Equal(t, 15, int(KeyAlgorithmED25519))
	assert.Equal(t, 16, int(KeyAlgorithmED448))
}

func TestKeyFlags(t *testing.T) {
	assert.Equal(t, 0, int(KeyFlagNone))
	assert.Equal(t, 256, int(KeyFlagZSK))
	assert.Equal(t, 257, int(KeyFlagKSK))
}

func TestDnsEntriesEncoding(t *testing.T) {
	entries := Entries{
		{
			Name:    "www",
			Content: "1.2.3.4",
			TTL:     3600,
			Type:    EntryTypeA,
		},
		{
			Name:    "@",
			Content: "ns0.transip.net",
			TTL:     86400,
			Type:    EntryTypeNS,
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

func TestDnsSecEntriesEncoding(t *testing.T) {
	entries := KeyEntries{
		{
			KeyTag:    1337,
			Flags:     KeyFlagKSK,
			Algorithm: KeyAlgorithmRSASHA512,
			PublicKey: "emFpcmFvcHUzdm9...cGhvYWwK",
		},
		{
			KeyTag:    12,
			Flags:     KeyFlagZSK,
			Algorithm: KeyAlgorithmECDSAP384SHA384,
			PublicKey: "dWl4YWl4MHBoZWV...ZXhpZTAK",
		},
	}

	fixtArgs := `<dnsSecEntries SOAP-ENC:arrayType="ns1:DnsSecEntry[2]" xsi:type="ns1:ArrayOfDnsSecEntry">
	<item xsi:type="ns1:DnsSecEntry">
		<keyTag xsi:type="xsd:int">1337</keyTag>
		<flags xsi:type="xsd:int">257</flags>
		<algorithm xsi:type="xsd:int">10</algorithm>
		<publicKey xsi:type="xsd:string">emFpcmFvcHUzdm9...cGhvYWwK</publicKey>
	</item>
	<item xsi:type="ns1:DnsSecEntry">
		<keyTag xsi:type="xsd:int">12</keyTag>
		<flags xsi:type="xsd:int">256</flags>
		<algorithm xsi:type="xsd:int">14</algorithm>
		<publicKey xsi:type="xsd:string">dWl4YWl4MHBoZWV...ZXhpZTAK</publicKey>
	</item>
</dnsSecEntries>`
	assert.Equal(t, fixtArgs, entries.EncodeArgs("dnsSecEntries"))

	prm := gotransip.TestParamsContainer{}
	entries.EncodeParams(&prm, "")
	assert.Equal(t, "00[0][keyTag]=1337&180[0][flags]=257&360[0][algorithm]=10&570[0][publicKey]=emFpcmFvcHUzdm9...cGhvYWwK&1020[1][keyTag]=12&1210[1][flags]=256&1400[1][algorithm]=14&1620[1][publicKey]=dWl4YWl4MHBoZWV...ZXhpZTAK", prm.Prm)
}
