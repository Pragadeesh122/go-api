package main

import (
	"fmt"
	"go_api/db"
	"go_api/routes/events"
	"go_api/routes/users"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("Go API Server")
	server := gin.Default()

	users.RegisterUserRoutes(server)
	events.RegisterEventRoutes(server)

	server.Run(":3000")
}
