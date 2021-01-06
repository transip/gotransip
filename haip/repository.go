package haip

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net"
	"net/url"
)

// Repository can be used to get a list of your Haips
// order new ones, changing specific Haip properties
// Updating the attached ssl certificates and attaching/detaching IP addresses to the HAIP
type Repository repository.RestRepository

// GetAll returns an array of all Haips in your account
func (r *Repository) GetAll() ([]Haip, error) {
	var response haipsWrapper
	err := r.Client.Get(rest.Request{Endpoint: "/haips"}, &response)

	return response.Haips, err
}

// GetSelection returns a limited list of your Haips,
// specify how many and which page/chunk of Haips you want to retrieve
func (r *Repository) GetSelection(page int, itemsPerPage int) ([]Haip, error) {
	var response haipsWrapper
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/haips", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.Haips, err
}

// GetByName returns information on a specific Haip by name
func (r *Repository) GetByName(haipName string) (Haip, error) {
	var response haipWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s", haipName)}
	err := r.Client.Get(restRequest, &response)

	return response.Haip, err
}

// Order allows you to order a new Haip
func (r *Repository) Order(productName string, description string) error {
	requestBody := haipOrderWrapper{ProductName: productName, Description: description}

	return r.Client.Post(rest.Request{Endpoint: "/haips", Body: requestBody})
}

// Update allows you to alter your Haip in several ways outlined below:
//   - Set the description of a HA-IP;
//   - Set the PTR record;
//   - Set the httpHealthCheckPath, must start with a /;
//   - Set the httpHealthCheckPort, the port must be configured on the HA-IP PortConfigurations.
//
// Load balancing options (loadBalancingMode):
//   - roundrobin: forward to next address everytime;
//   - cookie: forward to a fixed server, based on the cookie;
//   - source: choose a server to forward to based on the source address.
//
// IP setup options (ipSetup):
//   - both: accept ipv4 and ipv6 and forward them to separate ipv4 and ipv6 addresses;
//   - noipv6: do not accept ipv6 traffic;
//   - ipv6to4: forward ipv6 traffic to ipv4.
//   - ipv4to6: forward ipv4 traffic to ipv6.
//
// TLS options (tlsMode):
//   - tls10_11_12: only allow incoming tls traffic with versions 1.0, 1.1 and 1.2;
//   - tls11_12: only allow incoming tls traffic with version 1.1 or 1.2;
//   - tls12: only allow incoming traffic with tls version 1.2.
//
// For more information see: https://api.transip.nl/rest/docs.html#ha-ip-ha-ip-put
func (r *Repository) Update(haip Haip) error {
	requestBody := haipWrapper{Haip: haip}

	return r.Client.Put(rest.Request{Endpoint: fmt.Sprintf("/haips/%s", haip.Name), Body: requestBody})
}

