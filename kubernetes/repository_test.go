package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestRepository_GetClusters(t *testing.T) {
	const apiResponse = `{"clusters":[{"name":"k888k","description":"production cluster","isLocked":true,"isBlocked": false},{"name":"aiceayoo","description":"development cluster","isLocked":false,"isBlocked":true}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetClusters()
	require.NoError(t, err)

	if assert.Equal(t, 2, len(all)) {
		assert.Equal(t, "k888k", all[0].Name)
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
	const apiResponse = `{"cluster":{"name":"k888k","description":"production cluster","isLocked":true,"isBlocked": false}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/k888k", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	cluster, err := repo.GetClusterByName("k888k")
	require.NoError(t, err)

	assert.Equal(t, "k888k", cluster.Name)
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

	err := repo.CreateCluster(order)
	require.NoError(t, err)
}

func TestRepository_UpdateCluster(t *testing.T) {
	const expectedRequest = `{"cluster":{"name":"k888k","description":"staging cluster"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/k888k", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	clusterToUpdate := Cluster{
		Name:        "k888k",
		Description: "staging cluster",
	}

	err := repo.UpdateCluster(clusterToUpdate)

	require.NoError(t, err)
}

func TestRepository_HandoverCluster(t *testing.T) {
	const expectedRequest = `{"action":"handover","targetCustomerName":"bobexample"}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/k888k", ExpectedMethod: "PATCH", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.HandoverCluster("k888k", "bobexample")
	require.NoError(t, err)
}

func TestRepository_CancelCluster(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/k888k", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RemoveCluster("k888k")
	require.NoError(t, err)
}

func TestRepository_GetKubeConfig(t *testing.T) {
	const apiResponse = `{"kubeConfig": {"encodedYaml": "YXBpVmVyc2lvbjogdjEKY2x1c3RlcnM6IFtdCmNvbnRleHRzOiBbXQpraW5kOiBDb25maWcKcHJlZmVyZW5jZXM6IHt9CnVzZXJzOiBbXQoK"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/clusters/k888k/kubeconfig", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	config, err := repo.GetKubeConfig("k888k")
	require.NoError(t, err)

	assert.Contains(t, config, "apiVersion: v1")
}

func TestRepository_GetNodePools(t *testing.T) {
	const apiResponse = `{"nodePools":[{"uuid":"402c2f84-c37d-9388-634d-00002b7c6a82","description":"frontend","desiredNodeCount":3,"nodeSpec":"vps-bladevps-x4","nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","nodePoolUuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","status":"active"}]}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/node-pools", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodePools()
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].UUID)
		assert.Equal(t, "frontend", all[0].Description)
		assert.Equal(t, 3, all[0].DesiredNodeCount)
		assert.Equal(t, "vps-bladevps-x4", all[0].NodeSpec)
		if assert.Equal(t, 1, len(all[0].Nodes)) {
			assert.Equal(t, NodeStatusActive, all[0].Nodes[0].Status)
			assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].Nodes[0].UUID)
			assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].Nodes[0].NodePoolUUID)
			assert.Equal(t, "k888k", all[0].Nodes[0].ClusterName)
		}
	}
}

