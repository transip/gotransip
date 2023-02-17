package action

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
	"github.com/transip/gotransip/v6/rest"
)

func TestActionRepository_GetActions(t *testing.T) {
	const apiResponse = `{"actions":[{"uuid": "6c7fa1c1-f509-4999-a513-bdf4e7a0cebb", "name":"snapshot revert", "actionStartTime":"2023-02-01 17:01:51", "status":"running", "metadata": {"progress": 1337}, "parentActionUuid": ""}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/actions", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}
	actions, err := repo.GetActions()
	assert.Len(t, actions, 1)
	require.NoError(t, err)

	type Metadata struct {
		Progress int
	}

	var metadata Metadata

	err = json.Unmarshal([]byte(actions[0].Metadata), &metadata)
	assert.NoError(t, err)

	assert.Equal(t, "6c7fa1c1-f509-4999-a513-bdf4e7a0cebb", actions[0].UUID)
	assert.Equal(t, "snapshot revert", actions[0].Name)
	assert.Equal(t, "2023-02-01 17:01:51", actions[0].ActionStartTime)
	assert.Equal(t, "running", actions[0].Status)
	assert.Equal(t, 1337, metadata.Progress)
	assert.Equal(t, "", actions[0].ParentActionUUID)
}

func TestActionRepository_ParseActionFromResponse(t *testing.T) {
	const apiResponse = `{"action":{"uuid": "6c7fa1c1-f509-4999-a513-bdf4e7a0cebb", "name":"snapshot revert", "actionStartTime":"2023-02-01 17:01:51", "status":"running", "metadata": {"progress": 1337}, "parentActionUuid": ""}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/actions/6c7fa1c1-f509-4999-a513-bdf4e7a0cebb", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	var resp rest.Response
	resp.ContentLocation = "/v6/actions/6c7fa1c1-f509-4999-a513-bdf4e7a0cebb"
	action, err := repo.ParseActionFromResponse(resp)
	require.NoError(t, err)

	type Metadata struct {
		Progress int
	}

	var metadata Metadata

	err = json.Unmarshal([]byte(action.Metadata), &metadata)
	assert.NoError(t, err)

	assert.Equal(t, "6c7fa1c1-f509-4999-a513-bdf4e7a0cebb", action.UUID)
	assert.Equal(t, "snapshot revert", action.Name)
	assert.Equal(t, "2023-02-01 17:01:51", action.ActionStartTime)
	assert.Equal(t, "running", action.Status)
	assert.Equal(t, 1337, metadata.Progress)
	assert.Equal(t, "", action.ParentActionUUID)
}
