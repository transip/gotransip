package mailservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestRepository_AddDNSEntriesDomains(t *testing.T) {
	expectedRequestBody := `{"domainNames":["example.com","another.com"]}`
	server := testutil.MockServer{T: t, ExpectedMethod: "POST", ExpectedURL: "/mail-service", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()

	repo := Repository{Client: *client}
	exmapleDomains := []string{"example.com", "another.com"}
	err := repo.AddDNSEntriesDomains(exmapleDomains)

	require.NoError(t, err)
}

func TestRepository_GetInformation(t *testing.T) {
	responseBody := `{"mailServiceInformation":{ "username": "test@vps.transip.email", "password": "KgDseBsmWJNTiGww", "usage": 54, "quota": 1000, "dnsTxt": "782d28c2fa0b0bdeadf979e7155a83a15632fcddb0149d510c09fb78a470f7d3" } }`
	server := testutil.MockServer{T: t, ExpectedMethod: "GET", ExpectedURL: "/mail-service", StatusCode: 200, Response: responseBody}
	client, tearDown := server.GetClient()
	defer tearDown()

	repo := Repository{Client: *client}
	mailServiceInfo, err := repo.GetInformation()

	require.NoError(t, err)
	assert.Equal(t, "test@vps.transip.email", mailServiceInfo.Username)
	assert.Equal(t, "KgDseBsmWJNTiGww", mailServiceInfo.Password)
	assert.Equal(t, float32(54), mailServiceInfo.Usage)
	assert.Equal(t, float32(1000), mailServiceInfo.Quota)
	assert.Equal(t, "782d28c2fa0b0bdeadf979e7155a83a15632fcddb0149d510c09fb78a470f7d3", mailServiceInfo.DNSTxt)
}

func TestRepository_RegeneratePassword(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedMethod: "PATCH", ExpectedURL: "/mail-service", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()

	repo := Repository{Client: *client}
	err := repo.RegeneratePassword()

	require.NoError(t, err)
}
