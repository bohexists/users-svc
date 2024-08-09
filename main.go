package main

import (
	"github.com/bohexists/cache-library/cache"
)

type CacheStorage struct {
	cache *cache.Cache
}

func NewCacheStorage() *CacheStorage {
	return &CacheStorage{
		cache: cache.New(cache.CacheConfig{MaxSize: 100, DefaultTTL: 0}),
	}
}
