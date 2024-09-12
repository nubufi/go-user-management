package main

import (
	"os"

	"goUserManagement/config"
	"goUserManagement/repository"
	"goUserManagement/routers"
	"goUserManagement/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Initialize the application
func init() {
	godotenv.Load()
	db := config.ConnectToDb()
	config.Migrate(db)
	repository.InitDatabase(db)
	utils.InitLogger()
}

func main() {
	app := gin.New()
	gin.SetMode(gin.ReleaseMode)
	routers.UserRoutes(app)

	port := os.Getenv("HTTP_PORT")
	app.Run(":" + port)
}
