package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6/ipaddress"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// FirewallRepository allows you to get information on the current Vps firewall
// and to update it
type FirewallRepository repository.RestRepository

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
	// The start port of this firewall rule
	StartPort int `json:"startPort"`
	// The end port of this firewall rule
	EndPort int `json:"endPort"`
	// The protocol `tcp` ,  `udp` or `tcp_udp`
	Protocol string `json:"protocol"`
	// Whitelisted IP's or ranges that are allowed to connect, empty to allow all
	Whitelist []ipaddress.IPRange `json:"whitelist"`
}

// GetFirewall returns the state of the current VPS firewall
func (r *FirewallRepository) GetFirewall(vpsName string) (Firewall, error) {
	var response firewallWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/firewall", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Firewall, err
}

// UpdateFirewall allows you to update the state of the firewall
// Enabling it, disabling it
// Adding / removing of ruleSets, updating the whitelists
func (r *FirewallRepository) UpdateFirewall(vpsName string, firewall Firewall) error {
	requestBody := firewallWrapper{Firewall: firewall}
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/firewall", vpsName), Body: &requestBody}

	return r.Client.Put(restRequest)
}
