package email

// Mailforward struct of a mailforward
type Mailforward struct {
	ID        int    `json:"id"`
	LocalPart string `json:"localPart"`
	Domain    string `json:"domain"`
	ForwardTo string `json:"forwardTo"`
	Status    string `json:"status"`
}

// mailforwardWrapper struct contains a Mailforward in it,
// this is solely used for unmarshalling/marshalling
type mailforwardWrapper struct {
	Mailforward Mailforward `json:"forward"`
}

// mailforwardsWrapper struct contains a list of Mailforward in it,
// this is solely used for unmarshalling/marshalling
type mailforwardsWrappper struct {
	Mailforwards []Mailforward `json:"forwards"`
}

// CreateMailforwardRequest struct of a mailforward
type CreateMailforwardRequest struct {
	ForwardTo string `json:"forwardTo"`
	LocalPart string `json:"localPart"`
}

// UpdateMailforwardRequest struct of a mailforward
type UpdateMailforwardRequest struct {
	ForwardTo string `json:"forwardTo"`
	LocalPart string `json:"localPart"`
}
