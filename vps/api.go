package vps

import (
	"fmt"
	"net"

	"github.com/transip/gotransip"
	"github.com/transip/gotransip/util"
)

// This file holds all VpsService methods directly ported from TransIP API

// GetAvailableProducts returns a list of al available VPS products
func GetAvailableProducts(c gotransip.Client) ([]Product, error) {
	var v struct {
		V []Product `xml:"item"`
	}
	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailableProducts",
	}, &v)

	return v.V, err
}

// GetAvailableAddons returns a list of al available VPS addons
func GetAvailableAddons(c gotransip.Client) ([]Product, error) {
	var v struct {
		V []Product `xml:"item"`
	}
	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailableAddons",
	}, &v)

	return v.V, err
}

// GetActiveAddonsForVps returns a list of all the active addons for given Vps
func GetActiveAddonsForVps(c gotransip.Client, vpsName string) ([]Product, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getActiveAddonsForVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []Product `xml:"item"`
	}
	err := c.Call(sr, &v)

	return v.V, err
}

// GetAvailableUpgrades returns list of all available upgrades for given Vps
func GetAvailableUpgrades(c gotransip.Client, vpsName string) ([]Product, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailableUpgrades",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []Product `xml:"item"`
	}
	err := c.Call(sr, &v)

	return v.V, err
}

// GetAvailableAddonsForVps returns list of all available addons for a given Vps
func GetAvailableAddonsForVps(c gotransip.Client, vpsName string) ([]Product, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailableAddonsForVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []Product `xml:"item"`
	}
	err := c.Call(sr, &v)

	return v.V, err
}

// GetCancellableAddonsForVps returns a list of all cancellable addons for given Vps
func GetCancellableAddonsForVps(c gotransip.Client, vpsName string) ([]Product, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getCancellableAddonsForVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []Product `xml:"item"`
	}
	err := c.Call(sr, &v)

	return v.V, err
}

// OrderVps orders a VPS with optional addons
func OrderVps(c gotransip.Client, productName string, addons []string, operatingSystemName, hostname string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "orderVps",
	}
	sr.AddArgument("productName", productName)
	sr.AddArgument("addons", addons)
	sr.AddArgument("operatingSystemName", operatingSystemName)
	sr.AddArgument("hostname", hostname)

	return c.Call(sr, nil)
}

// OrderVpsInAvailabilityZone orders a VPS with optional addons in a specific
// availability zone
func OrderVpsInAvailabilityZone(c gotransip.Client, productName string, addons []string, operatingSystemName, hostname, availabilityZone string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "orderVpsInAvailabilityZone",
	}
	sr.AddArgument("productName", productName)
	sr.AddArgument("addons", addons)
	sr.AddArgument("operatingSystemName", operatingSystemName)
	sr.AddArgument("hostname", hostname)
	sr.AddArgument("availabilityZone", availabilityZone)

	return c.Call(sr, nil)
}

// CloneVps clones a Vps
func CloneVps(c gotransip.Client, vpsName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "cloneVps",
	}
	sr.AddArgument("vpsName", vpsName)

	return c.Call(sr, nil)
}

// CloneVpsToAvailabilityZone clones a Vps to a specific availability zone
func CloneVpsToAvailabilityZone(c gotransip.Client, vpsName, availabilityZone string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "cloneVpsToAvailabilityZone",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("availabilityZone", availabilityZone)

	return c.Call(sr, nil)
}

// OrderAddon orders addons to a VPS
func OrderAddon(c gotransip.Client, vpsName string, addons []string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "orderAddon",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("addons", addons)

	return c.Call(sr, nil)
}

// OrderPrivateNetwork orders a private Network
func OrderPrivateNetwork(c gotransip.Client) error {
	return c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "orderPrivateNetwork",
	}, nil)
}

// UpgradeVps upgrades given Vps to new contract name
func UpgradeVps(c gotransip.Client, vpsName, upgradeToProductName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "upgradeVps",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("upgradeToProductName", upgradeToProductName)

	return c.Call(sr, nil)
}

// CancelVps cancels contract for Vps per endTime
func CancelVps(c gotransip.Client, vpsName string, endTime gotransip.CancellationTime) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "cancelVps",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("endTime", string(endTime))

	return c.Call(sr, nil)
}

// CancelAddon cancels contract for Vps's given addon
func CancelAddon(c gotransip.Client, vpsName, addonName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "cancelAddon",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("addonName", addonName)

	return c.Call(sr, nil)
}

