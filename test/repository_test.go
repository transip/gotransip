package test

import (
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/authenticator"
	"testing"
)

func TestRepository_Test(t *testing.T) {
	config := gotransip.ClientConfiguration{Token: authenticator.DemoToken}
	client, err := gotransip.NewClient(config)
	require.NoError(t, err)

	repo := Repository{client}

	require.NoError(t, repo.Test())
}
