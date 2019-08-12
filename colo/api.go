package colo

import (
	"net"

	"github.com/transip/gotransip/v5"
)

// This file holds all ColocationService methods directly ported from TransIP API

// RequestAccess requests access to the data-center
func RequestAccess(c gotransip.Client, when string, duration int, visitors []string, phoneNumber string) ([]DatacenterVisitor, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "requestAccess",
	}
	sr.AddArgument("when", when)
	sr.AddArgument("duration", duration)
	sr.AddArgument("visitors", visitors)
	sr.AddArgument("phoneNumber", phoneNumber)

	var v struct {
		V []DatacenterVisitor `xml:"item"`
	}

	err := c.Call(sr, &v)
	return v.V, err
}

// RequestRemoteHands requests remote hands to the data-center
func RequestRemoteHands(c gotransip.Client, coloName, contactName, phoneNumber string, expectedDuration int, instructions string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "requestRemoteHands",
	}
	sr.AddArgument("coloName", coloName)
	sr.AddArgument("contactName", contactName)
	sr.AddArgument("phoneNumber", phoneNumber)
	sr.AddArgument("expectedDuration", expectedDuration)
	sr.AddArgument("instructions", instructions)

	return c.Call(sr, nil)
}

// GetColoNames get coloNames for customer
func GetColoNames(c gotransip.Client) ([]string, error) {
	var v struct {
		V []string `xml:"item"`
	}

	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getColoNames",
	}, &v)

	return v.V, err
}

// GetIPAddresses get IpAddresses that are active and assigned to a Colo
func GetIPAddresses(c gotransip.Client, coloName string) ([]net.IP, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getIpAddresses",
	}
	sr.AddArgument("coloName", coloName)

	var v struct {
		V []net.IP `xml:"item"`
	}

	err := c.Call(sr, &v)
	return v.V, err
}

// GetIPRanges Get ipranges that are assigned to a Colo
func GetIPRanges(c gotransip.Client, coloName string) ([]net.IPNet, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getIpRanges",
	}
	sr.AddArgument("coloName", coloName)

	var v struct {
		V []string `xml:"item"`
	}

	err := c.Call(sr, &v)
	lst := make([]net.IPNet, len(v.V))
	for i, x := range v.V {
		_, yy, _ := net.ParseCIDR(x)
		lst[i] = *yy
	}
	return lst, err
}

// CreateIPAddress adds a new IpAddress, either an ipv6 or an ipv4 address
func CreateIPAddress(c gotransip.Client, ipAddress net.IP, reverseDNS string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "createIpAddress",
	}
	sr.AddArgument("ipAddress", ipAddress)
	sr.AddArgument("reverseDns", reverseDNS)

	return c.Call(sr, nil)
}

// DeleteIPAddress deletes an IpAddress currently in use this account
func DeleteIPAddress(c gotransip.Client, ipAddress net.IP) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deleteIpAddress",
	}
	sr.AddArgument("ipAddress", ipAddress)

	return c.Call(sr, nil)
}

// GetReverseDNS get the Reverse DNS for an IpAddress assigned to the user
func GetReverseDNS(c gotransip.Client, ipAddress net.IP) (string, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getReverseDns",
	}
	sr.AddArgument("ipAddress", ipAddress)

	var v string

	err := c.Call(sr, &v)
	return v, err
}

// SetReverseDNS sets the Reverse DNS name for an ipAddress
func SetReverseDNS(c gotransip.Client, ipAddress net.IP, reverseDNS string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setReverseDns",
	}
	sr.AddArgument("ipAddress", ipAddress)
	sr.AddArgument("reverseDns", reverseDNS)

	return c.Call(sr, nil)
}