// CancelPrivateNetwork cancels contract for PrivateNetwork per endTime
func CancelPrivateNetwork(c gotransip.Client, privateNetworkName string, endTime gotransip.CancellationTime) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "cancelPrivateNetwork",
	}
	sr.AddArgument("privateNetworkName", privateNetworkName)
	sr.AddArgument("endTime", string(endTime))

	return c.Call(sr, nil)
}

// GetPrivateNetworksByVps returns a list of private networks for given Vps
func GetPrivateNetworksByVps(c gotransip.Client, vpsName string) ([]PrivateNetwork, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getPrivateNetworksByVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []PrivateNetwork `xml:"item"`
	}
	err := c.Call(sr, &v)

	return v.V, err
}

// GetAllPrivateNetworks returns list of all private networks
func GetAllPrivateNetworks(c gotransip.Client) ([]PrivateNetwork, error) {
	var v struct {
		V []PrivateNetwork `xml:"item"`
	}
	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAllPrivateNetworks",
	}, &v)

	return v.V, err
}

// AddVpsToPrivateNetwork connects given Vps to PrivateNetwork
func AddVpsToPrivateNetwork(c gotransip.Client, vpsName, privateNetworkName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "addVpsToPrivateNetwork",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("privateNetworkName", privateNetworkName)

	return c.Call(sr, nil)
}

// RemoveVpsFromPrivateNetwork disconnects given Vps from PrivateNetwork
func RemoveVpsFromPrivateNetwork(c gotransip.Client, vpsName, privateNetworkName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "removeVpsFromPrivateNetwork",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("privateNetworkName", privateNetworkName)

	return c.Call(sr, nil)
}

// GetTrafficInformationForVps returns traffic information for this contract
// period for given Vps
func GetTrafficInformationForVps(c gotransip.Client, vpsName string) (TrafficInformation, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getTrafficInformationForVps",
		Padding: []string{"item"},
	}
	sr.AddArgument("vpsName", vpsName)

	var v util.KeyValueXML
	err := c.Call(sr, &v)

	return keyValueXMLToTrafficInformation(v), err
}

// GetPooledTrafficInformation returns PooledTraffic information for the account
func GetPooledTrafficInformation(c gotransip.Client) (TrafficInformation, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getPooledTrafficInformation",
		Padding: []string{"item"},
	}

	var v util.KeyValueXML
	err := c.Call(sr, &v)

	return keyValueXMLToTrafficInformation(v), err
}

// Start starts given Vps
func Start(c gotransip.Client, vpsName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "start",
	}
	sr.AddArgument("vpsName", vpsName)

	return c.Call(sr, nil)
}

// Stop stops given Vps
func Stop(c gotransip.Client, vpsName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "stop",
	}
	sr.AddArgument("vpsName", vpsName)

	return c.Call(sr, nil)
}

// Reset resets given Vps
func Reset(c gotransip.Client, vpsName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "reset",
	}
	sr.AddArgument("vpsName", vpsName)

	return c.Call(sr, nil)
}

// CreateSnapshot creates a snapshot for Vps with given description
func CreateSnapshot(c gotransip.Client, vpsName, description string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "createSnapshot",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("description", description)

	return c.Call(sr, nil)
}

// RevertSnapshot reverts snapshot with snapshotName for Vps
func RevertSnapshot(c gotransip.Client, vpsName, snapshotName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "revertSnapshot",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("snapshotName", snapshotName)

	return c.Call(sr, nil)
}

// RevertSnapshotToOtherVps uses a snapshot from sourceVpsName to revert to
// destinationVpsName
func RevertSnapshotToOtherVps(c gotransip.Client, sourceVpsName, snapshotName, destinationVpsName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "revertSnapshotToOtherVps",
	}
	sr.AddArgument("sourceVpsName", sourceVpsName)
	sr.AddArgument("snapshotName", snapshotName)
	sr.AddArgument("destinationVpsName", destinationVpsName)

	return c.Call(sr, nil)
}

// RemoveSnapshot removes snapshot by snapshotName for Vps
func RemoveSnapshot(c gotransip.Client, vpsName, snapshotName string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "removeSnapshot",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("snapshotName", snapshotName)

	return c.Call(sr, nil)
}

// RevertVpsBackup reverts backup for Vps
func RevertVpsBackup(c gotransip.Client, vpsName string, vpsBackupID int64) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "revertVpsBackup",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("vpsBackupId", fmt.Sprintf("%d", vpsBackupID))

	return c.Call(sr, nil)
}

// GetVps returns Vps for given name or error if when it failed to retrieve
// Vps from API
func GetVps(c gotransip.Client, vpsName string) (Vps, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v Vps
	err := c.Call(sr, &v)

	return v, err
}

