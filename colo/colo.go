package colo

const (
	serviceName string = "ColocationService"
)

// DatacenterVisitor represents a Transip_DataCenterVisitor object as described
// at https://api.transip.nl/docs/transip.nl/class-Transip_DataCenterVisitor.html
type DatacenterVisitor struct {
	Name                    string `xml:"name"`
	ReservationNumber       string `xml:"reservationNumber"`
	AccessCode              string `xml:"accessCode"`
	HasBeenRegisteredBefore bool   `xml:"hasBeenRegisteredBefore"`
}
