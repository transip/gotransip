package sshkey

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockServer struct is used to test the how the client sends a request
// and responds to a servers response
type mockServer struct {
	t               *testing.T
	expectedURL     string
	expectedMethod  string
	statusCode      int
	expectedRequest string
	response        string
	skipRequestBody bool
}

func (m *mockServer) getHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(m.t, m.expectedURL, req.URL.String()) // check if right expectedURL is called

		if m.skipRequestBody == false && req.ContentLength != 0 {
			// get the request body
			// and check if the body matches the expected request body
			body, err := ioutil.ReadAll(req.Body)
			require.NoError(m.t, err)
			assert.Equal(m.t, m.expectedRequest, string(body))
		}

		assert.Equal(m.t, m.expectedMethod, req.Method) // check if the right expectedRequest expectedMethod is used
		rw.WriteHeader(m.statusCode)                    // respond with given status code

		if m.response != "" {
			_, err := rw.Write([]byte(m.response))
			require.NoError(m.t, err, "error when writing mock response")
		}
	}))
}

func (m *mockServer) getClient() (*repository.Client, func()) {
	httpServer := m.getHTTPServer()
	config := gotransip.DemoClientConfiguration
	config.URL = httpServer.URL
	client, err := gotransip.NewClient(config)
	require.NoError(m.t, err)

	// return tearDown method with which will close the test server after the test
	tearDown := func() {
		httpServer.Close()
	}

	return &client, tearDown
}

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
	server := mockServer{t: t, expectedURL: "/ssh-keys", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 123, all[0].ID)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", all[0].Key)
	assert.Equal(t, "Jim key", all[0].Description)
	assert.Equal(t, "2020-12-01 15:25:01", all[0].CreationDate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1", all[0].Fingerprint)
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
	server := mockServer{t: t, expectedURL: "/ssh-keys?page=1&pageSize=25", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	all, err := repo.GetSelection(1, 25)
	require.NoError(t, err)
	require.Equal(t, 1, len(all))

	assert.EqualValues(t, 123, all[0].ID)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", all[0].Key)
	assert.Equal(t, "Jim key", all[0].Description)
	assert.Equal(t, "2020-12-01 15:25:01", all[0].CreationDate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1", all[0].Fingerprint)
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
	server := mockServer{t: t, expectedURL: "/ssh-keys/123", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	sshKey, err := repo.GetByID(123)
	require.NoError(t, err)

	assert.EqualValues(t, 123, sshKey.ID)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", sshKey.Key)
	assert.Equal(t, "Jim key", sshKey.Description)
	assert.Equal(t, "2020-12-01 15:25:01", sshKey.CreationDate.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "bb:22:43:69:2b:0d:3e:16:58:91:27:8a:62:29:97:d1", sshKey.Fingerprint)
}

func TestRepository_Add(t *testing.T) {
	const expectedRequestBody = `{"description":"Jim key","sshKey":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example"}`
	server := mockServer{t: t, expectedURL: "/ssh-keys", expectedMethod: "POST", statusCode: 201, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Add("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDf2pxWX/yhUBDyk2LPhvRtI0LnVO8PyR5Zt6AHrnhtLGqK+8YG9EMlWbCCWrASR+Q1hFQG example", "Jim key")
	require.NoError(t, err)
}

func TestRepository_Update(t *testing.T) {
	const expectedRequestBody = `{"description":"Jim key"}`
	server := mockServer{t: t, expectedURL: "/ssh-keys/123", expectedMethod: "PUT", statusCode: 204, expectedRequest: expectedRequestBody}
	client, tearDown := server.getClient()
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
	server := mockServer{t: t, expectedURL: "/ssh-keys/123", expectedMethod: "DELETE", statusCode: 204}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := Repository{Client: *client}

	err := repo.Remove(123)
	require.NoError(t, err)
}
