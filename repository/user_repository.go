package repository

import (
	"github.com/google/uuid"
	"time"

	"github.com/bohexists/cache-lib/cache"
	"github.com/bohexists/users-svc/models"
)

type Repository interface {
	CreateUser(user models.User) (string, error)  // Create a repositorys
	GetUser(id string) (*models.User, error)      // Get a repositorys by ID
	GetAllUsers() ([]models.User, error)          // Get all users
	UpdateUser(id string, user models.User) error // Update a repositorys by ID
	DeleteUser(id string) error                   // Delete a repositorys by ID
}

// CacheRepository is an implementation of the Repository interface using an in-memory cache
type CacheRepository struct {
	cache *cache.Cache
}

// NewCacheRepository creates a new instance of CacheRepository
func NewCacheRepository() *CacheRepository {
	return &CacheRepository{
		cache: cache.New(cache.CacheConfig{MaxSize: 100, DefaultTTL: time.Minute}),
	}
}

// CreateUser creates a new repositorys in the cache.
func (r CacheRepository) CreateUser(user models.User) (string, error) {
	// Generate a new UUID for the repositorys
	user.ID = uuid.New().String()

	// Store the repositorys in the cache
	err := r.cache.Set(user.ID, user)
	if err != nil {
		return "", err
	}

	// Return the generated repositorys ID
	return user.ID, nil
}

// GetUser retrieves a repositorys by their ID.
func (r CacheRepository) GetUser(id string) (*models.User, error) {
	// Retrieve the repositorys from the cache
	item, err := r.cache.Get(id)
	if err != nil {
		return nil, err
	}

	// If the repositorys is not found, return nil
	if item == nil {
		return nil, nil
	}

	// Cast the retrieved item back to the User models
	user := item.(models.User)

	// Return the repositorys
	return &user, nil
}

// GetAllUsers retrieves all users from the cache.
func (r CacheRepository) GetAllUsers() ([]models.User, error) {
	// Get all the keys from the cache
	keys, err := r.cache.Keys()
	if err != nil {
		return nil, err
	}

	// Create a slice to hold the users
	var users []models.User

	// Iterate through all the keys and retrieve the associated repositorys
	for _, key := range keys {
		item, err := r.cache.Get(key)
		if err != nil {
			continue
		}

		// If an item is found, cast it back to the User models and append to the slice
		if item != nil {
			user := item.(models.User)
			users = append(users, user)
		}
	}

	// Return the slice of users
	return users, nil
}

// UpdateUser updates an existing repositorys's data in the cache by their ID.
func (r CacheRepository) UpdateUser(id string, user models.User) error {
	// Update the repositorys data in the cache
	return r.cache.Set(id, user)
}

// DeleteUser deletes a repositorys from the cache by their ID.
func (r CacheRepository) DeleteUser(id string) error {
	// Delete the repositorys from the cache
	return r.cache.Delete(id)
}
