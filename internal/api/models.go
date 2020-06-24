package api

import "github.com/jinzhu/gorm"

// Thing Model
type User struct {
	gorm.Model
	FirstName string `json:"first_name,omitempty" validate:"required"`
	LastName  string `json:"last_name,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Username  string `json:"username,omitempty"`
}
