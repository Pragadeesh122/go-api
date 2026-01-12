package events

import (
	"go_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
