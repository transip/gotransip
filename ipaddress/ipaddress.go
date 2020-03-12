package ipaddress

import (
	"net"
)

// IPAddress struct for IPAddress
type IPAddress struct {
	// The IP address
	Address net.IP `json:"address"`
	// The TransIP DNS resolvers you can use
	DnsResolvers []net.IP `json:"dnsResolvers,omitempty"`
	// Gateway
	Gateway net.IP `json:"gateway,omitempty"`
	// Reverse DNS, also known as the PTR record
	ReverseDns string `json:"reverseDns"`
	// Subnet mask
	SubnetMask SubnetMask `json:"subnetMask,omitempty"`
}

// IPAddressesWrapper struct for IPAddressesWrapper
type IPAddressesWrapper struct {
	// array of IP Addresses
	IPAddresses []IPAddress `json:"ipAddresses,omitempty"`
}
