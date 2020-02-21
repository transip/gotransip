package inline_objects

import (
	"github.com/transip/gotransip/v6/colocation"
	"github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/haip"
	"github.com/transip/gotransip/v6/invoice"
	"github.com/transip/gotransip/v6/ipaddress"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/vps"
	"github.com/transip/gotransip/v6/vps/firewall"
	"github.com/transip/gotransip/v6/vps/privatenetwork"
	"github.com/transip/gotransip/v6/vps/tcpmonitor"
	"net"
)

// InlineObject32 struct for InlineObject32
type InlineObject32 struct {
	PrivateNetwork privatenetwork.PrivateNetwork `json:"privateNetwork,omitempty"`
}

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Usage []vps.Usage `json:"usage,omitempty"`
}

// InlineResponse20024 struct for InlineResponse20024
type InlineResponse20024 struct {
	Haip InlineResponse20024Haip `json:"haip,omitempty"`
}

// InlineObject11 struct for InlineObject11
type InlineObject11 struct {
	// Cancellation time, either 'end' or 'immediately'
	EndTime string `json:"endTime,omitempty"`
}

// InlineObject17 struct for InlineObject17
type InlineObject17 struct {
	DnsEntry domain.DomainsDomainNameDnsDnsEntry `json:"dnsEntry,omitempty"`
}

// InlineResponse20017 struct for InlineResponse20017
type InlineResponse20017 struct {
	DnsEntries []map[string]interface{} `json:"dnsEntries,omitempty"`
}

// InlineResponse20040 struct for InlineResponse20040
type InlineResponse20040 struct {
	Tld InlineResponse20040Tld `json:"tld,omitempty"`
}

// InlineObject26 struct for InlineObject26
type InlineObject26 struct {
	// The mode determining how traffic between our load balancers and your VPS is protected: ‘off’, ‘on’, ‘strict’
	EndpointSslMode string `json:"endpointSslMode,omitempty"`
	// The mode determining how traffic is processed and forwarded: ‘tcp’, ‘http’, ‘https’, ‘proxy’
	Mode string `json:"mode,omitempty"`
	// Name of the port configuration
	Name string `json:"name,omitempty"`
	// The port at which traffic arrives on your HA-IP
	SourcePort float32 `json:"sourcePort,omitempty"`
	// The port at which traffic arrives on your attached IP address(es)
	TargetPort float32 `json:"targetPort,omitempty"`
}

// InlineResponse20045VpsFirewall struct for InlineResponse20045VpsFirewall
type InlineResponse20045VpsFirewall struct {
	// Whether the firewall is enabled for this VPS
	IsEnabled bool `json:"isEnabled"`
	// Ruleset of the VPS
	RuleSet []map[string]interface{} `json:"ruleSet"`
}

// InlineObject42 struct for InlineObject42
type InlineObject42 struct {
	IpAddress string `json:"ipAddress,omitempty"`
}

// InlineResponse20048 struct for InlineResponse20048
type InlineResponse20048 struct {
	Snapshot InlineResponse20048Snapshot `json:"snapshot,omitempty"`
}

// InlineObject48 struct for InlineObject48
type InlineObject48 struct {
	TcpMonitor tcpmonitor.TcpMonitor `json:"tcpMonitor,omitempty"`
}

// InlineResponse2009IpAddress struct for InlineResponse2009IpAddress
type InlineResponse2009IpAddress struct {
	// The IP address
	Address string `json:"address,omitempty"`
	// The TransIP DNS resolvers you can use
	DnsResolvers []net.IP `json:"dnsResolvers,omitempty"`
	// Gateway
	Gateway string `json:"gateway,omitempty"`
	// Reverse DNS, also known as the PTR record
	ReverseDns string `json:"reverseDns"`
	// Subnet mask
	SubnetMask string `json:"subnetMask,omitempty"`
}

// InlineObject21 struct for InlineObject21
type InlineObject21 struct {
	// Optional description for the the HA-IP, this you can use to identify your HA-IP once created
	Description string `json:"description,omitempty"`
	// Required HA-IP product name to order
	ProductName string `json:"productName,omitempty"`
}

