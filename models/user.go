package models

import "errors"

// User represents the user model
type User struct {
	ID         string `json:"id" bson:"_id"`
	FirstName  string `json:"first_name" form:"first_name" bson:"first_name"`
	SecondName string `json:"second_name" form:"second_name" bson:"second_name"`
	Email      string `json:"email" form:"email" bson:"email"`
	Password   string `json:"password" form:"password" bson:"password"`
}

// Validate validates the user model
func (u *User) Validate() error {
	if u.FirstName == "" || u.Email == "" {
		return errors.New("first name and email are required")
	}
	return nil
}
