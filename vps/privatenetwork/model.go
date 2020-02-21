package privatenetwork

// PrivateNetwork struct for PrivateNetwork
type PrivateNetwork struct {
	// The custom name that can be set by customer
	Description string `json:"description"`
	// If the Private Network is administratively blocked
	IsBlocked bool `json:"isBlocked,omitempty"`
	// When locked, another process is already working with this private network
	IsLocked bool `json:"isLocked,omitempty"`
	// The unique private network name
	Name string `json:"name,omitempty"`
	// The VPSes in this private network
	VpsNames []string `json:"vpsNames,omitempty"`
}

// PrivateNetworks struct for PrivateNetworks
type PrivateNetworks struct {
	// array of private networks
	PrivateNetworks []PrivateNetwork `json:"privateNetworks,omitempty"`
}
