package mailservice

// mailServiceInformationWrapper struct contains the MailServiceInformation struct in it,
// this is solely used for unmarshalling
type mailServiceInformationWrapper struct {
	MailServiceInformation MailServiceInformation `json:"mailServiceInformation"`
}

// domainNamesWrapper struct contains the domainNames in it which will be posted to the api
// this is solely used for marshalling
type domainNamesWrapper struct {
	DomainNames []string `json:"domainNames"`
}

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
