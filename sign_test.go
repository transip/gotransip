package gotransip

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
)

func TestSignWithKey(t *testing.T) {
	key, err := ioutil.ReadFile("testdata/signature.key")
	require.NoError(t, err)

	params := &soapParams{}
	params.Add("__method", "getHaip")
	params.Add("__service", "HaipService")
	params.Add("__hostname", "api.transip.nl")
	params.Add("__timestamp", 1534839460)
	params.Add("__nonce", "5b7bcab97f1a98.77032926")

	signature, err := signWithKey(params, []byte(key))
	require.NoError(t, err)

	fixture := "VBALqck%2BScVciJb%2BN2uBaPT1GaW7uLvKDsC9GfLPZT5sAIh2jhQ9Dc5mDIW9czrxxzOY3Vl1AWQI%2FMDbcIHAT4s8umpBYs8ZH4ORqiMZn4FOcypKRdPZOIdeHsqF%2FMv0Yb5YwBR6lBJrXAdh8DM%2BWt%2Fi8ZfPZV8KtXOyFb1zna9xEVmco6TSDL%2BpjUHDvzRkgJsYLZeZEOrvM7qDOxCVWRz1BKuf7wgsam6VGA6QFMA1mcjWdRd89X55075WwNXm4tEOGOMq%2BN5cf6N%2BERQMVPUF9w3EIv50bJCayWJduuk73sHUJn80tquJ6eVjky%2FS%2FDG1hUhyvngGUhwFaAWcJw%3D%3D"
	assert.Equal(t, fixture, signature)
}
