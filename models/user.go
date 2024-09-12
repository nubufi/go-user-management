package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	gorm.Model

	// Email is the email address of the user.
	Email string `json:"email" gorm:"unique" binding:"required"`

	// Name is the name of the user.
	Name string `json:"name" binding:"required"`

	// Age is the age of the user.
	Age int `json:"age" binding:"required"`
}
