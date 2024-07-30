package vps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestRescueImageRepository_GetAll(t *testing.T) {
	const apiResponse = `{"rescueImages":[{"name":"RescueLinux"},{"name":"RescueBSD"}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/rescue-images", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := RescueImageRepository{Client: *client}

	settings, err := repo.GetAll("example-vps")

	assert.NoError(t, err)
	assert.Len(t, settings, 2)

	assert.ElementsMatch(t, settings, []RescueImage{
		{
			Name: "RescueBSD",
		},
		{
			Name: "RescueLinux",
		},
	})
}

func TestRescueImageRepository_BootRescueImage(t *testing.T) {
	const apiRequest = `{"name":"RescueLinux"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/rescue-images", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: apiRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := RescueImageRepository{Client: *client}

	err := repo.BootRescueImage("example-vps", RescueImageLinux)

	assert.NoError(t, err)
}
