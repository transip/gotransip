package sslcertificate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestSslcertificateRepository_GetAll(t *testing.T) {
	const apiResponse = `{"certificates":[{"certificateId":1,"commonName":"example.com","expirationDate":"0000-00-00 00:00:00","status":"active"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssl-certificates", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.Equal(t, 1, all[0].CertificateID)
	assert.Equal(t, "example.com", all[0].CommonName)
	assert.Equal(t, "0000-00-00 00:00:00", all[0].ExpirationDate)
	assert.Equal(t, "active", all[0].Status)
}

func TestSslcertificateRepository_GetById(t *testing.T) {
	const apiResponse = `{"certificate":{"certificateId":1,"commonName":"example.com","expirationDate":"0000-00-00 00:00:00","status":"active"},"_links":[{"rel":"self","link":"https:\/\/127.0.0.1\/v6\/ssl-certificates\/2"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssl-certificates/1", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	sslcertificate, err := repo.GetByID(1)
	require.NoError(t, err)

	assert.Equal(t, 1, sslcertificate.CertificateID)
	assert.Equal(t, "example.com", sslcertificate.CommonName)
	assert.Equal(t, "0000-00-00 00:00:00", sslcertificate.ExpirationDate)
	assert.Equal(t, "active", sslcertificate.Status)
}

func TestSslcertificateRepository_GetDetails(t *testing.T) {
	const apiResponse = `{"certificateDetails":{"company":"Company B.V.","department":"IT","postbox":"springfieldroad 123","address":"springfieldroad 123","zipcode":"2345 BB","city":"The Hague","state":"Noord-Holland","countryCode":"NL","firstName":"Johnny","lastName":"Sins","email":"test@example.com","phoneNumber":"+316 12345678","expirationDate":"1970-01-01 00:00:00","name":"/CN=*.example.com","hash":"12abc4567","version":2,"serialNumber":"0x03DEDF0EBA8C8BE53B082CB002579BCC134E","serialNumberHex":"03DEDF0EBA8C8BE53B082CB002579BCC134E","validFrom":"211012092403Z","validTo":"220110092402Z","validFromTimestamp":1647443313,"validToTimestamp":1647443313,"signatureTypeSN":"RSA-SHA256","signatureTypeLN":"sha256WithRSAEncryption","signatureTypeNID":123,"subjectCommonName":"*.example.com","issuerCountry":"US","issuerOrganization":"Let's Encrypt","issuerCommonName":"R3","keyUsage":"Digital Signature, Key Encipherment","basicConstraints":"CA:FALSE","enhancedKeyUsage":"TLS Web Server Authentication, TLS Web Client Authentication","subjectKeyIdentifier":"A1:B2:C3:D4:E5:F6:G7:H8:I9:J1:K2:L3:M4:N5:O6:P7:Q8:R9:S0:T1","authorityKeyIdentifier":"keyid:A1:B2:C3:D4:E5:F6:G7:H8:I9:J1:K2:L3:M4:N5:O6:P7:Q8:R9:S0:T1","authorityInformationAccess":"OCSP - URI:http://r3.o.lencr.org CA Issuers - URI:http://r3.i.lencr.org/","subjectAlternativeName":"DNS:*.example.com, DNS:example.com","certificatePolicies":"Policy: 1.23.456.7.8.9 Policy: 1.2.3.4.5.6.12345.3.2.1 CPS: http://cps.letsencrypt.org","signedCertificateTimestamp":"Signed Certificate Timestamp: Version : v1 (0x0) Log ID..."}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssl-certificates/1/details", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	sslcertificateDetails, err := repo.GetDetails(1)
	require.NoError(t, err)

	assert.Equal(t, "Company B.V.", sslcertificateDetails.Company)
	assert.Equal(t, "IT", sslcertificateDetails.Department)
	assert.Equal(t, "springfieldroad 123", sslcertificateDetails.Postbox)
	assert.Equal(t, "springfieldroad 123", sslcertificateDetails.Address)
	assert.Equal(t, "2345 BB", sslcertificateDetails.Zipcode)
	assert.Equal(t, "The Hague", sslcertificateDetails.City)
	assert.Equal(t, "Noord-Holland", sslcertificateDetails.State)
	assert.Equal(t, "NL", sslcertificateDetails.CountryCode)
	assert.Equal(t, "Johnny", sslcertificateDetails.FirstName)
	assert.Equal(t, "Sins", sslcertificateDetails.LastName)
	assert.Equal(t, "test@example.com", sslcertificateDetails.Email)
	assert.Equal(t, "+316 12345678", sslcertificateDetails.PhoneNumber)
	assert.Equal(t, "1970-01-01 00:00:00", sslcertificateDetails.ExpirationDate)
	assert.Equal(t, "/CN=*.example.com", sslcertificateDetails.Name)
	assert.Equal(t, "12abc4567", sslcertificateDetails.Hash)
	assert.Equal(t, 2, sslcertificateDetails.Version)
	assert.Equal(t, "0x03DEDF0EBA8C8BE53B082CB002579BCC134E", sslcertificateDetails.SerialNumber)
	assert.Equal(t, "03DEDF0EBA8C8BE53B082CB002579BCC134E", sslcertificateDetails.SerialNumberHex)
	assert.Equal(t, "211012092403Z", sslcertificateDetails.ValidFrom)
	assert.Equal(t, "220110092402Z", sslcertificateDetails.ValidTo)
	assert.Equal(t, 1647443313, sslcertificateDetails.ValidFromTimestamp)
	assert.Equal(t, 1647443313, sslcertificateDetails.ValidToTimestamp)
	assert.Equal(t, "RSA-SHA256", sslcertificateDetails.SignatureTypeSN)
	assert.Equal(t, "sha256WithRSAEncryption", sslcertificateDetails.SignatureTypeLN)
	assert.Equal(t, 123, sslcertificateDetails.SignatureTypeNID)
	assert.Equal(t, "*.example.com", sslcertificateDetails.SubjectCommonName)
	assert.Equal(t, "US", sslcertificateDetails.IssuerCountry)
	assert.Equal(t, "Let's Encrypt", sslcertificateDetails.IssuerOrganization)
	assert.Equal(t, "R3", sslcertificateDetails.IssuerCommonName)
	assert.Equal(t, "Digital Signature, Key Encipherment", sslcertificateDetails.KeyUsage)
	assert.Equal(t, "CA:FALSE", sslcertificateDetails.BasicContraints)
	assert.Equal(t, "TLS Web Server Authentication, TLS Web Client Authentication", sslcertificateDetails.EnhancedKeyUsage)
	assert.Equal(t, "A1:B2:C3:D4:E5:F6:G7:H8:I9:J1:K2:L3:M4:N5:O6:P7:Q8:R9:S0:T1", sslcertificateDetails.SubjectKeyIdentifier)
	assert.Equal(t, "keyid:A1:B2:C3:D4:E5:F6:G7:H8:I9:J1:K2:L3:M4:N5:O6:P7:Q8:R9:S0:T1", sslcertificateDetails.AuthorityKeyIdentifier)
	assert.Equal(t, "OCSP - URI:http://r3.o.lencr.org CA Issuers - URI:http://r3.i.lencr.org/", sslcertificateDetails.AuthorityInformationAccess)
	assert.Equal(t, "DNS:*.example.com, DNS:example.com", sslcertificateDetails.SubjectAlternativeName)
	assert.Equal(t, "Policy: 1.23.456.7.8.9 Policy: 1.2.3.4.5.6.12345.3.2.1 CPS: http://cps.letsencrypt.org", sslcertificateDetails.CertificatePolicies)
	assert.Equal(t, "Signed Certificate Timestamp: Version : v1 (0x0) Log ID...", sslcertificateDetails.SignedCertificateTimestamp)
}

func TestSslcertificateRepository_Order(t *testing.T) {
	const expectedRequestBody = `{"productName":"ssl-certificate-comodo-ev","commonName":"example.com","approverFirstName":"John","approverLastName":"Doe","approverEmail":"example@example.com","approverPhone":"+31 715241919","company":"Example B.V.","department":"Example","kvk":"83057825","address":"Easy street 12","city":"Leiden","zipCode":"1337 XD","countryCode":"nl"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssl-certificates", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	request := OrderSSLCertificateRequest{
		ProductName:       "ssl-certificate-comodo-ev",
		CommonName:        "example.com",
		ApproverFirstName: "John",
		ApproverLastName:  "Doe",
		ApproverEmail:     "example@example.com",
		ApproverPhone:     "+31 715241919",
		Company:           "Example B.V.",
		Department:        "Example",
		Kvk:               "83057825",
		Address:           "Easy street 12",
		City:              "Leiden",
		ZipCode:           "1337 XD",
		CountryCode:       "nl",
	}

	err := repo.Order(request)
	require.NoError(t, err)
}

func TestSslcertificateRepository_Download(t *testing.T) {
	const apiResponse = `{"certificateData": {"caBundleCrt": "ca-bundle-crt","certificateCrt": "certificate-crt","certificateP7b": "certificate-p7b","certificateKey": "certificate-key"}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssl-certificates/1/download", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	sslcertificate, err := repo.Download(1)
	require.NoError(t, err)

	assert.Equal(t, "ca-bundle-crt", sslcertificate.CaBundleCrt)
	assert.Equal(t, "certificate-crt", sslcertificate.CertificateCrt)
	assert.Equal(t, "certificate-p7b", sslcertificate.CertificateP7b)
	assert.Equal(t, "certificate-key", sslcertificate.CertificateKey)
}
