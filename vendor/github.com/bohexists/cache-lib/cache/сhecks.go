package cache

import (
	"container/list"
	"errors"
	"time"
)

// validateKey checks if the key is valid.
func validateKey(key string) error {
	if key == "" {
		return errors.New("key empty")
	}
	return nil
}

// isExpired checks if the object is expired.
func isExpired(object *cacheObject) bool {
	return time.Now().UnixNano() > object.expired
}

// checkCacheSize checks if the cache size exceeds the maximum limit.
func checkCacheSize(data map[string]*list.Element, maxSize int) error {
	if maxSize > 0 && len(data) >= maxSize {
		return errors.New("cache size exceeds maximum limit")
	}
	return nil
}
