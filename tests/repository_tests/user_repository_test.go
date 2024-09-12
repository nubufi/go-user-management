package repository_test

import (
	"strconv"
	"testing"

	"goUserManagement/models"
	"goUserManagement/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// Initialize an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	// Initialize the database connection in the repository package
	repository.InitDatabase(db)

	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Define a new user
	user := &models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}

	// Create the user
	err = repository.CreateUser(user)
	assert.NoError(t, err, "Expected no error when creating user")
	assert.NotZero(t, user.ID, "Expected user ID to be set after creation")

	// Check if the user exists in the database
	var foundUser models.User
	err = db.Where("email = ?", user.Email).First(&foundUser).Error
	assert.NoError(t, err, "Expected no error when finding user")
	assert.Equal(t, user.Email, foundUser.Email, "Expected found user to have the same email")
}

func TestCreateUserDuplicateEmail(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create the first user
	user1 := &models.User{
		Name:  "John Doe",
		Email: "john.doe2@example.com",
		Age:   30,
	}
	err = repository.CreateUser(user1)
	assert.NoError(t, err, "Expected no error when creating first user")

	// Try to create a second user with the same email
	user2 := &models.User{
		Name:  "Jane Doe",
		Email: "john.doe2@example.com", // Duplicate email
		Age:   25,
	}
	err = repository.CreateUser(user2)
	assert.Error(t, err, "Expected error when creating user with duplicate email")
}

func TestDeleteUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create a user to delete
	user := &models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}
	err = repository.CreateUser(user)
	assert.NoError(t, err, "Expected no error when creating user")

	id := strconv.FormatUint(uint64(user.ID), 10)
	err = repository.DeleteUser(id)
	assert.NoError(t, err, "Expected no error when deleting user")

	// Check that the user no longer exists
	var foundUser models.User
	err = db.Where("email = ?", user.Email).First(&foundUser).Error
	assert.Error(t, err, "Expected error when looking for a deleted user")
	assert.Equal(t, gorm.ErrRecordNotFound, err, "Expected record not found error after deletion")
}

func TestGetUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create a user to retrieve
	user := &models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}
	err = repository.CreateUser(user)
	assert.NoError(t, err, "Expected no error when creating user")

	id := strconv.FormatUint(uint64(user.ID), 10)
	retrievedUser := repository.GetUser(id)
	assert.Equal(t, user.Email, retrievedUser.Email, "Expected retrieved user to have the same email")
}

func TestUpdateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create a user to update
	user := &models.User{
		Name:  "John Doe",
		Email: "john.doe3@example.com",
		Age:   30,
	}

	err = repository.CreateUser(user)
	assert.NoError(t, err, "Expected no error when creating user")

	// Update the user
	user.Name = "Jane Doe"
	_, err = repository.UpdateUser("1", user)
	assert.NoError(t, err, "Expected no error when updating user")

	// Check if the user has been updated
	var updatedUser models.User
	err = db.Where("email = ?", user.Email).First(&updatedUser).Error
	assert.NoError(t, err, "Expected no error when finding user")
	assert.Equal(t, user.Name, updatedUser.Name, "Expected updated user to have the new name")
}
