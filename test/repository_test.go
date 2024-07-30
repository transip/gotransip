package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/internal/testutil"
	"github.com/transip/gotransip/v6/rest"
)

func getRepository(t *testing.T, url string, responseStatusCode int, response string) (Repository, func()) {
	server := testutil.GetMockServer(t, url, "GET", responseStatusCode, response)
	config := gotransip.DemoClientConfiguration
	config.URL = server.URL
	client, err := gotransip.NewClient(config)
	require.NoError(t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		server.Close()
	}

	return Repository{Client: client}, tearDown
}

func TestRepository_Test(t *testing.T) {
	const apiResponse = `{ "ping":"pong" }`

	repo, tearDown := getRepository(t, "/api-test", 200, apiResponse)
	defer tearDown()

	require.NoError(t, repo.Test())
}

func TestRepository_TestError(t *testing.T) {
	const apiResponse = `{ "error":"blablabla" }`

	repo, tearDown := getRepository(t, "/api-test", 409, apiResponse)
	defer tearDown()

	err := repo.Test()

	if assert.Errorf(t, err, "server response error not returned") {
		assert.Equal(t, &rest.Error{Message: "blablabla", StatusCode: 409}, err)
	}
}
