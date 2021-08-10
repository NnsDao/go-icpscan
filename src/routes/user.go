package routes

import (
	"github.com/MatheusMeloAntiquera/api-go/src/controllers"
)

func addUserRoutes() {
	UserRoutes := Router.Group("/user")
	{
		// UserRoutes.GET("/", controllers.UserIndex)
		// UserRoutes.GET("/:id", controllers.UserShow)
		// UserRoutes.POST("/", controllers.UserCreate)
		// UserRoutes.PUT("/:id", controllers.UserUpdate)
		// UserRoutes.DELETE("/:id", controllers.UserDelete)
		UserRoutes.GET("/login", controllers.Login)
	}
}
