package events

import (
	"github.com/gin-gonic/gin"
	"go_api/models"
	"net/http"
	"strconv"
	"time"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event data"})
		return
	}

	event.Date = time.Now()
	event.UserID = 3

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, event)
}

func getEvent(context *gin.Context) {
	eventID := context.Param("eventID")

	id, err := strconv.ParseInt(eventID, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "The eventID cannot be parsed! Make sure it is an integer"})
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	update_err := event.UpdateEvent()

	if update_err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "The event was successfully updated"})

}

func updateEvents(context *gin.Context) {
	var events []models.Event
	err := context.ShouldBindBodyWithJSON(&events)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	for _, event := range events {
		err := event.UpdateEvent()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	context.JSON(http.StatusOK, gin.H{"messgae": "The events are successfully updated"})

}
