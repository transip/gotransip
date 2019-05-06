package dns

import (
	"github.com/transip/gotransip"
)

// This file holds all DnsService methods directly ported from TransIP API

// SetDNSEntries sets the DNS entries for this domain, will replace all existing DNS
// entries with the new entries
func SetDNSEntries(c gotransip.Client, domainName string, dnsEntries Entries) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setDnsEntries",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("dnsEntries", dnsEntries)

	return c.Call(sr, nil)
}

// CanEditDNSSec checks if the DNSSec entries of a domain can be updated.
func CanEditDNSSec(c gotransip.Client, domainName string) (bool, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "canEditDnsSec",
	}
	sr.AddArgument("domainName", domainName)

	var v bool
	err := c.Call(sr, &v)
	return v, err
}

// GetDNSSecEntries returns DNSSec entries for given domain name
func GetDNSSecEntries(c gotransip.Client, domainName string) (KeyEntries, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getDnsSecEntries",
	}
	sr.AddArgument("domainName", domainName)

	var v struct {
		V KeyEntries `xml:"item"`
	}
	err := c.Call(sr, &v)
	return v.V, err
}

// SetDNSSecEntries sets new DNSSec entries for a domain, replacing the current ones.
func SetDNSSecEntries(c gotransip.Client, domainName string, dnssecKeyEntrySet KeyEntries) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setDnsSecEntries",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("dnssecKeyEntrySet", dnssecKeyEntrySet)

	return c.Call(sr, nil)
}

// RemoveAllDNSSecEntries removes all the DNSSec entries from a domain.
func RemoveAllDNSSecEntries(c gotransip.Client, domainName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "removeAllDnsSecEntries",
	}
	sr.AddArgument("domainName", domainName)

	return c.Call(sr, nil)
}
