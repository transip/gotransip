package authenticator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/jwt"
	"os"
	"testing"
)

func TestFileTokenCache_New(t *testing.T) {
	tmpFile := os.TempDir() + "/gotransip_test123"
	defer os.Remove(tmpFile)

	_, err := NewFileTokenCache(tmpFile)
	require.NoError(t, err)
}

func TestFileTokenCache_Set(t *testing.T) {
	tmpFile := os.TempDir() + "/gotransip_test123"
	defer os.Remove(tmpFile)

	cache, err := NewFileTokenCache(tmpFile)
	require.NoError(t, err)

	tokenToCache := jwt.Token{ExpiryDate: 2118745550, RawToken: DemoToken}
	err = cache.Set("testkey", tokenToCache)
	require.NoError(t, err)

	dataFromCache, err := cache.Get("testkey")
	require.NoError(t, err)
	assert.Equal(t, tokenToCache, dataFromCache)
}

func TestFileTokenCache_SetGetFromFile(t *testing.T) {
	cacheLocation := os.TempDir() + "/gotransip_cache_setgetfromfile"
	defer os.Remove(cacheLocation)

	cache, err := NewFileTokenCache(cacheLocation)
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

	cache, err = NewFileTokenCache(cacheLocation)
	require.NoError(t, err)

	dataFromCache, err := cache.Get("testkey")
	require.NoError(t, err)
	assert.Equal(t, tokenToCache, dataFromCache)
}
