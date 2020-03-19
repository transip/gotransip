package product

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
)

// Repository should be used to select the products you want to use in the other repositories
// for example which product to order when ordering a new Vps
type Repository repository.RestRepository

// GetAll returns Products with a list Product per product group in it
func (r *Repository) GetAll() (Products, error) {
	var response ProductsResponse
	productsRequest := rest.Request{Endpoint: "/products"}
	err := r.Client.Get(productsRequest, &response)

	return response.Products, err
}

// GetSpecificationsForProduct returns the ProductElements for a given Product
func (r *Repository) GetSpecificationsForProduct(product Product) ([]Element, error) {
	var response productElementsResponse
	productRequest := rest.Request{Endpoint: fmt.Sprintf("/products/%s/elements", product.Name)}
	err := r.Client.Get(productRequest, &response)

	return response.ProductElements, err
}
