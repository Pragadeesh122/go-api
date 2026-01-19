package users

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(server *gin.Engine) {

	server.POST("/signup", signup)
	server.POST("/login", login)
}
