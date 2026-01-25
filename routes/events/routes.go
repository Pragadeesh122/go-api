package events

import (
	"go_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:eventID", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events", updateEvent)
	authenticated.PUT("/events/batch", updateEvents)
	authenticated.DELETE("/events/:eventID", deleteEvent)
	authenticated.DELETE("/events/batch", deleteEvents)

}
