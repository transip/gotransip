package colocation

// Colocation struct for Colocation
type Colocation struct {
	// List of IP ranges
	IpRanges []map[string]interface{} `json:"ipRanges"`
	// Colocation name
	Name string `json:"name"`
}

// Colocations struct for Colocations
type Colocations struct {
	// array of Colocations
	Colocations []Colocation `json:"colocations"`
}

// RemoteHands struct for RemoteHands
type RemoteHands struct {
	// Name of the colocation contract
	ColoName string `json:"coloName,omitempty"`
	// Name of the person that created the remote hands request
	ContactName string `json:"contactName,omitempty"`
	// Expected duration in minutes
	ExpectedDuration float32 `json:"expectedDuration,omitempty"`
	// The instructions for the datacenter engineer to perform
	Instructions string `json:"instructions,omitempty"`
	// Phonenumber to contact in case of questions regarding the remotehands request
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// DataCenterVisitor struct for DataCenterVisitor
type DataCenterVisitor struct {
	// The accesscode of the visitor.
	AccessCode string `json:"accessCode"`
	// True iff this visitor been registered before at the datacenter. If true, does not need the accesscode
	HasBeenRegisteredBefore string `json:"hasBeenRegisteredBefore"`
	// The name of the visitor
	Name string `json:"name"`
	// The reservation number of the visitor.
	ReservationNumber string `json:"reservationNumber"`
}
