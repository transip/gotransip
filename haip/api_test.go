package haip

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v5"
)

func TestGetCertificatesByHaip(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getcertificatesbyhaip.xml")
	require.NoError(t, err)

	lst, err := GetCertificatesByHaip(c, "example-haip")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Certificate{}, lst)
	assert.Equal(t, "example.org", lst[0].CommonName)
	x, _ := time.Parse("2006-01-02 15:04:05", "2018-11-21 20:07:33")
	assert.Equal(t, x, lst[0].ExpirationDate)
	assert.Equal(t, int64(484739), lst[0].ID)
	assert.Equal(t, int64(485554), lst[1].ID)
}

func TestGetAvailableCertificatesByHaip(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailablecertificatesbyhaip.xml")
	require.NoError(t, err)

	lst, err := GetAvailableCertificatesByHaip(c, "example-haip")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Certificate{}, lst)
	assert.Equal(t, "example.com", lst[0].CommonName)
	x, _ := time.Parse("2006-01-02 15:04:05", "2019-11-21 20:07:33")
	assert.Equal(t, x, lst[0].ExpirationDate)
	assert.Equal(t, int64(487739), lst[0].ID)
	assert.Equal(t, int64(495554), lst[1].ID)
}

func TestGetHaip(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/gethaip.xml")
	require.NoError(t, err)

	h, err := GetHaip(c, "example-haip")
	require.NoError(t, err)
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
	require.NoError(t, err)

	lst, err := GetHaips(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Haip{}, lst)
	assert.Equal(t, "example-haip", lst[0].Name)
	assert.Equal(t, "example-haip2", lst[1].Name)
}

func TestGetPortConfigurations(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getportconfigurations.xml")
	require.NoError(t, err)

	lst, err := GetPortConfigurations(c, "example-haip")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
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
	require.NoError(t, err)

	ptr, err := GetPtrForHaip(c, "example-haip")
	require.NoError(t, err)
	assert.Equal(t, "haip.example.net", ptr)
}

func TestAddCertificateToHaip(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/addcertificatetohaip.xml")
	require.NoError(t, err)

	err = AddCertificateToHaip(c, "transip-haip", 1234)
	require.NoError(t, err)

	err = AddCertificateFromHaip(c, "transip-haip", 1234)
	require.NoError(t, err)
}

func TestAddPortConfiguration(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/addportconfiguration.xml")
	require.NoError(t, err)

	err = AddPortConfiguration(c, "transip-haip", PortConfiguration{Name: "http"})
	require.NoError(t, err)
}

func TestCancelHaip(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/cancelhaip.xml")
	require.NoError(t, err)

	err = CancelHaip(c, "transip-haip", gotransip.CancellationTimeImmediately)
	require.NoError(t, err)
}

func TestDeleteCertificateFromHaip(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/deletecertificatefromhaip.xml")
	require.NoError(t, err)

	err = DeleteCertificateFromHaip(c, "transip-haip", 1234)
	require.NoError(t, err)
}

func TestDeletePortConfiguration(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/deleteportconfiguration.xml")
	require.NoError(t, err)

	err = DeletePortConfiguration(c, "transip-haip", 1234)
	require.NoError(t, err)
}

func TestSetBalancingMode(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setbalancingmode.xml")
	require.NoError(t, err)

	err = SetBalancingMode(c, "transip-haip", BalancingModeSource, "")
	require.NoError(t, err)
}

func TestSetDefaultPortConfiguration(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setdefaultportconfiguration.xml")
	require.NoError(t, err)

	err = SetDefaultPortConfiguration(c, "transip-haip")
	require.NoError(t, err)
}

func TestSetHaipDescription(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/sethaipdescription.xml")
	require.NoError(t, err)

	err = SetHaipDescription(c, "transip-haip", "My Highly Available IP Address")
	require.NoError(t, err)
}

func TestSetHaipVpses(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/sethaipvpses.xml")
	require.NoError(t, err)

	err = SetHaipVpses(c, "transip-haip", []string{"transip-vps", "transip-vps2", "transip-vps3"})
	require.NoError(t, err)
}

func TestSetHTTPHealthCheck(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/sethttphealthcheck.xml")
	require.NoError(t, err)

	err = SetHTTPHealthCheck(c, "transip-haip", "/healthz", 8080)
	require.NoError(t, err)
}

func TestSetIPSetup(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setipsetup.xml")
	require.NoError(t, err)

	err = SetIPSetup(c, "transip-haip", IPSetupIPv4To6)
	require.NoError(t, err)
}

func TestSetPtrForHaip(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setptrforhaip.xml")
	require.NoError(t, err)

	err = SetPtrForHaip(c, "transip-haip", "www.example.org")
	require.NoError(t, err)
}

func TestSetTCPHealthCheck(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/settcphealthcheck.xml")
	require.NoError(t, err)

	err = SetTCPHealthCheck(c, "transip-haip")
	require.NoError(t, err)
}

func TestStartHaipLetsEncryptCertificateIssue(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/starthaipletsencryptcertificateissue.xml")
	require.NoError(t, err)

	err = StartHaipLetsEncryptCertificateIssue(c, "transip-haip", "www.example.org")
	require.NoError(t, err)
}

func TestUpdatePortConfiguration(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/updateportconfiguration.xml")
	require.NoError(t, err)

	err = UpdatePortConfiguration(c, "transip-haip", PortConfiguration{Name: "http", TargetPort: 8080})
	require.NoError(t, err)
}

func TestGetStatusReport(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getstatusreport.xml")
	require.NoError(t, err)

	sr, err := GetStatusReport(c, "example-haip")
	require.NoError(t, err)

	// go over returned body and see if all essential structs are there
	require.Equal(t, 2, len(sr.PortConfiguration))
	assert.Equal(t, 80, sr.PortConfiguration[0].Port)
	assert.Equal(t, 443, sr.PortConfiguration[1].Port)
	require.Equal(t, 2, len(sr.PortConfiguration[0].VPS))
	assert.Equal(t, "example-vps", sr.PortConfiguration[0].VPS[0].Name)
	assert.Equal(t, "example-vps2", sr.PortConfiguration[0].VPS[1].Name)
	require.Equal(t, 2, len(sr.PortConfiguration[0].VPS[0].IPVersion))
	assert.Equal(t, "ipv4", sr.PortConfiguration[0].VPS[0].IPVersion[0].Version)
	assert.Equal(t, "ipv6", sr.PortConfiguration[0].VPS[0].IPVersion[1].Version)
	require.Equal(t, 2, len(sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer))
	assert.Equal(t, "lb0.ams0", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].Name)
	assert.Equal(t, "lb0.rtm0", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[1].Name)
	assert.Equal(t, net.IP{1, 2, 3, 4}, sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].IPAddress)
	assert.Equal(t, "up", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].State)
	assert.Equal(t, time.Unix(1535653953, 0), sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].LastChange)
}
