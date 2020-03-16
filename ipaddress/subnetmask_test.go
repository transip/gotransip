package ipaddress

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestSubnetMask_UnmarshalText(t *testing.T) {
	expects := []expect{
		{in: []byte("/48"), result: "ffffffffffff00000000000000000000", error: nil},
		{in: []byte("/128"), result: "ffffffffffffffffffffffffffffffff", error: nil},
		{in: []byte("/1"), result: "80000000000000000000000000000000", error: nil},
		{in: []byte("/0"), result: "00000000000000000000000000000000", error: nil},
		{in: []byte("48"), result: "<nil>", error: &net.ParseError{Type: "IP mask", Text: "48"}},
		{in: []byte("/129"), result: "<nil>", error: errors.New("Invalid prefixLength provided")},
		{in: []byte("255.255.255.0"), result: "00000000000000000000ffffffffff00", error: nil},
		{in: []byte("255.255.255.d"), result: "<nil>", error: &net.ParseError{Type: "IP mask", Text: "255.255.255.d"}},
		{in: []byte("255.0.255.0"), result: "00000000000000000000ffffff00ff00", error: nil},
		{in: []byte("255.0.128.0"), result: "00000000000000000000ffffff008000", error: nil},
	}

	for _, expect := range expects {
		mask := SubnetMask{}
		err := mask.UnmarshalText(expect.in)

		assert.Equal(t, expect.error, err)
		assert.Equal(t, expect.result, mask.IPMask.String())
	}
}

func TestSubnetMask_MarshalText(t *testing.T) {
	expects := []expect{
		{in: []byte("/48"), out: []byte("/48"), error: nil},
		{in: []byte("/128"), out: []byte("/128"), error: nil},
		{in: []byte("/1"), out: []byte("/1"), error: nil},
		{in: []byte("/0"), out: []byte("/0"), error: nil},
		{in: []byte("48"), out: []byte(""), error: nil},
		{in: []byte("/129"), out: []byte(""), error: nil},
		{in: []byte("255.255.255.0"), out: []byte("255.255.255.0"), error: nil},
		{in: []byte("255.255.255.d"), out: []byte(""), error: nil},
		{in: []byte("255.0.255.0"), out: []byte("255.0.255.0"), error: nil},
		{in: []byte("255.0.128.0"), out: []byte("255.0.128.0"), error: nil},
	}

	for _, expect := range expects {
		mask := SubnetMask{}
		// lets get some data, empty masks are also allowed
		mask.UnmarshalText(expect.in)

		out, err := mask.MarshalText()

		assert.Equal(t, expect.error, err)
		require.Equal(t, expect.out, out, fmt.Sprintf("Expecting output: '%s' for '%s'", string(expect.out), string(expect.in)))
	}
}

func Test_InvalidPrefixLength(t *testing.T) {
	mask := net.IPMask([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	subnetMask := SubnetMask{IPType: 6, IPMask: mask}
	_, err := subnetMask.MarshalText()

	require.Error(t, err)
	assert.Equal(t, err.Error(), "IPv6 mask can't be bigger than 128 bits")
}

func Test_parseIPv6Mask(t *testing.T) {
	_, err := parseIPv6Mask("d")
	assert.Error(t, err, "Error should be returned when no integer is provided")

	mask, err := parseIPv6Mask("48")
	assert.NoError(t, err)
	assert.Equal(t, "ffffffffffff00000000000000000000", mask.String())

	mask, err = parseIPv6Mask("128")
	assert.NoError(t, err)
	assert.Equal(t, "ffffffffffffffffffffffffffffffff", mask.String())

	mask, err = parseIPv6Mask("-1")
	assert.Error(t, err)
	assert.Nil(t, mask)

	mask, err = parseIPv6Mask("129")
	assert.Error(t, err)
	assert.Nil(t, mask)
}
