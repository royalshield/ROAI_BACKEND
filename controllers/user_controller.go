package controllers

import (
	"github.com/gin-gonic/gin"
	"roai.global/config"
)

func GetAllUsers(c *gin.Context) {
	var users []map[string]interface{}
	config.DB.Model(&[]map[string]interface{}{}).Table("users").Find(&users)
	c.JSON(200, users)
}
