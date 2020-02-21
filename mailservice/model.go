package mailservice

// MailServiceInformation struct for MailServiceInformation
type MailServiceInformation struct {
	// x-transip-mail-auth DNS TXT record Value
	DnsTxt string `json:"dnsTxt,omitempty"`
	// The password of the mail service
	Password string `json:"password,omitempty"`
	// The quota of the mail service
	Quota float32 `json:"quota,omitempty"`
	// The usage of the mail service
	Usage float32 `json:"usage,omitempty"`
	// The username of the mail service
	Username string `json:"username,omitempty"`
}

// MailService struct for MailService
type MailService struct {
	MailServiceInformation MailServiceInformation `json:"mailServiceInformation,omitempty"`
}
