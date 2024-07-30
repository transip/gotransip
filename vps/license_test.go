package vps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestLicenseRepository_GetLicenses(t *testing.T) {
	const apiResponse = `{ "licenses": { "active": [ { "id": 42, "name": "cpanel-admin", "price": 1050, "recurringPrice": 1050, "type": "addon", "quantity": 1, "maxQuantity": 1, "keys": [ { "name": "Cpanel license key", "key": "XXXXXXXXXXX" } ] } ], "cancellable": [ { "id": 42, "name": "cpanel-admin", "price": 1050, "recurringPrice": 1050, "type": "addon", "quantity": 1, "maxQuantity": 1, "keys": [ { "name": "Cpanel license key", "key": "XXXXXXXXXXX" } ] } ], "available": [ { "name": "installatron", "price": 1050, "recurringPrice": 1050, "type": "addon", "maxQuantity": 1 } ] } }` //nolint
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/licenses", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := LicenseRepository{Client: *client}

	all, err := repo.GetAll("example-vps")
	require.NoError(t, err)
	require.Equal(t, 1, len(all.Active))
	require.Equal(t, 1, len(all.Available))
	require.Equal(t, 1, len(all.Cancellable))

	assert.Equal(t, "cpanel-admin", all.Active[0].Name)
	assert.EqualValues(t, 42, all.Active[0].ID)
	assert.EqualValues(t, 1050, all.Active[0].Price)
	assert.EqualValues(t, 1050, all.Active[0].RecurringPrice)
	assert.EqualValues(t, 1, all.Active[0].Quantity)
	assert.EqualValues(t, 1, all.Active[0].MaxQuantity)
	require.Equal(t, 1, len(all.Active[0].Keys))
	assert.Equal(t, "Cpanel license key", all.Active[0].Keys[0].Name)
	assert.Equal(t, "XXXXXXXXXXX", all.Active[0].Keys[0].Key)

	assert.Equal(t, "cpanel-admin", all.Cancellable[0].Name)
	assert.EqualValues(t, 42, all.Cancellable[0].ID)
	assert.EqualValues(t, 1050, all.Cancellable[0].Price)
	assert.EqualValues(t, 1050, all.Cancellable[0].RecurringPrice)
	assert.EqualValues(t, 1, all.Cancellable[0].Quantity)
	assert.EqualValues(t, 1, all.Cancellable[0].MaxQuantity)
	require.Equal(t, 1, len(all.Cancellable[0].Keys))
	assert.Equal(t, "Cpanel license key", all.Cancellable[0].Keys[0].Name)
	assert.Equal(t, "XXXXXXXXXXX", all.Cancellable[0].Keys[0].Key)

	assert.Equal(t, "installatron", all.Available[0].Name) //nolint
	assert.EqualValues(t, 1050, all.Available[0].Price)
	assert.EqualValues(t, 1050, all.Available[0].RecurringPrice)
	assert.EqualValues(t, 1, all.Available[0].MaxQuantity)
}

func TestLicenseRepository_OrderLicense(t *testing.T) {
	const expectedRequest = `{"licenseName":"cpanel-pro","quantity":1}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/licenses", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := LicenseRepository{Client: *client}

	order := LicenseOrder{LicenseName: "cpanel-pro", Quantity: 1}
	err := repo.Order("example-vps", order)

	require.NoError(t, err)
}

func TestLicenseRepository_ReplaceLicense(t *testing.T) {
	const expectedRequest = `{"newLicenseName":"cpanel-pro"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/licenses/1337", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := LicenseRepository{Client: *client}

	replaceRequest := ReplaceLicenseRequest{LicenseID: 1337, NewLicenseName: "cpanel-pro"}
	err := repo.Replace("example-vps", replaceRequest)

	require.NoError(t, err)
}

func TestLicenseRepository_CancelLicense(t *testing.T) {
	const expectedRequest = `{"endTime":"end"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/licenses/1337", ExpectedMethod: "DELETE", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := LicenseRepository{Client: *client}

	err := repo.Cancel("example-vps", 1337)
	require.NoError(t, err)
}
