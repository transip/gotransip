package vps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRescueImageRepository_GetAll(t *testing.T) {
	const apiResponse = `{"rescueImages":[{"name":"RescueLinux"},{"name":"RescueBSD"}]}`
	server := mockServer{t: t, expectedURL: "/vps/example-vps/rescue-images", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
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
	server := mockServer{t: t, expectedURL: "/vps/example-vps/rescue-images", expectedMethod: "PATCH", statusCode: 204, expectedRequest: apiRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := RescueImageRepository{Client: *client}

	err := repo.BootRescueImage("example-vps", RescueImageLinux)

	assert.NoError(t, err)
}
