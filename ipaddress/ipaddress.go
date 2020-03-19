package ipaddress

import (
	"net"
)

// IPAddress struct for an IPAddress
type IPAddress struct {
	// The IP address
	Address net.IP `json:"address"`
	// The TransIP DNS resolvers you can use
	DNSResolvers []net.IP `json:"dnsResolvers,omitempty"`
	// Gateway
	Gateway net.IP `json:"gateway,omitempty"`
	// Reverse DNS, also known as the PTR record
	ReverseDNS string `json:"reverseDns"`
	// Subnet mask
	SubnetMask SubnetMask `json:"subnetMask,omitempty"`
}

// IPAddressesWrapper struct wraps an IPAddress struct,
// this is mainly used in other subpackages that need to unmarshal a ipAddresses: [] server response
type IPAddressesWrapper struct {
	// array of IP Addresses
	IPAddresses []IPAddress `json:"ipAddresses,omitempty"`
}
