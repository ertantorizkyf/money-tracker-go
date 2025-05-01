package router

import (
	"github.com/ertantorizkyf/money-tracker-go/handlers"
	"github.com/gin-gonic/gin"
)

func userRoutes(apiGroup *gin.RouterGroup, userHandler *handlers.UserHandler) {
	userGroup := apiGroup.Group("/users")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.POST("/login", userHandler.Login)
	}
}
