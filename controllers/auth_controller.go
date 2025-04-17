package controllers

import (
	"github.com/gin-gonic/gin"
	"roai.global/config"
	"roai.global/models"
	"roai.global/utils"
)

func WalletLogin(c *gin.Context) {
	var req struct {
		Wallet string `json:"wallet" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Wallet required"})
		return
	}

	var user models.User
	if err := config.DB.Where("wallet = ?", req.Wallet).First(&user).Error; err != nil {
		if req.Wallet == config.ADMIN_WALLET_ADDRESS {
			user = models.User{Wallet: req.Wallet, Role: "admin"}
		} else {
			user = models.User{Wallet: req.Wallet, Role: "user"}
		}
		config.DB.Create(&user)
	}
	token, _ := utils.GenerateJWT(user.Wallet, user.Role)
	c.JSON(200, gin.H{"token": token})
}
