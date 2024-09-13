package models

import "errors"

// User represents the user model
type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// Validate validates the user model
func (u *User) Validate() error {
	if u.FirstName == "" || u.Email == "" {
		return errors.New("first name and email are required")
	}
	return nil
}
