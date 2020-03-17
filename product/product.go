package product

// ProductsResponse object for when calling /products
// used to unpack the rest response and return the encapsulated Products
// this is just used internally for unpacking, this should not be exported
// we want to return Products object not a ProductsResponse
type ProductsResponse struct {
	// Products contains all products as described in struct Products
	Products Products `json:"products"`
}

// Products struct for Products
type Products struct {
	// A list of big storage products
	BigStorage []Product `json:"bigStorage,omitempty"`
	// A list of haip products
	Haip []Product `json:"haip,omitempty"`
	// A list of private network products
	PrivateNetworks []Product `json:"privateNetworks,omitempty"`
	// A list of vps products
	Vps []Product `json:"vps,omitempty"`
	// A list of vps addons
	VpsAddon []Product `json:"vpsAddon,omitempty"`
}

// Product struct for Product
type Product struct {
	// Describes this product
	Description string `json:"description,omitempty"`
	// Name of the product
	Name string `json:"name,omitempty"`
	// Price in cents
	Price int `json:"price,omitempty"`
	// The recurring price for the product in cents
	RecurringPrice int `json:"recurringPrice,omitempty"`
}

// ProductElementsResponse object contains a list of ProductElements in it
// used to unpack the rest response and return the encapsulated ProductElements
// this is just used internal for unpacking, this should not be exported
// we want to return ProductElement objects not a ProductElementsResponse
type ProductElementsResponse struct {
	ProductElements []ProductElement `json:"productElements,omitempty"`
}

// ProductElement struct for ProductElement
type ProductElement struct {
	// Amount
	Amount int64 `json:"amount,omitempty"`
	// Describes this product element
	Description string `json:"description,omitempty"`
	// Name of the product element
	Name string `json:"name,omitempty"`
}
