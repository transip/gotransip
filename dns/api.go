package dns

import (
	"github.com/transip/gotransip"
)

// This file holds all DnsService methods directly ported from TransIP API

// SetDnsEntries sets the Dns entries for this domain, will replace all existing dns
// entries with the new entries
func SetDnsEntries(c gotransip.Client, domainName string, dnsEntries DNSEntries) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setDnsEntries",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("dnsEntries", dnsEntries)

	return c.Call(sr, nil)
}

// CanEditDnsSec checks if the DNSSec entries of a domain can be updated.
func CanEditDnsSec(c gotransip.Client, domainName string) (bool, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "canEditDnsSec",
	}
	sr.AddArgument("domainName", domainName)

	var v bool
	err := c.Call(sr, &v)
	return v, err
}

// GetDnsSecEntries returns DNSSec entries for given domain name
func GetDnsSecEntries(c gotransip.Client, domainName string) (DNSSecEntries, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getDnsSecEntries",
	}
	sr.AddArgument("domainName", domainName)

	var v struct {
		V DNSSecEntries `xml:"item"`
	}
	err := c.Call(sr, &v)
	return v.V, err
}

// SetDnsSecEntries sets new DNSSec entries for a domain, replacing the current ones.
func SetDnsSecEntries(c gotransip.Client, domainName string, dnssecKeyEntrySet DNSSecEntries) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setDnsSecEntries",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("dnssecKeyEntrySet", dnssecKeyEntrySet)

	return c.Call(sr, nil)
}

// RemoveAllDnsSecEntries removes all the DNSSec entries from a domain.
func RemoveAllDnsSecEntries(c gotransip.Client, domainName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "removeAllDnsSecEntries",
	}
	sr.AddArgument("domainName", domainName)

	return c.Call(sr, nil)
}