// Cancel will cancel the Haip, thus deleting it
func (r *Repository) Cancel(haipName string, endTime gotransip.CancellationTime) error {
	var requestBody gotransip.CancellationRequest
	requestBody.EndTime = endTime
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s", haipName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetAllCertificates will return a list of certificates currently attached to the given Haip
func (r *Repository) GetAllCertificates(haipName string) ([]Certificate, error) {
	var response certificatesWrapper
	err := r.Client.Get(rest.Request{Endpoint: fmt.Sprintf("/haips/%s/certificates", haipName)}, &response)

	return response.Certificates, err
}

// AddCertificate allows you to add a DV, OV or EV Certificate to Haip for SSL offloading
// Enable HTTPS mode in Configuration to use these certificates
func (r *Repository) AddCertificate(haipName string, sslCertificateID int64) error {
	requestBody := addCertificateRequest{SslCertificateID: sslCertificateID}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/certificates", haipName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// AddLetsEncryptCertificate allows you to add a LetsEncrypt certificate to your HA-IP.
// We will take care of all the validation and renewals.
//
// In order to provide free LetsEncrypt certificates for the domains on your HA-IP,
// some requirements must be met in order to complete the certificate request:
//   - DNS: the given CommonName must resolve to the HA-IP IP.
//       IPv6 is not required, but when set, it must resolve to the HA-IP IPv6;
//   - Configuration: LetsEncrypt verifies domains with a HTTP call to /.well-know.
//       When requesting a LetsEncrypt certificate, our proxies will handle all ACME requests
//       to automatically verify the certificate.
//       To achieve this, the HA-IP must have a HTTP portConfiguration on port 80.
//       When using this, you will also no longer be able to verify your own LetsEncrypt certificates via HA-IP.
//
// For more information, see: https://api.transip.nl/rest/docs.html#ha-ip-ha-ip-certificates-post-1
func (r *Repository) AddLetsEncryptCertificate(haipName string, commonName string) error {
	requestBody := addCertificateRequest{CommonName: commonName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/certificates", haipName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// DetachCertificate detaches a certificate from a Haip by certificateId
func (r *Repository) DetachCertificate(haipName string, certificateID int64) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/certificates/%d", haipName, certificateID)}

	return r.Client.Delete(restRequest)
}

// GetAttachedIPAddresses returns a list of currently attached IP address(es) to your Haip
func (r *Repository) GetAttachedIPAddresses(haipName string) ([]net.IP, error) {
	var response ipAddressesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/ip-addresses", haipName)}
	err := r.Client.Get(restRequest, &response)

	return response.IPAddresses, err
}

// SetAttachedIPAddresses allows you to replace the IP address(es) attached your Haip
func (r *Repository) SetAttachedIPAddresses(haipName string, ipAddresses []net.IP) error {
	requestBody := ipAddressesWrapper{IPAddresses: ipAddresses}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/ip-addresses", haipName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// DetachIPAddresses allows you to detach all IP Addresses from a Haip
func (r *Repository) DetachIPAddresses(haipName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/ip-addresses", haipName)}

	return r.Client.Delete(restRequest)
}

// GetPortConfigurations returns a list of all PortConfigurations on the given Haip
func (r *Repository) GetPortConfigurations(haipName string) ([]PortConfiguration, error) {
	var response portConfigurationsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/port-configurations", haipName)}
	err := r.Client.Get(restRequest, &response)

	return response.PortConfigurations, err
}

// GetPortConfiguration returns the Configuration struct for a given Configuration by id
func (r *Repository) GetPortConfiguration(haipName string, portConfigurationID int64) (PortConfiguration, error) {
	var response portConfigurationWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/port-configurations/%d", haipName, portConfigurationID)}
	err := r.Client.Get(restRequest, &response)

	return response.Configuration, err
}

// AddPortConfiguration allows you to Add PortConfigurations to your HA-IP to route traffic to your attached IP address(es)
//
// Mode options:
//   - http: appends a X-Forwarded-For header to HTTP requests with the original remote IP;
//   - https: same as HTTP, with SSL Certificate offloading;
//   - http2_https: same as HTTPS, with http/2 support;
//   - tcp: plain TCP forward to your attached IP address(es);
//   - proxy: proxy protocol is also a way to retain the original remote IP, but also works for non HTTP traffic
//     (note: the receiving application has to support this).
//
// Endpoint SSL mode options:
//
//   - off: no SSL connection is established between our load balancers and your attached IP address(es);
//   - on: an SSL connection is established between our load balancers your attached IP address(es),
//     but the certificate is not validated;
//   - strict: an SSL connection is established between our load balancers your attached IP address(es),
//     and the certificate must signed by a trusted Certificate Authority.
//
// For more information, see https://api.transip.nl/rest/docs.html#ha-ip-ha-ip-port-configurations-post
func (r *Repository) AddPortConfiguration(haipName string, configuration PortConfiguration) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/port-configurations", haipName), Body: &configuration}

	return r.Client.Post(restRequest)
}

// UpdatePortConfiguration allows you to update:
//   Name, SourcePort, TargetPort, Mode, or EndpointSslMode of a Configuration
// For more information on these fields see the AddPortConfiguration method and: https://api.transip.nl/rest/docs.html#ha-ip-ha-ip-port-configurations-put
func (r *Repository) UpdatePortConfiguration(haipName string, configuration PortConfiguration) error {
	requestBody := portConfigurationWrapper{Configuration: configuration}
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/haips/%s/port-configurations/%d", haipName, configuration.ID),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// RemovePortConfiguration allows you to remove a port configuration
func (r *Repository) RemovePortConfiguration(haipName string, portConfigurationID int64) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/port-configurations/%d", haipName, portConfigurationID)}

	return r.Client.Delete(restRequest)
}

// GetStatusReport returns a StatusReport per attached IP address, IP version, port and load balancer.
// You can use this method to monitor / verify the status of your HA-IP and attached IP addresses
func (r *Repository) GetStatusReport(haipName string) ([]StatusReport, error) {
	var response statusReportsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/haips/%s/status-reports", haipName)}
	err := r.Client.Get(restRequest, &response)

	return response.StatusReports, err
}
