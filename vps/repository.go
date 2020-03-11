package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/ipaddress"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest/request"
	"github.com/transip/gotransip/v6/vps/bigstorage"
	"github.com/transip/gotransip/v6/vps/firewall"
	"github.com/transip/gotransip/v6/vps/privatenetwork"
	"github.com/transip/gotransip/v6/vps/tcpmonitor"
	"net"
	"strings"
	"time"
)

// Repository is the vps repository
// this repository allows you to manage all VPS services for your TransIP account
type Repository repository.RestRepository

// GetAll returns a list of all your VPSs
func (r *Repository) GetAll() ([]Vps, error) {
	var response vpssWrapper
	restRequest := request.RestRequest{Endpoint: "/vps"}
	err := r.Client.Get(restRequest, &response)

	return response.Vpss, err
}

// GetByName returns information on a specific VPS by name
func (r *Repository) GetByName(vpsName string) (Vps, error) {
	var response vpsWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Vps, err
}

// Order allows you to order a new VPS
func (r *Repository) Order(vpsOrder VpsOrder) error {
	restRequest := request.RestRequest{Endpoint: "/vps", Body: &vpsOrder}

	return r.Client.Post(restRequest)
}

// Allows you to order multiple vpses at the same time
func (r *Repository) OrderMultiple(orders []VpsOrder) error {
	requestBody := vpssOrderWrapper{Orders: orders}
	restRequest := request.RestRequest{Endpoint: "/vps", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Clone allows you to clone an existing VPS
// There are a few things to take into account when you want to clone an existing VPS to a new VPS:
// - If the original VPS (which you’re going to clone) is currently locked, the clone will fail;
// - Cloned control panels can be used on the VPS, but as the IP address changes, this does require you to synchronise the new license on the new VPS (licenses are often IP-based);
// - Possibly, your VPS has its network interface(s) configured using (a) static IP(‘s) rather than a dynamic allocation using DHCP. If this is the case, you have to configure the new IP(‘s) on the new VPS. Do note that this is not the case with our pre-installed control panel images;
// - VPS add-ons such as Big Storage aren’t affected by cloning - these will stay attached to the original VPS and can’t be swapped automatically
func (r *Repository) Clone(vpsName string) error {
	requestBody := cloneRequest{VpsName: vpsName}
	restRequest := request.RestRequest{Endpoint: "/vps", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// CloneToAvailabilityZone allows you to clone a vps to a specific availability zone, identified by name
func (r *Repository) CloneToAvailabilityZone(vpsName string, availabilityZone string) error {
	requestBody := cloneRequest{VpsName: vpsName, AvailabilityZone: availabilityZone}
	restRequest := request.RestRequest{Endpoint: "/vps", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Update allows you to lock/unlock a VPS, update a VPS description, and add/remove tags
// For locking the VPS, set isCustomerLocked to true. Set the value to false for unlocking the VPS
// You can change your VPS description by simply changing the description attribute
// To add/remove tags, you must update the tags attribute
func (r *Repository) Update(vps Vps) error {
	requestBody := vpsWrapper{Vps: vps}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s", vps.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Start allows you to start a VPS, given that it’s currently in a stopped state
func (r *Repository) Start(vpsName string) error {
	requestBody := actionWrapper{Action: "start"}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Stop allows you to stop a VPS
func (r *Repository) Stop(vpsName string) error {
	requestBody := actionWrapper{Action: "stop"}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Reset allows you to reset a VPS, a reset is essentially the stop and start command combined into one
func (r *Repository) Reset(vpsName string) error {
	requestBody := actionWrapper{Action: "reset"}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Handover will handover a VPS to another TransIP Account. This call will initiate the handover process
// The actual handover will be done when the target customer accepts the handover
func (r *Repository) Handover(vpsName string, targetCustomerName string) error {
	requestBody := handoverRequest{Action: "handover", TargetCustomerName: targetCustomerName}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Cancel will cancel the VPS, thus deleting it
func (r *Repository) Cancel(vpsName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetUsage will allow you to request your vps usage for a specified period and usage type
// for convenience you can also use the GetUsages or GetUsagesLast24Hours
func (r *Repository) GetUsageDataByVps(vpsName string, usageTypes []VpsUsageType, period UsagePeriod) (Usage, error) {
	var response usageWrapper
	var types []string
	for _, usageType := range usageTypes {
		types = append(types, string(usageType))
	}
	requestBody := vpsUsageRequest{Types: strings.Join(types, ","), UsagePeriod: period}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/usage", vpsName), Body: &requestBody}
	err := r.Client.Get(restRequest, &response)

	return response.Usage, err
}

// GetUsages
func (r *Repository) GetAllUsageDataByVps(vpsName string, period UsagePeriod) (Usage, error) {
	return r.GetUsageDataByVps(
		vpsName,
		[]VpsUsageType{VpsUsageTypeCpu, VpsUsageTypeDisk, VpsUsageTypeNetwork},
		period,
	)
}

func (r *Repository) GetAllUsageDataByVps24Hours(vpsName string) (Usage, error) {
	// always define a period body, this way we don't have to depend on the empty body logic on the api server
	period := UsagePeriod{TimeStart: time.Now().Unix() - 24*3600, TimeEnd: time.Now().Unix()}

	return r.GetAllUsageDataByVps(vpsName, period)
}

// GetVNCData will return VncData about your vps
// It allows you to get the location, token and password in order to connect directly to the VNC console of your VPS
func (r *Repository) GetVNCData(vpsName string) (VncData, error) {
	var response vncDataWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/vnc-data", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.VncData, err
}

// RegenerateVNCToken allows you to regenerate the VNC credentials for a VPS
func (r *Repository) RegenerateVNCToken(vpsName string) error {
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/vnc-data", vpsName)}

	return r.Client.Patch(restRequest)
}

// GetAddons returns a struct with 'cancellable', 'available' and 'active' addons in it for the given VPS
func (r *Repository) GetAddons(vpsName string) (Addons, error) {
	var response addonsWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/addons", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Addons, err
}

// OrderAddons allows you to expand VPS specs with a given list of addons to order
func (r *Repository) OrderAddons(vpsName string, addons []string) error {
	response := addonOrderRequest{Addons: addons}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/addons", vpsName), Body: &response}

	return r.Client.Post(restRequest)
}

// CancelAddon allows you to cancel an add-on by name, specifying the VPS name as well
// Due to technical restrictions (possible dataloss) storage add-ons cannot be cancelled
func (r *Repository) CancelAddon(vpsName string, addon string) error {
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/addons/%s", vpsName, addon)}

	return r.Client.Delete(restRequest)
}

// GetUpgrades returns all available product upgrades for a VPS
func (r *Repository) GetUpgrades(vpsName string) ([]product.Product, error) {
	var response upgradesWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/upgrades", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Upgrades, err
}

// Upgrade allows you to upgrade a VPS by name and productName
func (r *Repository) Upgrade(vpsName string, productName string) error {
	requestBody := upgradeRequest{ProductName: productName}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/upgrades", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// GetOperatingSystems
func (r *Repository) GetOperatingSystems() ([]OperatingSystem, error) {
	var response operatingSystemsWrapper
	restRequest := request.RestRequest{Endpoint: "/vps/placeholder/operating-systems"}
	err := r.Client.Get(restRequest, &response)

	return response.OperatingSystems, err
}

func (r *Repository) InstallOperatingSystem(vpsName string, operatingSystemName string, hostname string, base64InstallText string) error {
	requestBody := installRequest{OperatingSystemName: operatingSystemName, Hostname: hostname, Base64InstallText: base64InstallText}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/operating-systems", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// GetIPAddresses returns will return all IPv4 and IPv6 addresses attached to the VPS
func (r *Repository) GetIPAddresses(vpsName string) ([]ipaddress.IPAddress, error) {
	var response ipaddress.IPAddressesWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.IPAddresses, err
}

// GetIPAddressByAddress returns network information for the specified IP address
func (r *Repository) GetIPAddressByAddress(vpsName string, address net.IP) (ipaddress.IPAddress, error) {
	var response ipAddressWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses/%s", vpsName, address.String())}
	err := r.Client.Get(restRequest, &response)

	return response.IPAddress, err
}

// AddIPv6Address allows you to add an ipv6 address to your vps
// After adding an IPv6 address, you cam set the reverse DNS for this address using the UpdateReverseDNS function
func (r *Repository) AddIPv6Address(vpsName string, address net.IP) error {
	requestBody := addIpRequest{IPAddress: address.String()}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateReverseDNS allows you to update the reverse dns for IPv4 addresses as wal as IPv6 addresses
func (r *Repository) UpdateReverseDNS(vpsName string, ip ipaddress.IPAddress) error {
	requestBody := ipAddressWrapper{IPAddress: ip}
	restRequest := request.RestRequest{
		Endpoint: fmt.Sprintf("/vps/%s/ip-addresses/%s", vpsName, ip.Address.String()),
		Body: &requestBody,
	}

	return r.Client.Put(restRequest)
}

// RemoveIPv6Address allows you to remove an ipv6 address from the registered list of IPv6 address within your VPS's `/64` range.
func (r *Repository) RemoveIPv6Address(vpsName string, address net.IP) error {
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses/%s", vpsName, address.String())}

	return r.Client.Delete(restRequest)
}

// GetSnapshots returns a list of Snapshots for a given VPS
func (r *Repository) GetSnapshots(vpsName string) ([]Snapshot, error) {
	var response snapshotsWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/snapshots", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Snapshots, err
}

// GetSnapshotByName returns a Snapshot for a VPS given its snapshotName and vpsName
func (r *Repository) GetSnapshotByName(vpsName string, snapshotName string) (Snapshot, error) {
	var response snapshotWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName)}
	err := r.Client.Get(restRequest, &response)

	return response.Snapshot, err
}

// CreateSnapshot allows you to create a snapshot for restoring it at a later time or restoring it to another VPS
// See the function RevertSnapshot for this
func (r *Repository) CreateSnapshot(vpsName string, description string, shouldStartVps bool) error {
	requestBody := createSnapshotRequest{Description: description, ShouldStartVps: shouldStartVps}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/snapshots", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// RevertSnapshot allows you to revert a snapshot of a vps,
// if you want to revert a snapshot to a different vps you can use the RevertSnapshotToOtherVps method
func (r *Repository) RevertSnapshot(vpsName string, snapshotName string) error {
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName)}

	return r.Client.Patch(restRequest)
}

// RevertSnapshotToOtherVps allows you to revert a snapshot to a different vps
func (r *Repository) RevertSnapshotToOtherVps(vpsName string, snapshotName string, destinationVps string) error {
	requestBody := revertSnapshotRequest{DestinationVpsName: destinationVps}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// DeleteSnapshot allows you to remove a snapshot from a given VPS
func (r *Repository) DeleteSnapshot(vpsName string, snapshotName string) error {
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName)}

	return r.Client.Delete(restRequest)
}

// GetBackups allows you to get a list of backups for a given VPS which you can use to revert or convert to snapshot
func (r *Repository) GetBackups(vpsName string) ([]Backup, error) {
	var response backupsWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/backups", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Backups, err
}

// RevertBackup allows you to revert a backup
func (r *Repository) RevertBackup(vpsName string, backupId int64) error {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/backups/%d", vpsName, backupId), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// ConvertBackupToSnapshot allows you to convert a backup to a snapshot
func (r *Repository) ConvertBackupToSnapshot(vpsName string, backupId int64, snapshotDescription string) error {
	requestBody := convertBackupRequest{SnapshotDescription: snapshotDescription, Action: "convert"}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/backups/%d", vpsName, backupId), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// GetFirewall returns the state of the current VPS firewall
func (r *Repository) GetFirewall(vpsName string) (firewall.Firewall, error) {
	var response firewallWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/firewall", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Firewall, err
}

// UpdateFirewall allows you to update the state of the firewall
// Enabling it, disabling it
// Adding / removing of ruleSets, updating the whitelists
func (r *Repository) UpdateFirewall(vpsName string, firewall firewall.Firewall) error {
	requestBody := firewallWrapper{Firewall: firewall}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/firewall", vpsName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetPrivateNetworks returns a list of all your private networks
func (r *Repository) GetPrivateNetworks() ([]privatenetwork.PrivateNetwork, error) {
	var response privateNetworksWrapper
	restRequest := request.RestRequest{Endpoint: "/private-networks"}
	err := r.Client.Get(restRequest, &response)

	return response.PrivateNetworks, err
}

// GetPrivateNetworkByName allows you to get a specific PrivateNetwork by name
func (r *Repository) GetPrivateNetworkByName(privateNetworkName string) (privatenetwork.PrivateNetwork, error) {
	var response privateNetworkWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName)}
	err := r.Client.Get(restRequest, &response)

	return response.PrivateNetwork, err
}

// OrderPrivateNetwork allows you to order new private network with a given description
func (r *Repository) OrderPrivateNetwork(description string) error {
	requestBody := privateNetworkOrderRequest{Description: description}
	restRequest := request.RestRequest{Endpoint: "/private-networks", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdatePrivateNetwork allows you to update the private network
// you can change the description by changing the Description field
// on the PrivateNetwork struct Updating it using this function
func (r *Repository) UpdatePrivateNetwork(privateNetwork privatenetwork.PrivateNetwork) error {
	requestBody := privateNetworkWrapper{PrivateNetwork: privateNetwork}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetwork.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// AttachVpsToPrivateNetwork allows you to attach a VPS to a PrivateNetwork
func (r *Repository) AttachVpsToPrivateNetwork(vpsName string, privateNetworkName string) error {
	requestBody := privateNetworkActionwrapper{Action: "attachvps", VpsName: vpsName}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// DetachVpsFromPrivateNetwork allows you to detachvps a VPS from a PrivateNetwork
func (r *Repository) DetachVpsFromPrivateNetwork(vpsName string, privateNetworkName string) error {
	requestBody := privateNetworkActionwrapper{Action: "detachvps", VpsName: vpsName}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// CancelPrivateNetwork allows you to cancel a private network
func (r *Repository) CancelPrivateNetwork(privateNetworkName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/private-networks/%s", privateNetworkName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetBigStorages returns a list of your bigstorages
func (r *Repository) GetBigStorages() ([]bigstorage.BigStorage, error) {
	var response bigStoragesWrapper
	restRequest := request.RestRequest{Endpoint: "/big-storages"}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorages, err
}

// GetBigStorageByName returns a specific BigStorage struct by name
func (r *Repository) GetBigStorageByName(bigStorageName string) (bigstorage.BigStorage, error) {
	var response bigStorageWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorage, err
}

// OrderBigStorage allows you to order a new bigstorage
func (r *Repository) OrderBigStorage(order bigstorage.Order) error {
	restRequest := request.RestRequest{Endpoint: "/big-storages", Body: &order}

	return r.Client.Post(restRequest)
}

// UpgradeBigStorage allows you to upgrade a BigStorage's size or/and to enable off-site backups
func (r *Repository) UpgradeBigStorage(bigStorageName string, size int, offsiteBackups bool) error {
	requestBody := bigStorageUpgradeRequest{Size: size, OffsiteBackups: offsiteBackups}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorageName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// UpdateBigStorage allows you to alter the BigStorage in several ways outlined below:
// - Changing the description of a Big Storage;
// - One Big Storages can only be attached to one VPS at a time;
// - One VPS can have a maximum of 10 bigstorages attached;
// - Set the vpsName property to the VPS name to attach to for attaching Big Storage;
// - Set the vpsName property to null to detach the Big Storage from the currently attached VPS.
func (r *Repository) UpdateBigStorage(bigStorage bigstorage.BigStorage) error {
	requestBody := bigStorageWrapper{BigStorage: bigStorage}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorage.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// DetachVpsFromBigStorage allows you to detach a bigstorage from the vps it is attached to
func (r *Repository) DetachVpsFromBigStorage(bigStorage bigstorage.BigStorage) error {
	bigStorage.VpsName = ""

	return r.UpdateBigStorage(bigStorage)
}

// AttachVpsToBigStorage allows you to attach a given VPS by name to a BigStorage
func (r *Repository) AttachVpsToBigStorage(vpsName string, bigStorage bigstorage.BigStorage) error {
	bigStorage.VpsName = vpsName

	return r.UpdateBigStorage(bigStorage)
}

// CancelBigStorage cancels a bigstorage for the specified endTime.
// You can set the endTime to end or immediately, this has the following implications:
// - end: The Big Storage will be terminated from the end date of the agreement as can be found in the applicable quote;
// - immediately: The Big Storage will be terminated immediately.
func (r *Repository) CancelBigStorage(bigStorageName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s", bigStorageName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetBigStorageBackups returns a list of backups for a specific bigstorage
func (r *Repository) GetBigStorageBackups(bigStorageName string) ([]bigstorage.BigStorageBackup, error) {
	var response bigStorageBackupsWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s/backups", bigStorageName)}
	err := r.Client.Get(restRequest, &response)

	return response.BigStorageBackups, err
}

// RevertBigStorageBackup allows you to revert a bigstorage by bigstorage name and backupId
func (r *Repository) RevertBigStorageBackup(bigStorageName string, backupId int64) error {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s/backups/%d", bigStorageName, backupId), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// GetBigStorageUsage allows you to query your bigstorage usage within a certain period
func (r *Repository) GetBigStorageUsage(bigStorageName string, period UsagePeriod) (UsageDataDisk, error) {
	var response usageDataDiskWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/big-storages/%s/usage", bigStorageName), Body: &period}

	err := r.Client.Get(restRequest, &response)

	return response.Usage, err
}

// This method allows you to get usage statistics for the last 24 hours
func (r *Repository) GetBigStorageUsageLast24Hours(bigStorageName string) (UsageDataDisk, error) {
	// always define a period body, this way we don't have to depend on the empty body logic on the api server
	period := UsagePeriod{TimeStart: time.Now().Unix() - 24*3600, TimeEnd: time.Now().Unix()}

	return r.GetBigStorageUsage(bigStorageName, period)
}

// GetTCPMonitors returns an overview of all existing monitors attached to a VPS
func (r *Repository) GetTCPMonitors(vpsName string) ([]tcpmonitor.TcpMonitor, error) {
	var response tcpMonitorsWrapper
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.TcpMonitors, err
}

// CreateTCPMonitor allows you to create a tcp monitor and specify which ports you would like to monitor
// to get a better grip on which fields exist and which can be changes have a look at the TcpMonitor struct
// or see the documentation: https://api.transip.nl/rest/docs.html#vps-tcp-monitors-post
func (r *Repository) CreateTCPMonitor(vpsName string, tcpMonitor tcpmonitor.TcpMonitor) error {
	requestBody := tcpMonitorWrapper{TcpMonitor: tcpMonitor}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateTCPMonitor allows you to update your monitor settings for a given tcp monitored ip
func (r *Repository) UpdateTCPMonitor(vpsName string, tcpMonitor tcpmonitor.TcpMonitor) error {
	requestBody := tcpMonitorWrapper{TcpMonitor: tcpMonitor}
	restRequest := request.RestRequest{
		Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors/%s", vpsName, tcpMonitor.IPAddress.String()),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// DeleteTCPMonitor allows you to remove a tcp monitor for a specific ip address on a specifc VPS
func (r *Repository) DeleteTCPMonitor(vpsName string, ip net.IP) error {
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/vps/%s/tcp-monitors/%s", vpsName, ip.String())}

	return r.Client.Delete(restRequest)
}

// GetContacts returns a list of all your monitoring contacts
func (r *Repository) GetContacts() ([]tcpmonitor.MonitoringContact, error) {
	var response contactsWrapper
	restRequest := request.RestRequest{Endpoint: "/monitoring-contacts"}
	err := r.Client.Get(restRequest, &response)

	return response.Contacts, err
}

// CreateContact allows you to add a new contact which could be used by the tcp monitoring
func (r *Repository) CreateContact(contact tcpmonitor.MonitoringContact) error {
	requestBody := contactWrapper{Contact: contact}
	restRequest := request.RestRequest{Endpoint: "/monitoring-contacts", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateContact updates the specified contact
func (r *Repository) UpdateContact(contact tcpmonitor.MonitoringContact) error {
	requestBody := contactWrapper{Contact: contact}
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/monitoring-contacts/%d", contact.Id), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// DeleteContact allows you to delete a specific contact by id
func (r *Repository) DeleteContact(contactId int64) error {
	restRequest := request.RestRequest{Endpoint: fmt.Sprintf("/monitoring-contacts/%d", contactId)}

	return r.Client.Delete(restRequest)
}
