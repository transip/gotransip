package ipaddress

import "net"

// IPAddress struct for IPAddress
type IPAddress struct {
	// The IP address
	Address net.IP `json:"address,omitempty"`
	// The TransIP DNS resolvers you can use
	DnsResolvers []string `json:"dnsResolvers,omitempty"`
	// Gateway
	Gateway net.IP `json:"gateway,omitempty"`
	// Reverse DNS, also known as the PTR record
	ReverseDns string `json:"reverseDns"`
	// Subnet mask
	SubnetMask net.IPMask `json:"subnetMask,omitempty"`
}

// IPAddresses struct for IPAddresses
type IPAddresses struct {
	// array of IP Addresses
	IPAddresses []IPAddress `json:"ipAddresses,omitempty"`
}
