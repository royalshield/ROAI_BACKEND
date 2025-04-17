package routes

import (
	"github.com/gin-gonic/gin"
	"roai.global/controllers"
	"roai.global/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/auth/wallet", controllers.WalletLogin)

	api.GET("/users", middlewares.AuthMiddleware("admin"), controllers.GetAllUsers)

	api.POST("/notifications", middlewares.AuthMiddleware("admin"), controllers.CreateNotification)
	api.GET("/notifications", middlewares.AuthMiddleware("admin", "user"), controllers.GetNotifications)
	api.POST("/notifications/:id/delete", middlewares.AuthMiddleware("admin"), controllers.DeleteNotification)

	api.POST("/events", middlewares.AuthMiddleware("admin"), controllers.CreateEvent)
	api.GET("/events", middlewares.AuthMiddleware("admin", "user"), controllers.GetAllEvents)
	api.POST("/events/:id/register", middlewares.AuthMiddleware("user"), controllers.RegisterEvent)
	api.POST("/events/:id/delete", middlewares.AuthMiddleware("admin"), controllers.DeleteEvent)
	api.POST("/events/:id/attendees", middlewares.AuthMiddleware("admin"), controllers.GetEventAttendees)
}
