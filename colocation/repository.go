package colocation

import (
	"fmt"
	"github.com/transip/gotransip/v6/ipaddress"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net"
)

// Repository can be used to get a list of your colocations, create a remote hands request
// and edit/show/update colocation IP address data
type Repository repository.RestRepository

// GetAll returns a list of your colocations
func (r *Repository) GetAll() ([]Colocation, error) {
	var response colocationsWrapper
	restRequest := rest.RestRequest{Endpoint: "/colocations"}
	err := r.Client.Get(restRequest, &response)

	return response.Colocations, err

}

// GetByName returns a specific colocation by name
func (r *Repository) GetByName(coloName string) (Colocation, error) {
	var response colocationWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/colocations/%s", coloName)}
	err := r.Client.Get(restRequest, &response)

	return response.Colocation, err
}

// CreateRemoteHandsRequest allows you to request a remote task from an engineer
// It sends a request to a datacenter engineer to perform simple task on your server, e.g. a 'powercycle'
func (r *Repository) CreateRemoteHandsRequest(remoteHandsRequest RemoteHandsRequest) error {
	requestBody := remoteHandsRequestWrapper{RemoteHands: remoteHandsRequest}
	restRequest := rest.RestRequest{
		Endpoint: fmt.Sprintf("/colocations/%s/remote-hands", remoteHandsRequest.ColoName),
		Body:     &requestBody,
	}

	return r.Client.Post(restRequest)
}

// GetIPAddresses returns all IP addresses attached to your Colocation
func (r *Repository) GetIPAddresses(coloName string) ([]ipaddress.IPAddress, error) {
	var response ipaddress.IPAddressesWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/colocations/%s/ip-addresses", coloName)}
	err := r.Client.Get(restRequest, &response)

	return response.IPAddresses, err
}

// GetIPAddressByAddress returns network information for the specified IP address of the specified Colocation
func (r *Repository) GetIPAddressByAddress(coloName string, address net.IP) (ipaddress.IPAddress, error) {
	var response ipAddressWrapper
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/colocations/%s/ip-addresses/%s", coloName, address.String())}
	err := r.Client.Get(restRequest, &response)

	return response.IPAddress, err
}

// AddIPAddress allows you to add an IP address to your Colocation by specifying the Colocation name and the IP to add
// Note: the IP address you want to add should be in a range you own.
// Optionally, you can also set the reverse dns by setting the reverseDns argument
func (r *Repository) AddIPAddress(coloName string, address net.IP, reverseDns string) error {
	requestBody := addIpRequest{IPAddress: address, ReverseDns: reverseDns}
	restRequest := rest.RestRequest{
		Endpoint: fmt.Sprintf("/colocations/%s/ip-addresses", coloName),
		Body:     &requestBody,
	}

	return r.Client.Post(restRequest)
}

// UpdateReverseDNS allows you to update the reverse dns for IPv4 addresses as wal as IP addresses
func (r *Repository) UpdateReverseDNS(coloName string, ip ipaddress.IPAddress) error {
	requestBody := ipAddressWrapper{IPAddress: ip}
	restRequest := rest.RestRequest{
		Endpoint: fmt.Sprintf("/colocations/%s/ip-addresses/%s", coloName, ip.Address.String()),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// RemoveIPAddress allows you to remove an IP address from the registered list of IP address within your Colocation's range.
func (r *Repository) RemoveIPAddress(coloName string, address net.IP) error {
	restRequest := rest.RestRequest{Endpoint: fmt.Sprintf("/colocations/%s/ip-addresses/%s", coloName, address.String())}

	return r.Client.Delete(restRequest)
}
