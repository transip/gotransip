package haip

import (
	"fmt"
	"strconv"

	"github.com/transip/gotransip"
	"github.com/transip/gotransip/util"
)

// This file holds all HaipService methods directly ported from TransIP API

// GetHaip returns Haip for given name or error if when it failed to retrieve
// Haip from API
func GetHaip(c gotransip.Client, haipName string) (Haip, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getHaip",
	}
	sr.AddArgument("haipName", haipName)

	var haip Haip
	err := c.Call(sr, &haip)

	return haip, err
}

// GetHaips returns slice of Haip structs or error when it failed to retrieve
// list of Haips from API
func GetHaips(c gotransip.Client) ([]Haip, error) {
	var h struct {
		H []Haip `xml:"item"`
	}
	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getHaips",
	}, &h)

	return h.H, err
}

// NOTE: TransIP API's changeHaipVps is not implemented, as it was superseded by
// SetHaipVpses

// SetHaipVpses replaces currently attached VPSes to the HA-IP with the provided
// list of VPSes
func SetHaipVpses(c gotransip.Client, haipName string, vpsNames []string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setHaipVpses",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("vpsNames", vpsNames)

	return c.Call(sr, nil)
}

// SetIPSetup sets the IP setup for the HA-IP
func SetIPSetup(c gotransip.Client, haipName string, setup IPSetup) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setIPSetup",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("ipSetup", string(setup))

	return c.Call(sr, nil)
}

// SetBalancingMode sets the provided balancing mode for the HA-IP. The
// cookieName argument may be an empty string unless the balancing mode is set
// to 'cookie'.
func SetBalancingMode(c gotransip.Client, haipName string, balancingMode BalancingMode, cookieName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setBalancingMode",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("balancingMode", string(balancingMode))
	sr.AddArgument("cookieName", cookieName)

	return c.Call(sr, nil)
}

// SetHTTPHealthCheck configures a HTTP health check for the HA-IP. To disable
// a HTTP health check use SetTCPHealthCheck()
func SetHTTPHealthCheck(c gotransip.Client, haipName, path string, port int) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setHttpHealthCheck",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("path", path)
	sr.AddArgument("port", fmt.Sprintf("%d", port))

	return c.Call(sr, nil)
}

// SetTCPHealthCheck configures a TCP health check for the HA-IP (this is the
// default health check)
func SetTCPHealthCheck(c gotransip.Client, haipName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setTcpHealthCheck",
	}
	sr.AddArgument("haipName", haipName)

	return c.Call(sr, nil)
}

// GetStatusReport returns status report for given HA-IP
func GetStatusReport(c gotransip.Client, haipName string) (StatusReport, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getStatusReport",
	}
	sr.AddArgument("haipName", haipName)

	var v statusXMLOuter
	err := c.Call(sr, &v)
	if err != nil {
		return StatusReport{}, err
	}

	return parseStatusReportBody(v)
}

// GetCertificatesByHaip returns all certificates attached given HA-IP
func GetCertificatesByHaip(c gotransip.Client, haipName string) ([]Certificate, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getCertificatesByHaip",
	}
	sr.AddArgument("haipName", haipName)

	var h util.KeyValueXML
	err := c.Call(sr, &h)

	if err != nil {
		return nil, err
	}

	return keyValueXMLToCertificates(h), nil
}

// GetAvailableCertificatesByHaip returns all available certificates ready to
// attach to given HA-IP
func GetAvailableCertificatesByHaip(c gotransip.Client, haipName string) ([]Certificate, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailableCertificatesByHaip",
	}
	sr.AddArgument("haipName", haipName)

	var h util.KeyValueXML
	err := c.Call(sr, &h)

	if err != nil {
		return nil, err
	}

	return keyValueXMLToCertificates(h), nil
}

// AddCertificateToHaip adds a certificate to given HA-IP
func AddCertificateToHaip(c gotransip.Client, haipName string, certificateID int64) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "addCertificateToHaip",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("certificateId", fmt.Sprintf("%d", certificateID))

	return c.Call(sr, nil)
}

// AddCertificateFromHaip adds a certificate to given HA-IP
//
// Deprecated: use AddCertificateToHaip instead
func AddCertificateFromHaip(c gotransip.Client, haipName string, certificateID int64) error {
	return AddCertificateToHaip(c, haipName, certificateID)
}

// DeleteCertificateFromHaip removes a certificate from given HA-IP
func DeleteCertificateFromHaip(c gotransip.Client, haipName string, certificateID int64) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deleteCertificateFromHaip",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("certificateId", fmt.Sprintf("%d", certificateID))

	return c.Call(sr, nil)
}