// InlineResponse20041TrafficInformation struct for InlineResponse20041TrafficInformation
type InlineResponse20041TrafficInformation struct {
	// The end date in 'Y-m-d' format
	EndDate string `json:"endDate"`
	// The maximum amount of bytes that can be used in this period
	MaxInBytes float32 `json:"maxInBytes"`
	// The start date in 'Y-m-d' format
	StartDate string `json:"startDate"`
	// The usage in bytes for this period
	UsedInBytes float32 `json:"usedInBytes"`
	// The usage in bytes
	UsedTotalBytes float32 `json:"usedTotalBytes"`
}

// InlineResponse20045 struct for InlineResponse20045
type InlineResponse20045 struct {
	VpsFirewall InlineResponse20045VpsFirewall `json:"vpsFirewall,omitempty"`
}

// InlineObject33 struct for InlineObject33
type InlineObject33 struct {
	// Cancellation time, either 'end' (default) or 'immediately'
	EndTime string `json:"endTime,omitempty"`
}

// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	IpAddress ipaddress.IpAddress `json:"ipAddress,omitempty"`
}

// InlineResponse20011 struct for InlineResponse20011
type InlineResponse20011 struct {
	Availability InlineResponse20011Availability `json:"availability,omitempty"`
}

// InlineResponse20035 struct for InlineResponse20035
type InlineResponse20035 struct {
	PrivateNetworks []map[string]interface{} `json:"privateNetworks,omitempty"`
}

// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// The IP address to register to the colocation
	IpAddress string `json:"ipAddress,omitempty"`
	// Reverse DNS, also known as the PTR record
	ReverseDns string `json:"reverseDns,omitempty"`
}

// InlineResponse20023 struct for InlineResponse20023
type InlineResponse20023 struct {
	Haips []map[string]interface{} `json:"haips,omitempty"`
}

// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// The end date of the usage statistics
	DateTimeEnd float32 `json:"dateTimeEnd,omitempty"`
	// The start date of the usage statistics
	DateTimeStart float32 `json:"dateTimeStart,omitempty"`
}

// InlineResponse2007 struct for InlineResponse2007
type InlineResponse2007 struct {
	Colocation InlineResponse2007Colocation `json:"colocation,omitempty"`
}

// InlineObject struct for InlineObject
type InlineObject struct {
	// The name of the bigstorage to upgrade
	BigStorageName string `json:"bigStorageName,omitempty"`
	// Whether to order offsite backups, omit this to use current value
	OffsiteBackups bool `json:"offsiteBackups,omitempty"`
	// The size of the big storage in TB's, use a multitude of 2. The maximum size is 40.
	Size float32 `json:"size,omitempty"`
}

// InlineResponse20014Action struct for InlineResponse20014Action
type InlineResponse20014Action struct {
	// If this action has failed, this field will be true.
	HasFailed bool `json:"hasFailed,omitempty"`
	// If this action has failed, this field will contain an descriptive message.
	Message string `json:"message,omitempty"`
	// The name of this DomainAction.
	Name string `json:"name"`
}

// InlineObject37 struct for InlineObject37
type InlineObject37 struct {
	// Cancellation time, either 'end' (default) or 'immediately'
	EndTime string `json:"endTime,omitempty"`
}

// InlineResponse20021 struct for InlineResponse20021
type InlineResponse20021 struct {
	Certificate InlineResponse20021Certificate `json:"certificate,omitempty"`
}

// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
	AuthCode    string              `json:"authCode,omitempty"`
	Contacts    []domain.Contact    `json:"contacts,omitempty"`
	DnsEntries  []domain.DnsEntry   `json:"dnsEntries,omitempty"`
	Nameservers []domain.Nameserver `json:"nameservers,omitempty"`
}

// InlineResponse20046 struct for InlineResponse20046
type InlineResponse20046 struct {
	OperatingSystems []map[string]interface{} `json:"operatingSystems,omitempty"`
}

