package events

import "github.com/gin-gonic/gin"

func RegisterEventRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)

	server.GET("/events/:eventID", getEvent)

	server.POST("/events", createEvent)

}
