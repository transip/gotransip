package colocation

import (
	"github.com/transip/gotransip/v6/ipaddress"
	"net"
)

// Colocation struct for a Colocation
type Colocation struct {
	// List of IP ranges
	IPRanges []ipaddress.IPRange `json:"ipRanges"`
	// Colocation name
	Name string `json:"name"`
}

// colocationsWrapper struct contains a list of Colocations in it,
// this is solely used for unmarshalling
type colocationsWrapper struct {
	// array of Colocations
	Colocations []Colocation `json:"colocations"`
}

// colocationWrapper struct contains a Colocation in it,
// this is solely used for unmarshalling
type colocationWrapper struct {
	// array of Colocations
	Colocation Colocation `json:"colocation"`
}

// ipAddressWrapper struct contains an IPAddress in it,
// this is solely used for unmarshalling
type ipAddressWrapper struct {
	IPAddress ipaddress.IPAddress `json:"ipAddress"`
}

// addIpRequest struct contains an IPAddress in it,
// this is solely used for marshalling
type addIPRequest struct {
	// The IP address to register to the colocation
	IPAddress net.IP `json:"ipAddress"`
	// Reverse DNS, also known as the PTR record
	ReverseDNS string `json:"reverseDns,omitempty"`
}
