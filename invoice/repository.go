package invoice

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net/url"
)

// Repository can be used to get a list of your invoices, invoice subitems (a specific product)
// or an invoice as pdf
type Repository repository.RestRepository

// GetAll returns a list of all invoices attached to your TransIP account
func (r *Repository) GetAll() ([]Invoice, error) {
	var response invoicesResponse
	restRequest := rest.Request{Endpoint: "/invoices"}
	err := r.Client.Get(restRequest, &response)

	return response.Invoices, err
}

// GetSelection returns a limited list of invoices,
// specify how many and which page/chunk of invoices you want to retrieve
func (r *Repository) GetSelection(page int, itemsPerPage int) ([]Invoice, error) {
	var response invoicesResponse
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/invoices", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.Invoices, err
}

// GetByInvoiceNumber returns an Invoice object for the given invoice number
// invoiceNumber corresponds to the InvoiceNumber property on a Invoice struct
func (r *Repository) GetByInvoiceNumber(invoiceNumber string) (Invoice, error) {
	var response invoiceResponse
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/invoices/%s", invoiceNumber)}
	err := r.Client.Get(restRequest, &response)

	return response.Invoice, err
}

// GetInvoiceItems returns a list of InvoiceItems
// detailing what specific products or services are on this invoice
// invoiceNumber corresponds to the InvoiceNumber property on a Invoice struct
func (r *Repository) GetInvoiceItems(invoiceNumber string) ([]Item, error) {
	var response invoiceItemsResponse
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/invoices/%s/invoice-items", invoiceNumber)}
	err := r.Client.Get(restRequest, &response)

	return response.InvoiceItems, err
}

// GetInvoicePdf returns a pdf struct containing the contents of the pdf encoded in base64
// there are specific Pdf struct functions that help you decode the contents of the pdf to a file
// invoiceNumber corresponds to the InvoiceNumber property on a Invoice struct
func (r *Repository) GetInvoicePdf(invoiceNumber string) (Pdf, error) {
	var response Pdf
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/invoices/%s/pdf", invoiceNumber)}
	err := r.Client.Get(restRequest, &response)

	return response, err
}
