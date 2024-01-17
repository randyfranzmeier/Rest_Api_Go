package routes

import (
	"Rest_Api_Go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)

	server.GET("events/:id", GetSingleEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", NewEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", SignUp)

	server.POST("/login", UserLogin)

	authenticated.PUT("/changePassword", ChangePassword)
}