// InlineResponse2007Colocation struct for InlineResponse2007Colocation
type InlineResponse2007Colocation struct {
	// List of IP ranges
	IpRanges []map[string]interface{} `json:"ipRanges"`
	// Colocation name
	Name string `json:"name"`
}

// InlineObject45 struct for InlineObject45
type InlineObject45 struct {
	Description string `json:"description,omitempty"`
	// Specify whether the VPS should be started immediately after the snapshot was created, default is `true`
	ShouldStartVps bool `json:"shouldStartVps,omitempty"`
}

// InlineObject41 struct for InlineObject41
type FirewallResponse struct {
	VpsFirewall firewall.Firewall `json:"vpsFirewall,omitempty"`
}

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// Cancellation time, either 'end' or 'immediately'
	EndTime string `json:"endTime,omitempty"`
}

// InlineResponse20014 struct for InlineResponse20014
type InlineResponse20014 struct {
	Action InlineResponse20014Action `json:"action,omitempty"`
}

// InlineObject20 struct for InlineObject20
type InlineObject20 struct {
	Nameservers []domain.Nameserver `json:"nameservers,omitempty"`
}

// InlineResponse20022 struct for InlineResponse20022
type InlineResponse20022 struct {
	Whois string `json:"whois,omitempty"`
}

// InlineObject22 struct for InlineObject22
type InlineObject22 struct {
	Haip haip.Haip `json:"haip,omitempty"`
}

// InlineObject31 struct for InlineObject31
type InlineObject31 struct {
	Description string `json:"description,omitempty"`
}

// InlineObject34 struct for InlineObject34
type InlineObject34 struct {
	Action string `json:"action,omitempty"`
	// Name of the vps that you want to detach
	VpsName string `json:"vpsName,omitempty"`
}

// InlineObject43 struct for InlineObject43
type InlineObject43 struct {
	IpAddress ipaddress.IpAddress `json:"ipAddress,omitempty"`
}

// InlineObject28 struct for InlineObject28
type InlineObject28 struct {
	// The domain names to which the DNS entries should be added
	DomainNames []string `json:"domainNames,omitempty"`
}

// InlineObject19 struct for InlineObject19
type InlineObject19 struct {
	DnsSecEntries []domain.DnsEntry `json:"dnsSecEntries,omitempty"`
}

// InlineResponse20044 struct for InlineResponse20044
type InlineResponse20044 struct {
	Addons vps.Addons `json:"addons,omitempty"`
}

// InlineObject35 struct for InlineObject35
type InlineObject35 struct {
	// The name of the availability zone where the clone should be created
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The vps name of the VPS to clone.
	VpsName string `json:"vpsName,omitempty"`
}

// InlineObject39 struct for InlineObject39
type InlineObject39 struct {
	// Addons to be added
	Addons []string `json:"addons,omitempty"`
}

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	BigStorage InlineResponse2003BigStorage `json:"bigStorage,omitempty"`
}

// InlineResponse20020 struct for InlineResponse20020
type InlineResponse20020 struct {
	Certificates []map[string]interface{} `json:"certificates,omitempty"`
}

// InlineResponse20013Domain struct for InlineResponse20013Domain
type InlineResponse20013Domain struct {
	// The custom tags added to this domain.
	Tags []map[string]interface{} `json:"tags"`
	// The authcode for this domain as generated by the registry.
	AuthCode *string `json:"authCode,omitempty"`
	// Cancellation data, in YYYY-mm-dd h:i:s format, null if the domain is active.
	CancellationDate *string `json:"cancellationDate,omitempty"`
	// Cancellation status, null if the domain is active, 'cancelled' when the domain is cancelled.
	CancellationStatus *string `json:"cancellationStatus,omitempty"`
	// Whether this domain is DNS only
	IsDnsOnly bool `json:"isDnsOnly,omitempty"`
	// If this domain supports transfer locking, this flag is true when the domains ability to transfer is locked at the registry.
	IsTransferLocked bool `json:"isTransferLocked"`
	// If this domain is added to your whitelabel.
	IsWhitelabel bool `json:"isWhitelabel"`
	// The name, including the tld of this domain
	Name string `json:"name"`
	// Registration date of the domain, in YYYY-mm-dd format.
	RegistrationDate string `json:"registrationDate,omitempty"`
	// Next renewal date of the domain, in YYYY-mm-dd format.
	RenewalDate string `json:"renewalDate,omitempty"`
}

