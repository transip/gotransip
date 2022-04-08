package email

// Maillist struct of a maillist
type Maillist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	EmailAddress string   `json:"emailAddress"`
	Entries      []string `json:"entries"`
}

// maillistWrapper struct contains a Maillist in it,
// this is solely used for unmarshalling/marshalling
type maillistWrapper struct {
	Maillist Maillist `json:"list"`
}

// maillistsWrapper struct contains a list of Maillist in it,
// this is solely used for unmarshalling/marshalling
type maillistsWrapper struct {
	Maillists []Maillist `json:"lists"`
}

// CreateMaillistRequest struct of a maillist
type CreateMaillistRequest struct {
	EmailAddress string   `json:"emailAddress"`
	Entries      []string `json:"entries"`
	Name         string   `json:"name"`
}

// UpdateMaillistRequest struct of a maillist
type UpdateMaillistRequest struct {
	EmailAddress string   `json:"emailAddress"`
	Entries      []string `json:"entries"`
}
