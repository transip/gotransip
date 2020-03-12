package ipaddress

import "net"

// IPRange is a wrapper around IPNet as it has no default marshallers
type IPRange struct {
	net.IPNet
}

// MarshalText MarshalText implements the encoding.TextMarshaler interface.
// The encoding is the same as returned by net.IPNet.String
func (r *IPRange) MarshalText() ([]byte, error) {
	return []byte(r.IPNet.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The IPRange is expected in a form accepted by net.ParseCIDR function
// e.g.: '192.168.0.1/24' or '2a01::1/48' for IPv6
func (r *IPRange) UnmarshalText(input []byte) error {
	_, ipNet, err := net.ParseCIDR(string(input))
	if err != nil {
		return err
	}

	r.IPNet = *ipNet
	return nil
}
