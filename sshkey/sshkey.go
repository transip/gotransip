package sshkey

import (
	"github.com/transip/gotransip/v6/rest"
)

// addSSHKeyRequest is used to marshal to a request that adds an SSH key
type addSSHKeyRequest struct {
	Description string `json:"description,omitempty"`
	SSHKey      string `json:"sshKey,omitempty"`
}

// modifySSHKeyRequest is used to marshal to a request that modifies an SSH key
// it is only possible to modify the description
type modifySSHKeyRequest struct {
	Description string `json:"description,omitempty"`
}

// sshKeyWrapper will be used to unmarshal an SSH keys
type sshKeyWrapper struct {
	SSHKey SSHKey `json:"sshKey"`
}

// sshKeysWrapper will be used to unmarshal a list of SSH keys
type sshKeysWrapper struct {
	SSHKeys []SSHKey `json:"sshKeys"`
}

// SSHKey struct
type SSHKey struct {
	// The SSH key id
	ID int64 `json:"id,omitempty"`
	// Description
	Description string `json:"description"`
	// SSH key
	Key string `json:"key"`
	// SSH key fingerprint
	Fingerprint string `json:"fingerprint"`
	// Creation date of the SSH key
	CreationDate rest.Time `json:"creationDate"`
}
