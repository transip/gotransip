package invoice

type InvoicesResponse struct {
	Invoices []Invoice `json:"invoices,omitempty"`
}

type InvoiceResponse struct {
	Invoice Invoice `json:"invoice,omitempty"`
}

// InlineResponse20030Invoice struct for InlineResponse20030Invoice
type Invoice struct {
	// Invoice creation date
	CreationDate string `json:"creationDate"`
	// Currency used for this invoice
	Currency string `json:"currency"`
	// Invoice deadline
	DueDate string `json:"dueDate"`
	// Invoice number
	InvoiceNumber string `json:"invoiceNumber"`
	// Invoice status
	InvoiceStatus string `json:"invoiceStatus"`
	// Invoice paid date
	PayDate string `json:"payDate"`
	// Invoice total (displayed in cents)
	TotalAmount float32 `json:"totalAmount"`
	// Invoice total including VAT (displayed in cents)
	TotalAmountInclVat float32 `json:"totalAmountInclVat"`
}

// InvoiceItem struct for InvoiceItem
type InvoiceItem struct {
	// Date when the order line item was up for invoicing
	Date string `json:"date"`
	// Product description
	Description string `json:"description"`
	// Applied discounts
	Discounts []map[string]interface{} `json:"discounts"`
	// Payment is recurring
	IsRecurring bool `json:"isRecurring"`
	// Price excluding VAT (displayed in cents)
	Price float32 `json:"price"`
	// Price including VAT (displayed in cents)
	PriceInclVat float32 `json:"priceInclVat"`
	// Product name
	Product string `json:"product"`
	// Quantity
	Quantity float32 `json:"quantity"`
	// Amount of VAT charged
	Vat float32 `json:"vat"`
	// Percentage used to calculate the VAT
	VatPercentage float32 `json:"vatPercentage"`
}

// InvoiceItemDiscount struct for InvoiceItemDiscount
type InvoiceItemDiscount struct {
	// Discounted amount (in cents)
	Amount float32 `json:"amount"`
	// Applied discount description
	Description string `json:"description"`
}

// InvoiceItems struct for InvoiceItems
type InvoiceItems struct {
	// array of invoice items
	InvoiceItems []InvoiceItem `json:"invoiceItems"`
}

// Invoices struct for Invoices
type Invoices struct {
	// list of invoices
	Invoices []Invoice `json:"invoices"`
}

// Pdf struct for Pdf
type Pdf struct {
	Content string `json:"pdf"`
}