// InlineResponse20018 struct for InlineResponse20018
type InlineResponse20018 struct {
	DnsSecEntries []map[string]interface{} `json:"dnsSecEntries,omitempty"`
}

// InlineResponse2006 struct for InlineResponse2006
type InlineResponse2006 struct {
	Colocations []map[string]interface{} `json:"colocations,omitempty"`
}

// InlineResponse20032 struct for InlineResponse20032
type InlineResponse20032 struct {
	Pdf string `json:"pdf,omitempty"`
}

// InlineResponse20052VncData struct for InlineResponse20052VncData
type InlineResponse20052VncData struct {
	// Location of the VNC Proxy
	Host string `json:"host,omitempty"`
	// Password to setup up the VNC connection (changes dynamically)
	Password string `json:"password,omitempty"`
	// Websocket path including the token
	Path string `json:"path,omitempty"`
	// Token to identify the VPS to connect to (changes dynamically)
	Token string `json:"token,omitempty"`
	// Complete websocket URL
	Url string `json:"url,omitempty"`
}

// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	IpAddresses []map[string]interface{} `json:"ipAddresses,omitempty"`
}

// InlineResponse20033 struct for InlineResponse20033
type InlineResponse20033 struct {
	MailServiceInformation InlineResponse20033MailServiceInformation `json:"mailServiceInformation,omitempty"`
}

// InlineObject24 struct for InlineObject24
type InlineObject24 struct {
	// The domain name that the SSL certificate is added to. Start with ‘*.’ when the certificate is a wildcard
	CommonName string `json:"commonName,omitempty"`
}

// InlineObject8 struct for InlineObject8
type InlineObject8 struct {
	// array of domainNames to check
	DomainNames []string `json:"domainNames,omitempty"`
}

// InlineResponse20025 struct for InlineResponse20025
type InlineResponse20025 struct {
	// List of IP addresses attached to this HA-IP
	IpAddresses []map[string]interface{} `json:"ipAddresses,omitempty"`
}

// InlineObject46 struct for InlineObject46
type InlineObject46 struct {
	// When set, revert the snapshot to this VPS
	DestinationVpsName string `json:"destinationVpsName,omitempty"`
}

// InlineObject10 struct for InlineObject10
type InlineObject10 struct {
	Domain domain.DomainsDomainNameDomain `json:"domain,omitempty"`
}

// InlineObject44 struct for InlineObject44
type InlineObject44 struct {
	// Base64 encoded preseed / kickstart instructions, when installing unattended
	Base64InstallText string `json:"base64InstallText,omitempty"`
	// Hostname is required for preinstallable web controlpanels
	Hostname string `json:"hostname,omitempty"`
	// The name of the operating system
	OperatingSystemName string `json:"operatingSystemName,omitempty"`
}

// InlineResponse20047 struct for InlineResponse20047
type InlineResponse20047 struct {
	Snapshots []map[string]interface{} `json:"snapshots,omitempty"`
}

// InlineResponse20043Vps struct for InlineResponse20043Vps
type InlineResponse20043Vps struct {
	// The custom tags added to this VPS
	Tags []map[string]interface{} `json:"tags"`
	// The name of the availability zone the VPS is in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// The VPS cpu count
	Cpus float32 `json:"cpus,omitempty"`
	// The amount of snapshots that is used on this VPS
	CurrentSnapshots float32 `json:"currentSnapshots,omitempty"`
	// The name that can be set by customer
	Description *string `json:"description,omitempty"`
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
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// The product name
	ProductName string `json:"productName,omitempty"`
	// The VPS status, either 'created', 'installing', 'running', 'stopped' or 'paused'
	Status string `json:"status,omitempty"`
}

