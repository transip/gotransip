package sshkey

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestRepository_GetAll(t *testing.T) {
	const apiResponse = `{
  "sshKeys": [
    {
      "id": 123,
      "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example",
      "description": "Jim key",
      "creationDate": "2020-12-01 15:25:01",
      "fingerprint": "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1"
    }
  ]
}
`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssh-keys", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 123, all[0].ID)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", all[0].Key)
	assert.Equal(t, "Jim key", all[0].Description)
	assert.Equal(t, "2020-12-01 15:25:01", all[0].CreationDate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1", all[0].MD5Fingerprint)
}

func TestRepository_GetSelection(t *testing.T) {
	const apiResponse = `{
  "sshKeys": [
    {
      "id": 123,
      "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example",
      "description": "Jim key",
      "creationDate": "2020-12-01 15:25:01",
      "fingerprint": "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1"
    }
  ]
}
`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssh-keys?page=1&pageSize=25", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetSelection(1, 25)
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 123, all[0].ID)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", all[0].Key)
	assert.Equal(t, "Jim key", all[0].Description)
	assert.Equal(t, "2020-12-01 15:25:01", all[0].CreationDate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1", all[0].MD5Fingerprint)
}

func TestRepository_GetById(t *testing.T) {
	const apiResponse = `{
  "sshKey": {
    "id": 123,
    "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example",
    "description": "Jim key",
    "creationDate": "2020-12-01 15:25:01",
    "fingerprint": "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1"
  }
}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssh-keys/123", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	sshKey, err := repo.GetByID(123)
	require.NoError(t, err)

	assert.EqualValues(t, 123, sshKey.ID)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", sshKey.Key)
	assert.Equal(t, "Jim key", sshKey.Description)
	assert.Equal(t, "2020-12-01 15:25:01", sshKey.CreationDate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1", sshKey.MD5Fingerprint)
}

func TestRepository_Add(t *testing.T) {
	const expectedRequestBody = `{"description":"Jim key","sshKey":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssh-keys", ExpectedMethod: "POST", StatusCode: 201, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Add("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", "Jim key")
	require.NoError(t, err)
}

func TestRepository_Update(t *testing.T) {
	const expectedRequestBody = `{"description":"Jim key"}`
	server := testutil.MockServer{T: t, ExpectedURL: "/ssh-keys/123", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: expectedRequestBody}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	sshKey := SSHKey{
		ID:          123,
		Description: "Jim key",
	}

	err := repo.Update(sshKey)
	require.NoError(t, err)
}

func TestRepository_Remove(t *testing.T) {
	server := testutil.MockServer{T: t, ExpectedURL: "/ssh-keys/123", ExpectedMethod: "DELETE", StatusCode: 204}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Remove(123)
	require.NoError(t, err)
}
