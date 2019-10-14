package haip

import (
	"net"
	"strconv"
	"time"

	"github.com/transip/gotransip/v5/util"
	"github.com/transip/gotransip/v5/vps"
)

const (
	serviceName string = "HaipService"
)

// Status represents the possibles states a HA-IP can be in
type Status string

var (
	// StatusCreated means the HA-IP was provisioned but not yet configured
	StatusCreated Status = "created"
	// StatusActive means the HA-IP is active
	StatusActive Status = "active"
	// StatusInactive means the HA-IP is deactiveated
	StatusInactive Status = "inactive"
)

// StatusReport represents the structured data that is returned from
// GetStatusReport
type StatusReport struct {
	PortConfiguration []statusReportPortConfiguration
}

// IPSetup represents the possible ipSetups configurable for HA-IP
type IPSetup string

var (
	// IPSetupBoth forwards IPv4 and IPv6 to the VPS's respective addresses
	IPSetupBoth IPSetup = "both"
	// IPSetupNoIPv6 only forwards IPv4
	IPSetupNoIPv6 IPSetup = "noipv6"
	// IPSetupIPv6To4 forwards both IPv4 and IPv6 to the VPS's IPv4 address
	IPSetupIPv6To4 IPSetup = "ipv6to4"
	// IPSetupIPv4To6 forwards both IPv4 and IPv6 to the VPS's IPv6 address
	IPSetupIPv4To6 IPSetup = "ipv4to6"
)

// Mode represents possible modes a HA-IP can operate in
type Mode string

var (
	// ModeTCP operates in TCP mode
	ModeTCP Mode = "tcp"
	// ModeHTTP operates in HTTP mode
	ModeHTTP Mode = "http"
	// ModeHTTPS operates in HTTPS mode
	ModeHTTPS Mode = "https"
	// ModeProxy operates in haproxy's PROXY protocol
	ModeProxy Mode = "proxy"
)

// EndpointSSLMode represents possible modes of encryption between HA-IP and VPSes
type EndpointSSLMode string

var (
	// EndpointSSLModeOff means traffic to VPSes is sent unencrypted
	EndpointSSLModeOff EndpointSSLMode = "off"
	// EndpointSSLModeOn means traffic to VPSes is sent encrypted, certificate can
	// be self-signed, no verification is applied
	EndpointSSLModeOn EndpointSSLMode = "on"
	// EndpointSSLModeStrict means traffic to VPSes is sent encrypted, certificate's
	// CA should be trusted
	EndpointSSLModeStrict EndpointSSLMode = "strict"
)

// BalancingMode represents the possible load balancing modes configurable for HA-IP
type BalancingMode string

var (
	// BalancingModeRoundRobin balances on a round-robin base
	BalancingModeRoundRobin BalancingMode = "roundrobin"
	// BalancingModeCookie balances based on a predefined cookie sent by the client
	// from the second request onwards
	BalancingModeCookie BalancingMode = "cookie"
	// BalancingModeSource balances based on the source address of the client
	BalancingModeSource BalancingMode = "source"
)

// HealthCheckMode represents the possible health checking modes configurable for
// HA-IP
type HealthCheckMode string

var (
	// HealthCheckModeTCP checks whether the specified TCP port of a Vps is reachable
	HealthCheckModeTCP HealthCheckMode = "tcp"
	// HealthCheckModeHTTP performs HTTP requests based on HTTPHealthCheckPath and
	// httpHealthCheckPort
	HealthCheckModeHTTP HealthCheckMode = "http"
)

// Certificate represents a HA-IP certificate
type Certificate struct {
	ID             int64
	CommonName     string
	ExpirationDate time.Time
}

// PortConfiguration represents a HA-IP's port configuration
type PortConfiguration struct {
	ID              int64
	Name            string
	Mode            Mode
	SourcePort      int64
	TargetPort      int64
	EndpointSSLMode EndpointSSLMode
}

// Haip represents a Transip_Haip object
// https://api.transip.nl/docs/transip.nl/class-Transip_Haip.html
type Haip struct {
	Name                   string          `xml:"name"`
	Status                 Status          `xml:"status"`
	IsBlocked              bool            `xml:"isBlocked"`
	IsLoadBalancingEnabled bool            `xml:"isLoadBalancingEnabled"`
	LoadBalancingMode      BalancingMode   `xml:"loadBalancingMode"`
	StickyCookieName       string          `xml:"stickyCookieName"`
	HealthCheckMode        HealthCheckMode `xml:"healthCheckMode"`
	HTTPHealthCheckPath    string          `xml:"httpHealthCheckPath"`
	HTTPHealthCheckPort    int             `xml:"httpHealthCheckPort"`
	IPv4Address            net.IP          `xml:"ipv4Address"`
	IPv6Address            net.IP          `xml:"ipv6Address"`
	IPSetup                IPSetup         `xml:"ipSetup"`
	AttachedVpses          []vps.Vps       `xml:"attachedVpses>item"`
}

// keyValueXMLToCertificates converts KeyValueXML to an array of Certificates
func keyValueXMLToCertificates(k util.KeyValueXML) (certs []Certificate) {

	// convert SOAP result into Certificates
	certs = make([]Certificate, len(k.Cont))
	for i, ci := range k.Cont {
		crt := Certificate{}
		for _, x := range ci.Item {
			switch x.Key {
			case "id":
				crt.ID, _ = strconv.ParseInt(x.Value, 10, 64)
			case "commonName":
				crt.CommonName = x.Value
			case "expirationDate":
				crt.ExpirationDate, _ = time.Parse("2006-01-02 15:04:05", x.Value)
			}
		}

		certs[i] = crt
	}

	return
}
