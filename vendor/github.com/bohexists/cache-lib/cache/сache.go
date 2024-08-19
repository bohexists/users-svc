package cache

import (
	"container/list"
	"sync"
	"time"
)

const (
	FILO = iota
	LRU
	FIFO
)

// CacheConfig holds the configuration for the Cache.
type CacheConfig struct {
	MaxSize      int           // Maximum number of Object in the Cache.
	DefaultTTL   time.Duration // Default TTL for cache objects.
	EvictionType int           // Eviction strategy type.
}

// Cache is a basic in-memory storage for data.
type Cache struct {
	data         map[string]*list.Element
	mu           sync.RWMutex
	maxSize      int
	defaultTTL   time.Duration
	ll           *list.List
	evictionType int
}

// cacheObject struct to store value and expiration in Cache.
type cacheObject struct {
	key     string
	value   interface{}
	expired int64
}

// New creates and returns a new Cache.
func New(config CacheConfig) *Cache {
	return &Cache{
		data:         make(map[string]*list.Element),
		maxSize:      config.MaxSize,
		defaultTTL:   config.DefaultTTL,
		ll:           list.New(),
		evictionType: config.EvictionType,
	}
}

// Set adds a value.
func (c *Cache) Set(key string, value interface{}, ttl ...time.Duration) error {
	// Lock for writing
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := validateKey(key); err != nil {
		return err
	}

	expired := time.Now().Add(c.defaultTTL)
	if len(ttl) > 0 {
		expired = time.Now().Add(ttl[0])
	}

	// Remove the oldest element if the cache size exceeds maxSize
	if c.maxSize > 0 && c.ll.Len() >= c.maxSize {
		c.evict()
	}

	if elem, exists := c.data[key]; exists {
		c.ll.MoveToFront(elem)
		elem.Value.(*cacheObject).value = value
		elem.Value.(*cacheObject).expired = expired.UnixNano()
	} else {
		elem := c.ll.PushFront(&cacheObject{
			key:     key,
			value:   value,
			expired: expired.UnixNano(),
		})
		c.data[key] = elem
	}

	return nil
}

// Get retrieves a value.
func (c *Cache) Get(key string) (interface{}, error) {
	// Lock for reading
	c.mu.RLock()
	defer c.mu.RUnlock()

	if err := validateKey(key); err != nil {
		return nil, err
	}

	elem, exists := c.data[key]
	if !exists || isExpired(elem.Value.(*cacheObject)) {
		return nil, nil // or a specific error indicating the key does not exist or is expired
	}

	c.ll.MoveToFront(elem)
	return elem.Value.(*cacheObject).value, nil
}

// Delete removes a value.
func (c *Cache) Delete(key string) error {
	// Lock for writing
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := validateKey(key); err != nil {
		return err
	}

	if elem, exists := c.data[key]; exists {
		c.ll.Remove(elem)
		delete(c.data, key)
	}

	return nil
}

// Exists checks if a key exists.
func (c *Cache) Exists(key string) bool {
	// Lock for reading
	c.mu.RLock()
	defer c.mu.RUnlock()

	if err := validateKey(key); err != nil {
		return false
	}
	_, exists := c.data[key]
	return exists
}

// Keys returns a list of keys.
func (c *Cache) Keys() ([]string, error) {
	// Lock for reading
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]string, 0, len(c.data))
	for key := range c.data {
		result = append(result, key)
	}
	return result, nil
}

// Clear removes all object from cache.
func (c *Cache) Clear() {
	// Lock for writing
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string]*list.Element)
	c.ll.Init()
}

func (c *Cache) Size() int {
	// Lock for reading
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.data)
}
