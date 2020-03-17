package haip

import (
	"github.com/transip/gotransip/v6/rest"
	"net"
)

// HaipStatus is one of the following strings
// 'active', 'inactive', 'creating'
type HaipStatus string

// define all of the possible haip statuses
const (
	// HaipStatusActive is the status field for an active Haip, ready to use
	HaipStatusActive HaipStatus = "active"
	// HaipStatusInactive is the status field for an inactive Haip, not usable, please contact support
	HaipStatusInactive HaipStatus = "inactive"
	// HaipStatusCreating is the status field for a Haip that is being created
	HaipStatusCreating HaipStatus = "creating"
)

// LoadBalancingMode is one of the following strings
// 'roundrobin', 'cookie', 'source'
type LoadBalancingMode string

// define all of the possible load balancing modes
const (
	// LoadBalancingModeRoundRobin is the LoadBalancing mode roundrobin for a Haip, forward to next address everytime
	LoadBalancingModeRoundRobin LoadBalancingMode = "roundrobin"
	// LoadBalancingModeCookie is the LoadBalancing mode cookie for a Haip, forward to a fixed server, based on the cookie
	LoadBalancingModeCookie LoadBalancingMode = "cookie"
	// LoadBalancingModeSource is the LoadBalancing mode source for a Haip, choose a server to forward based on the source address
	LoadBalancingModeSource LoadBalancingMode = "source"
)

// IpSetup is one of the following strings
// 'both', 'noipv6', 'ipv6to4'
type IpSetup string

// define all of the possible ip setup options
const (
	// IpSetupBoth accept ipv4 and ipv6 and forward them to seperate ipv4 and ipv6 addresses
	IpSetupBoth IpSetup = "both"
	// IpSetupNoIpv6 do not accept ipv6 traffic
	IpSetupNoIpv6 IpSetup = "noipv6"
	// IpSetupIpv6to4 forward ipv6 traffic to ipv4
	IpSetupIpv6to4 IpSetup = "ipv6to4"
)

// PortConfigurationMode is one of the following strings
// 'tcp', 'http', 'https', 'proxy', 'http2_https'
type PortConfigurationMode string

// define all of the possible port configuration modes
const (
	// PortConfigurationModeTcp plain TCP forward to your VPS
	PortConfigurationModeTcp PortConfigurationMode = "tcp"
	// PortConfigurationModeHttp appends a X-Forwarded-For header to HTTP requests with the original remote IP
	PortConfigurationModeHttp PortConfigurationMode = "http"
	// PortConfigurationModeHttps same as HTTP, with SSL Certificate offloading
	PortConfigurationModeHttps PortConfigurationMode = "https"
	// PortConfigurationModeProxy proxy protocol is also a way to retain the original remote IP, but also works for non HTTP traffic (note: the receiving application has to support this)
	PortConfigurationModeProxy PortConfigurationMode = "proxy"
	// PortConfigurationModeHttp2Https same as HTTPS, with http/2 support
	PortConfigurationModeHttp2Https PortConfigurationMode = "http2_https"
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
	Certificates []HaipCertificate `json:"certificates,certificates"`
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
	SslCertificateId int64  `json:"sslCertificateId,omitempty"`
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

// Haip struct for Haip
type Haip struct {
	// HA-IP name
	Name string `json:"name"`
	// The description that can be set by the customer
	Description string `json:"description"`
	// HA-IP status, either 'active', 'inactive', 'creating'
	Status HaipStatus `json:"status"`
	// Whether load balancing is enabled for this HA-IP
	IsLoadBalancingEnabled bool `json:"isLoadBalancingEnabled"`
	// HA-IP load balancing mode: 'roundrobin', 'cookie', 'source'
	LoadBalancingMode LoadBalancingMode `json:"loadBalancingMode,omitempty"`
	// Cookie name to pin sessions on when using cookie balancing mode
	StickyCookieName string `json:"stickyCookieName,omitempty"`
	// The interval in milliseconds at which health checks are performed. The interval may not be smaller than 2000ms.
	HealthCheckInterval int64 `json:"healthCheckInterval,omitempty"`
	// The path (URI) of the page to check HTTP status code on
	HttpHealthCheckPath string `json:"httpHealthCheckPath,omitempty"`
	// The port to perform the HTTP check on
	HttpHealthCheckPort int `json:"httpHealthCheckPort,omitempty"`
	// Whether to use SSL when performing the HTTP check
	HttpHealthCheckSsl bool `json:"httpHealthCheckSsl"`
	// HA-IP IPv4 address
	Ipv4Address net.IP `json:"ipv4Address,omitempty"`
	// HA-IP IPv6 address
	Ipv6Address net.IP `json:"ipv6Address,omitempty"`
	// HA-IP IP setup: 'both', 'noipv6', 'ipv6to4'
	IpSetup IpSetup `json:"ipSetup"`
	// The PTR record for the HA-IP
	PtrRecord string `json:"ptrRecord,omitempty"`
	// The IPs attached to this haip
	IpAddresses []net.IP `json:"ipAddresses,omitempty"`
}

// HaipCertificate struct for HaipCertificate
type HaipCertificate struct {
	// The common name of the certificate, usually a domain name
	CommonName string `json:"commonName,omitempty"`
	// The expiration date of the certificate in 'Y-m-d' format
	ExpirationDate string `json:"expirationDate,omitempty"`
	// The domain ssl certificate id
	Id int64 `json:"id,omitempty"`
}

// StatusReport struct for StatusReport
type StatusReport struct {
	// Attached IP address this status report is for
	IpAddress net.IP `json:"ipAddress,omitempty"`
	// IP Version 4,6
	IpVersion int `json:"ipVersion,omitempty"`
	// Last change in the state in Europe/Amsterdam timezone
	LastChange rest.Time `json:"lastChange,omitempty"`
	// The IP address of the HA-IP load balancer
	LoadBalancerIp net.IP `json:"loadBalancerIp,omitempty"`
	// The name of the load balancer
	LoadBalancerName string `json:"loadBalancerName,omitempty"`
	// HA-IP Configuration port
	Port int `json:"port,omitempty"`
	// The state of the load balancer, either 'up' or 'down'
	State string `json:"state,omitempty"`
}

// PortConfiguration struct for PortConfiguration
type PortConfiguration struct {
	// The port configuration Id
	Id int64 `json:"id,omitempty"`
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

// HaipIpAddressesResponse object contains a list of IPAddresses in it
// used to unpack the rest response and return the encapsulated net.IPs
// this is just used internal for unpacking, this should not be exported
// we want to return net.IP objects not a HaipIpAddressesResponse
type HaipIpAddressesResponse struct {
	// Set of IP addresses to attach, replaces the current set of IP addresses
	IpAddresses []net.IP `json:"ipAddresses,omitempty"`
}
