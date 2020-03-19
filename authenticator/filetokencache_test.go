package authenticator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/jwt"
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

	tokenToCache := jwt.Token{ExpiryDate: 2118745550, RawToken: DemoToken}
	err = cache.Set("testkey", tokenToCache)
	require.NoError(t, err)

	dataFromCache, err := cache.Get("testkey")
	require.NoError(t, err)
	assert.Equal(t, tokenToCache, dataFromCache)
}

func TestFileTokenCache_SetGetFromFile(t *testing.T) {
	defer os.Remove("/tmp/gotransip_cache_setgetfromfile")

	cache, err := NewFileTokenCache("/tmp/gotransip_cache_setgetfromfile")
	require.NoError(t, err)

	tokenToCache := jwt.Token{RawToken: DemoToken + DemoToken}
	err = cache.Set("testkey", tokenToCache)
	require.NoError(t, err)

	// write again so we know the File gets overridden
	tokenToCache = jwt.Token{ExpiryDate: 2118745550, RawToken: DemoToken}
	err = cache.Set("testkey", tokenToCache)
	require.NoError(t, err)

	// close the File so we know we will fetch it with a new File token cache
	err = cache.File.Close()
	require.NoError(t, err)

	cache, err = NewFileTokenCache("/tmp/gotransip_cache_setgetfromfile")
	require.NoError(t, err)

	dataFromCache, err := cache.Get("testkey")
	require.NoError(t, err)
	assert.Equal(t, tokenToCache, dataFromCache)
}
