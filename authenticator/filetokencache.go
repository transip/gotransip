package authenticator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// cacheItem is one named item inside the filesystem cache
type cacheItem struct {
	// Key of the cache item, containing
	Key string `json:"key"`
	// Data containing the content of the cache item
	Data []byte `json:"data"`
}

// FileTokenCache is a cache that takes a path and writes a json marshalled file to it,
// it decodes it when created with the NewFileTokenCache method.
// It has a Set method to save a token by name as byte array
// and a Get method one to get a previously acquired token by name returned as byte array
type FileTokenCache struct {
	file *os.File
	// CacheItems contains a list of cache items, all of them have a key
	CacheItems []cacheItem `json:"items"`
	// prevent simultaneous cache writes
	writeLock sync.RWMutex
}

// NewFileTokenCache creates a filesystem cache file on the specified path
func NewFileTokenCache(path string) (*FileTokenCache, error) {
	// open the file or create a new one on the given location
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return &FileTokenCache{}, fmt.Errorf("error opening cache file: %w", err)
	}

	cache := FileTokenCache{file: file}

	// try to read the file
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return &FileTokenCache{}, fmt.Errorf("error reading cache file: %w", err)
	}

	if len(fileContent) > 0 {
		// read the cached file data as json
		if err := json.Unmarshal(fileContent, &cache); err != nil {
			return &FileTokenCache{}, fmt.Errorf("error unmarshalling cache file: %w", err)
		}
	}

	return &cache, nil
}

// Set will save a token by name as byte array
func (f *FileTokenCache) Set(key string, data []byte) error {
	for idx, item := range f.CacheItems {
		if item.Key == key {
			f.CacheItems[idx].Data = data

			// persist this change to the cache file
			return f.writeCacheToFile()
		}
	}

	// if the key did not exist before, we append a new item to the cache item list
	f.CacheItems = append(f.CacheItems, cacheItem{Key: key, Data: data})

	return f.writeCacheToFile()
}

func (f *FileTokenCache) writeCacheToFile() error {
	// try to convert the cache to json, so we can write it to file
	cacheData, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("error marshalling cache file: %w", err)
	}

	f.writeLock.Lock()
	defer f.writeLock.Unlock()
	// write the cache data to the file cache
	if err := f.file.Truncate(0); err != nil {
		return fmt.Errorf("error while truncating cache file: %w", err)
	}
	_, err = f.file.WriteAt(cacheData, 0)

	return err
}

// Get a previously acquired token by name returned as byte array
func (f *FileTokenCache) Get(key string) ([]byte, error) {
	for _, item := range f.CacheItems {
		if item.Key == key {
			return item.Data, nil
		}
	}

	return nil, nil
}