func TestRepository_GetNodePoolsByClusterName(t *testing.T) {
	const apiResponse = `{"nodePools":[{"uuid":"402c2f84-c37d-9388-634d-00002b7c6a82","description":"frontend","desiredNodeCount":3,"nodeSpec":"vps-bladevps-x4","nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","nodePoolUuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","status":"active"}]}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/node-pools?clusterName=k888k", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodePoolsByClusterName("k888k")
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].UUID)
		assert.Equal(t, "frontend", all[0].Description)
		assert.Equal(t, 3, all[0].DesiredNodeCount)
		assert.Equal(t, "vps-bladevps-x4", all[0].NodeSpec)
		if assert.Equal(t, 1, len(all[0].Nodes)) {
			assert.Equal(t, NodeStatusActive, all[0].Nodes[0].Status)
			assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].Nodes[0].UUID)
			assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].Nodes[0].NodePoolUUID)
			assert.Equal(t, "k888k", all[0].Nodes[0].ClusterName)
		}
	}
}

func TestRepository_GetNodePool(t *testing.T) {
	const apiResponse = `{"nodePool":{"uuid":"402c2f84-c37d-9388-634d-00002b7c6a82","description":"frontend","desiredNodeCount":3,"nodeSpec":"vps-bladevps-x4","nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","nodePoolUuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","status":"active"}]}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/node-pools/402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	nodePool, err := repo.GetNodePool("402c2f84-c37d-9388-634d-00002b7c6a82")
	require.NoError(t, err)

	assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", nodePool.UUID)
	assert.Equal(t, "frontend", nodePool.Description)
	assert.Equal(t, 3, nodePool.DesiredNodeCount)
	assert.Equal(t, "vps-bladevps-x4", nodePool.NodeSpec)
	if assert.Equal(t, 1, len(nodePool.Nodes)) {
		assert.Equal(t, NodeStatusActive, nodePool.Nodes[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", nodePool.Nodes[0].UUID)
		assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", nodePool.Nodes[0].NodePoolUUID)
		assert.Equal(t, "k888k", nodePool.Nodes[0].ClusterName)
	}
}

func TestRepository_OrderNodePool(t *testing.T) {
	const expectedRequestBody = `{"clusterName":"k888k","description":"frontend","desiredNodeCount":3,"nodeSpec":"vps-bladevps-x4"}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/node-pools", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	order := NodePoolOrder{
		ClusterName:      "k888k",
		Description:      "frontend",
		DesiredNodeCount: 3,
		NodeSpec:         "vps-bladevps-x4",
	}

	err := repo.AddNodePool(order)
	require.NoError(t, err)
}

func TestRepository_UpdateNodePool(t *testing.T) {
	const expectedRequest = `{"nodePool":{"uuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","description":"backend","desiredNodeCount":4,"nodeSpec":"vps-bladevps-x8"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/node-pools/402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	nodePoolToUpdate := NodePool{
		UUID:             "402c2f84-c37d-9388-634d-00002b7c6a82",
		ClusterName:      "k888k",
		Description:      "backend",
		DesiredNodeCount: 4,
		NodeSpec:         "vps-bladevps-x8",
	}

	err := repo.UpdateNodePool(nodePoolToUpdate)

	require.NoError(t, err)
}

func TestRepository_CancelNodePool(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/node-pools/402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RemoveNodePool("402c2f84-c37d-9388-634d-00002b7c6a82")
	require.NoError(t, err)
}

func TestRepository_GetNodes(t *testing.T) {
	const apiResponse = `{"nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","nodePoolUuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","status":"active"}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodes()
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, NodeStatusActive, all[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].UUID)
		assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].NodePoolUUID)
		assert.Equal(t, "k888k", all[0].ClusterName)
	}
}

func TestRepository_GetNodesByClusterName(t *testing.T) {
	const apiResponse = `{"nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","nodePoolUuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","status":"active"}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes?clusterName=k888k", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodesByClusterName("k888k")
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, NodeStatusActive, all[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].UUID)
		assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].NodePoolUUID)
		assert.Equal(t, "k888k", all[0].ClusterName)
	}
}

