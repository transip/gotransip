package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/internal/testutil"
	"github.com/transip/gotransip/v6/rest"
)

const (
	productsAPIResponse = `{
  "products": {
    "vps": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "vpsAddon": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "haip": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "bigStorage": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "privateNetworks": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ]
  }
}`
	elementsAPIResponse = `{
  "productElements": [
    { "name": "ipv4Addresses", "description": "description of ipv4Addresses", "amount": 1 },
	{ "name": "disk-size", "description": "description of disk-size", "amount": 157286400 }
  ]
}`
)

const errorResponse = `{"error":"errortest"}`

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

func TestRepository_GetAll(t *testing.T) {
	repo, tearDown := getRepository(t, "/products", 200, productsAPIResponse)
	defer tearDown()

	all, err := repo.GetAll()
	require.NoError(t, err)

	require.Equal(t, 1, len(all.Vps))
	require.Equal(t, 1, len(all.BigStorage))
	require.Equal(t, 1, len(all.Haip))
	require.Equal(t, 1, len(all.PrivateNetworks))
	require.Equal(t, 1, len(all.VpsAddon))

	assert.Equal(t, "example-product-name", all.Vps[0].Name)
	assert.Equal(t, "This is an example product", all.Vps[0].Description)
	assert.Equal(t, 499, all.Vps[0].Price)
	assert.Equal(t, 799, all.Vps[0].RecurringPrice)
}

func TestProductRepository_GetSpecificationsForProduct(t *testing.T) {
	repo, tearDown := getRepository(t, "/products/vps-bladevps-x4/elements", 200, elementsAPIResponse)
	defer tearDown()

	all, err := repo.GetSpecificationsForProduct(Product{Name: "vps-bladevps-x4"})
	require.NoError(t, err)

	require.Equal(t, 2, len(all))
	assert.Equal(t, "disk-size", all[1].Name)
	assert.EqualValues(t, uint64(0x9600000), all[1].Amount)
}

func TestRepository_GetAllError(t *testing.T) {
	repo, tearDown := getRepository(t, "/products", 409, errorResponse)
	defer tearDown()

	products, err := repo.GetAll()

	if assert.Errorf(t, err, "getall server response error not returned") {
		assert.Nil(t, products.Vps)
		assert.Equal(t, &rest.Error{Message: "errortest", StatusCode: 409}, err)
	}
}
