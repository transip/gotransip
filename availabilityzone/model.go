package availabilityzone

// AvailabilityZonesResponse object with a list of AvailabilityZones in it
// used to unpack the rest response and return the encapsulated AvailabilityZone objects
// this is just used internal for unpacking, this should not be exported
// we want to return a AvailabilityZones object not a AvailabilityZonesResponse
type AvailabilityZonesResponse struct {
	AvailabilityZones []AvailabilityZone `json:"availabilityZones"`
}

// AvailabilityZone struct for AvailabilityZone
type AvailabilityZone struct {
	// Name of AvailabilityZone
	Name string `json:"name,omitempty"`
	// The 2 letter code for the country the AvailabilityZone is in
	Country string `json:"country,omitempty"`
	// If true this is the default zone new VPSes and clones are created in
	IsDefault bool `json:"isDefault,omitempty"`
}
