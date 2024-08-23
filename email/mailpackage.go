package email

// Mailpackage struct of a mailpackage
type Mailpackage struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
}

// mailpackagesWrapper contains a list of Mailpackage in it.
// this is solely used for unmarshalling/marshalling
type mailpackagesWrapper struct {
	MailPackages []Mailpackage `json:"packages"`
}
