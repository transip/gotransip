package authenticator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestFileTokenCache_New(t *testing.T) {
	defer os.Remove("/tmp/gotransip_test123")

	_, err := NewFileTokenCache("/tmp/gotransip_test123")
	require.NoError(t, err)
}

func TestFileTokenCache_Set(t *testing.T) {
	defer os.Remove("/tmp/gotransip_test123")

	cache, err := NewFileTokenCache("/tmp/gotransip_test123")
	require.NoError(t, err)

	err = cache.Set("testkey", []byte{0x00, 0x20})
	require.NoError(t, err)

	dataFromCache, err := cache.Get("testkey")
	require.NoError(t, err)
	assert.Equal(t, []byte{0x00, 0x20}, dataFromCache)
}

func TestFileTokenCache_SetGetFromFile(t *testing.T) {
	defer os.Remove("/tmp/gotransip_cache_setgetfromfile")

	cache, err := NewFileTokenCache("/tmp/gotransip_cache_setgetfromfile")
	require.NoError(t, err)

	err = cache.Set("testkey", []byte{0x00, 0x20, 0x30, 0x40, 0x50, 0x60})
	require.NoError(t, err)

	// write again so we know the file gets overridden
	err = cache.Set("testkey", []byte{0x00, 0x20})
	require.NoError(t, err)

	// close the file so we know we will fetch it with a new file token cache
	err = cache.file.Close()
	require.NoError(t, err)

	cache, err = NewFileTokenCache("/tmp/gotransip_cache_setgetfromfile")
	require.NoError(t, err)

	dataFromCache, err := cache.Get("testkey")
	require.NoError(t, err)
	assert.Equal(t, []byte{0x00, 0x20}, dataFromCache)
}
