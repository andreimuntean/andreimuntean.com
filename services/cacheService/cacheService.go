package cacheService

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/memcache"
)

// Add adds a byte array to the cache.
func Add(c context.Context, key string, data []byte) {
	memcache.Add(c, &memcache.Item{
		Key:   key,
		Value: data,
	})
}

// Get retrieves a byte array from the cache.
func Get(c context.Context, key string) []byte {
	item, err := memcache.Get(c, key)

	if err != nil {
		return nil
	}

	return item.Value
}
