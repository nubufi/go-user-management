package controller_test

import (
	"bytes"
	"encoding/json"
	"goUserManagement/controllers"
	"goUserManagement/models"
	"goUserManagement/repository"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup a test Gin router with our controller routes
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	return router
}

// Setup a test database (in-memory SQLite)
func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	repository.InitDatabase(db)
	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	router := setupRouter()

	// Mock request body for user creation
	user := models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}
	body, _ := json.Marshal(user)

	// Create an HTTP request
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert that the response is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body to validate the created user
	var createdUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, createdUser.Email, "Expected created user to have the same email")
}

func TestGetUser(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	router := setupRouter()

	// Create a test user directly in the database
	user := models.User{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
		Age:   25,
	}
	db.Create(&user)

	// Create an HTTP request to get the user
	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body to validate the retrieved user
	var retrievedUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &retrievedUser)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, retrievedUser.Email, "Expected retrieved user to have the same email")
}

func TestUpdateUser(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	router := setupRouter()

	// Create a test user
	user := models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}
	db.Create(&user)
	// Mock request body to update the user's name
	updatedUser := models.User{
		Email: "john.doe@example.com",
		Name:  "John Updated",
		Age:   35,
	}
	body, _ := json.Marshal(updatedUser)

	// Create an HTTP request to update the user
	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body to validate the updated user
	var retrievedUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &retrievedUser)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser.Name, retrievedUser.Name, "Expected updated user to have the new name")
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	router := setupRouter()

	// Create a test user
	user := models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}
	db.Create(&user)

	// Create an HTTP request to delete the user
	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify the user has been deleted
	var retrievedUser models.User
	err := db.Where("id = ?", user.ID).First(&retrievedUser).Error
	assert.Error(t, err, "Expected error when retrieving a deleted user")
	assert.Equal(t, gorm.ErrRecordNotFound, err, "Expected record not found error after deletion")
}