// InlineResponse20034Contact struct for InlineResponse20034Contact
type InlineResponse20034Contact struct {
	// Email address of the contact
	Email string `json:"email"`
	// Id number of the contact
	Id float32 `json:"id"`
	// Name of the contact
	Name string `json:"name"`
	// Telephone number of the contact
	Telephone string `json:"telephone"`
}

// InlineObject50 struct for InlineObject50
type InlineObject50 struct {
	// The end date of the usage statistics
	DateTimeEnd float32 `json:"dateTimeEnd,omitempty"`
	// The start date of the usage statistics
	DateTimeStart float32 `json:"dateTimeStart,omitempty"`
	// The types of statistics that can be returned, `cpu`, `disk` and `network` can be specified in a comma seperated way. If not specified, all data will be returned.
	Types string `json:"types,omitempty"`
}

// InlineObject15 struct for InlineObject15
type DnsEntriesResponse struct {
	DnsEntries []domain.DnsEntry `json:"dnsEntries,omitempty"`
}

// InlineResponse20011Availability struct for InlineResponse20011Availability
type InlineResponse20011Availability struct {
	// List of available actions to perform on this domain. Possible actions are: 'register', 'transfer', 'internalpull' and 'internalpush'
	Actions []map[string]interface{} `json:"actions"`
	// The name of the domain
	DomainName string `json:"domainName"`
	// The status for this domain. Possible statuses are: 'inyouraccount', 'unavailable', 'notfree', 'free', 'internalpull' and 'internalpush'
	Status string `json:"status"`
}

// InlineResponse20049 struct for InlineResponse20049
type InlineResponse20049 struct {
	TcpMonitors []map[string]interface{} `json:"tcpMonitors,omitempty"`
}

// InlineResponse20044Addons struct for InlineResponse20044Addons
type InlineResponse20044Addons struct {
	// A list of non cancellable active addons
	Active []map[string]interface{} `json:"active,omitempty"`
	// A list of available addons that you can order
	Available []map[string]interface{} `json:"available,omitempty"`
	// A list of addons that you can cancel
	Cancellable []map[string]interface{} `json:"cancellable,omitempty"`
}

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	IpAddress ipaddress.IpAddress `json:"ipAddress,omitempty"`
}

// InlineObject16 struct for InlineObject16
type InlineObject16 struct {
	DnsEntry domain.DomainsDomainNameDnsDnsEntry `json:"dnsEntry,omitempty"`
}

// InlineObject47 struct for InlineObject47
type InlineObject47 struct {
	TcpMonitor tcpmonitor.TcpMonitor `json:"tcpMonitor,omitempty"`
}

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	Action string `json:"action,omitempty"`
}

// InlineResponse20028 struct for InlineResponse20028
type InlineResponse20028 struct {
	StatusReport haip.HaipStatusReport `json:"statusReport,omitempty"`
}

// InlineResponse20036 struct for InlineResponse20036
type InlineResponse20036 struct {
	PrivateNetwork privatenetwork.PrivateNetwork `json:"privateNetwork,omitempty"`
}

// InlineObject30 struct for InlineObject30
type InlineObject30 struct {
	Contact InlineResponse20034Contact `json:"contact,omitempty"`
}

// InlineResponse20027PortConfiguration struct for InlineResponse20027PortConfiguration
type InlineResponse20027PortConfiguration struct {
	// The mode determining how traffic between our load balancers and your attached IP address(es) is encrypted: 'off', 'on', 'strict'
	EndpointSslMode string `json:"endpointSslMode"`
	// The port configuration Id
	Id float32 `json:"id,omitempty"`
	// The mode determining how traffic is processed and forwarded: 'tcp', 'http', 'https', 'proxy'
	Mode string `json:"mode"`
	// A name describing the port
	Name string `json:"name"`
	// The port at which traffic arrives on your HA-IP
	SourcePort float32 `json:"sourcePort"`
	// The port at which traffic arrives on your attached IP address(es)
	TargetPort float32 `json:"targetPort"`
}

// InlineResponse20010 struct for InlineResponse20010
type InlineResponse20010 struct {
	Availability []map[string]interface{} `json:"availability,omitempty"`
}

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	BigStorage InlineResponse2003BigStorage `json:"bigStorage,omitempty"`
}

