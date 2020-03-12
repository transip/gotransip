package vps

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

type expect struct {
	in     []byte
	out    []byte
	result string
	error  error
}

func TestIPRange_MarshalText(t *testing.T) {
	expects := []expect{
		{in: []byte("2a01::1/48"), result: "2a01::/48"},
		{in: []byte("2a01::1/128"), result: "2a01::1/128"},
		{in: []byte("48"), error: &net.ParseError{Type:"CIDR address", Text:"48"}},
		{in: []byte("/129"), error: &net.ParseError{Type:"CIDR address", Text:"/129"}},
		{in: []byte("255.255.255.0"), error: &net.ParseError{Type:"CIDR address", Text:"255.255.255.0"}},
		{in: []byte("255.255.255.d"), error: &net.ParseError{Type:"CIDR address", Text:"255.255.255.d"}},
		{in: []byte("255.0.255.0/24"), result: "255.0.255.0/24"},
		{in: []byte("255.0.128.0/32"), result: "255.0.128.0/32"},
		{in: []byte("255.0.128.0/33"), error: &net.ParseError{Type:"CIDR address", Text:"255.0.128.0/33"}},
		{in: []byte(""), error: &net.ParseError{Type:"CIDR address", Text:""}},
	}

	for _, expect := range expects {
		ipRange := IPRange{}
		err := ipRange.UnmarshalText(expect.in)

		assert.Equal(t, expect.error, err)
		if expect.error == nil {
			assert.Equal(t, expect.result, ipRange.IPNet.String())
		}
	}
}

func TestIPRange_UnmarshalText(t *testing.T) {
	expects := []expect{
		{in: []byte("2a01::1/48"), result: "2a01::/48"},
		{in: []byte("2a01::1/128"), result: "2a01::1/128"},
		{in: []byte("255.0.255.0/24"), result: "255.0.255.0/24"},
		{in: []byte("255.0.128.0/32"), result: "255.0.128.0/32"},
		{in: []byte("192.168.0.1/24"), result: "192.168.0.0/24"},
	}

	for _, expect := range expects {
		_, ipNet, err := net.ParseCIDR(string(expect.in))
		require.NoError(t, err)
		ipRange := IPRange{IPNet: *ipNet}

		result, err := ipRange.MarshalText()

		assert.Equal(t, expect.error, err)
		assert.Equal(t, expect.result, string(result))
	}
}
