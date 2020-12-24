package haip

import (
	"github.com/transip/gotransip/v6/rest"
	"net"
)

// Status is one of the following strings
// 'active', 'inactive', 'creating'
type Status string

// Definition of all of the possible haip statuses
const (
	// HaipStatusActive is the status field for an active Haip, ready to use
	HaipStatusActive Status = "active"
	// HaipStatusInactive is the status field for an inactive Haip, not usable, please contact support
	HaipStatusInactive Status = "inactive"
	// HaipStatusCreating is the status field for a Haip that is being created
	HaipStatusCreating Status = "creating"
)

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
// 'both', 'noipv6', 'ipv6to4'
type IPSetup string

// Definition of all of the possible ip setup options
const (
	// IPSetupBoth accept ipv4 and ipv6 and forward them to separate ipv4 and ipv6 addresses
	IPSetupBoth IPSetup = "both"
	// IPSetupNoIPv6 do not accept ipv6 traffic
	IPSetupNoIPv6 IPSetup = "noipv6"
	// IPSetupIPv6to4 forward ipv6 traffic to ipv4
	IPSetupIPv6to4 IPSetup = "ipv6to4"
)

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

// haipsWrapper is a wrapper used to unpack the server response
// it contains a list of haips
type haipsWrapper struct {
	// list of HA-IPs
	Haips []Haip `json:"haips,omitempty"`
}

// haipWrapper struct contains a haip in it,
// this is solely used for unmarshalling/marshalling
type haipWrapper struct {
	Haip Haip `json:"haip"`
}

// certificatesWrapper is a wrapper used to unpack the server response
// it contains a list of haip certificates in it
type certificatesWrapper struct {
	// list of HA-IPs
	Certificates []Certificate `json:"certificates"`
}

// haipOrderWrapper is used to marshal an order request for a new Haip
type haipOrderWrapper struct {
	ProductName string `json:"productName"`
	Description string `json:"description,omitempty"`
}

// addCertificateRequest is used to marshal an add certificateRequest on a Haip
// this can either be an existing certificate or a to be ordered lets encrypt certificate
type addCertificateRequest struct {
	CommonName       string `json:"commonName,omitempty"`
	SslCertificateID int64  `json:"sslCertificateId,omitempty"`
}

// ipAddressesWrapper will be used to marshal/unmarshal ipAddresses that are or will be attached to a Haip
type ipAddressesWrapper struct {
	// list of IP addresses
	IPAddresses []net.IP `json:"ipAddresses"`
}

// portConfigurationsWrapper will be used to unmarshal PortConfigurations currently on a Haip
type portConfigurationsWrapper struct {
	// list of HA-IP port configurations
	PortConfigurations []PortConfiguration `json:"portConfigurations"`
}

// portConfigurationWrapper will be used to marshal/unmarshal a Configuration
type portConfigurationWrapper struct {
	Configuration PortConfiguration `json:"portConfiguration"`
}

// statusReportsWrapper will be used to unmarshal a list of status reports for a Haip
type statusReportsWrapper struct {
	StatusReports []StatusReport `json:"statusReport"`
}

// Haip struct for a Haip
type Haip struct {
	// HA-IP name
	Name string `json:"name"`
	// The description that can be set by the customer
	Description string `json:"description"`
	// HA-IP status, either 'active', 'inactive', 'creating'
	Status Status `json:"status"`
	// Whether load balancing is enabled for this HA-IP
	IsLoadBalancingEnabled bool `json:"isLoadBalancingEnabled"`
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
	HTTPHealthCheckSsl bool `json:"httpHealthCheckSsl"`
	// HA-IP IPv4 address
	IPv4Address net.IP `json:"ipv4Address,omitempty"`
	// HA-IP IPv6 address
	IPv6Address net.IP `json:"ipv6Address,omitempty"`
	// HA-IP IP setup: 'both', 'noipv6', 'ipv6to4'
	IPSetup IPSetup `json:"ipSetup"`
	// The PTR record for the HA-IP
	PtrRecord string `json:"ptrRecord,omitempty"`
	// The IPs attached to this haip
	IPAddresses []net.IP `json:"ipAddresses,omitempty"`
	// HA-IP TLS Mode: 'tls10_11_12', 'tls11_12', 'tls12'
	TLSMode TLSMode `json:"tlsMode"`
}

// Certificate struct for haip certificates it contains an ID, expiration date and common name
type Certificate struct {
	// The common name of the certificate, usually a domain name
	CommonName string `json:"commonName,omitempty"`
	// The expiration date of the certificate in 'Y-m-d' format
	ExpirationDate string `json:"expirationDate,omitempty"`
	// The domain ssl certificate id
	ID int64 `json:"id,omitempty"`
}

// StatusReport struct for a StatusReport
type StatusReport struct {
	// Attached IP address this status report is for
	IPAddress net.IP `json:"ipAddress,omitempty"`
	// IP Version 4,6
	IPVersion int `json:"ipVersion,omitempty"`
	// Last change in the state in Europe/Amsterdam timezone
	LastChange rest.Time `json:"lastChange,omitempty"`
	// The IP address of the HA-IP load balancer
	LoadBalancerIP net.IP `json:"loadBalancerIp,omitempty"`
	// The name of the load balancer
	LoadBalancerName string `json:"loadBalancerName,omitempty"`
	// HA-IP Configuration port
	Port int `json:"port,omitempty"`
	// The state of the load balancer, either 'up' or 'down'
	State string `json:"state,omitempty"`
}

// PortConfiguration struct for a PortConfiguration
type PortConfiguration struct {
	// The port configuration ID
	ID int64 `json:"id,omitempty"`
	// A name describing the port
	Name string `json:"name"`
	// The port at which traffic arrives on your HA-IP
	SourcePort int `json:"sourcePort"`
	// The port at which traffic arrives on your attached IP address(es)
	TargetPort int `json:"targetPort"`
	// The mode determining how traffic is processed and forwarded: 'tcp', 'http', 'https', 'proxy', 'http2_https'
	Mode PortConfigurationMode `json:"mode"`
	// The mode determining how traffic between our load balancers and your attached IP address(es) is encrypted: 'off', 'on', 'strict'
	EndpointSslMode string `json:"endpointSslMode"`
}