// InlineResponse20052 struct for InlineResponse20052
type InlineResponse20052 struct {
	VncData InlineResponse20052VncData `json:"vncData,omitempty"`
}

// InlineResponse20051Usage struct for InlineResponse20051Usage
type InlineResponse20051Usage struct {
	Cpu     []vps.VpsUsageDataCpu     `json:"cpu"`
	Disk    []vps.VpsUsageDataDisk    `json:"disk"`
	Network []vps.VpsUsageDataNetwork `json:"network"`
}

// InlineResponse20031 struct for InlineResponse20031
type InlineResponse20031 struct {
	InvoiceItems []invoice.InvoiceItem `json:"invoiceItems,omitempty"`
}

// InlineObject7 struct for InlineObject7
type RemoteHandsResponse struct {
	RemoteHands colocation.RemoteHands `json:"remoteHands,omitempty"`
}

// InlineObject18 struct for InlineObject18
type InlineObject18 struct {
	DnsEntry domain.DomainsDomainNameDnsDnsEntry `json:"dnsEntry,omitempty"`
}

// InlineResponse20050 struct for InlineResponse20050
type InlineResponse20050 struct {
	Upgrades []map[string]interface{} `json:"upgrades,omitempty"`
}

// InlineObject13 struct for InlineObject13
type InlineObject13 struct {
	Branding InlineResponse20015Branding `json:"branding,omitempty"`
}

// InlineResponse20040Tld struct for InlineResponse20040Tld
type InlineResponse20040Tld struct {
	// Number of days a domain needs to be canceled before the renewal date.
	CancelTimeFrame float32 `json:"cancelTimeFrame,omitempty"`
	// A list of the capabilities that this Tld has (the things that can be done with a domain under this tld). Possible capabilities are: 'requiresAuthCode', 'canRegister', 'canTransferWithOwnerChange', 'canTransferWithoutOwnerChange', 'canSetLock', 'canSetOwner', 'canSetContacts', 'canSetNameservers', 'supportsDnsSec'
	Capabilities []map[string]interface{} `json:"capabilities,omitempty"`
	// The maximum amount of characters need for registering a domain under this TLD.
	MaxLength float32 `json:"maxLength,omitempty"`
	// The minimum amount of characters need for registering a domain under this TLD.
	MinLength float32 `json:"minLength,omitempty"`
	// The name of this TLD, including the starting dot. E.g. .nl or .com.
	Name string `json:"name"`
	// Price of the TLD in cents
	Price float32 `json:"price,omitempty"`
	// Price for renewing the TLD in cents
	RecurringPrice float32 `json:"recurringPrice,omitempty"`
	// Length in months of each registration or renewal period.
	RegistrationPeriodLength float32 `json:"registrationPeriodLength,omitempty"`
}

// InlineResponse20027 struct for InlineResponse20027
type InlineResponse20027 struct {
	PortConfiguration InlineResponse20027PortConfiguration `json:"portConfiguration,omitempty"`
}

// InlineResponse20019 struct for InlineResponse20019
type InlineResponse20019 struct {
	Nameservers []map[string]interface{} `json:"nameservers,omitempty"`
}

// InlineResponse20034 struct for InlineResponse20034
type InlineResponse20034 struct {
	Contact InlineResponse20034Contact `json:"contact,omitempty"`
}

// InlineResponse2003BigStorage struct for InlineResponse2003BigStorage
type InlineResponse2003BigStorage struct {
	// The availability zone the bigstorage is located in
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// Name that can be set by customer
	Description string `json:"description"`
	// Disk size of the big storage in kB
	DiskSize float32 `json:"diskSize,omitempty"`
	// Lock status of the big storage, when it is locked, it cannot be attached or detached.
	IsLocked bool `json:"isLocked,omitempty"`
	// Name of the big storage
	Name string `json:"name,omitempty"`
	// Whether a bigstorage has backups
	OffsiteBackups bool `json:"offsiteBackups,omitempty"`
	// Status of the big storage can be 'active', 'attaching' or 'detachting'
	Status string `json:"status,omitempty"`
	// The VPS that the big storage is attached to
	VpsName string `json:"vpsName"`
}

