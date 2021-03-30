package vps

import (
	"github.com/transip/gotransip/v6/ipaddress"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/rest"
	"net"
)

// BackupStatus is one of the following strings
// 'active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving'
type BackupStatus string

// Definition of all of the possible backup statuses
const (
	// BackupStatusActive is the status field for a ready to use backup
	BackupStatusActive BackupStatus = "active"
	// BackupStatusCreating is the status field for a backup that is still in creation
	BackupStatusCreating BackupStatus = "creating"
	// BackupStatusReverting is the status field for a currently used backup for a revert
	BackupStatusReverting BackupStatus = "reverting"
	// BackupStatusDeleting is the status field for a backup that is about to be deleted
	BackupStatusDeleting BackupStatus = "deleting"
	// BackupStatusPendingDeletion is the status field for a backup that has a pending deletion
	BackupStatusPendingDeletion BackupStatus = "pendingDeletion"
	// BackupStatusSyncing is the status field for a backup that is still syncing
	BackupStatusSyncing BackupStatus = "syncing"
	// BackupStatusMoving is the status field for a moving backup, this means that the backup is under migration
	BackupStatusMoving BackupStatus = "moving"
)

// SnapshotStatus is one of the following strings
// 'active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving'
type SnapshotStatus string

// Definition of all of the possible snapshot statuses
const (
	// SnapshotStatusActive is the status field for an active snapshot, ready to use
	SnapshotStatusActive SnapshotStatus = "active"
	// SnapshotStatusCreating is the status field for a snapshot that is going to be created
	SnapshotStatusCreating SnapshotStatus = "creating"
	// SnapshotStatusReverting is the status field for snapshot that is being reverted
	SnapshotStatusReverting SnapshotStatus = "reverting"
	// SnapshotStatusDeleting is the status field for snapshot that is going to be deleted
	SnapshotStatusDeleting SnapshotStatus = "deleting"
	// SnapshotStatusPendingDeletion is the status field for snapshot that is going to be deleted in the near future
	SnapshotStatusPendingDeletion SnapshotStatus = "pendingDeletion"
	// SnapshotStatusSyncing is the status field for a snapshot that is still under creation
	SnapshotStatusSyncing SnapshotStatus = "syncing"
	// SnapshotStatusMoving is the status field for a snapshot that is moving to another location,
	// this means that the snapshot is under migration
	SnapshotStatusMoving SnapshotStatus = "moving"
)

// Status is one of the following strings
// 'created', 'installing', 'running', 'stopped', 'paused'
type Status string

// Definition of all of the possible vps statuses
const (
	// VpsStatusCreated is the status field for a vps that is created but not yet used
	VpsStatusCreated Status = "created"
	// VpsStatusInstalling is the status field for a vps that is going to be installed
	VpsStatusInstalling Status = "installing"
	// VpsStatusRunning is the status field for a vps that is currently turned on
	VpsStatusRunning Status = "running"
	// VpsStatusStopped is the status field for a vps that is in stopped state
	VpsStatusStopped Status = "stopped"
	// VpsStatusPaused is the status field for a vps that is in paused state
	VpsStatusPaused Status = "paused"
)

// UsageType can be one of the following strings
// 'cpu', 'disk', 'network'
type UsageType string

const (
	// UsageTypeCPU is used to request the cpu usage data of a VPS
	UsageTypeCPU UsageType = "cpu"
	// UsageTypeDisk is used to request the disk usage data of a VPS
	UsageTypeDisk UsageType = "disk"
	// UsageTypeNetwork is used to request the network usage data of a VPS
	UsageTypeNetwork UsageType = "network"
)

// InstallFlavour can be one of the following strings
// 'installer', 'preinstallable', 'cloudinit'
type InstallFlavour string

