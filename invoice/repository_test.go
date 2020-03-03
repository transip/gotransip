package invoice

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	invoicesApiResponse = `{
  "invoices": [
    {
      "invoiceNumber": "F0000.1911.0000.0004",
      "creationDate": "2020-01-01",
      "payDate": "2020-01-01",
      "dueDate": "2020-02-01",
      "invoiceStatus": "waitsforpayment",
      "currency": "EUR",
      "totalAmount": 1000,
      "totalAmountInclVat": 1240
    }
  ]
}`
	invoiceApiResponse = `{
  "invoice": {
    "invoiceNumber": "F0000.1911.0000.0004",
    "creationDate": "2020-01-01",
    "payDate": "2020-01-01",
    "dueDate": "2020-02-01",
    "invoiceStatus": "waitsforpayment",
    "currency": "EUR",
    "totalAmount": 1000,
    "totalAmountInclVat": 1240
  }
}`
	errorResponse           = `{ "error": "errortest" }`
	error404Response        = `{ "error": "Invoice with number 'F0000.1911.0000.0004' not found" }`
	invoiceItemsApiResponse = `{
  "invoiceItems": [
    {
      "product": "Big Storage Disk 2000 GB",
      "description": "Big Storage Disk 2000 GB (example-bigstorage)",
      "isRecurring": false,
      "date": "2020-01-01",
      "quantity": 1,
      "price": 1000,
      "priceInclVat": 1210,
      "vat": 210,
      "vatPercentage": 21,
      "discounts": [
        {
          "description": "Korting (20% Black Friday)",
          "amount": -500
        }
      ]
    }
  ]
}`
	invoicePdfResponse = `{ "pdf": "dGVzdDEyMw==" }`
)

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
	repo, tearDown := getRepository(t, "/invoices", 200, invoicesApiResponse)
	defer tearDown()

	all, err := repo.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 1, len(all))
}

func TestRepository_GetAllError(t *testing.T) {
	repo, tearDown := getRepository(t, "/invoices", 500, errorResponse)
	defer tearDown()

	all, err := repo.GetAll()
	require.Error(t, err)
	require.Nil(t, all)
	assert.Equal(t, errors.New("errortest"), err)
}

func TestRepository_GetByInvoiceNumber(t *testing.T) {
	invoiceNumber := "F0000.1911.0000.0004"
	repo, tearDown := getRepository(t, "/invoices/"+invoiceNumber, 200, invoiceApiResponse)
	defer tearDown()

	invoice, err := repo.GetByInvoiceNumber(invoiceNumber)
	require.NoError(t, err)

	assert.Equal(t, invoiceNumber, invoice.InvoiceNumber)
	assert.Equal(t, "2020-01-01 00:00:00 +0100 CET", invoice.CreationDate.String())
	assert.Equal(t, "2020-01-01 00:00:00 +0100 CET", invoice.PayDate.String())
	assert.Equal(t, "2020-02-01 00:00:00 +0100 CET", invoice.DueDate.String())
	assert.Equal(t, "waitsforpayment", invoice.InvoiceStatus)
	assert.Equal(t, "EUR", invoice.Currency)
	assert.Equal(t, 1000, invoice.TotalAmount)
	assert.Equal(t, 1240, invoice.TotalAmountInclVat)
}

func TestRepository_GetByInvoiceNumberError(t *testing.T) {
	invoiceNumber := "throwmea404"
	repo, tearDown := getRepository(t, fmt.Sprintf("/invoices/%s", invoiceNumber), 404, error404Response)
	defer tearDown()

	all, err := repo.GetByInvoiceNumber(invoiceNumber)
	require.Error(t, err)
	require.Empty(t, all.InvoiceNumber)
	assert.Equal(t, errors.New("Invoice with number 'F0000.1911.0000.0004' not found"), err)
}

func TestRepository_GetInvoiceItems(t *testing.T) {
	invoiceNumber := "F0000.1911.0000.0004"
	repo, tearDown := getRepository(t, fmt.Sprintf("/invoices/%s/invoice-items", invoiceNumber), 200, invoiceItemsApiResponse)
	defer tearDown()

	all, err := repo.GetInvoiceItems(invoiceNumber)
	require.NoError(t, err)

	require.Equal(t, 1, len(all))
	assert.Equal(t, "Big Storage Disk 2000 GB", all[0].Product)
	assert.Equal(t, "Big Storage Disk 2000 GB (example-bigstorage)", all[0].Description)
	assert.Equal(t, false, all[0].IsRecurring)
	assert.Equal(t, "2020-01-01 00:00:00 +0100 CET", all[0].Date.String())
	assert.Equal(t, 1, all[0].Quantity)
	assert.Equal(t, 1000, all[0].Price)
	assert.Equal(t, 1210, all[0].PriceInclVat)
	assert.Equal(t, 210, all[0].Vat)
	assert.Equal(t, 21, all[0].VatPercentage)

	require.Equal(t, 1, len(all[0].Discounts))
	assert.Equal(t, "Korting (20% Black Friday)", all[0].Discounts[0].Description)
	assert.Equal(t, -500, all[0].Discounts[0].Amount)
}

func TestRepository_GetInvoiceItemsError(t *testing.T) {
	invoiceNumber := "throwmea404"
	repo, tearDown := getRepository(t, fmt.Sprintf("/invoices/%s/invoice-items", invoiceNumber), 404, error404Response)
	defer tearDown()

	all, err := repo.GetInvoiceItems(invoiceNumber)
	require.Error(t, err)
	require.Nil(t, all)
	assert.Equal(t, errors.New("Invoice with number 'F0000.1911.0000.0004' not found"), err)
}

func TestRepository_GetInvoicePdf(t *testing.T) {
	invoiceNumber := "F0000.1911.0000.0004"
	repo, tearDown := getRepository(t, fmt.Sprintf("/invoices/%s/pdf", invoiceNumber), 200, invoicePdfResponse)
	defer tearDown()

	pdf, err := repo.GetInvoicePdf(invoiceNumber)
	require.NoError(t, err)
	assert.Equal(t, 12, len(pdf.Content))

	reader := pdf.GetIoReader()

	bytes, err := ioutil.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, []byte("test123"), bytes)
}

func TestRepository_GetInvoicePdfError(t *testing.T) {
	invoiceNumber := "throwmea404"
	repo, tearDown := getRepository(t, fmt.Sprintf("/invoices/%s/pdf", invoiceNumber), 404, error404Response)
	defer tearDown()

	pdf, err := repo.GetInvoicePdf(invoiceNumber)
	require.Error(t, err)
	require.Empty(t, pdf.Content)
	assert.Equal(t, errors.New("Invoice with number 'F0000.1911.0000.0004' not found"), err)
}
