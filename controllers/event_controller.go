package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"roai.global/config"
	"roai.global/models"
)

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(event)
	config.DB.Create(&event)
	c.JSON(201, event)
}

func GetAllEvents(c *gin.Context) {
	var events []models.Event
	config.DB.Preload("Attendees").Find(&events)
	c.JSON(200, events)
}

func RegisterEvent(c *gin.Context) {
	eventID := c.Param("id")
	wallet := c.MustGet("wallet").(string)

	var event models.Event
	if err := config.DB.Preload("Attendees").First(&event, eventID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}
	if len(event.Attendees) >= event.MembersLimit {
		c.JSON(400, gin.H{"error": "Event is full"})
		return
	}

	var user models.User
	config.DB.Where("wallet = ?", wallet).First(&user)
	config.DB.Model(&event).Association("Attendees").Append(&user)
	c.JSON(200, gin.H{"message": "Registered successfully"})
}

func DeleteEvent(c *gin.Context) {
	eventID := c.Param("id")
	config.DB.Delete(&models.Event{}, eventID)
	c.JSON(200, gin.H{"message": "Deleted"})
}

func GetEventAttendees(c *gin.Context) {
	eventID := c.Param("id")
	var event models.Event
	config.DB.Preload("Attendees").First(&event, eventID)
	c.JSON(200, event.Attendees)
}
