package cache

import "time"

// LaunchCleaner periodically removes expired cache data.
func (c *Cache) LaunchCleaner(interval time.Duration) {
	go func() {
		for range time.Tick(interval) {
			c.mu.Lock()
			for key, elem := range c.data {
				if isExpired(elem.Value.(*cacheObject)) {
					c.ll.Remove(elem)
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
