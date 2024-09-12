package repository

import (
	"goUserManagement/models"

	"gorm.io/gorm"
)

var db *gorm.DB

// InitDatabase initializes the database connection and stores it in the repository package
func InitDatabase(d *gorm.DB) {
	db = d
}

// CreateUser creates a new user in the repository.
// It takes a pointer to a User model as a parameter and returns an error.
// The user parameter represents the user to be created.
// If the user is successfully created, it returns nil.
// Otherwise, it returns an error indicating the reason for the failure.
func CreateUser(user *models.User) error {
	return db.Create(user).Error
}

// GetUser retrieves a user from the repository based on the provided ID.
// It returns a pointer to the User object if found, otherwise it returns nil.
func GetUser(id string) *models.User {
	var user models.User
	db.First(&user, id)
	return &user
}

// UpdateUser updates the given user in the repository.
// It saves the user to the database and returns an error if any occurred.
func UpdateUser(id string, user *models.User) (models.User, error) {
	var existingUser models.User

	// Find the user by ID
	if err := db.First(&existingUser, id).Error; err != nil {
		return models.User{}, err
	}

	// Update user fields
	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Age = user.Age

	// Save the updated user back to the database
	if err := db.Save(&existingUser).Error; err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

// DeleteUser deletes a user from the repository based on the given ID.
// It returns an error if the user is not found or if there is an error during deletion.
func DeleteUser(id string) error {
	user := &models.User{}
	err := db.First(user, id).Error
	if err != nil {
		return err
	}

	err = db.Delete(user).Error
	if err != nil {
		return err
	}

	return err
}