// InlineObject23 struct for InlineObject23
type InlineObject23 struct {
	// Cancellation time, either 'end' or 'immediately'
	EndTime string `json:"endTime,omitempty"`
}

// InlineObject49 struct for InlineObject49
type InlineObject49 struct {
	ProductName string `json:"productName,omitempty"`
}

// InlineObject29 struct for InlineObject29
type InlineObject29 struct {
	// Email address of the contact
	Email string `json:"email,omitempty"`
	// Name of the contact
	Name string `json:"name,omitempty"`
	// Telephone number of the contact
	Telephone string `json:"telephone,omitempty"`
}

// InlineResponse20016 struct for InlineResponse20016
type InlineResponse20016 struct {
	Contacts []map[string]interface{} `json:"contacts,omitempty"`
}

// InlineResponse20038 struct for InlineResponse20038
type ProductElementsResponse struct {
	ProductElements []product.ProductElement `json:"productElements,omitempty"`
}

// InlineResponse20037 struct for InlineResponse20037
type InlineResponse20037 struct {
	Products InlineResponse20037Products `json:"products,omitempty"`
}

// InlineResponse20037Products struct for InlineResponse20037Products
type InlineResponse20037Products struct {
	// A list of big storage products
	BigStorage []map[string]interface{} `json:"bigStorage,omitempty"`
	// A list of haip products
	Haip []map[string]interface{} `json:"haip,omitempty"`
	// A list of private network products
	PrivateNetworks []map[string]interface{} `json:"privateNetworks,omitempty"`
	// A list of vps products
	Vps []map[string]interface{} `json:"vps,omitempty"`
	// A list of vps addons
	VpsAddon []map[string]interface{} `json:"vpsAddon,omitempty"`
}

// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	BigStorages []map[string]interface{} `json:"bigStorages,omitempty"`
}

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	Backups []map[string]interface{} `json:"backups,omitempty"`
}

// InlineResponse20015 struct for InlineResponse20015
type InlineResponse20015 struct {
	Branding InlineResponse20015Branding `json:"branding,omitempty"`
}

// InlineObject27 struct for InlineObject27
type InlineObject27 struct {
	PortConfiguration InlineResponse20027PortConfiguration `json:"portConfiguration,omitempty"`
}

// InlineObject38 struct for InlineObject38
type InlineObject38 struct {
	Action             string `json:"action,omitempty"`
	TargetCustomerName string `json:"targetCustomerName,omitempty"`
}

// InlineResponse20015Branding struct for InlineResponse20015Branding
type InlineResponse20015Branding struct {
	// The first generic bannerLine displayed in whois-branded whois output.
	BannerLine1 string `json:"bannerLine1"`
	// The second generic bannerLine displayed in whois-branded whois output.
	BannerLine2 string `json:"bannerLine2"`
	// The third generic bannerLine displayed in whois-branded whois output.
	BannerLine3 string `json:"bannerLine3"`
	// The company name displayed in transfer-branded e-mails
	CompanyName string `json:"companyName"`
	// The company url displayed in transfer-branded e-mails
	CompanyUrl string `json:"companyUrl"`
	// The support email used for transfer-branded e-mails
	SupportEmail string `json:"supportEmail"`
	// The terms of usage url as displayed in transfer-branded e-mails
	TermsOfUsageUrl *string `json:"termsOfUsageUrl"`
}

// InlineResponse20033MailServiceInformation struct for InlineResponse20033MailServiceInformation
type InlineResponse20033MailServiceInformation struct {
	// x-transip-mail-auth DNS TXT record Value
	DnsTxt string `json:"dnsTxt,omitempty"`
	// The password of the mail service
	Password string `json:"password,omitempty"`
	// The quota of the mail service
	Quota float32 `json:"quota,omitempty"`
	// The usage of the mail service
	Usage float32 `json:"usage,omitempty"`
	// The username of the mail service
	Username string `json:"username,omitempty"`
}

