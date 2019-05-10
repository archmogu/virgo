package routes

import (
	"app/modules/user/controllers"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes ...
func RegisterRoutes(router *gin.Engine) {
	userController := controllers.NewUserController()
	v1 := router.Group("api/v1")
	{
		v1.POST("/login", userController.Login)

		auth := v1.Group("", middlewares.Jwt())
		{
			auth.GET("/users", userController.Index)
			auth.POST("/users", userController.Store)
			auth.GET("/users/:id", userController.Show)
			auth.PUT("/users/:id", userController.Update)
			auth.DELETE("/users/:id", userController.Destroy)
		}
	}
}
