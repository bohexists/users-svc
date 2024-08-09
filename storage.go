package main

import "github.com/bohexists/cache-library/cache"

type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Storage interface {
	CreateUser(u User) error
	GetUser(id string) (*User, error)
	UpdateUser(id string, u User) error
	DeleteUser(id string) error
	GetAllUsers() ([]User, error)
}

type CacheStorage struct {
	cache *cache.Cache
}

func NewCacheStorage() *CacheStorage {
	return &CacheStorage{
		cache: cache.New(cache.CacheConfig{MaxSize: 100, DefaultTTL: 0}),
	}
}

func (s *CacheStorage) CreateUser(u User) error {
	return s.cache.Set(u.ID, u)
}

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

func (s *CacheStorage) UpdateUser(id string, u User) error {
	return s.cache.Set(id, u)
}

func (s *CacheStorage) DeleteUser(id string) error {
	return s.cache.Delete(id)
}

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
