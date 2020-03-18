package authenticator

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// See: https://api.transip.nl/v6/auth
func TestSignWithKey(t *testing.T) {
	key, err := ioutil.ReadFile("../testdata/signature.key")
	require.NoError(t, err)

	requestBody := &AuthRequest{
		Login: "test-user",
		Nonce: "98475920834",
	}

	bodyToSign, err := json.Marshal(requestBody)
	require.NoError(t, err)

	signature, err := signWithKey(bodyToSign, []byte(key))
	require.NoError(t, err)

	fixture := "TKjrjkdRqJLTQJI9QtI3JETV554bnrCmWUbUNdpUg/9OwOYHmtK76gjGs5nyWHVgOBHO9KZ15bCjkup/mzZP2sBnUtqfXxqXBfSh6bn/7a/1gOJzK71RtO84S0q1x7+DGago1OuYSMOdj8mgEMBUtY4aHHpHEy7eCahsCJCTEfMUb05Cq87mhE4XrjjGN2BJ8tEHPMxpWHjEtX1Z8uyaL0XY5l6dBmy1QP+ChyISrNe1n3gYZs9tyyPA9vgW+TqEVgq7mHL8l+g2Va1BwxR+rChoa5gTiJcA9fKJ8evVIBfcocXjRduMFzQW/SMp/yp3I4P7J0lUO0vDWVjEO8LX4A=="
	assert.Equal(t, fixture, signature)
}
