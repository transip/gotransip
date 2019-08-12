package vps

import (
	"net"
	"strconv"
	"time"

	"github.com/transip/gotransip/v5/util"
)

const (
	serviceName string = "VpsService"
)

// Status represents the possibles states a Vps can be in
type Status string

var (
	// StatusCreated means the VPS was provisioned and is waiting for install
	StatusCreated Status = "created"
	// StatusInstalling means the VPS is currently installing it's OS
	StatusInstalling Status = "installing"
	// StatusRunning means the VPS is running and operating normally
	StatusRunning Status = "running"
	// StatusStopped means the VPS was stopped
	StatusStopped Status = "stopped"
	// StatusPaused means the VPS is currently in a suspended state
	StatusPaused Status = "paused"
	// StatusFailPaused means something went wrong and the VPS got suspended
	StatusFailPaused Status = "failPaused"
)

// TrafficInformation holds information about a VPS's traffic usage for a specific
// period
type TrafficInformation struct {
	From  time.Time
	End   time.Time
	Used  int64
	Total int64
	Max   int64
}

// Vps represents a Transip_Vps object
// https://api.transip.nl/docs/transip.nl/class-Transip_Vps.html
type Vps struct {
	Name             string `xml:"name"`
	Description      string `xml:"description"`
	OperatingSystem  string `xml:"operatingSystem"`
	DiskSize         int64  `xml:"diskSize"`
	MemorySize       int64  `xml:"memorySize"`
	Processors       int64  `xml:"cpus"`
	Status           Status `xml:"status"`
	IPv4Address      net.IP `xml:"ipAddress"`
	IPv6Address      net.IP `xml:"ipv6Address"`
	MACAddress       string `xml:"macAddress"`
	IsBlocked        bool   `xml:"isBlocked"`
	IsCustomerLocked bool   `xml:"isCustomerLocked"`
	AvailabilityZone string `xml:"availabilityZone"`
}

// PrivateNetwork represents a Transip_PrivateNetwork object
// https://api.transip.nl/docs/transip.nl/class-Transip_PrivateNetwork.html
type PrivateNetwork struct {
	Name string `xml:"name"`
}

// AvailabilityZone represents an availability zone as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_AvailabilityZone.html
type AvailabilityZone struct {
	Name      string `xml:"name"`
	Country   string `xml:"country"`
	IsDefault bool   `xml:"isDefault"`
}

// Snapshot represents a Transip_Snapshot object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_Snapshot.html
type Snapshot struct {
	Name             string       `xml:"name"`
	Description      string       `xml:"description"`
	Created          util.XMLTime `xml:"dateTimeCreate"`
	AvailabilityZone string       `xml:"availabilityZone"`
}

// Product represents a Transip_Product object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_Product.html
type Product struct {
	Name         string  `xml:"name"`
	Description  string  `xml:"description"`
	Price        float64 `xml:"price"`
	RenewalPrice float64 `xml:"renewalPrice"`
}

// OperatingSystem represents a Transip_OperatingSystem object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_OperatingSystem.html
type OperatingSystem struct {
	Name                  string `xml:"name"`
	Description           string `xml:"description"`
	IsPreinstallableImage bool   `xml:"isPreinstallableImage"`
}

// Backup represents a Transip_VpsBackup object
// https://api.transip.nl/docs/transip.nl/class-Transip_VpsBackup.html
type Backup struct {
	ID               int64        `xml:"id"`
	Created          util.XMLTime `xml:"dateTimeCreate"`
	DiskSize         int64        `xml:"diskSize"`
	OperatingSystem  string       `xml:"operatingSystem"`
	AvailabilityZone string       `xml:"availabilityZone"`
}

// keyValueXMLToTrafficInformation converts KeyValueXML to a object of TrafficInformation
func keyValueXMLToTrafficInformation(k util.KeyValueXML) (ti TrafficInformation) {
	// convert SOAP result into Certificates
	if len(k.Cont) == 0 {
		return
	}

	var err error
	for _, x := range k.Cont[0].Item {
		switch x.Key {
		case "startDate":
			ti.From, err = time.Parse("2006-01-02", x.Value)
			if err != nil {
				ti.From, _ = time.Parse("2006-01-02 15:04:05", x.Value)
			}
		case "endDate":
			ti.End, err = time.Parse("2006-01-02", x.Value)
			if err != nil {
				ti.End, _ = time.Parse("2006-01-02 15:04:05", x.Value)
			}
		case "usedInBytes":
			ti.Used, _ = strconv.ParseInt(x.Value, 10, 64)
		case "usedTotalBytes":
			ti.Total, _ = strconv.ParseInt(x.Value, 10, 64)
		case "maxBytes":
			ti.Max, _ = strconv.ParseInt(x.Value, 10, 64)
		}
	}

	return
}
