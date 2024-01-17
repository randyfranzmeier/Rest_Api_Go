package routes

import (
	"Rest_Api_Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetALlEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetSingleEvent(context *gin.Context) {
	//extract id:
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventID."})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func NewEvent(context *gin.Context) {
	var event models.Event                //create event
	err := context.ShouldBindJSON(&event) //populate event per user request

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Unable to parse data!"})
		return
	}

	userID := context.GetInt64("userID")
	event.UserID = userID
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"msg": "Event Created", "event": event})
}

func UpdateEvent(context *gin.Context) {
	//extract id:
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventID."})
		return
	}
	userID := context.GetInt64("userID")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"mg": "not authorized!"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) //populate event per user request
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Unable to parse data!"})
		return
	}
	updatedEvent.ID = eventID //use newfound event id

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event!"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"msg": "Event updated successfully!!!"})

}

func DeleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventID."})
		return
	}

	userID := context.GetInt64("userID")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"mg": "not authorized to delete event!"})
		return
	}

	err = models.DeleteEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg:": "Unable to delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg:": "Event deleted successfully!"})
}
