package email

// MailList struct of a maillist
type MailList struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	EmailAddress string   `json:"emailAddress"`
	Entries      []string `json:"entries"`
}

// mailListWrapper struct contains a MailList in it,
// this is solely used for unmarshalling/marshalling
type mailListWrapper struct {
	MailList MailList `json:"mailList"`
}

// mailListsWrapper struct contains a list of MailList in it,
// this is solely used for unmarshalling/marshalling
type mailListsWrapper struct {
	MailLists []MailList `json:"mailLists"`
}

// CreateMailListRequest struct of a maillist
type CreateMailListRequest struct {
	EmailAddress string   `json:"emailAddress"`
	Entries      []string `json:"entries"`
	Name         string   `json:"name"`
}

// UpdateMailListRequest struct of a maillist
type UpdateMailListRequest struct {
	MailList MailList `json:"mailList"`
}
