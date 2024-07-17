package email

// MailAddon is the struct used to unmarshall the addon response.
type MailAddon struct {
	// ID is the ID of the addon
	ID int `json:"id"`
	// DiskSpace is the amount of disk space the addon provides
	DiskSpace int `json:"diskSpace"`
	// Mailboxes is the amount of extra mailboxes the addon provides
	Mailboxes int `json:"mailboxes"`
	// LinkedMailBox is the mailbox the addon is currently linked to, if any
	LinkedMailBox string `json:"linkedMailBox"`
	// CanBeLinked is whether this Addon is allowed to be linked
	CanBeLinked bool `json:"canBeLinked"`
}

// mailAddonWrapper struct contains a MailAddon in it,
// this is solely used for unmarshalling/marshalling
type mailAddonWrapper struct {
	MailAddons []MailAddon `json:"addons"`
}

// LinkAddonRequest is used to generate a json body that contains the Action, AddonID, and Mailbox properties
type LinkAddonRequest struct {
	// Action could be linkmailbox or unlinkmailbox
	Action string `json:"action"`
	// AddonID is the id of the addon to link
	AddonID int `json:"addonId"`
	// Mailbox is the email address associated with the mailbox
	Mailbox string `json:"mailbox"`
}