const (
	// InstallFlavourInstaller is used to explicitly specify that the
	// operating system will be provided through the standard installer
	InstallFlavourInstaller InstallFlavour = "installer"
	// InstallFlavourPreinstallable is used to explicitly specify that the
	// operating system will be provided by a preinstalled image
	InstallFlavourPreinstallable InstallFlavour = "preinstallable"
	// InstallFlavourCloudInit is used to explicitly specify that the
	// operating system will be provided by a cloudinit-enabled image
	InstallFlavourCloudInit InstallFlavour = "cloudinit"
)

// vpsWrapper struct contains a Vps in it,
// this is solely used for unmarshalling/marshalling
type vpsWrapper struct {
	Vps Vps `json:"vps"`
}

// vpssWrapper struct contains a list of Vpses in it,
// this is solely used for unmarshalling/marshalling
type vpssWrapper struct {
	Vpss []Vps `json:"vpss"`
}

// vpssOrderWrapper struct contains a list of VpsOrders in it,
// this is solely used for marshalling
type vpssOrderWrapper struct {
	Orders []Order `json:"vpss"`
}

// cloneRequest is solely used for marshalling a vpsName and an availabilityZone
type cloneRequest struct {
	VpsName          string `json:"vpsName"`
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// actionWrapper struct contains an action in it,
// this is solely used for marshalling
type actionWrapper struct {
	Action string `json:"action"`
}

// handoverRequest is used to request a handover, this is solely used for marshalling
type handoverRequest struct {
	Action             string `json:"action"`
	TargetCustomerName string `json:"targetCustomerName"`
}

// convertBackupRequest is used to request a backup conversion to snapshot,
// this is solely used for marshalling
type convertBackupRequest struct {
	Action              string `json:"action"`
	SnapshotDescription string `json:"description"`
}

// usageWrapper struct contains Usage in it,
// this is solely used for unmarshalling
type usageWrapper struct {
	Usage Usage `json:"usage"`
}

// vncDataWrapper struct contains VncData in it,
// this is solely used for unmarshalling
type vncDataWrapper struct {
	VncData VncData `json:"vncData"`
}

// addonsWrapper struct contains a list with Addons in it,
// this is solely used for unmarshalling
type addonsWrapper struct {
	Addons Addons `json:"addons"`
}

// addonOrderRequest struct contains a list with Addons in it,
// this is solely used for marshalling
type addonOrderRequest struct {
	Addons []string `json:"addons"`
}

// upgradeRequest struct contains a Product Name in it,
// this is solely used for marshalling
type upgradeRequest struct {
	ProductName string `json:"productName"`
}

// upgradesWrapper struct contains a list with Products in it,
// this is solely used for marshalling
type upgradesWrapper struct {
	Upgrades []product.Product `json:"upgrades"`
}

// operatingSystemsWrapper struct contains a list with OperatingSystems in it,
// this is solely used for marshalling
type operatingSystemsWrapper struct {
	OperatingSystems []OperatingSystem `json:"operatingSystems"`
}

// ipAddressWrapper struct contains an IPAddress in it,
// this is solely used for unmarshalling
type ipAddressWrapper struct {
	IPAddress ipaddress.IPAddress `json:"ipAddress"`
}

// snapshotWrapper struct contains a Snapshot in it,
// this is solely used for unmarshalling
type snapshotWrapper struct {
	Snapshot Snapshot `json:"snapshot"`
}

// snapshotWrapper struct contains a list of Snapshots in it,
// this is solely used for unmarshalling
type snapshotsWrapper struct {
	Snapshots []Snapshot `json:"snapshots"`
}

// backupsWrapper struct contains a list of Backups in it,
// this is solely used for unmarshalling
type backupsWrapper struct {
	Backups []Backup `json:"backups"`
}

// firewallWrapper struct contains a Firewall in it,
// this is solely used for marshalling/unmarshalling
type firewallWrapper struct {
	Firewall Firewall `json:"vpsFirewall"`
}

// privateNetworkWrapper struct contains a PrivateNetwork in it,
// this is solely used for marshalling/unmarshalling
type privateNetworkWrapper struct {
	PrivateNetwork PrivateNetwork `json:"privateNetwork"`
}

// privateNetworksWrapper struct contains a PrivateNetwork in it,
// this is solely used for unmarshalling
type privateNetworksWrapper struct {
	PrivateNetworks []PrivateNetwork `json:"privateNetworks"`
}

// privateNetworkActionWrapper struct is used to attach/detach a vps with a private network,
// this is solely used for marshalling
type privateNetworkActionwrapper struct {
	Action  string `json:"action"`
	VpsName string `json:"vpsName"`
}

// privateNetworkOrderRequest struct contains a description in it,
// this is solely used for ordering a private network and encapsulating the description
type privateNetworkOrderRequest struct {
	Description string `json:"description"`
}

// addIPRequest struct contains an IPAddress in it,
// this is solely used for marshalling
type addIPRequest struct {
	IPAddress net.IP `json:"ipAddress"`
}

// createSnapshotRequest is used to marshal a request for creating a snapshot on a vps
// this is solely used for marshalling
type createSnapshotRequest struct {
	Description    string `json:"description"`
	ShouldStartVps bool   `json:"shouldStartVps"`
}

// revertSnapshotRequest is used to marshal a request for reverting a snapshot to a vps
// this is solely used for marshalling
type revertSnapshotRequest struct {
	DestinationVpsName string `json:"destinationVpsName"`
}

// installRequest struct contains a list with OperatingSystems in it,
// this is used for marshalling and aliased by InstallOptions
type installRequest struct {
	// The name of the operating system to install
	OperatingSystemName string `json:"operatingSystemName"`
	// Install flavour to use for providing the operating system (optional)
	// If the field is left empty, the default install flavour for the operating system will be used
	InstallFlavour InstallFlavour `json:"installFlavour,omitempty"`
	// The name for the host, only needed for the preinstallable and cloudinit install flavours
	Hostname string `json:"hostname,omitempty"`
	// Username used for account creating during cloudinit installation (max 32 chars)
	Username string `json:"username,omitempty"`
	// Array of public SSH keys to use for account creating during installation
	// (currently only supported with the cloudinit flavour)
	SSHKeys []string `json:"sshKeys,omitempty"`
	// Base64 encoded preseed / kickstart / cloudinit instructions, when installing unattended
	Base64InstallText string `json:"base64InstallText,omitempty"`
}

// InstallOptions can be used to provide options to
// the InstallOperatingSystemWithOptions method
type InstallOptions installRequest

// bigStorageWrapper struct contains a BigStorage in it,
// this is solely used for marshalling/unmarshalling
type bigStorageWrapper struct {
	BigStorage BigStorage `json:"bigStorage"`
}

// bigStoragesWrapper struct contains a list of BigStorages in it,
// this is solely used for unmarshalling
type bigStoragesWrapper struct {
	BigStorages []BigStorage `json:"bigStorages"`
}

// bigStorageUpgradeRequest struct is used upon when upgrading a bigstorage
// this struct is used for marshalling the request
type bigStorageUpgradeRequest struct {
	BigStorageName string `json:"bigStorageName"`
	Size           int    `json:"size"`
	OffsiteBackups bool   `json:"offsiteBackups"`
}

// bigStorageBackupsWrapper struct contains a list of BigStorageBackups in it,
// this is solely used for unmarshalling
type bigStorageBackupsWrapper struct {
	BigStorageBackups []BigStorageBackup `json:"backups"`
}

// usageDataDiskWrapper struct contains UsageDataDisk struct in it
type usageDataDiskWrapper struct {
	Usage []UsageDataDisk `json:"usage"`
}

// tcpMonitorsWrapper struct is used for unmarshalling a []TCPMonitor list
type tcpMonitorsWrapper struct {
	TCPMonitors []TCPMonitor `json:"tcpMonitors"`
}

// tcpMonitorWrapper struct is used for marshalling/unmarshalling the TCPMonitor struct
type tcpMonitorWrapper struct {
	TCPMonitor TCPMonitor `json:"tcpMonitor"`
}

// contactsWrapper struct is used for unmarshalling a []MonitoringContact list
type contactsWrapper struct {
	Contacts []MonitoringContact `json:"contacts"`
}

// contactWrapper struct is used for marshalling/unmarshalling a MonitoringContact
type contactWrapper struct {
	Contact MonitoringContact `json:"contact"`
}

// Vps struct for a Vps
type Vps struct {
	// The unique VPS name
	Name string `json:"name"`
	// The unique identifier for the VPS
	UUID string `json:"uuid"`
	// The name that can be set by customer
	Description string `json:"description"`
	// The product name
	ProductName string `json:"productName,omitempty"`
	// The VPS OperatingSystem
	OperatingSystem string `json:"operatingSystem,omitempty"`
	// The VPS disk size in kB
	DiskSize int64 `json:"diskSize,omitempty"`
	// The VPS memory size in kB
	MemorySize int64 `json:"memorySize,omitempty"`
	// The VPS cpu count
	CPUs int `json:"cpus,omitempty"`
	// The VPS status, either 'created', 'installing', 'running', 'stopped' or 'paused'
	Status Status `json:"status,omitempty"`
	// The VPS main ipAddress
	IPAddress string `json:"ipAddress,omitempty"`
	// The VPS macaddress
	MacAddress string `json:"macAddress,omitempty"`
	// The amount of snapshots that is used on this VPS
	CurrentSnapshots int `json:"currentSnapshots,omitempty"`
	// The maximum amount of snapshots for this VPS
	MaxSnapshots int `json:"maxSnapshots,omitempty"`
	// Whether or not another process is already doing stuff with this VPS
	IsLocked bool `json:"isLocked,omitempty"`
	// If the VPS is administratively blocked
	IsBlocked bool `json:"isBlocked,omitempty"`
	// If this VPS is locked by the customer
	IsCustomerLocked bool `json:"isCustomerLocked"`
	// The name of the availability zone the VPS is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The custom tags added to this VPS
	Tags []string `json:"tags,omitempty"`
}

// VncData struct for the vps vnc data
type VncData struct {
	// Location of the VNC Proxy
	Host string `json:"host,omitempty"`
	// Password to setup up the VNC connection (changes dynamically)
	Password string `json:"password,omitempty"`
	// Websocket path including the token
	Path string `json:"path,omitempty"`
	// token to identify the VPS to connect to (changes dynamically)
	Token string `json:"token,omitempty"`
	// Complete websocket URL
	URL string `json:"url,omitempty"`
}

// UsageDataNetwork struct for UsageDataNetwork
type UsageDataNetwork struct {
	// Date of the entry, by default in UNIX timestamp format
	Date float32 `json:"date"`
	// The amount of inbound traffic in Mbps for this usage entry
	MbitIn float32 `json:"mbitIn"`
	// The amount of outbound traffic in Mbps for this usage entry
	MbitOut float32 `json:"mbitOut"`
}

// UsagePeriod is struct that can be used to query usage statistics for a certain period
type UsagePeriod struct {
	// TimeStart contains a unix timestamp for the start of the period
	TimeStart int64 `json:"dateTimeStart"`
	// TimeEnd contains a unix timestamp for the end of the period
	TimeEnd int64 `json:"dateTimeEnd"`
}

// UsageDataDisk struct contains disk usage for a certain date
type UsageDataDisk struct {
	// Date of the entry, by default in UNIX timestamp format
	Date int64 `json:"date"`
	// The read IOPS for this entry
	IopsRead float32 `json:"iopsRead"`
	// The write IOPS for this entry
	IopsWrite float32 `json:"iopsWrite"`
}

// UsageDataCPU struct contains cpu usage percentage for a certain date
type UsageDataCPU struct {
	// Date of the entry, by default in UNIX timestamp format
	Date int64 `json:"date"`
	// The percentage of CPU usage for this entry
	Percentage float32 `json:"percentage"`
}

// Order struct can be used to order a new VPS
type Order struct {
	// Name of the product
	ProductName string `json:"productName"`
	// The name of the operating system to install
	OperatingSystem string `json:"operatingSystem"`
	// The name of the availability zone where the vps should be created
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The description of the VPS
	Description string `json:"description,omitempty"`
	// Array with additional addons
	Addons []string `json:"addons,omitempty"`
	// Install flavour to use for providing the operating system (optional)
	// If the field is left empty, the default install flavour for the operating system will be used
	InstallFlavour InstallFlavour `json:"installFlavour,omitempty"`
	// The name for the host, only needed for the preinstallable and cloudinit install flavours
	Hostname string `json:"hostname,omitempty"`
	// Username used for account creating during cloudinit installation (max 32 chars)
	Username string `json:"username,omitempty"`
	// Array of public SSH keys to use for account creating during installation
	// (currently only supported with the cloudinit flavour)
	SSHKeys []string `json:"sshKeys,omitempty"`
	// Base64 encoded preseed / kickstart / cloudinit instructions, when installing unattended
	Base64InstallText string `json:"base64InstallText,omitempty"`
}

// Addons struct for an Addons
type Addons struct {
	// A list of all active addons
	Active []product.Product `json:"active,omitempty"`
	// A list of available addons that you can order
	Available []product.Product `json:"available,omitempty"`
	// A list of addons that you can cancel
	Cancellable []product.Product `json:"cancellable,omitempty"`
}

// Backup struct for a Backup
type Backup struct {
	// The backup id
	ID int64 `json:"id"`
	// Status of the backup ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status BackupStatus `json:"status"`
	// The backup creation date
	DateTimeCreate rest.Time `json:"dateTimeCreate"`
	// The backup disk size in kB
	DiskSize int64 `json:"diskSize"`
	// The backup operatingSystem
	OperatingSystem string `json:"operatingSystem"`
	// The name of the availability zone the backup is in
	AvailabilityZone string `json:"availabilityZone"`
}

// Snapshot struct for a Snapshot
type Snapshot struct {
	// The snapshot creation date
	DateTimeCreate string `json:"dateTimeCreate,omitempty"`
	// The snapshot description
	Description string `json:"description,omitempty"`
	// The size of the snapshot in kB
	DiskSize int64 `json:"diskSize,omitempty"`
	// The snapshot name
	Name string `json:"name,omitempty"`
	// The snapshot OperatingSystem
	OperatingSystem string `json:"operatingSystem,omitempty"`
	// The snapshot status ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status SnapshotStatus `json:"status,omitempty"`
}

// OperatingSystem struct for an OperatingSystem
type OperatingSystem struct {
	// Description
	Description string `json:"description,omitempty"`
	// Is a preinstallable image
	// Deprecated: Use the InstallFlavours field to determine this instead
	IsPreinstallableImage bool `json:"isPreinstallableImage,omitempty"`
	// List of supported install flavours for this operating system
	InstallFlavours []InstallFlavour `json:"installFlavours"`
	// The operating system name
	Name string `json:"name"`
	// The monthly price of the operating system in cents
	Price int `json:"price,omitempty"`
	// The version of the operating system
	Version string `json:"version,omitempty"`
	// available licenses for this operating system
	Licenses []LicenseProduct `json:"licenses,omitempty"`
}

// Usage struct for an Usage
type Usage struct {
	CPU     []UsageDataCPU     `json:"cpu"`
	Disk    []UsageDataDisk    `json:"disk"`
	Network []UsageDataNetwork `json:"network"`
}
