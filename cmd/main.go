package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"roai.global/config"
	"roai.global/models"
	"roai.global/routes"
	"time"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Everyone
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // must be false when AllowOrigins is "*"
		MaxAge:           12 * time.Hour,
	}))
	routes.RegisterRoutes(r)

	r.Run(":" + os.Getenv("PORT"))
}
