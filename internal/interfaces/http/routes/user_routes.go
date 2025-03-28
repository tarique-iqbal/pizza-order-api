package routes

import (
	"pizza-order-api/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *http.UserHandler) {
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/signup", userHandler.CreateUser)
			users.POST("/signin", userHandler.SignIn)
		}
	}
}
