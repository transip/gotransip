package traffic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestRepository_GetTrafficInformationForVps(t *testing.T) {
	const apiResponse = `{ "trafficInformation": { "startDate": "2019-06-22", "endDate": "2019-07-22", "usedInBytes": 7860253754, "usedTotalBytes": 11935325369, "maxInBytes": 1073741824000 } }`
	server := testutil.MockServer{T: t, ExpectedURL: "/traffic", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}
	pool, err := repo.GetTrafficPool()
	require.NoError(t, err)

	assert.Equal(t, "2019-06-22", pool.StartDate.Format("2006-01-02"))
	assert.Equal(t, "2019-07-22", pool.EndDate.Format("2006-01-02"))
	assert.EqualValues(t, 7860253754, pool.UsedInBytes)
	assert.EqualValues(t, 11935325369, pool.UsedTotalBytes)
	assert.EqualValues(t, 1073741824000, pool.MaxInBytes)
}

func TestRepository_GetTrafficPool(t *testing.T) {
	const apiResponse = `{ "trafficInformation": { "startDate": "2019-06-22", "endDate": "2019-07-22", "usedInBytes": 7860253754, "usedTotalBytes": 11935325369, "maxInBytes": 1073741824000 } }`
	server := testutil.MockServer{T: t, ExpectedURL: "/traffic/test-vps", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}
	pool, err := repo.GetTrafficInformationForVps("test-vps")
	require.NoError(t, err)

	assert.Equal(t, "2019-06-22", pool.StartDate.Format("2006-01-02"))
	assert.Equal(t, "2019-07-22", pool.EndDate.Format("2006-01-02"))
	assert.EqualValues(t, 7860253754, pool.UsedInBytes)
	assert.EqualValues(t, 11935325369, pool.UsedTotalBytes)
	assert.EqualValues(t, 1073741824000, pool.MaxInBytes)
}
