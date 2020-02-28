package product

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/authenticator"
	"testing"
)

func getRepository(t *testing.T) Repository {
	config := gotransip.ClientConfiguration{Token: authenticator.DemoToken}
	client, err := gotransip.NewClient(config)
	require.NoError(t, err)

	return Repository{Client: client}
}

func TestRepository_GetAll(t *testing.T) {
	repo := getRepository(t)

	all, err := repo.GetAll()
	require.NoError(t, err)

	assert.Equal(t, 6, len(all.Vps))
}

func TestProductRepository_GetSpecificationsForProduct(t *testing.T) {
	repo := getRepository(t)

	all, err := repo.GetSpecificationsForProduct(Product{Name:"vps-bladevps-x4"})
	require.NoError(t, err)

	require.Equal(t, 7, len(all))
	assert.Equal(t, "disk-size", all[0].Name)
	assert.Equal(t, uint64(0x9600000), all[0].Amount)
}
