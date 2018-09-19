package haip

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
)

func TestGetCertificatesByHaip(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getcertificatesbyhaip.xml")
	assert.NoError(t, err)

	lst, err := GetCertificatesByHaip(c, "example-haip")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []Certificate{}, lst)
	assert.Equal(t, "example.org", lst[0].CommonName)
	x, _ := time.Parse("2006-01-02 15:04:05", "2018-11-21 20:07:33")
	assert.Equal(t, x, lst[0].ExpirationDate)
	assert.Equal(t, int64(484739), lst[0].ID)
	assert.Equal(t, int64(485554), lst[1].ID)
}

func TestGetHaip(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/gethaip.xml")
	assert.NoError(t, err)

	h, err := GetHaip(c, "example-haip")
	assert.NoError(t, err)
	assert.IsType(t, Haip{}, h)
	assert.Equal(t, "example-haip", h.Name)
	assert.Equal(t, StatusActive, h.Status)
	assert.Equal(t, true, h.IsBlocked)
	assert.Equal(t, true, h.IsLoadBalancingEnabled)
	assert.Equal(t, BalancingModeRoundRobin, h.LoadBalancingMode)
	assert.Equal(t, "foobar", h.StickyCookieName)
	assert.Equal(t, HealthCheckModeTCP, h.HealthCheckMode)
	assert.Equal(t, "/", h.HTTPHealthCheckPath)
	assert.Equal(t, 8443, h.HTTPHealthCheckPort)
	assert.Equal(t, net.ParseIP("89.41.170.108"), h.IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8:e001:f00f::f00f"), h.IPv6Address)
	assert.Equal(t, IPSetup("ipv6to4"), h.IPSetup)
	assert.Equal(t, 2, len(h.AttachedVpses))
}

func TestGetHaips(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/gethaips.xml")
	assert.NoError(t, err)

	lst, err := GetHaips(c)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []Haip{}, lst)
	assert.Equal(t, "example-haip", lst[0].Name)
	assert.Equal(t, "example-haip2", lst[1].Name)
}

func TestGetPortConfigurations(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getportconfigurations.xml")
	assert.NoError(t, err)

	lst, err := GetPortConfigurations(c, "example-haip")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lst))
	assert.IsType(t, []PortConfiguration{}, lst)
	assert.Equal(t, EndpointSSLModeOff, lst[0].EndpointSSLMode)
	assert.Equal(t, int64(4740), lst[0].ID)
	assert.Equal(t, ModeTCP, lst[0].Mode)
	assert.Equal(t, "a9e05b317b2a311e893aa525400dd557", lst[0].Name)
	assert.Equal(t, int64(80), lst[0].SourcePort)
	assert.Equal(t, int64(32036), lst[0].TargetPort)
	assert.Equal(t, int64(4843), lst[1].ID)
}

func TestGetPtrForHaip(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getptrforhaip.xml")
	assert.NoError(t, err)

	ptr, err := GetPtrForHaip(c, "example-haip")
	assert.NoError(t, err)
	assert.Equal(t, "haip.example.net", ptr)
}
