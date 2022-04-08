package sslcertificate

// SSLCertificate struct of a sslcertificate
type SSLCertificate struct {
	CertificateID  int    `json:"certificateId"`
	CommonName     string `json:"commonName"`
	ExpirationDate string `json:"expirationDate"`
	Status         string `json:"status"`
}

// Details struct of SSLCertificate details
type Details struct {
	Company                    string `json:"company"`
	Department                 string `json:"department"`
	Postbox                    string `json:"postbox"`
	Address                    string `json:"address"`
	Zipcode                    string `json:"zipcode"`
	City                       string `json:"city"`
	State                      string `json:"state"`
	CountryCode                string `json:"countryCode"`
	FirstName                  string `json:"firstName"`
	LastName                   string `json:"lastName"`
	Email                      string `json:"email"`
	PhoneNumber                string `json:"phoneNumber"`
	ExpirationDate             string `json:"expirationDate"`
	Name                       string `json:"name"`
	Hash                       string `json:"hash"`
	Version                    int    `json:"version"`
	SerialNumber               string `json:"serialNumber"`
	SerialNumberHex            string `json:"serialNumberHex"`
	ValidFrom                  string `json:"validFrom"`
	ValidTo                    string `json:"validTo"`
	ValidFromTimestamp         int    `json:"validFromTimestamp"`
	ValidToTimestamp           int    `json:"validToTimestamp"`
	SignatureTypeSN            string `json:"signatureTypeSN"`
	SignatureTypeLN            string `json:"signatureTypeLN"`
	SignatureTypeNID           int    `json:"signatureTypeNID"`
	SubjectCommonName          string `json:"subjectCommonName"`
	IssuerCountry              string `json:"issuerCountry"`
	IssuerOrganization         string `json:"issuerOrganization"`
	IssuerCommonName           string `json:"issuerCommonName"`
	KeyUsage                   string `json:"keyUsage"`
	BasicContraints            string `json:"basicConstraints"`
	EnhancedKeyUsage           string `json:"enhancedKeyUsage"`
	SubjectKeyIdentifier       string `json:"subjectKeyIdentifier"`
	AuthorityKeyIdentifier     string `json:"authorityKeyIdentifier"`
	AuthorityInformationAccess string `json:"authorityInformationAccess"`
	SubjectAlternativeName     string `json:"subjectAlternativeName"`
	CertificatePolicies        string `json:"certificatePolicies"`
	SignedCertificateTimestamp string `json:"signedCertificateTimestamp"`
}

// Wrapper struct contains a Sslcertificate in it,
// this is solely used for unmarshalling/marshalling
type wrapper struct {
	Sslcertificate SSLCertificate `json:"certificate"`
}

// SslcertificatesWrapper struct contains a list of Sslcertificates in it,
// this is solely used for unmarshalling/marshalling
type sslcertificatesWrapper struct {
	Sslcertificates []SSLCertificate `json:"certificates"`
}

// DetailsWrapper struct contains a Details in it,
// this is solely used for unmarshalling/marshalling
type detailsWrapper struct {
	Details Details `json:"certificateDetails"`
}

// OrderSSLCertificateRequest struct of a sslcertificate
type OrderSSLCertificateRequest struct {
	ProductName       string `json:"productName"`
	CommonName        string `json:"commonName"`
	ApproverFirstName string `json:"approverFirstName"`
	ApproverLastName  string `json:"approverLastName"`
	ApproverEmail     string `json:"approverEmail"`
	ApproverPhone     string `json:"approverPhone"`
	Company           string `json:"company"`
	Department        string `json:"department"`
	Kvk               string `json:"kvk"`
	Address           string `json:"address"`
	City              string `json:"city"`
	ZipCode           string `json:"zipCode"`
	CountryCode       string `json:"countryCode"`
}
