package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"roai.global/utils"
	"strings"
)

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		roleAllowed := false
		for _, r := range roles {
			if claims.Role == r {
				roleAllowed = true
				break
			}
		}
		if !roleAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Set("wallet", claims.Wallet)
		c.Set("role", claims.Role)
		c.Next()
	}
}
