package availabilityzone

// AvailabilityZone struct for AvailabilityZone
type AvailabilityZone struct {
	// The 2 letter code for the country the AvailabilityZone is in
	Country string `json:"country,omitempty"`
	// If true this is the default zone new VPSes and clones are created in
	IsDefault bool `json:"isDefault,omitempty"`
	// Name of AvailabilityZone
	Name string `json:"name,omitempty"`
}
