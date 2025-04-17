package controllers

import (
	"github.com/gin-gonic/gin"
	"roai.global/config"
	"roai.global/models"
)

func CreateNotification(c *gin.Context) {
	var notif models.Notification
	if err := c.ShouldBindJSON(&notif); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&notif)
	c.JSON(201, notif)
}

func DeleteNotification(c *gin.Context) {
	notificationID := c.Param("id")
	config.DB.Delete(&models.Notification{}, notificationID)
	c.JSON(200, gin.H{"message": "Deleted"})
}

func GetNotifications(c *gin.Context) {
	var notifs []models.Notification
	config.DB.Order("created_at desc").Find(&notifs)
	c.JSON(200, notifs)
}