// GetVpses returns Vpses or error when it failed to retrieve list of Vpses
// from API
func GetVpses(c gotransip.Client) ([]Vps, error) {
	var v struct {
		V []Vps `xml:"item"`
	}
	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getVpses",
	}, &v)

	return v.V, err
}

// GetSnapshotsByVps gets all Snapshots for a Vps
func GetSnapshotsByVps(c gotransip.Client, vpsName string) ([]Snapshot, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getSnapshotsByVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []Snapshot `xml:"item"`
	}

	err := c.Call(sr, &v)

	return v.V, err
}

// GetVpsBackupsByVps gets all Backups for a Vps
func GetVpsBackupsByVps(c gotransip.Client, vpsName string) ([]Backup, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getVpsBackupsByVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []Backup `xml:"item"`
	}

	err := c.Call(sr, &v)

	return v.V, err
}

// GetOperatingSystems returns a list of all OperatingSystem available for
// installation
func GetOperatingSystems(c gotransip.Client) ([]OperatingSystem, error) {
	var v struct {
		V []OperatingSystem `xml:"item"`
	}

	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getOperatingSystems",
	}, &v)

	return v.V, err
}

// InstallOperatingSystem installs given operatingSystem on Vps, preseeded by
// hostname
func InstallOperatingSystem(c gotransip.Client, vpsName, operatingSystemName, hostname string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "installOperatingSystem",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("operatingSystemName", operatingSystemName)
	sr.AddArgument("hostname", hostname)

	return c.Call(sr, nil)
}

// InstallOperatingSystemUnattended installs given operatingSystem on Vps, preseeded
// by contents of the installText
func InstallOperatingSystemUnattended(c gotransip.Client, vpsName, operatingSystemName, base64InstallText string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "installOperatingSystemUnattended",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("operatingSystemName", operatingSystemName)
	sr.AddArgument("base64InstallText", base64InstallText)

	return c.Call(sr, nil)
}

// GetIpsForVps returns all IP addresses bound to given VPS
func GetIpsForVps(c gotransip.Client, vpsName string) ([]net.IP, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getIpsForVps",
	}
	sr.AddArgument("vpsName", vpsName)

	var v struct {
		V []net.IP `xml:"item"`
	}

	err := c.Call(sr, &v)

	return v.V, err
}

// GetAllIps returns IP addresses for all VPSes
func GetAllIps(c gotransip.Client) ([]net.IP, error) {
	var v struct {
		V []net.IP `xml:"item"`
	}

	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAllIps",
	}, &v)

	return v.V, err
}

// AddIpv6ToVps adds IpV6 address to Vps
func AddIpv6ToVps(c gotransip.Client, vpsName string, ipv6Address net.IP) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "addIpv6ToVps",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("ipv6Address", ipv6Address.String())

	return c.Call(sr, nil)
}

// UpdatePtrRecord updates the PTR record for given ipAddress
func UpdatePtrRecord(c gotransip.Client, ipAddress net.IP, ptrRecord string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "updatePtrRecord",
	}
	sr.AddArgument("ipAddress", ipAddress.String())
	sr.AddArgument("ptrRecord", ptrRecord)

	return c.Call(sr, nil)
}

// SetCustomerLock locks or unlocks any actions on given Vps
func SetCustomerLock(c gotransip.Client, vpsName string, enabled bool) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setCustomerLock",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("enabled", fmt.Sprintf("%t", enabled))

	return c.Call(sr, nil)
}

// HandoverVps hands Vps over to target TransIP account
func HandoverVps(c gotransip.Client, vpsName, targetAccountname string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "handoverVps",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("targetAccountname", targetAccountname)

	return c.Call(sr, nil)
}

// GetAvailableAvailabilityZones returns a list of available availability zones
func GetAvailableAvailabilityZones(c gotransip.Client) ([]AvailabilityZone, error) {
	var v struct {
		V []AvailabilityZone `xml:"item"`
	}
	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailableAvailabilityZones",
	}, &v)

	return v.V, err
}

// ConvertBackupToSnapshot converts a backup to a separate snapshot for a VPS
func ConvertBackupToSnapshot(c gotransip.Client, vpsName, description string, vpsBackupID int64) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "convertVpsBackupToSnapshot",
	}
	sr.AddArgument("vpsName", vpsName)
	sr.AddArgument("description", description)
	sr.AddArgument("vpsBackupID", fmt.Sprintf("%d", vpsBackupID))

	return c.Call(sr, nil)
}
