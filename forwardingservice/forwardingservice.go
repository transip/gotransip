package forwardingservice

// ForwardMethod is one of the following strings
// 'iframe', 'location', 'movedPermanently'
type ForwardMethod string

// Definition of all of the possible forward methods
const (
	// ForwardMethodIframe shows the destination in a frame, FrameTitle and FrameIcon are used
	ForwardMethodIframe ForwardMethod = "iframe"
	// ForwardMethodLocation redirects the visitor using a temporary redirect
	ForwardMethodLocation ForwardMethod = "location"
	// ForwardMethodMovedPermanently redirects the visitor using a permanent redirect
	ForwardMethodMovedPermanently ForwardMethod = "movedPermanently"
)

// SSLStatus is one of the following strings
// 'pending', 'active', 'disabled'
type SSLStatus string

// Definition of all of the possible SSL statuses
const (
	// SSLStatusPending is the status of a forwarding service for which SSL is enabled, but the certificate has not been issued yet
	SSLStatusPending SSLStatus = "pending"
	// SSLStatusActive is the status of a forwarding service that can be reached over HTTPS
	SSLStatusActive SSLStatus = "active"
	// SSLStatusDisabled is the status of a forwarding service for which SSL is disabled
	SSLStatusDisabled SSLStatus = "disabled"
)

// sslAction is the action that is sent when enabling or disabling SSL
type sslAction string

const (
	sslActionEnable  sslAction = "enable"
	sslActionDisable sslAction = "disable"
)

// forwardingServiceWrapper struct contains a ForwardingService in it,
// this is solely used for unmarshalling/marshalling
type forwardingServiceWrapper struct {
	ForwardingService ForwardingService `json:"forwardingService"`
}

// forwardingServicesWrapper struct contains a list of ForwardingServices in it,
// this is solely used for unmarshalling/marshalling
type forwardingServicesWrapper struct {
	ForwardingServices []ForwardingService `json:"forwardingServices"`
}

// sslStatusWrapper is used to unmarshal the SSL status of a forwarding service
type sslStatusWrapper struct {
	SSLStatus SSLStatus `json:"sslStatus"`
}

// sslRequest is used to marshal to a request that enables or disables SSL
type sslRequest struct {
	Action sslAction `json:"action"`
}

// emailForwardToWrapper is used to marshal and unmarshal the address email is forwarded to,
// ForwardTo is empty when email is not forwarded
type emailForwardToWrapper struct {
	ForwardTo string `json:"forwardTo"`
}

// ForwardingService struct for a forwarding service,
// which forwards HTTP requests for a domain name to another URL
type ForwardingService struct {
	// Hostname of the forwarding service, this is the domain name
	Hostname string `json:"hostname"`
	// IsBlocked is true if the forwarding service is administratively blocked
	IsBlocked bool `json:"isBlocked,omitempty"`
	// IsLocked is true while another action is running for the forwarding service, it is read only
	IsLocked bool `json:"isLocked,omitempty"`
	// ForwardTo is the URL to forward HTTP requests to
	ForwardTo string `json:"forwardTo"`
	// ForwardMethod is the method used to forward HTTP requests
	ForwardMethod ForwardMethod `json:"forwardMethod"`
	// FrameTitle is the title of the page, only used when ForwardMethod is ForwardMethodIframe
	FrameTitle string `json:"frameTitle"`
	// FrameIcon is the URL of the icon of the page, only used when ForwardMethod is ForwardMethodIframe
	FrameIcon string `json:"frameIcon"`
	// ForwardEverything states whether to forward request methods, headers and the request path
	ForwardEverything bool `json:"forwardEverything"`
	// ForwardSubdomains states whether to forward subdomains to the same subdomains on the destination
	ForwardSubdomains bool `json:"forwardSubdomains"`
	// SSLStatus is the current SSL status of the forwarding service, it is read only
	SSLStatus SSLStatus `json:"sslStatus,omitempty"`
}
