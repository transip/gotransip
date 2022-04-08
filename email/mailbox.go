package email

// Mailbox struct of a mailbox
type Mailbox struct {
	Identifier         string `json:"identifier"`
	LocalPart          string `json:"localPart"`
	Domain             string `json:"domain"`
	ForwardTo          string `json:"forwardTo"`
	AvailableDiskSpace int    `json:"availableDiskSpace"`
	UsedDiskSpace      int    `json:"usedDiskSpace"`
	Status             string `json:"status"`
	IsLocked           bool   `json:"isLocked"`
	ImapServer         string `json:"imapServer"`
	ImapPort           int    `json:"imapPort"`
	SMTPServer         string `json:"smtpServer"`
	SMTPPort           int    `json:"smtpPort"`
	Pop3Server         string `json:"pop3Server"`
	Pop3Port           int    `json:"pop3Port"`
}

// mailboxWrapper struct contains a Mailbox in it,
// this is solely used for unmarshalling/marshalling
type mailboxWrapper struct {
	Mailbox Mailbox `json:"Mailbox"`
}

// mailboxesWrapper struct contains a list of Mailboxes in it,
// this is solely used for unmarshalling/marshalling
type mailboxesWrapper struct {
	Mailboxes []Mailbox `json:"Mailboxes"`
}

// CreateMailboxRequest struct of a mailbox
type CreateMailboxRequest struct {
	LocalPart    string `json:"localPart"`
	Password     string `json:"password"`
	MaxDiskUsage int    `json:"maxDiskUsage"`
	ForwardTo    string `json:"forwardTo"`
}

// UpdateMailboxRequest struct of a mailbox
type UpdateMailboxRequest struct {
	Password     string `json:"password"`
	MaxDiskUsage int    `json:"maxDiskUsage"`
	ForwardTo    string `json:"forwardTo"`
}
