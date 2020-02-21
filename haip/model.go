package haip

import (
	"github.com/transip/gotransip/v6/ipaddress"
	"net"
)

type Haips struct {
	// list of HA-IPs
	Haips []Haip `json:"haips,omitempty"`
}

// Haip struct for Haip
type Haip struct {
	// The description that can be set by the customer
	Description string `json:"description,omitempty"`
	// The interval in milliseconds at which health checks are performed. The interval may not be smaller than 2000ms.
	HealthCheckInterval int64 `json:"healthCheckInterval,omitempty"`
	// The path (URI) of the page to check HTTP status code on
	HttpHealthCheckPath string `json:"httpHealthCheckPath,omitempty"`
	// The port to perform the HTTP check on
	HttpHealthCheckPort uint16 `json:"httpHealthCheckPort,omitempty"`
	// Whether to use SSL when performing the HTTP check
	HttpHealthCheckSsl bool `json:"httpHealthCheckSsl,omitempty"`
	// The IPs attached to this haip
	IpAddresses []net.IP `json:"ipAddresses,omitempty"`
	// HA-IP IP setup: 'both', 'noipv6', 'ipv6to4'
	IpSetup string `json:"ipSetup,omitempty"`
	// HA-IP IPv4 address
	Ipv4Address net.IP `json:"ipv4Address,omitempty"`
	// HA-IP IPv6 address
	Ipv6Address net.IP `json:"ipv6Address,omitempty"`
	// Whether load balancing is enabled for this HA-IP
	IsLoadBalancingEnabled bool `json:"isLoadBalancingEnabled,omitempty"`
	// HA-IP load balancing mode: 'roundrobin', 'cookie', 'source'
	LoadBalancingMode string `json:"loadBalancingMode,omitempty"`
	// HA-IP name
	Name string `json:"name,omitempty"`
	// The PTR record for the HA-IP
	PtrRecord string `json:"ptrRecord,omitempty"`
	// HA-IP status, either 'active', 'inactive', 'creating'
	Status string `json:"status,omitempty"`
	// Cookie name to pin sessions on when using cookie balancing mode
	StickyCookieName string `json:"stickyCookieName,omitempty"`
}

// HaipCertificates struct for HaipCertificates
type HaipCertificates struct {
	// list of HA-IP certificates
	HaipCertificates []HaipCertificate `json:"haipCertificates,omitempty"`
}

// HaipIpAddresses struct for HaipIpAddresses
type HaipIpAddresses struct {
	// list of HA-IPs
	IpAddresses []net.IP `json:"ipAddresses,omitempty"`
}

// HaipPortConfigurations struct for HaipPortConfigurations
type HaipPortConfigurations struct {
	// list of HA-IP port configurations
	PortConfigurations []PortConfiguration `json:"haipPortConfiguration,omitempty"`
}

// HaipStatusReports struct for HaipStatusReports
type HaipStatusReports struct {
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

// HaipStatusReport struct for HaipStatusReport
type HaipStatusReport struct {
	// Attached IP address this status report is for
	IpAddress string `json:"ipAddress,omitempty"`
	// IP Version 4,6
	IpVersion uint8 `json:"ipVersion,omitempty"`
	// Last change in the state in Europe/Amsterdam timezone
	LastChange string `json:"lastChange,omitempty"`
	// The IP address of the HA-IP load balancer
	LoadBalancerIp net.IP `json:"loadBalancerIp,omitempty"`
	// The name of the load balancer
	LoadBalancerName string `json:"loadBalancerName,omitempty"`
	// HA-IP PortConfiguration port
	Port uint16 `json:"port,omitempty"`
	// The state of the load balancer, either 'up' or 'down'
	State string `json:"state,omitempty"`
}

// PortConfiguration struct for PortConfiguration
type PortConfiguration struct {
	// The mode determining how traffic between our load balancers and your attached IP address(es) is encrypted: 'off', 'on', 'strict'
	EndpointSslMode string `json:"endpointSslMode"`
	// The port configuration Id
	Id int64 `json:"id,omitempty"`
	// The mode determining how traffic is processed and forwarded: 'tcp', 'http', 'https', 'proxy'
	Mode string `json:"mode"`
	// A name describing the port
	Name string `json:"name"`
	// The port at which traffic arrives on your HA-IP
	SourcePort uint16 `json:"sourcePort"`
	// The port at which traffic arrives on your attached IP address(es)
	TargetPort uint16 `json:"targetPort"`
}

type HaipIpAddressesResponse struct {
	// Set of IP addresses to attach, replaces the current set of IP addresses
	IpAddresses []ipaddress.IpAddress `json:"ipAddresses,omitempty"`
}
