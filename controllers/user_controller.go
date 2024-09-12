package controllers

import (
	"goUserManagement/models"
	"goUserManagement/repository"
	"goUserManagement/utils"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user.
// It binds the JSON data from the request body to a User struct,
// then calls the CreateUser method of the UserRepository to persist the user in the database.
// If there is an error during the binding or creation process, it returns an appropriate error response.
// Otherwise, it returns a JSON response with the created user.
// This function also logs the time taken to create a user using a goroutine and a wait group.
func CreateUser(c *gin.Context) {
	// Initialize a wait group and channel for logging
	var wg sync.WaitGroup
	ch := make(chan time.Duration, 1)

	// Start a goroutine to log the time taken to create a CreateUser
	wg.Add(1)
	go utils.LogRequestDuration(&wg, ch, "POST /users")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repository.CreateUser(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)

	// Log the time taken to create a CreateUser
	wg.Wait()
	close(ch)
}

// GetUser gets user details by ID.
// It extracts the ID from the request parameters, then calls the GetUser method of the UserRepository to fetch the user details.
// If the user is not found, it returns a 404 Not Found response.
// Otherwise, it returns a JSON response with the user details.
// This function also logs the time taken to get user details using a goroutine and a wait group.
func GetUser(c *gin.Context) {
	// Initialize a wait group and channel for logging
	var wg sync.WaitGroup
	ch := make(chan time.Duration, 1)

	// Start a goroutine to log the time taken to create a CreateUser
	wg.Add(1)
	go utils.LogRequestDuration(&wg, ch, "GET /users/:id")

	id := c.Param("id")
	user := repository.GetUser(id)
	c.JSON(http.StatusOK, user)

	// Log the time taken to create a CreateUser
	wg.Wait()
	close(ch)
}

// UpdateUser updates user details by ID.
// It binds the JSON data from the request body to a User struct,
// then calls the UpdateUser method of the UserRepository to update the user details in the database.
// If there is an error during the binding or update process, it returns an appropriate error response.
// Otherwise, it returns a JSON response with the updated user details.
// This function also logs the time taken to update user details using a goroutine and a wait group.
func UpdateUser(c *gin.Context) {
	// Initialize a wait group and channel for logging
	var wg sync.WaitGroup
	ch := make(chan time.Duration, 1)

	// Start a goroutine to log the time taken to create a CreateUser
	wg.Add(1)
	go utils.LogRequestDuration(&wg, ch, "UPDATE /users/:id")

	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr, err := repository.UpdateUser(id, &user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, usr)

	// Log the time taken to create a CreateUser
	wg.Wait()
	close(ch)
}

// DeleteUser deletes a user by ID.
// It extracts the ID from the request parameters, then calls the DeleteUser method of the UserRepository to delete the user.
// If there is an error during the deletion process, it returns an appropriate error response.
// Otherwise, it returns a JSON response indicating that the user was deleted successfully.
// This function also logs the time taken to delete a user using a goroutine and a wait group.
func DeleteUser(c *gin.Context) {
	// Initialize a wait group and channel for logging
	var wg sync.WaitGroup
	ch := make(chan time.Duration, 1)

	// Start a goroutine to log the time taken to create a CreateUser
	wg.Add(1)
	go utils.LogRequestDuration(&wg, ch, "DELETE /users/:id")

	id := c.Param("id")

	// Check if the ID is valid and convert it to an integer
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the repository to delete the user by ID
	err := repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

	// Log the time taken to create a CreateUser
	wg.Wait()
	close(ch)
}
