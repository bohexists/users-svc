package main

import (
	"github.com/bohexists/cache-library/cache"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// Storage defines the interвface for user data storage operations
type Storage interface {
	CreateUser(u User) (string, error)
	GetUser(id string) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id string, u User) error
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
func (s *CacheStorage) CreateUser(u User) (string, error) {
	u.ID = uuid.New().String()
	if err := s.cache.Set(u.ID, u); err != nil {
		return "", err
	}
	return u.ID, nil
}

// GetUser retrieves a user from the cache
func (s *CacheStorage) GetUser(id string) (*User, error) {
	item, err := s.cache.Get(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}
	u := item.(User)
	return &u, nil
}

// GetAllUsers retrieves all users from the cache
func (s *CacheStorage) GetAllUsers() ([]User, error) {
	keys, err := s.cache.Keys()
	if err != nil {
		return nil, err
	}
	var users []User
	for _, key := range keys {
		item, err := s.cache.Get(key)
		if err != nil {
			continue
		}
		if item != nil {
			u := item.(User)
			users = append(users, u)
		}
	}
	return users, nil
}

// UpdateUser updates an existing user in the cache
func (s *CacheStorage) UpdateUser(id string, u User) error {
	return s.cache.Set(id, u)
}

// DeleteUser removes a user from the cache
func (s *CacheStorage) DeleteUser(id string) error {
	return s.cache.Delete(id)
}
