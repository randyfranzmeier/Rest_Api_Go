package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)

	server.GET("events/:id", GetSingleEvent)

	server.POST("/events", NewEvent)

	server.PUT("/events/:id", UpdateEvent)

	server.DELETE("/events/:id", DeleteEvent)

	server.POST("/signup", SignUp)

	server.POST("/login", UserLogin)
}
