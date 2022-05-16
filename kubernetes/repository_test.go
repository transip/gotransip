package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestRepository_GetClusters(t *testing.T) {
	const apiResponse = `{"clusters":[{"name":"paeweeth","description":"production cluster","isLocked":true,"isBlocked": false},{"name":"aiceayoo","description":"development cluster","isLocked":false,"isBlocked":true}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetClusters()
	require.NoError(t, err)

	if assert.Equal(t, 2, len(all)) {
		assert.Equal(t, "paeweeth", all[0].Name)
		assert.Equal(t, "production cluster", all[0].Description)
		assert.True(t, all[0].IsLocked)
		assert.False(t, all[0].IsBlocked)
		assert.Equal(t, "aiceayoo", all[1].Name)
		assert.Equal(t, "development cluster", all[1].Description)
		assert.False(t, all[1].IsLocked)
		assert.True(t, all[1].IsBlocked)
	}
}

func TestRepository_GetClusterByName(t *testing.T) {
	const apiResponse = `{"cluster":{"name":"paeweeth","description":"production cluster","isLocked":true,"isBlocked": false}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	cluster, err := repo.GetClusterByName("paeweeth")
	require.NoError(t, err)

	assert.Equal(t, "paeweeth", cluster.Name)
	assert.Equal(t, "production cluster", cluster.Description)
	assert.True(t, cluster.IsLocked)
	assert.False(t, cluster.IsBlocked)
}

func TestRepository_OrderCluster(t *testing.T) {
	const expectedRequestBody = `{"description":"production cluster"}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	order := ClusterOrder{
		Description: "production cluster",
	}

	err := repo.OrderCluster(order)
	require.NoError(t, err)
}

func TestRepository_UpdateCluster(t *testing.T) {
	const expectedRequest = `{"cluster":{"name":"paeweeth","description":"staging cluster"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	vpsToUpdate := Cluster{
		Name:        "paeweeth",
		Description: "staging cluster",
	}

	err := repo.UpdateCluster(vpsToUpdate)

	require.NoError(t, err)
}

func TestRepository_HandoverCluster(t *testing.T) {
	const expectedRequest = `{"action":"handover","targetCustomerName":"bobexample"}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.HandoverCluster("paeweeth", "bobexample")
	require.NoError(t, err)
}

func TestRepository_CancelCluster(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.CancelCluster("paeweeth")
	require.NoError(t, err)
}

func TestRepository_GetKubeConfig(t *testing.T) {
	const apiResponse = `{"kubeConfig": {"yaml": "YXBpVmVyc2lvbjogdjEKY2x1c3RlcnM6IFtdCmNvbnRleHRzOiBbXQpraW5kOiBDb25maWcKcHJlZmVyZW5jZXM6IHt9CnVzZXJzOiBbXQoK"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth/kubeconfig", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	config, err := repo.GetKubeConfig("paeweeth")
	require.NoError(t, err)

	assert.Contains(t, config, "apiVersion: v1")
}
