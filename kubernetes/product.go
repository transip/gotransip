package kubernetes

// Product is a Kubernetes product that can be ordered, such as a worker node
// specification, block storage or a load balancer. It describes the resources
// a product of the given type provides through its Specs.
type Product struct {
	// Type of the product, e.g. workerNode, performanceWorkerNode, blockStorage or loadBalancer
	Type string `json:"type"`
	// Name of the product, e.g. node-k8. For worker nodes this matches NodePool.NodeSpec
	Name string `json:"name"`
	// Description is a human readable description of the product
	Description string `json:"description"`
	// Specs describe the resources this product provides
	Specs []ProductSpec `json:"specs"`
	// PeriodPrices contains the prices for the available billing periods
	PeriodPrices []ProductPeriodPrice `json:"periodPrices"`
}

// ProductSpec describes a single resource specification of a Product, such as
// the amount of cpu, memory or disk it provides.
type ProductSpec struct {
	// Name of the spec, e.g. cpu, memory, disk or traffic
	Name string `json:"name"`
	// Amount of the resource, interpreted according to Unit
	Amount int64 `json:"amount"`
	// Unit of the amount, e.g. core, KiB or GiB
	Unit string `json:"unit"`
	// Description is a human readable description of the spec
	Description string `json:"description"`
}

// ProductPeriodPrice describes the price of a Product for a specific billing period.
type ProductPeriodPrice struct {
	// PeriodUnit is the unit of the billing period, e.g. month or hour
	PeriodUnit string `json:"periodUnit"`
	// PeriodLength is the length of the billing period expressed in PeriodUnit
	PeriodLength int `json:"periodLength"`
	// IsExact indicates whether the price is exact for the given period
	IsExact bool `json:"isExact"`
	// Currency is the currency the price is expressed in, e.g. EUR
	Currency string `json:"currency"`
	// CostCents is the price in cents, expressed as a string
	CostCents string `json:"costCents"`
}

type productWrapper struct {
	Product Product `json:"product"`
}

type productsWrapper struct {
	Products []Product `json:"products"`
}
