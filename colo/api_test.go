package colo

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
)

func TestGetColoNames(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getcolonames.xml")
	assert.NoError(t, err)

	lst, err := GetColoNames(c)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []string{}, lst)
	assert.Equal(t, "example", lst[0])
	assert.Equal(t, "example2", lst[1])
}

func TestGetIPAddresses(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getipaddresses.xml")
	assert.NoError(t, err)

	lst, err := GetIPAddresses(c, "example")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []net.IP{}, lst)
	assert.Equal(t, net.IP{1, 2, 3, 4}.To16(), lst[0])
	assert.Equal(t, net.IP{2, 3, 4, 5}.To16(), lst[1])
}

func TestGetIPRanges(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getipranges.xml")
	assert.NoError(t, err)

	lst, err := GetIPRanges(c, "example")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []net.IPNet{}, lst)
	_, x, _ := net.ParseCIDR("1.2.3.0/25")
	assert.Equal(t, *x, lst[0])
	_, x, _ = net.ParseCIDR("2a01:7c8::/64")
	assert.Equal(t, *x, lst[1])
}

func TestGetReverseDNS(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getreversedns.xml")
	assert.NoError(t, err)

	ptr, err := GetReverseDNS(c, net.IP{1, 2, 3, 4})
	assert.NoError(t, err)
	assert.Equal(t, "example.org", ptr)
}
