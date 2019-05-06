package dns

import (
	"fmt"

	"github.com/transip/gotransip"
)

const (
	serviceName string = "DnsService"
)

type DNSEntryType string

var (
	// DNSEntryTypeA represents an A-record
	DNSEntryTypeA DNSEntryType = "A"
	// DNSEntryTypeAAAA represents an AAAA-record
	DNSEntryTypeAAAA DNSEntryType = "AAAA"
	// DNSEntryTypeCNAME represents a CNAME-record
	DNSEntryTypeCNAME DNSEntryType = "CNAME"
	// DNSEntryTypeMX represents an MX-record
	DNSEntryTypeMX DNSEntryType = "MX"
	// DNSEntryTypeNS represents an NS-record
	DNSEntryTypeNS DNSEntryType = "NS"
	// DNSEntryTypeTXT represents a TXT-record
	DNSEntryTypeTXT DNSEntryType = "TXT"
	// DNSEntryTypeSRV represents an SRV-record
	DNSEntryTypeSRV DNSEntryType = "SRV"
)

// DNSEntry represents a Transip_DnsEntry object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_DnsEntry.html
type DNSEntry struct {
	Name    string       `xml:"name"`
	TTL     int64        `xml:"expire"`
	Type    DNSEntryType `xml:"type"`
	Content string       `xml:"content"`
}

// DNSEntries is just an array of DNSEntry
// basically only here so it can implement paramsEncoder
type DNSEntries []DNSEntry

// EncodeParams returns DNSEntries parameters ready to be used for constructing a signature
// the order of parameters added here has to match the order in the WSDL
// as described at http://api.transip.nl/wsdl/?service=DomainService
func (d DNSEntries) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
	if len(d) == 0 {
		prm.Add("anything", nil)
		return
	}

	if len(prefix) == 0 {
		prefix = fmt.Sprintf("%d", prm.Len())
	}

	for i, e := range d {
		prm.Add(fmt.Sprintf("%s[%d][name]", prefix, i), e.Name)
		prm.Add(fmt.Sprintf("%s[%d][expire]", prefix, i), fmt.Sprintf("%d", e.TTL))
		prm.Add(fmt.Sprintf("%s[%d][type]", prefix, i), string(e.Type))
		prm.Add(fmt.Sprintf("%s[%d][content]", prefix, i), e.Content)
	}
}

// EncodeArgs returns DNSEntries XML body ready to be passed in the SOAP call
func (d DNSEntries) EncodeArgs(key string) string {
	output := fmt.Sprintf(`<%s SOAP-ENC:arrayType="ns1:DnsEntry[%d]" xsi:type="ns1:ArrayOfDnsEntry">`, key, len(d)) + "\n"
	for _, e := range d {
		output += fmt.Sprintf(`	<item xsi:type="ns1:DnsEntry">
		<name xsi:type="xsd:string">%s</name>
		<expire xsi:type="xsd:int">%d</expire>
		<type xsi:type="xsd:string">%s</type>
		<content xsi:type="xsd:string">%s</content>
	</item>`, e.Name, e.TTL, e.Type, e.Content) + "\n"
	}

	return fmt.Sprintf("%s</%s>", output, key)
}

type DNSSecAlgorithm int

const (
	DNSSecAlgorithmDSA DNSSecAlgorithm = iota + 3
	_
	DNSSecAlgorithmRSASHA1
	DNSSecAlgorithmDSANSEC3SHA1
	DNSSecAlgorithmRSASHA1NSEC3SHA1
	DNSSecAlgorithmRSASHA256
	DNSSecAlgorithmRSASHA512 DNSSecAlgorithm = iota + 4
	_
	DNSSecAlgorithmECCGOST
	DNSSecAlgorithmECDSAP256SHA256
	DNSSecAlgorithmECDSAP384SHA384
	DNSSecAlgorithmED25519
	DNSSecAlgorithmED448
)

type DNSSecFlag int

const (
	DNSSecFlagNone DNSSecFlag = 0
	DNSSecFlagZSK  DNSSecFlag = 256
	DNSSecFlagKSK  DNSSecFlag = 257
)

// DNSSecEntry represents a Transip_DnsSecEntry object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_DnsSecEntry.html
type DNSSecEntry struct {
	KeyTag    int             `xml:"keyTag"`
	Flags     DNSSecFlag      `xml:"flags"`
	Algorithm DNSSecAlgorithm `xml:"algorithm"`
	PublicKey string          `xml:"publicKey"`
}

// DNSSecEntry is just an array of DNSSecEntries
// basically only here so it can implement paramsEncoder
type DNSSecEntries []DNSSecEntry

// EncodeParams returns DNSSecEntries parameters ready to be used for constructing
// a signature
// the order of parameters added here has to match the order in the WSDL as
// described at http://api.transip.nl/wsdl/?service=DnsService
func (d DNSSecEntries) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
	if len(d) == 0 {
		prm.Add("anything", nil)
		return
	}

	if len(prefix) == 0 {
		prefix = fmt.Sprintf("%d", prm.Len())
	}

	for i, e := range d {
		prm.Add(fmt.Sprintf("%s[%d][keyTag]", prefix, i), fmt.Sprintf("%d", e.KeyTag))
		prm.Add(fmt.Sprintf("%s[%d][flags]", prefix, i), fmt.Sprintf("%d", e.Flags))
		prm.Add(fmt.Sprintf("%s[%d][algorithm]", prefix, i), fmt.Sprintf("%d", e.Algorithm))
		prm.Add(fmt.Sprintf("%s[%d][publicKey]", prefix, i), e.PublicKey)
	}
}

// EncodeArgs returns DNSEntries XML body ready to be passed in the SOAP call
func (d DNSSecEntries) EncodeArgs(key string) string {
	output := fmt.Sprintf(`<%s SOAP-ENC:arrayType="ns1:DnsSecEntry[%d]" xsi:type="ns1:ArrayOfDnsSecEntry">`, key, len(d)) + "\n"
	for _, e := range d {
		output += fmt.Sprintf(`	<item xsi:type="ns1:DnsSecEntry">
		<keyTag xsi:type="xsd:int">%d</keyTag>
		<flags xsi:type="xsd:int">%d</flags>
		<algorithm xsi:type="xsd:int">%d</algorithm>
		<publicKey xsi:type="xsd:string">%s</publicKey>
	</item>`, e.KeyTag, e.Flags, e.Algorithm, e.PublicKey) + "\n"
	}

	return fmt.Sprintf("%s</%s>", output, key)
}
