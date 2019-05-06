package dns

import (
	"fmt"

	"github.com/transip/gotransip"
)

const (
	serviceName string = "DnsService"
)

// EntryType represents the possible types of DNS entries
type EntryType string

var (
	// EntryTypeA represents an A-record
	EntryTypeA EntryType = "A"
	// EntryTypeAAAA represents an AAAA-record
	EntryTypeAAAA EntryType = "AAAA"
	// EntryTypeCNAME represents a CNAME-record
	EntryTypeCNAME EntryType = "CNAME"
	// EntryTypeMX represents an MX-record
	EntryTypeMX EntryType = "MX"
	// EntryTypeNS represents an NS-record
	EntryTypeNS EntryType = "NS"
	// EntryTypeTXT represents a TXT-record
	EntryTypeTXT EntryType = "TXT"
	// EntryTypeSRV represents an SRV-record
	EntryTypeSRV EntryType = "SRV"
)

// Entry represents a Transip_DnsEntry object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_DnsEntry.html
type Entry struct {
	Name    string    `xml:"name"`
	TTL     int64     `xml:"expire"`
	Type    EntryType `xml:"type"`
	Content string    `xml:"content"`
}

// Entries is just an array of Entry
// basically only here so it can implement paramsEncoder
type Entries []Entry

// EncodeParams returns Entries parameters ready to be used for constructing a signature
// the order of parameters added here has to match the order in the WSDL
// as described at http://api.transip.nl/wsdl/?service=DomainService
func (d Entries) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
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

// EncodeArgs returns Entries XML body ready to be passed in the SOAP call
func (d Entries) EncodeArgs(key string) string {
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

// KeyAlgorithm represents the possible types of DNSSec algorithms
type KeyAlgorithm int

const (
	// KeyAlgorithmDSA represents DSA
	KeyAlgorithmDSA KeyAlgorithm = iota + 3
	_
	// KeyAlgorithmRSASHA1 represents RSASHA1
	KeyAlgorithmRSASHA1
	// KeyAlgorithmDSANSEC3SHA1 represents DSANSEC3SHA1
	KeyAlgorithmDSANSEC3SHA1
	// KeyAlgorithmRSASHA1NSEC3SHA1 represents RSASHA1NSEC3SHA1
	KeyAlgorithmRSASHA1NSEC3SHA1
	// KeyAlgorithmRSASHA256 represents RSASHA256
	KeyAlgorithmRSASHA256
	// KeyAlgorithmRSASHA512 represents RSASHA512
	KeyAlgorithmRSASHA512 KeyAlgorithm = iota + 4
	_
	// KeyAlgorithmECCGOST represents ECCGOST
	KeyAlgorithmECCGOST
	// KeyAlgorithmECDSAP256SHA256 represents ECDSAP256SHA256
	KeyAlgorithmECDSAP256SHA256
	// KeyAlgorithmECDSAP384SHA384 represents ECDSAP384SHA384
	KeyAlgorithmECDSAP384SHA384
	// KeyAlgorithmED25519 represents ED25519
	KeyAlgorithmED25519
	// KeyAlgorithmED448 represents ED448
	KeyAlgorithmED448
)

// KeyFlag represents the possible types of DNSSec flags
type KeyFlag int

const (
	// KeyFlagNone means no flag is set
	KeyFlagNone KeyFlag = 0
	// KeyFlagZSK means this is a Zone Signing Key
	KeyFlagZSK KeyFlag = 256
	// KeyFlagKSK means this is a Key Signing Key
	KeyFlagKSK KeyFlag = 257
)

// KeyEntry represents a Transip_DnsSecEntry object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_DnsSecEntry.html
type KeyEntry struct {
	KeyTag    int          `xml:"keyTag"`
	Flags     KeyFlag      `xml:"flags"`
	Algorithm KeyAlgorithm `xml:"algorithm"`
	PublicKey string       `xml:"publicKey"`
}

// KeyEntries is just an array of KeyEntry
// basically only here so it can implement paramsEncoder
type KeyEntries []KeyEntry

// EncodeParams returns KeyEntries parameters ready to be used for constructing
// a signature
// the order of parameters added here has to match the order in the WSDL as
// described at http://api.transip.nl/wsdl/?service=DnsService
func (d KeyEntries) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
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

// EncodeArgs returns Entries XML body ready to be passed in the SOAP call
func (d KeyEntries) EncodeArgs(key string) string {
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
