package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"roai.global/config"
	"roai.global/models"
	"roai.global/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	config.ConnectDB()

	err = config.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Notification{})
	if err != nil {
		return
	}

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":" + os.Getenv("PORT"))
}
