package routers

import (
	"goUserManagement/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes sets up the user-related routes for the given router.
// It registers the following routes:
//
// - POST /users: Create a new user
//
// - GET /users/:id: Get user details by ID
//
// - PUT /users/:id: Update user details by ID
//
// - DELETE /users/:id: Delete a user by ID
//
// Parameters:
// - router: The gin.Engine instance to register the routes with.
//
// Example usage:
//
//	router := gin.Default()
//	UserRoutes(router)
func UserRoutes(router *gin.Engine) {
	// Route to create a new user
	router.POST("/users", controllers.CreateUser)

	// Route to get user details by ID
	router.GET("/users/:id", controllers.GetUser)

	// Route to update user details by ID
	router.PUT("/users/:id", controllers.UpdateUser)

	// Route to delete a user by ID
	router.DELETE("/users/:id", controllers.DeleteUser)
}
