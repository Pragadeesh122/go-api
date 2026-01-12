package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api/db"
	"go_api/routes/events"
)

func main() {
	db.InitDB()
	fmt.Println("Go API Server")
	server := gin.Default()

	events.RegisterEventRoutes(server)

	server.Run(":3000")
}
