package vps

import "github.com/transip/gotransip/v6/product"

// VpsResponse object with a Vps in it
// used to unpack the rest response and return the encapsulated Vps object
// this is just used internal for unpacking, this should not be exported
// we want to return a Vps object not a VpsResponse
type VpsResponse struct {
	Vps Vps `json:"vps"`
}

// Vps struct for Vps
type Vps struct {
	// The custom tags added to this VPS
	Tags []string `json:"tags,omitempty"`
	// The name of the availability zone the VPS is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The VPS cpu count
	Cpus float32 `json:"cpus,omitempty"`
	// The amount of snapshots that is used on this VPS
	CurrentSnapshots float32 `json:"currentSnapshots,omitempty"`
	// The name that can be set by customer
	Description string `json:"description,omitempty"`
	// The VPS disk size in kB
	DiskSize float32 `json:"diskSize,omitempty"`
	// The VPS main ipAddress
	IpAddress string `json:"ipAddress,omitempty"`
	// If the VPS is administratively blocked
	IsBlocked bool `json:"isBlocked,omitempty"`
	// If this VPS is locked by the customer
	IsCustomerLocked bool `json:"isCustomerLocked,omitempty"`
	// Whether or not another process is already doing stuff with this VPS
	IsLocked bool `json:"isLocked,omitempty"`
	// The VPS macaddress
	MacAddress string `json:"macAddress,omitempty"`
	// The maximum amount of snapshots for this VPS
	MaxSnapshots float32 `json:"maxSnapshots,omitempty"`
	// The VPS memory size in kB
	MemorySize float32 `json:"memorySize,omitempty"`
	// The unique VPS name
	Name string `json:"name"`
	// The VPS OperatingSystem
	OperatingSystem string `json:"operatingSystem,omitempty"`
	// The product name
	ProductName string `json:"productName,omitempty"`
	// The VPS status, either 'created', 'installing', 'running', 'stopped' or 'paused'
	Status string `json:"status,omitempty"`
}

// VncData struct for VpsVncData
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
	Url string `json:"url,omitempty"`
}

// VpsUsageDataNetwork struct for VpsUsageDataNetwork
type VpsUsageDataNetwork struct {
	// Date of the entry, by default in UNIX timestamp format
	Date float32 `json:"date"`
	// The amount of inbound traffic in Mbps for this usage entry
	MbitIn float32 `json:"mbitIn"`
	// The amount of outbound traffic in Mbps for this usage entry
	MbitOut float32 `json:"mbitOut"`
}

// VpsUsageDataDisk struct for VpsUsageDataDisk
type VpsUsageDataDisk struct {
	// Date of the entry, by default in UNIX timestamp format
	Date float32 `json:"date"`
	// The read IOPS for this entry
	IopsRead float32 `json:"iopsRead"`
	// The write IOPS for this entry
	IopsWrite float32 `json:"iopsWrite"`
}

// VpsUsageDataCpu struct for VpsUsageDataCpu
type VpsUsageDataCpu struct {
	// Date of the entry, by default in UNIX timestamp format
	Date float32 `json:"date"`
	// The percentage of CPU usage for this entry
	Percentage float32 `json:"percentage"`
}

// VpsOrder struct for VpsOrder
type VpsOrder struct {
	// Array with additional addons
	Addons []string `json:"addons,omitempty"`
	// The name of the availability zone where the vps should be created
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// Base64 encoded preseed / kickstart instructions, when installing unattended
	Base64InstallText string `json:"base64InstallText,omitempty"`
	// The description of the VPS
	Description string `json:"description,omitempty"`
	// The name for the host, only needed for installing a preinstallable control panel image
	Hostname string `json:"hostname,omitempty"`
	// The name of the operating system to install
	OperatingSystem string `json:"operatingSystem"`
	// Name of the product
	ProductName string `json:"productName"`
}

// Addons struct for Addons
type Addons struct {
	// A list of non cancellable active addons
	Active []string `json:"active,omitempty"`
	// A list of available addons that you can order
	Available []string `json:"available,omitempty"`
	// A list of addons that you can cancel
	Cancellable []string `json:"cancellable,omitempty"`
}

// VpsBackup struct for VpsBackup
type VpsBackup struct {
	// The name of the availability zone the backup is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The backup creation date
	DateTimeCreate string `json:"dateTimeCreate"`
	// The backup disk size in kB
	DiskSize float32 `json:"diskSize"`
	// The backup id
	Id float32 `json:"id"`
	// The backup operatingSystem
	OperatingSystem string `json:"operatingSystem"`
	// Status of the backup ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status string `json:"status,omitempty"`
}

// Snapshots struct for Snapshots
type Snapshots struct {
	// The snapshot creation date
	Snapshots []Snapshot `json:"snapshots,omitempty"`
}

// Snapshot struct for Snapshot
type Snapshot struct {
	// The snapshot creation date
	DateTimeCreate string `json:"dateTimeCreate,omitempty"`
	// The snapshot description
	Description string `json:"description,omitempty"`
	// The size of the snapshot in kB
	DiskSize float32 `json:"diskSize,omitempty"`
	// The snapshot name
	Name string `json:"name,omitempty"`
	// The snapshot OperatingSystem
	OperatingSystem string `json:"operatingSystem,omitempty"`
	// The snapshot status ('active', 'creating', 'reverting', 'deleting', 'pendingDeletion', 'syncing', 'moving')
	Status string `json:"status,omitempty"`
}

// OperatingSystems struct for OperatingSystems
type OperatingSystems struct {
	// OperatingSystems
	OperatingSystems string `json:"operatingSystems,omitempty"`
}

// OperatingSystem struct for OperatingSystem
type OperatingSystem struct {
	// Description
	Description string `json:"description,omitempty"`
	// Is a preinstallable image
	IsPreinstallableImage bool `json:"isPreinstallableImage,omitempty"`
	// The operating system name
	Name string `json:"name"`
	// The monthly price of the operating system in cents
	Price float32 `json:"price,omitempty"`
	// The version of the operating system
	Version string `json:"version,omitempty"`
}

// Usage struct for Usage
type Usage struct {
	Cpu     []VpsUsageDataCpu     `json:"cpu"`
	Disk    []VpsUsageDataDisk    `json:"disk"`
	Network []VpsUsageDataNetwork `json:"network"`
}

// Upgrades struct for Upgrades
type Upgrades struct {
	Upgrades []product.Product `json:"upgrades"`
}
