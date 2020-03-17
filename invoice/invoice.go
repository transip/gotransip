package invoice

import (
	"github.com/transip/gotransip/v6/rest"
)

// InvoiceStatus is one of the following strings
// 'opened', 'closed', 'waitsforpayment', 'overdue', 'problem' , 'paid', 'paymentpending'
type InvoiceStatus string

const (
	// InvoiceStatusOpened is the invoice status field for a opened invoice
	InvoiceStatusOpened InvoiceStatus = "opened"
	// InvoiceStatusClosed is the invoice status field for a closed invoice
	InvoiceStatusClosed InvoiceStatus = "closed"
	// InvoiceStatusWaitsforpayment is the invoice status field for when the invoice needs to be paid
	InvoiceStatusWaitsforpayment InvoiceStatus = "waitsforpayment"
	// InvoiceStatusOverdue is the invoice status field for when a payment is overdue
	InvoiceStatusOverdue InvoiceStatus = "overdue"
	// InvoiceStatusProblem is the invoice status field for when a problem occured during invoicing
	InvoiceStatusProblem InvoiceStatus = "problem"
	// InvoiceStatusPaid is the invoice status field for a paid invoice
	InvoiceStatusPaid InvoiceStatus = "paid"
	// InvoiceStatusPaymentpending is the invoice status field for when a payment is pending
	InvoiceStatusPaymentpending InvoiceStatus = "paymentpending"
)

// invoicesResponse object contains a list of Invoices in it
// used to unpack the rest response and return the encapsulated Invoices
// this is just used internal for unpacking, this should not be exported
// we want to return Invoice objects not a invoicesResponse
type invoicesResponse struct {
	Invoices []Invoice `json:"invoices"`
}

// invoiceResponse object contains a Invoice in it
// used to unpack the rest response and return the encapsulated Invoice
// this is just used internal for unpacking, this should not be exported
// we want to return a Invoice object not a invoiceResponse
type invoiceResponse struct {
	Invoice Invoice `json:"invoice"`
}

// invoiceItemsResponse object contains a list of InvoiceItems in it
// used to unpack the rest response and return the encapsulated InvoiceItems
// this is just used internal for unpacking, this should not be exported
// we want to return InvoiceItem objects not a invoicesItemsResponse
type invoiceItemsResponse struct {
	// array of invoice items
	InvoiceItems []InvoiceItem `json:"invoiceItems"`
}

// Invoice struct for a invoice
type Invoice struct {
	// Invoice creation date
	CreationDate rest.Date `json:"creationDate"`
	// Currency used for this invoice
	Currency string `json:"currency"`
	// Invoice deadline
	DueDate rest.Date `json:"dueDate"`
	// Invoice number
	InvoiceNumber string `json:"invoiceNumber"`
	// Invoice status
	InvoiceStatus InvoiceStatus `json:"invoiceStatus"`
	// Invoice paid date
	PayDate rest.Date `json:"payDate"`
	// Invoice total (displayed in cents)
	TotalAmount int `json:"totalAmount"`
	// Invoice total including VAT (displayed in cents)
	TotalAmountInclVat int `json:"totalAmountInclVat"`
}

// InvoiceItem struct for InvoiceItem
type InvoiceItem struct {
	// Date when the order line item was up for invoicing
	Date rest.Date `json:"date"`
	// Product description
	Description string `json:"description"`
	// Applied discounts
	Discounts []InvoiceItemDiscount `json:"discounts"`
	// Payment is recurring
	IsRecurring bool `json:"isRecurring"`
	// Price excluding VAT (displayed in cents)
	Price int `json:"price"`
	// Price including VAT (displayed in cents)
	PriceInclVat int `json:"priceInclVat"`
	// Product name
	Product string `json:"product"`
	// Quantity
	Quantity int `json:"quantity"`
	// Amount of VAT charged
	Vat int `json:"vat"`
	// Percentage used to calculate the VAT
	VatPercentage int `json:"vatPercentage"`
}

// InvoiceItemDiscount struct for InvoiceItemDiscount
type InvoiceItemDiscount struct {
	// Discounted amount (in cents)
	Amount int `json:"amount"`
	// Applied discount description
	Description string `json:"description"`
}
