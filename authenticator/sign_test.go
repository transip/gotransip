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

	fixture := "MjUq83SdzBrT4yqoEv5qQj5GOjC6xjnvj8wm6%2FQ5LWxNeQSc9yl8n44vE4mw0XkiL%2FRyKj1fJxoMu2lgw%2B7Wn5J7aTcHGiURbJksW%2BR1GyVU5czy9D9L3ZehfZbKkED4pCwMhsjLTUlbaLRN%2FKjgJjdTH76C6uJFJBNYCGH0FDQe0TTy8JTaIFX6OLU%2FOFywasffrmnn%2F7kR1ue0hjb5ghfoMcQg55klYsbUihprdWPerjMsMY%2B2QUClTVpfG%2FkBFDwLn1A6ViWCF9O9yWM8nOxFmBIjQnNrFkwwBU5jMbVL6eUaizd%2FemOBsCe1XWN%2B0Unx5Ph9vyQzh86PpxvV7Q%3D%3D"
	assert.Equal(t, fixture, signature)
}
