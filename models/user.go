package models

import "errors"

type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// Пример метода валидации модели
func (u *User) Validate() error {
	if u.FirstName == "" || u.Email == "" {
		return errors.New("first name and email are required")
	}
	return nil
}