func TestRepository_GetNodesByNodePoolUUID(t *testing.T) {
	const apiResponse = `{"nodes":[{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","nodePoolUuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","status":"active"}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes?nodePoolUuid=402c2f84-c37d-9388-634d-00002b7c6a82", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetNodesByNodePoolUUID("402c2f84-c37d-9388-634d-00002b7c6a82")
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, NodeStatusActive, all[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].UUID)
		assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", all[0].NodePoolUUID)
		assert.Equal(t, "k888k", all[0].ClusterName)
	}
}

func TestRepository_GetNode(t *testing.T) {
	const apiResponse = `{"node":{"uuid":"76743b28-f779-3e68-6aa1-00007fbb911d","nodePoolUuid":"402c2f84-c37d-9388-634d-00002b7c6a82","clusterName":"k888k","status":"active"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/nodes/76743b28-f779-3e68-6aa1-00007fbb911d", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	node, err := repo.GetNode("76743b28-f779-3e68-6aa1-00007fbb911d")
	require.NoError(t, err)

	assert.Equal(t, NodeStatusActive, node.Status)
	assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", node.UUID)
	assert.Equal(t, "402c2f84-c37d-9388-634d-00002b7c6a82", node.NodePoolUUID)
	assert.Equal(t, "k888k", node.ClusterName)
}

func TestRepository_GetBlockStorageVolumes(t *testing.T) {
	const apiResponse = `{"volumes":[{"uuid":"220887f0-db1a-76a9-2332-00004f589b19","name":"custom-2c3501ab-5a45-34e9-c289-00002b084a0c","sizeInGib":20,"type":"hdd","availabilityZone":"ams0","status":"available","nodeUuid":"76743b28-f779-3e68-6aa1-00007fbb911d","serial":"a4d857d3fe5e814f34bb"}]}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/block-storages", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetBlockStorageVolumes()
	require.NoError(t, err)

	if assert.Equal(t, 1, len(all)) {
		assert.Equal(t, "220887f0-db1a-76a9-2332-00004f589b19", all[0].UUID)
		assert.Equal(t, "custom-2c3501ab-5a45-34e9-c289-00002b084a0c", all[0].Name)
		assert.Equal(t, 20, all[0].SizeInGiB)
		assert.Equal(t, "hdd", all[0].Type)
		assert.Equal(t, "ams0", all[0].AvailabilityZone)
		assert.Equal(t, BlockStorageStatusAvailable, all[0].Status)
		assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", all[0].NodeUUID)
		assert.Equal(t, "a4d857d3fe5e814f34bb", all[0].Serial)
	}
}

func TestRepository_GetBlockStorageVolume(t *testing.T) {
	const apiResponse = `{"volume":{"uuid":"220887f0-db1a-76a9-2332-00004f589b19","name":"custom-2c3501ab-5a45-34e9-c289-00002b084a0c","sizeInGib":20,"type":"hdd","availabilityZone":"ams0","status":"available","nodeUuid":"76743b28-f779-3e68-6aa1-00007fbb911d","serial":"a4d857d3fe5e814f34bb"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/block-storages/custom-2c3501ab-5a45-34e9-c289-00002b084a0c", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	volume, err := repo.GetBlockStorageVolume("custom-2c3501ab-5a45-34e9-c289-00002b084a0c")
	require.NoError(t, err)

	assert.Equal(t, "220887f0-db1a-76a9-2332-00004f589b19", volume.UUID)
	assert.Equal(t, "custom-2c3501ab-5a45-34e9-c289-00002b084a0c", volume.Name)
	assert.Equal(t, 20, volume.SizeInGiB)
	assert.Equal(t, "hdd", volume.Type)
	assert.Equal(t, "ams0", volume.AvailabilityZone)
	assert.Equal(t, BlockStorageStatusAvailable, volume.Status)
	assert.Equal(t, "76743b28-f779-3e68-6aa1-00007fbb911d", volume.NodeUUID)
	assert.Equal(t, "a4d857d3fe5e814f34bb", volume.Serial)
}

func TestRepository_AddBlockStorageVolume(t *testing.T) {
	const expectedRequestBody = `{"name":"custom-2c3501ab-5a45-34e9-c289-00002b084a0c","sizeInGib":200,"type":"hdd","availabilityZone":"ams0"}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/block-storages", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	order := BlockStorageOrder{
		Name:             "custom-2c3501ab-5a45-34e9-c289-00002b084a0c",
		SizeInGiB:        200,
		Type:             "hdd",
		AvailabilityZone: "ams0",
	}

	err := repo.AddBlockStorageVolume(order)
	require.NoError(t, err)
}

func TestRepository_UpdateBlockStorageVolume(t *testing.T) {
	const expectedRequest = `{"volume":{"uuid":"220887f0-db1a-76a9-2332-00004f589b19","name":"custom-2c3501ab-5a45-34e9-c289-00002b084a0c","sizeInGib":20,"type":"hdd","availabilityZone":"ams0","status":"available","nodeUuid":"76743b28-f779-3e68-6aa1-00007fbb911d","serial":"a4d857d3fe5e814f34bb"}}`

	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/block-storages/custom-2c3501ab-5a45-34e9-c289-00002b084a0c", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.UpdateBlockStorageVolume(BlockStorage{
		UUID:             "220887f0-db1a-76a9-2332-00004f589b19",
		Name:             "custom-2c3501ab-5a45-34e9-c289-00002b084a0c",
		SizeInGiB:        20,
		Type:             "hdd",
		AvailabilityZone: "ams0",
		Status:           BlockStorageStatusAvailable,
		NodeUUID:         "76743b28-f779-3e68-6aa1-00007fbb911d",
		Serial:           "a4d857d3fe5e814f34bb",
	})

	require.NoError(t, err)
}

func TestRepository_RemoveBlockStorageVolume(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/kubernetes/block-storages/custom-2c3501ab-5a45-34e9-c289-00002b084a0c", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.RemoveBlockStorageVolume("custom-2c3501ab-5a45-34e9-c289-00002b084a0c")
	require.NoError(t, err)
}