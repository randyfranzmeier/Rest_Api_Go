package routes

import (
	"Rest_Api_Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg:": "Unable to parse request!"})
		return
	}
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg:": "could not fetch events!"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg:": "Unable to register, sorry!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"msg:": "registered"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg:": "Unable to parse request!"})
		return
	}
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg:": "could not fetch events!"})
		return
	}
	err = event.Cancel(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg:": "unable to delete registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg:": "event deleted successfully"})
}
