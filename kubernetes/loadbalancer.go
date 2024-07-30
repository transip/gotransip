package kubernetes

import (
	"net"

	"github.com/transip/gotransip/v6/rest"
)

// lbsWrapper struct contains a list of LoadBalancers in it,
// this is solely used for unmarshalling/marshalling
type lbsWrapper struct {
	LoadBalancers []LoadBalancer `json:"loadBalancers"`
}

// lbWrapper struct contains a LoadBalancer in it,
// this is solely used for unmarshalling/marshalling
type lbWrapper struct {
	LoadBalancer LoadBalancer `json:"loadBalancer"`
}

// lbOrder struct is used to wrap creating a new LoadBalancer
type lbOrder struct {
	Name string `json:"name"`
}

// lbcWrapper struct is used to wrap configuring a LoadBalancer
type lbcWrapper struct {
	Config LoadBalancerConfig `json:"loadBalancerConfig"`
}

// LoadBalancer struct for a kubernetes loadbalancer
type LoadBalancer struct {
	// The unique identifier for the loadbalancer
	UUID string `json:"uuid"`
	// User configurable unique identifier (max 64 chars)
	Name string `json:"name"`
	// LoadBalancer status, either ‘active’, ‘inactive’, ‘creating’
	Status LoadBalancerStatus `json:"status"`

	// HA-IP IPv4 address
	IPv4Address net.IP `json:"ipv4Address,omitempty"`
	// HA-IP IPv6 address
	IPv6Address net.IP `json:"ipv6Address,omitempty"`
}

// LoadBalancerStatus status, either ‘active’, ‘inactive’, ‘creating’
type LoadBalancerStatus string

// Definition of all of the possible loadbalancer statuses
const (
	// LoadBalancerStatusActive means the loadbalancer is active
	LoadBalancerStatusActive LoadBalancerStatus = "active"
	// LoadBalancerStatusInactive means the load balancer is inactive
	LoadBalancerStatusInactive LoadBalancerStatus = "inactive"
	// LoadBalancerStatusCreating means the load balancer is being created
	LoadBalancerStatusCreating LoadBalancerStatus = "creating"
)

// LoadBalancerConfig is a representation of all the options that can be configured for a Load Balancer
type LoadBalancerConfig struct {
	// HA-IP load balancing mode: 'roundrobin', 'cookie', 'source'
	LoadBalancingMode LoadBalancingMode `json:"loadBalancingMode,omitempty"`
	// Cookie name to pin sessions on when using cookie balancing mode
	StickyCookieName string `json:"stickyCookieName,omitempty"`
	// The interval in milliseconds at which health checks are performed. The interval may not be smaller than 2000ms.
	HealthCheckInterval int64 `json:"healthCheckInterval,omitempty"`
	// The path (URI) of the page to check HTTP status code on
	HTTPHealthCheckPath string `json:"httpHealthCheckPath,omitempty"`
	// The port to perform the HTTP check on
	HTTPHealthCheckPort int `json:"httpHealthCheckPort,omitempty"`
	// Whether to use SSL when performing the HTTP check
	HTTPHealthCheckSSL bool `json:"httpHealthCheckSsl"`
	// HA-IP IP setup: 'both', 'noipv6', 'ipv6to4', 'ipv4to6'
	IPSetup IPSetup `json:"ipSetup"`
	// The PTR record for the HA-IP
	PTRRecord string `json:"ptrRecord,omitempty"`
	// HA-IP TLS Mode: 'tls10_11_12', 'tls11_12', 'tls12'
	TLSMode TLSMode `json:"tlsMode"`
	// The IPs attached to this haip
	IPAddresses []net.IP `json:"ipAddresses,omitempty"`
	// Array with port configurations for this LoadBalancer
	PortConfigurations []PortConfiguration `json:"portConfiguration,omitempty"`
}

// LoadBalancingMode is one of the following strings
// 'roundrobin', 'cookie', 'source'
type LoadBalancingMode string

// Definition of all of the possible load balancing modes
const (
	// LoadBalancingModeRoundRobin is the LoadBalancing mode roundrobin for a Haip, forward to next address everytime
	LoadBalancingModeRoundRobin LoadBalancingMode = "roundrobin"
	// LoadBalancingModeCookie is the LoadBalancing mode cookie for a Haip, forward to a fixed server, based on the cookie
	LoadBalancingModeCookie LoadBalancingMode = "cookie"
	// LoadBalancingModeSource is the LoadBalancing mode source for a Haip, choose a server to forward based on the source address
	LoadBalancingModeSource LoadBalancingMode = "source"
)

// IPSetup is one of the following strings
// 'both', 'noipv6', 'ipv6to4', 'ipv4to6'
type IPSetup string

