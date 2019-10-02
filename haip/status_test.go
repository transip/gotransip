package haip

import (
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseStatusReportBody(t *testing.T) {

	// read example status report output
	data, err := ioutil.ReadFile("testdata/statusreportbody.xml")
	require.NoError(t, err)

	// parse status report body
	sr, err := parseStatusReportBody(data)
	require.NoError(t, err)

	// go over returned body and see if all essential structs are there
	assert.Equal(t, 2, len(sr.PortConfiguration))
	assert.Equal(t, 80, sr.PortConfiguration[0].Port)
	assert.Equal(t, 443, sr.PortConfiguration[1].Port)
	assert.Equal(t, 2, len(sr.PortConfiguration[0].VPS))
	assert.Equal(t, "example-vps", sr.PortConfiguration[0].VPS[0].Name)
	assert.Equal(t, "example-vps2", sr.PortConfiguration[0].VPS[1].Name)
	assert.Equal(t, 2, len(sr.PortConfiguration[0].VPS[0].IPVersion))
	assert.Equal(t, "ipv4", sr.PortConfiguration[0].VPS[0].IPVersion[0].Version)
	assert.Equal(t, "ipv6", sr.PortConfiguration[0].VPS[0].IPVersion[1].Version)
	assert.Equal(t, 2, len(sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer))
	assert.Equal(t, "lb0.ams0", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].Name)
	assert.Equal(t, "lb0.rtm0", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[1].Name)
	assert.Equal(t, net.IP{1, 2, 3, 4}, sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].IPAddress)
	assert.Equal(t, "up", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].State)
	assert.Equal(t, time.Unix(1535653953, 0), sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].LastChange)
}
