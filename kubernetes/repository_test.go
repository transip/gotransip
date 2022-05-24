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

	clusterToUpdate := Cluster{
		Name:        "paeweeth",
		Description: "staging cluster",
	}

	err := repo.UpdateCluster(clusterToUpdate)

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

func TestRepository_GetNodePools(t *testing.T) {
	const apiResponse = `{"nodePools":[{"uuid":"402c2f84-c37d-9388-634d-00002b7c6a82","description":"frontend","desiredNodeCount":3,"nodeSpec":"vps-bladevps-x4","nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","status":"active"}]}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth/node-pools", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodePools("paeweeth")
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].UUID)
		assert.Equal(t, "frontend", all[0].Description)
		assert.Equal(t, 3, all[0].DesiredNodeCount)
		assert.Equal(t, "vps-bladevps-x4", all[0].NodeSpec)
		if assert.Equal(t, 1, len(all[0].Nodes)) {
			assert.Equal(t, NodeStatusActive, all[0].Nodes[0].Status)
			assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].Nodes[0].UUID)
		}
	}
}

func TestRepository_GetNodePool(t *testing.T) {
	const apiResponse = `{"nodePool":{"uuid":"402c2f84-c37d-9388-634d-00002b7c6a82","description":"frontend","desiredNodeCount":3,"nodeSpec":"vps-bladevps-x4","nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","status":"active"}]}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth/node-pools/402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	nodePool, err := repo.GetNodePool("paeweeth", "402c2f84-c37d-9388-634d-00002b7c6a82")
	require.NoError(t, err)

	assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", nodePool.UUID)
	assert.Equal(t, "frontend", nodePool.Description)
	assert.Equal(t, 3, nodePool.DesiredNodeCount)
	assert.Equal(t, "vps-bladevps-x4", nodePool.NodeSpec)
	if assert.Equal(t, 1, len(nodePool.Nodes)) {
		assert.Equal(t, NodeStatusActive, nodePool.Nodes[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", nodePool.Nodes[0].UUID)
	}
}

func TestRepository_OrderNodePool(t *testing.T) {
	const expectedRequestBody = `{"description":"frontend","desiredNodeCount":3,"nodeSpec":"vps-bladevps-x4"}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth/node-pools", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	order := NodePoolOrder{
		Description:      "frontend",
		DesiredNodeCount: 3,
		NodeSpec:         "vps-bladevps-x4",
	}

	err := repo.OrderNodePool("paeweeth", order)
	require.NoError(t, err)
}

func TestRepository_UpdateNodePool(t *testing.T) {
	const expectedRequest = `{"nodePool":{"uuid":"402c2f84-c37d-9388-634d-00002b7c6a82","description":"backend","desiredNodeCount":4,"nodeSpec":"vps-bladevps-x8"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth/node-pools/402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	nodePoolToUpdate := NodePool{
		UUID:             "402c2f84-c37d-9388-634d-00002b7c6a82",
		Description:      "backend",
		DesiredNodeCount: 4,
		NodeSpec:         "vps-bladevps-x8",
	}

	err := repo.UpdateNodePool("paeweeth", nodePoolToUpdate)

	require.NoError(t, err)
}

func TestRepository_CancelNodePool(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/paeweeth/node-pools/402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.CancelNodePool("paeweeth", "402c2f84-c37d-9388-634d-00002b7c6a82")
	require.NoError(t, err)
}

func TestRepository_GetNodes(t *testing.T) {
	const apiResponse = `{"nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","status":"active"}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodes()
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, NodeStatusActive, all[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].UUID)
	}
}

func TestRepository_GetNodesByClusterName(t *testing.T) {
	const apiResponse = `{"nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","status":"active"}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes?clusterName=paeweeth", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodesByClusterName("paeweeth")
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, NodeStatusActive, all[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].UUID)
	}
}

func TestRepository_GetNodesByNodePoolUUID(t *testing.T) {
	const apiResponse = `{"nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","status":"active"}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes?nodePoolUuid=402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodesByNodePoolUUID("402c2f84-c37d-9388-634d-00002b7c6a82")
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, NodeStatusActive, all[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].UUID)
	}
}

func TestRepository_GetNode(t *testing.T) {
	const apiResponse = `{"node":{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","status":"active"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes/76743b28-f779-3e68-6aa1-00007fbb911d", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	node, err := repo.GetNode("76743b28-f779-3e68-6aa1-00007fbb911d")
	require.NoError(t, err)

	assert.Equal(t, NodeStatusActive, node.Status)
	assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", node.UUID)
}