// InlineResponse20039 struct for InlineResponse20039
type InlineResponse20039 struct {
	Tlds []map[string]interface{} `json:"tlds,omitempty"`
}

// InlineObject40 struct for InlineObject40
type InlineObject40 struct {
	Action      string `json:"action,omitempty"`
	Description string `json:"description,omitempty"`
}

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	AvailabilityZones []map[string]interface{} `json:"availabilityZones,omitempty"`
}

// InlineResponse20042 struct for InlineResponse20042
type InlineResponse20042 struct {
	Vpss []map[string]interface{} `json:"vpss,omitempty"`
}

// InlineResponse20048Snapshot struct for InlineResponse20048Snapshot
type InlineResponse20048Snapshot struct {
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

// InlineResponse20013 struct for InlineResponse20013
type InlineResponse20013 struct {
	Domain InlineResponse20013Domain `json:"domain,omitempty"`
}

// InlineResponse20024Haip struct for InlineResponse20024Haip
type InlineResponse20024Haip struct {
	// The description that can be set by the customer
	Description string `json:"description,omitempty"`
	// The interval in milliseconds at which health checks are performed. The interval may not be smaller than 2000ms.
	HealthCheckInterval float32 `json:"healthCheckInterval,omitempty"`
	// The path (URI) of the page to check HTTP status code on
	HttpHealthCheckPath string `json:"httpHealthCheckPath,omitempty"`
	// The port to perform the HTTP check on
	HttpHealthCheckPort float32 `json:"httpHealthCheckPort,omitempty"`
	// Whether to use SSL when performing the HTTP check
	HttpHealthCheckSsl bool `json:"httpHealthCheckSsl,omitempty"`
	// The IPs attached to this haip
	IpAddresses []map[string]interface{} `json:"ipAddresses,omitempty"`
	// HA-IP IP setup: 'both', 'noipv6', 'ipv6to4'
	IpSetup string `json:"ipSetup,omitempty"`
	// HA-IP IPv4 address
	Ipv4Address string `json:"ipv4Address,omitempty"`
	// HA-IP IPv6 address
	Ipv6Address string `json:"ipv6Address,omitempty"`
	// Whether load balancing is enabled for this HA-IP
	IsLoadBalancingEnabled bool `json:"isLoadBalancingEnabled,omitempty"`
	// HA-IP load balancing mode: 'roundrobin', 'cookie', 'source'
	LoadBalancingMode string `json:"loadBalancingMode,omitempty"`
	// HA-IP name
	Name string `json:"name,omitempty"`
	// The PTR record for the HA-IP
	PtrRecord string `json:"ptrRecord,omitempty"`
	// HA-IP status, either 'active', 'inactive', 'creating'
	Status string `json:"status,omitempty"`
	// Cookie name to pin sessions on when using cookie balancing mode
	StickyCookieName string `json:"stickyCookieName,omitempty"`
}

// InlineResponse20041 struct for InlineResponse20041
type InlineResponse20041 struct {
	TrafficInformation InlineResponse20041TrafficInformation `json:"trafficInformation,omitempty"`
}

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Ping string `json:"ping,omitempty"`
}

// InlineResponse20051 struct for InlineResponse20051
type InlineResponse20051 struct {
	Usage InlineResponse20051Usage `json:"usage,omitempty"`
}

// InlineResponse20043 struct for InlineResponse20043
type InlineResponse20043 struct {
	Vps InlineResponse20043Vps `json:"vps,omitempty"`
}

// InlineResponse20021Certificate struct for InlineResponse20021Certificate
type InlineResponse20021Certificate struct {
	// The id of the certificate, can be used to retrieve additional info
	CertificateId float32 `json:"certificateId"`
	// The domain name that the SSL certificate is added to. Start with '*.' when the certificate is a wildcard.
	CommonName string `json:"commonName"`
	// Expiration date
	ExpirationDate string `json:"expirationDate"`
	// The current status, either 'active', 'inactive' or 'expired'
	Status string `json:"status"`
}