// StartHaipLetsEncryptCertificateIssue adds Let's Encrypt SSL encryption to HA-IP
// for given commonName
func StartHaipLetsEncryptCertificateIssue(c gotransip.Client, haipName, commonName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "startHaipLetsEncryptCertificateIssue",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("commonName", commonName)

	return c.Call(sr, nil)
}

// GetPtrForHaip returns the current PTR for the given HA-IP
func GetPtrForHaip(c gotransip.Client, haipName string) (string, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getPtrForHaip",
	}
	sr.AddArgument("haipName", haipName)

	var ptr string
	err := c.Call(sr, &ptr)

	return ptr, err
}

// SetPtrForHaip updates the PTR records for the given HA-IP
func SetPtrForHaip(c gotransip.Client, haipName string, ptr string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setPtrForHaip",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("ptr", ptr)

	return c.Call(sr, nil)
}

// SetHaipDescription updates the description for HA-IP
func SetHaipDescription(c gotransip.Client, haipName string, description string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setHaipDescription",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("description", description)

	return c.Call(sr, nil)
}

// GetPortConfigurations returns all port configurations for given HA-IP
func GetPortConfigurations(c gotransip.Client, haipName string) ([]PortConfiguration, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getPortConfigurations",
	}
	sr.AddArgument("haipName", haipName)

	// complex struct for getting key/value pairs from Map
	var h struct {
		PortConf []struct {
			Item []struct {
				Key   string `xml:"key"`
				Value string `xml:"value"`
			} `xml:"item"`
		} `xml:"item"`
	}

	err := c.Call(sr, &h)
	if err != nil {
		return nil, err
	}

	// convert SOAP result into PortConfigurations
	pcs := make([]PortConfiguration, len(h.PortConf))
	for i, ci := range h.PortConf {
		pc := PortConfiguration{}
		for _, x := range ci.Item {
			switch x.Key {
			case "configurationId":
				pc.ID, _ = strconv.ParseInt(x.Value, 10, 64)
			case "name":
				pc.Name = x.Value
			case "mode":
				pc.Mode = Mode(x.Value)
			case "sourcePort":
				pc.SourcePort, _ = strconv.ParseInt(x.Value, 10, 64)
			case "targetPort":
				pc.TargetPort, _ = strconv.ParseInt(x.Value, 10, 64)
			case "endpointSslMode":
				pc.EndpointSSLMode = EndpointSSLMode(x.Value)
			}
		}

		pcs[i] = pc
	}

	return pcs, err
}

// SetDefaultPortConfiguration reverts the port configuration to the default
// for given HA-IP
func SetDefaultPortConfiguration(c gotransip.Client, haipName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setDefaultPortConfiguration",
	}
	sr.AddArgument("haipName", haipName)

	return c.Call(sr, nil)
}

// AddPortConfiguration adds a port configuration to given HA-IP
func AddPortConfiguration(c gotransip.Client, haipName string, pc PortConfiguration) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "addPortConfiguration",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("name", pc.Name)
	sr.AddArgument("sourcePort", fmt.Sprintf("%d", pc.SourcePort))
	sr.AddArgument("targetPort", fmt.Sprintf("%d", pc.TargetPort))
	sr.AddArgument("mode", string(pc.Mode))
	sr.AddArgument("endpointSslMode", string(pc.EndpointSSLMode))

	return c.Call(sr, nil)
}

// UpdatePortConfiguration updates an existsing port configuration to given HA-IP
func UpdatePortConfiguration(c gotransip.Client, haipName string, pc PortConfiguration) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "updatePortConfiguration",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("configurationId", fmt.Sprintf("%d", pc.ID))
	sr.AddArgument("name", pc.Name)
	sr.AddArgument("sourcePort", fmt.Sprintf("%d", pc.SourcePort))
	sr.AddArgument("targetPort", fmt.Sprintf("%d", pc.TargetPort))
	sr.AddArgument("mode", string(pc.Mode))
	sr.AddArgument("endpointSslMode", string(pc.EndpointSSLMode))

	return c.Call(sr, nil)
}

// DeletePortConfiguration deletes a port configuration to given HA-IP
func DeletePortConfiguration(c gotransip.Client, haipName string, ID int64) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deletePortConfiguration",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("configurationId", fmt.Sprintf("%d", ID))

	return c.Call(sr, nil)
}

// CancelHaip cancels contract for HA-IP per endTime
func CancelHaip(c gotransip.Client, haipName string, endTime gotransip.CancellationTime) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "cancelHaip",
	}
	sr.AddArgument("haipName", haipName)
	sr.AddArgument("endTime", string(endTime))

	return c.Call(sr, nil)
}
