package sshkey

import (
	"fmt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net/url"
)

// Repository can be used to add, modify, remove, or list SSH keys in your account
type Repository repository.RestRepository

// GetAll returns an array of all SSH keys in your account
func (r *Repository) GetAll() ([]SSHKey, error) {
	var response sshKeysWrapper
	err := r.Client.Get(rest.Request{Endpoint: "/ssh-keys"}, &response)

	return response.SSHKeys, err
}

// GetSelection returns a limited list of SSH keys,
// specify how many and which page/chunk of SSH keys you want to retrieve
func (r *Repository) GetSelection(page int, itemsPerPage int) ([]SSHKey, error) {
	var response sshKeysWrapper
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/ssh-keys", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.SSHKeys, err
}

// GetByID returns a specific SSH key struct by id
func (r *Repository) GetByID(sshKeyID int64) (SSHKey, error) {
	var response sshKeyWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/ssh-keys/%d", sshKeyID)}
	err := r.Client.Get(restRequest, &response)

	return response.SSHKey, err
}

// Add allows you add an SSH key
func (r *Repository) Add(key string, description string) error {
	requestBody := addSSHKeyRequest{
		SSHKey:      key,
		Description: description,
	}
	restRequest := rest.Request{Endpoint: "/ssh-keys", Body: &requestBody}

	return r.Client.Post(restRequest)
}

// Update allows you to modify the SSH key description
func (r *Repository) Update(sshKey SSHKey) error {
	requestBody := modifySSHKeyRequest{Description: sshKey.Description}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/ssh-keys/%d", sshKey.ID), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Remove will remove the SSH key
func (r *Repository) Remove(sshKeyID int64) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/ssh-keys/%d", sshKeyID)}

	return r.Client.Delete(restRequest)
}