// Definition of all of the possible ip setup options
const (
	// IPSetupBoth accept ipv4 and ipv6 and forward them to separate ipv4 and ipv6 addresses
	IPSetupBoth IPSetup = "both"
	// IPSetupNoIPv6 do not accept ipv6 traffic
	IPSetupNoIPv6 IPSetup = "noipv6"
	// IPSetupIPv6to4 forward ipv6 traffic to ipv4
	IPSetupIPv6to4 IPSetup = "ipv6to4"
	// IPSetupIPv4to6 forward ipv4 traffic to ipv6
	IPSetupIPv4to6 IPSetup = "ipv4to6"
)

// TLSMode is one of the following strings
// 'tls10_11_12', 'tls11_12', 'tls12'
type TLSMode string

// Definition of all of the possible tls mode options
const (
	// TLSModeMinTLS10 only allow incoming tls traffic with versions 1.0, 1.1 and 1.2
	TLSModeMinTLS10 TLSMode = "tls10_11_12"
	// TLSModeMinTLS11 only allow incoming tls traffic with version 1.1 or 1.2
	TLSModeMinTLS11 TLSMode = "tls11_12"
	// TLSModeMinTLS12 only allow incoming traffic with tls version 1.2
	TLSModeMinTLS12 TLSMode = "tls12"
)

// PortConfiguration struct for a PortConfiguration
type PortConfiguration struct {
	// A name describing the port
	Name string `json:"name"`
	// The port at which traffic arrives on your HA-IP
	SourcePort int `json:"sourcePort"`
	// The port at which traffic arrives on your attached IP address(es)
	TargetPort int `json:"targetPort"`
	// The mode determining how traffic is processed and forwarded: 'tcp', 'http', 'https', 'proxy', 'http2_https'
	Mode PortConfigurationMode `json:"mode"`
	// The mode determining how traffic between our load balancers and your attached IP address(es) is encrypted: 'off', 'on', 'strict'
	EndpointSSLMode PortConfigurationEndpointSSLMode `json:"endpointSslMode"`
}

// LoadBalancerStatusReport A status report for the laodbalancer
type LoadBalancerStatusReport struct {
	NodeUUID         string            `json:"nodeUuid"`
	NodeIPAddress    net.IP            `json:"nodeIpAddress"`
	Port             int               `json:"port"`
	IPVersion        int               `json:"ipVersion"`
	LoadBalancerName string            `json:"loadBalancerName"`
	LoadBalancerIP   net.IP            `json:"loadBalancerIp"`
	State            LoadBalancerState `json:"state"`
	LastChange       rest.Time         `json:"lastChange"`
}

// LoadBalancerState the state of the connection from the loadbalancer to the node
type LoadBalancerState string

const (
	// LoadBalancerStateUp the connection from the loadbalanacer to the node is up
	LoadBalancerStateUp LoadBalancerState = "up"
	// LoadBalancerStateDown the connection from the loadbalanacer to the node is down
	LoadBalancerStateDown LoadBalancerState = "down"
)

type loadBalancerStatusReportsWrapper struct {
	StatusReports []LoadBalancerStatusReport `json:"statusReports"`
}

// PortConfigurationMode is one of the following strings
// 'tcp', 'http', 'https', 'proxy', 'http2_https'
type PortConfigurationMode string

// Definition of all of the possible port configuration modes
const (
	// PortConfigurationModeTCP plain TCP forward to your VPS
	PortConfigurationModeTCP PortConfigurationMode = "tcp"
	// PortConfigurationModeHTTP appends a X-Forwarded-For header to HTTP requests with the original remote IP
	PortConfigurationModeHTTP PortConfigurationMode = "http"
	// PortConfigurationModeHTTPS same as HTTP, with SSL Certificate offloading
	PortConfigurationModeHTTPS PortConfigurationMode = "https"
	// PortConfigurationModePROXY proxy protocol is also a way to retain the original remote IP,
	// but also works for non HTTP traffic (note: the receiving application has to support this)
	PortConfigurationModePROXY PortConfigurationMode = "proxy"
	// PortConfigurationModeHTTP2HTTPS same as HTTPS, with http/2 support
	PortConfigurationModeHTTP2HTTPS PortConfigurationMode = "http2_https"
)

// PortConfigurationEndpointSSLMode is one of the following values
// 'off', 'on', 'strict'
type PortConfigurationEndpointSSLMode string

// Definition of all the possible port configuration endpoint SSL modes
const (
	// PortConfigurationEndpointSSLModeOff means the traffic to the backends is unencrypted
	PortConfigurationEndpointSSLModeOff PortConfigurationEndpointSSLMode = "off"
	// PortConfigurationEndpointSSLModeOn means the traffic to the backends is encrypted but not verified
	PortConfigurationEndpointSSLModeOn PortConfigurationEndpointSSLMode = "on"
	// PortConfigurationEndpointSSLModeStrict means the traffic to the backends is encrypted and verified
	PortConfigurationEndpointSSLModeStrict PortConfigurationEndpointSSLMode = "strict"
)
