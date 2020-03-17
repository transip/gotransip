package vps

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/ipaddress"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
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
	restRequest := rest.Request{Endpoint: "/vps"}
	err := r.Client.Get(restRequest, &response)

	return response.Vpss, err
}

// GetByName returns information on a specific VPS by name
func (r *Repository) GetByName(vpsName string) (Vps, error) {
	var response vpsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Vps, err
}

// Order allows you to order a new VPS
func (r *Repository) Order(vpsOrder VpsOrder) error {
	restRequest := rest.Request{Endpoint: "/vps", Body: &vpsOrder}

	return r.Client.Post(restRequest)
}

// OrderMultiple allows you to order multiple vpses at the same time
func (r *Repository) OrderMultiple(orders []VpsOrder) error {
	requestBody := vpssOrderWrapper{Orders: orders}
	restRequest := rest.Request{Endpoint: "/vps", Body: &requestBody}

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
	restRequest := rest.Request{Endpoint: "/vps", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// CloneToAvailabilityZone allows you to clone a vps to a specific availability zone, identified by name
func (r *Repository) CloneToAvailabilityZone(vpsName string, availabilityZone string) error {
	requestBody := cloneRequest{VpsName: vpsName, AvailabilityZone: availabilityZone}
	restRequest := rest.Request{Endpoint: "/vps", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Update allows you to lock/unlock a VPS, update a VPS description, and add/remove tags
// For locking the VPS, set isCustomerLocked to true. Set the value to false for unlocking the VPS
// You can change your VPS description by simply changing the description attribute
// To add/remove tags, you must update the tags attribute
func (r *Repository) Update(vps Vps) error {
	requestBody := vpsWrapper{Vps: vps}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s", vps.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Start allows you to start a VPS, given that it’s currently in a stopped state
func (r *Repository) Start(vpsName string) error {
	requestBody := actionWrapper{Action: "start"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Stop allows you to stop a VPS
func (r *Repository) Stop(vpsName string) error {
	requestBody := actionWrapper{Action: "stop"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Reset allows you to reset a VPS, a reset is essentially the stop and start command combined into one
func (r *Repository) Reset(vpsName string) error {
	requestBody := actionWrapper{Action: "reset"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Handover will handover a VPS to another TransIP Account. This call will initiate the handover process
// The actual handover will be done when the target customer accepts the handover
func (r *Repository) Handover(vpsName string, targetCustomerName string) error {
	requestBody := handoverRequest{Action: "handover", TargetCustomerName: targetCustomerName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// Cancel will cancel the VPS, thus deleting it
func (r *Repository) Cancel(vpsName string, endTime gotransip.CancellationTime) error {
	requestBody := gotransip.CancellationRequest{EndTime: endTime}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s", vpsName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetUsageDataByVps will allow you to request your vps usage for a specified period and usage type
// for convenience you can also use the GetUsages or GetUsagesLast24Hours
func (r *Repository) GetUsageDataByVps(vpsName string, usageTypes []UsageType, period UsagePeriod) (Usage, error) {
	var response usageWrapper
	var types []string
	for _, usageType := range usageTypes {
		types = append(types, string(usageType))
	}
	requestBody := vpsUsageRequest{Types: strings.Join(types, ","), UsagePeriod: period}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/usage", vpsName), Body: &requestBody}
	err := r.Client.Get(restRequest, &response)

	return response.Usage, err
}

// GetAllUsageDataByVps returns a Usage struct filled with all usage data for the given UsagePeriod
// UsagePeriod is struct containing a start and end unix timestamp
func (r *Repository) GetAllUsageDataByVps(vpsName string, period UsagePeriod) (Usage, error) {
	return r.GetUsageDataByVps(
		vpsName,
		[]UsageType{UsageTypeCpu, UsageTypeDisk, UsageTypeNetwork},
		period,
	)
}

// GetAllUsageDataByVps24Hours returns all usage data for a given Vps within the last 24 hours
func (r *Repository) GetAllUsageDataByVps24Hours(vpsName string) (Usage, error) {
	// always define a period body, this way we don't have to depend on the empty body logic on the api server
	period := UsagePeriod{TimeStart: time.Now().Unix() - 24*3600, TimeEnd: time.Now().Unix()}

	return r.GetAllUsageDataByVps(vpsName, period)
}

// GetVNCData will return VncData about your vps
// It allows you to get the location, token and password in order to connect directly to the VNC console of your VPS
func (r *Repository) GetVNCData(vpsName string) (VncData, error) {
	var response vncDataWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/vnc-data", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.VncData, err
}

// RegenerateVNCToken allows you to regenerate the VNC credentials for a VPS
func (r *Repository) RegenerateVNCToken(vpsName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/vnc-data", vpsName)}

	return r.Client.Patch(restRequest)
}

// GetAddons returns a struct with 'cancellable', 'available' and 'active' addons in it for the given VPS
func (r *Repository) GetAddons(vpsName string) (Addons, error) {
	var response addonsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/addons", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Addons, err
}

// OrderAddons allows you to expand VPS specs with a given list of addons to order
func (r *Repository) OrderAddons(vpsName string, addons []string) error {
	response := addonOrderRequest{Addons: addons}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/addons", vpsName), Body: &response}

	return r.Client.Post(restRequest)
}

// CancelAddon allows you to cancel an add-on by name, specifying the VPS name as well
// Due to technical restrictions (possible dataloss) storage add-ons cannot be cancelled
func (r *Repository) CancelAddon(vpsName string, addon string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/addons/%s", vpsName, addon)}

	return r.Client.Delete(restRequest)
}

// GetUpgrades returns all available product upgrades for a VPS
func (r *Repository) GetUpgrades(vpsName string) ([]product.Product, error) {
	var response upgradesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/upgrades", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Upgrades, err
}

// Upgrade allows you to upgrade a VPS by name and productName
func (r *Repository) Upgrade(vpsName string, productName string) error {
	requestBody := upgradeRequest{ProductName: productName}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/upgrades", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// GetOperatingSystems returns a list of operating systems that you can install on a vps
func (r *Repository) GetOperatingSystems(vpsName string) ([]OperatingSystem, error) {
	var response operatingSystemsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/operating-systems", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.OperatingSystems, err
}

// InstallOperatingSystem allows you to install an operating system to a Vps,
// optionally you can specify a hostname and a base64InstallText,
// which would be the automatic installation configuration of your Vps
// for more information, see: https://api.transip.nl/rest/docs.html#vps-operatingsystems-post
func (r *Repository) InstallOperatingSystem(vpsName string, operatingSystemName string, hostname string, base64InstallText string) error {
	requestBody := installRequest{OperatingSystemName: operatingSystemName, Hostname: hostname, Base64InstallText: base64InstallText}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/operating-systems", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// GetIPAddresses returns all IPv4 and IPv6 addresses attached to the VPS
func (r *Repository) GetIPAddresses(vpsName string) ([]ipaddress.IPAddress, error) {
	var response ipaddress.IPAddressesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.IPAddresses, err
}

// GetIPAddressByAddress returns network information for the specified IP address
func (r *Repository) GetIPAddressByAddress(vpsName string, address net.IP) (ipaddress.IPAddress, error) {
	var response ipAddressWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses/%s", vpsName, address.String())}
	err := r.Client.Get(restRequest, &response)

	return response.IPAddress, err
}

// AddIPv6Address allows you to add an IPv6 address to your VPS
// After adding an IPv6 address, you can set the reverse DNS for this address using the UpdateReverseDNS function
func (r *Repository) AddIPv6Address(vpsName string, address net.IP) error {
	requestBody := addIpRequest{IPAddress: address}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateReverseDNS allows you to update the reverse dns for IPv4 addresses as wal as IPv6 addresses
func (r *Repository) UpdateReverseDNS(vpsName string, ip ipaddress.IPAddress) error {
	requestBody := ipAddressWrapper{IPAddress: ip}
	restRequest := rest.Request{
		Endpoint: fmt.Sprintf("/vps/%s/ip-addresses/%s", vpsName, ip.Address.String()),
		Body:     &requestBody,
	}

	return r.Client.Put(restRequest)
}

// RemoveIPv6Address allows you to remove an IPv6 address from the registered list of IPv6 address within your VPS's `/64` range.
func (r *Repository) RemoveIPv6Address(vpsName string, address net.IP) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/ip-addresses/%s", vpsName, address.String())}

	return r.Client.Delete(restRequest)
}

// GetSnapshots returns a list of Snapshots for a given VPS
func (r *Repository) GetSnapshots(vpsName string) ([]Snapshot, error) {
	var response snapshotsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/snapshots", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Snapshots, err
}

// GetSnapshotByName returns a Snapshot for a VPS given its snapshotName and vpsName
func (r *Repository) GetSnapshotByName(vpsName string, snapshotName string) (Snapshot, error) {
	var response snapshotWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName)}
	err := r.Client.Get(restRequest, &response)

	return response.Snapshot, err
}

// CreateSnapshot allows you to create a snapshot for restoring it at a later time or restoring it to another VPS
// See the function RevertSnapshot for this
func (r *Repository) CreateSnapshot(vpsName string, description string, shouldStartVps bool) error {
	requestBody := createSnapshotRequest{Description: description, ShouldStartVps: shouldStartVps}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/snapshots", vpsName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// RevertSnapshot allows you to revert a snapshot of a vps,
// if you want to revert a snapshot to a different vps you can use the RevertSnapshotToOtherVps method
func (r *Repository) RevertSnapshot(vpsName string, snapshotName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName)}

	return r.Client.Patch(restRequest)
}

// RevertSnapshotToOtherVps allows you to revert a snapshot to a different vps
func (r *Repository) RevertSnapshotToOtherVps(vpsName string, snapshotName string, destinationVps string) error {
	requestBody := revertSnapshotRequest{DestinationVpsName: destinationVps}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// RemoveSnapshot allows you to remove a snapshot from a given VPS
func (r *Repository) RemoveSnapshot(vpsName string, snapshotName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/snapshots/%s", vpsName, snapshotName)}

	return r.Client.Delete(restRequest)
}

// GetBackups allows you to get a list of backups for a given VPS which you can use to revert or convert to snapshot
func (r *Repository) GetBackups(vpsName string) ([]Backup, error) {
	var response backupsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/backups", vpsName)}
	err := r.Client.Get(restRequest, &response)

	return response.Backups, err
}

// RevertBackup allows you to revert a backup
func (r *Repository) RevertBackup(vpsName string, backupId int64) error {
	requestBody := actionWrapper{Action: "revert"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/backups/%d", vpsName, backupId), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// ConvertBackupToSnapshot allows you to convert a backup to a snapshot
func (r *Repository) ConvertBackupToSnapshot(vpsName string, backupId int64, snapshotDescription string) error {
	requestBody := convertBackupRequest{SnapshotDescription: snapshotDescription, Action: "convert"}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/vps/%s/backups/%d", vpsName, backupId), Body: &requestBody}

	return r.Client.Patch(restRequest)
}
