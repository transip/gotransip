package vps

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
	"github.com/transip/gotransip/v6/ipaddress"
)

func TestFirewallRepository_GetFirewall(t *testing.T) {
	const apiResponse = `{"vpsFirewall":{"isEnabled":true,"ruleSet":[{"description":"HTTP","startPort":80,"endPort":80,"protocol":"tcp","whitelist":["80.69.69.80/32","80.69.69.100/32","2a01:7c8:3:1337::1/128"]}]}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/firewall", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := FirewallRepository{Client: *client}

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

func TestFirewallRepository_UpdateFirewall(t *testing.T) {
	const expectedRequest = `{"vpsFirewall":{"isEnabled":true,"ruleSet":[{"description":"HTTP","startPort":80,"endPort":80,"protocol":"tcp","whitelist":["80.69.69.80/32","80.69.69.100/32","2a01:7c8:3:1337::1/128"]}]}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/firewall", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := FirewallRepository{Client: *client}

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
