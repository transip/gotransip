package product

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

// Elements struct for Elements
type Elements struct {
	// different elements for a product with the amount that it comes with
	Elements []ProductElement `json:"elements"`
}

// Product struct for Product
type Product struct {
	// Describes this product
	Description string `json:"description,omitempty"`
	// Name of the product
	Name string `json:"name,omitempty"`
	// Price in cents
	Price int64 `json:"price,omitempty"`
	// The recurring price for the product in cents
	RecurringPrice int64 `json:"recurringPrice,omitempty"`
}

// ProductElement struct for ProductElement
type ProductElement struct {
	// Amount
	Amount uint64 `json:"amount,omitempty"`
	// Describes this product element
	Description string `json:"description,omitempty"`
	// Name of the product element
	Name string `json:"name,omitempty"`
}
