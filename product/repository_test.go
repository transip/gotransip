package product

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	productsApiResponse = `{
  "products": {
    "vps": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "vpsAddon": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "haip": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "bigStorage": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ],
    "privateNetworks": [ { "name": "example-product-name", "description": "This is an example product", "price": 499, "recurringPrice": 799 } ]
  }
}`
	elementsApiResponse = `{
  "productElements": [
    { "name": "ipv4Addresses", "description": "description of ipv4Addresses", "amount": 1 },
	{ "name": "disk-size", "description": "description of disk-size", "amount": 157286400 }
  ]
}`
)

const errorResponse = `{"error":"errortest"}`

func getMockServer(t *testing.T, url string, method string, statusCode int, response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, url, req.URL.String()) // check if right url is called
		assert.Equal(t, method, req.Method)    // check if the right request method is used
		rw.WriteHeader(statusCode)             // respond with given status code
		rw.Write([]byte(response))
	}))
}

func getRepository(t *testing.T, url string, responseStatusCode int, response string) (Repository, func()) {
	server := getMockServer(t, url, "GET", responseStatusCode, response)
	config := gotransip.ClientConfiguration{DemoMode: true, URL: server.URL}
	client, err := gotransip.NewClient(config)
	require.NoError(t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		server.Close()
	}

	return Repository{Client: client}, tearDown
}

func TestRepository_GetAll(t *testing.T) {
	repo, tearDown := getRepository(t, "/products", 200, productsApiResponse)
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
	repo, tearDown := getRepository(t, "/products/vps-bladevps-x4/elements", 200, elementsApiResponse)
	defer tearDown()

	all, err := repo.GetSpecificationsForProduct(Product{Name: "vps-bladevps-x4"})
	require.NoError(t, err)

	require.Equal(t, 2, len(all))
	assert.Equal(t, "disk-size", all[1].Name)
	assert.Equal(t, uint64(0x9600000), all[1].Amount)
}

func TestRepository_GetAllError(t *testing.T) {
	repo, tearDown := getRepository(t, "/products", 409, errorResponse)
	defer tearDown()

	products, err := repo.GetAll()
	require.Error(t, err)
	assert.Nil(t, products.Vps)
	assert.Equal(t, errors.New("errortest"), err)
}
