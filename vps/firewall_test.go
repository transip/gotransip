package vps

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/ipaddress"
	"net"
	"testing"
)

func TestRepository_GetFirewall(t *testing.T) {
	const apiResponse = `{"vpsFirewall":{"isEnabled":true,"ruleSet":[{"description":"HTTP","startPort":80,"endPort":80,"protocol":"tcp","whitelist":["80.69.69.80/32","80.69.69.100/32","2a01:7c8:3:1337::1/128"]}]}}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/firewall", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	firewall, err := repo.GetFirewall("example-vps")
	require.NoError(t, err)
	assert.True(t, firewall.IsEnabled)
	require.Equal(t, 1, len(firewall.RuleSet))
	rule := firewall.RuleSet[0]

	assert.Equal(t, "HTTP", rule.Description)
	assert.Equal(t, 80, rule.StartPort)
	assert.Equal(t, 80, rule.EndPort)
	assert.Equal(t, "tcp", rule.Protocol)

	require.Equal(t, 3, len(rule.Whitelist))
	assert.EqualValues(t, "80.69.69.80/32", rule.Whitelist[0].String())
	assert.EqualValues(t, "80.69.69.100/32", rule.Whitelist[1].String())
	assert.EqualValues(t, "2a01:7c8:3:1337::1/128", rule.Whitelist[2].String())
}

func TestRepository_UpdateFirewall(t *testing.T) {
	const expectedRequest = `{"vpsFirewall":{"isEnabled":true,"ruleSet":[{"description":"HTTP","startPort":80,"endPort":80,"protocol":"tcp","whitelist":["80.69.69.80/32","80.69.69.100/32","2a01:7c8:3:1337::1/128"]}]}}`
	server := mockServer{t: t, expectedUrl: "/vps/example-vps/firewall", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	testWhitelists := []string{"80.69.69.80/32", "80.69.69.100/32", "2a01:7c8:3:1337::1/128"}
	whiteListRanges := make([]ipaddress.IPRange, len(testWhitelists))
	for idx, ipRange := range testWhitelists {
		_, ipNet, err := net.ParseCIDR(ipRange)
		require.NoError(t, err)
		whiteListRanges[idx] = ipaddress.IPRange{IPNet: *ipNet}
	}

	firewall := Firewall{
		IsEnabled: true,
		RuleSet: []FirewallRule{{
			Description: "HTTP",
			EndPort:     80,
			Protocol:    "tcp",
			StartPort:   80,
			Whitelist:   whiteListRanges,
		}},
	}

	err := repo.UpdateFirewall("example-vps", firewall)
	require.NoError(t, err)
}
