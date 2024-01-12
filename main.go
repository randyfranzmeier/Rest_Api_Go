package main

import (
	"Rest_Api_Go/db"
	"Rest_Api_Go/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("db successfully initialized!!!")
	server := gin.Default() //configures http server

	server.GET("/events", func(context *gin.Context) {
		events, err := models.GetALlEvents()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
			return
		}
		context.JSON(http.StatusOK, events)
	})

	server.GET("events/:id", func(context *gin.Context) {
		//extract id:
		eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventID."})
			return
		}
		event, err := models.GetEventByID(eventID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		}
		context.JSON(http.StatusOK, event)

	})

	server.POST("/events", func(context *gin.Context) {
		var event models.Event                //create event
		err := context.ShouldBindJSON(&event) //works like scan function
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"msg": "Unable to parse data!"})
		}
		event.ID = 1
		event.UserID = 1
		err = event.Save()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event!"})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"msg": "Event Created", "event": event})
	})

	server.Run(":8080") //localhost 8080

}
