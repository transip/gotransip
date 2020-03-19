package ipaddress

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

// SubnetMask is a wrapper around net.IPMask.
// This is needed as the transip api returns subnetmask strings that are not unmarshallable by default.
// So we need to do a little bit of magic to unmarshal either '255.255.255.0' and '/48' as net.IPMask type
type SubnetMask struct {
	IPType int
	net.IPMask
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The SubnetMask is expected in a form accepted by parseIPv6Mask or net.ParseIP
// e.g.: '255.255.255.0' or '/48' for IPv6
func (mask *SubnetMask) UnmarshalText(input []byte) error {
	if len(input) < 2 {
		return errors.New("Subnet mask cannot be empty")
	}

	// if the string starts with a '/' we generate an ipv6 mask from the given prefix length,
	// parse it and return it
	if input[0] == '/' {
		parsedMask, err := parseIPv6Mask(string(input[1:]))
		mask.IPMask = parsedMask
		if err == nil {
			mask.IPType = 6
		}
		return err
	}

	// ipv4 mask is in the following form '255.255.255.0'
	ip := net.ParseIP(string(input))
	if ip == nil {
		return &net.ParseError{Type: "IP mask", Text: string(input)}
	}

	mask.IPMask = net.IPMask(ip)
	mask.IPType = 4

	return nil
}

// MarshalText MarshalText implements the encoding.TextMarshaler interface.
// The encoding is the same as returned by net.IPMask.String
func (mask *SubnetMask) MarshalText() ([]byte, error) {
	// if first leading bit is set in an ipv6 bitmask
	if mask.IPType == 6 {
		return mask.marshalIPv6Mask()
	}

	ip := net.IP(mask.IPMask)

	// when ip mask is unset, return an empty string
	if ip == nil {
		return []byte(""), nil
	}

	return []byte(ip.String()), nil
}

// This method returns a prefix '/48' for the amount of leading bits in the given SubnetMask
func (mask *SubnetMask) marshalIPv6Mask() ([]byte, error) {
	size, total := mask.IPMask.Size()

	// check on max ipv6 mask length
	if total > (net.IPv6len * 8) {
		return nil, errors.New("IPv6 mask can't be bigger than 128 bits")
	}

	prefix := fmt.Sprintf("/%d", size)

	return []byte(prefix), nil
}

// parseIPv6Mask takes a prefix length as string, decodes it and returns a ipv6 IPMask generated from the prefixLength
// using net.CIDRMask
func parseIPv6Mask(input string) (net.IPMask, error) {
	prefixLength, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}
	if prefixLength < 0 || prefixLength > 128 {
		return nil, errors.New("Invalid prefixLength provided")
	}

	return net.CIDRMask(prefixLength, 8*net.IPv6len), nil
}
