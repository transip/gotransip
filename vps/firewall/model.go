package firewall

// Firewall struct for the Vps Firewall
type Firewall struct {
	// Whether the firewall is enabled for this VPS
	IsEnabled bool `json:"isEnabled"`
	// Ruleset of the VPS
	RuleSet []FirewallRule `json:"ruleSet"`
}

// FirewallRule struct for a VpsFirewallRule
type FirewallRule struct {
	// The rule name
	Description string `json:"description,omitempty"`
	// The end port of this firewall rule
	EndPort float32 `json:"endPort"`
	// The protocol `tcp` ,  `udp` or `tcp_udp`
	Protocol string `json:"protocol"`
	// The start port of this firewall rule
	StartPort float32 `json:"startPort"`
	// Whitelisted IP's or ranges that are allowed to connect, empty to allow all
	Whitelist []map[string]interface{} `json:"whitelist"`
}
