package routes

import (
	"icpscan/src/controllers"
)

func addUserRoutes() {
	UserRoutes := Router.Group("/api/user")
	{
		// UserRoutes.GET("/", controllers.UserIndex)
		// UserRoutes.GET("/:id", controllers.UserShow)
		// UserRoutes.POST("/", controllers.UserCreate)
		// UserRoutes.PUT("/:id", controllers.UserUpdate)
		// UserRoutes.DELETE("/:id", controllers.UserDelete)
		UserRoutes.GET("/login", controllers.Login)
		UserRoutes.POST("/wallet_relation", controllers.WalletRelation)
	}
}
