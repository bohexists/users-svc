package internal

import (
	"time"

	"github.com/bohexists/cache-lib/cache"
	"github.com/bohexists/users-svc/models"
	"github.com/google/uuid"
)

type cacheObject struct {
	value   models.User
	expired time.Time
}

// Storage defines the interвface for user data storage operations
type Storage interface {
	CreateUser(u models.User) (string, error)
	GetUser(id string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id string, u models.User) error
	DeleteUser(id string) error
}

// CacheStorage is a storage implementation that uses an in-memory cache
type CacheStorage struct {
	cache *cache.Cache
}

// NewCacheStorage initializes a new CacheStorage instance
func NewCacheStorage() *CacheStorage {
	return &CacheStorage{
		cache: cache.New(cache.CacheConfig{MaxSize: 100, DefaultTTL: time.Minute}),
	}
}

// CreateUser save a new user in the cache
func (s *CacheStorage) CreateUser(u models.User) (string, error) {
	u.ID = uuid.New().String()
	if err := s.cache.Set(u.ID, u); err != nil {
		return "", err
	}
	return u.ID, nil
}

// GetUser retrieves a user from the cache
func (s *CacheStorage) GetUser(id string) (*models.User, error) {
	item, err := s.cache.Get(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}
	u := item.(models.User)
	return &u, nil
}

// GetAllUsers retrieves all users from the cache
func (s *CacheStorage) GetAllUsers() ([]models.User, error) {
	keys, err := s.cache.Keys()
	if err != nil {
		return nil, err
	}
	var users []models.User
	for _, key := range keys {
		item, err := s.cache.Get(key)
		if err != nil {
			continue
		}
		if item != nil {
			u := item.(models.User)
			users = append(users, u)
		}
	}
	return users, nil
}

// UpdateUser updates an existing user in the cache
func (s *CacheStorage) UpdateUser(id string, u models.User) error {
	return s.cache.Set(id, u)
}

// DeleteUser removes a user from the cache
func (s *CacheStorage) DeleteUser(id string) error {
	return s.cache.Delete(id)
}
