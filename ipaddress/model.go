package ipaddress

import "net"

// IpAddress struct for IpAddress
type IpAddress struct {
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

// IpAddresses struct for IpAddresses
type IpAddresses struct {
	// array of IP Addresses
	IpAddresses []IpAddress `json:"ipAddresses,omitempty"`
}
