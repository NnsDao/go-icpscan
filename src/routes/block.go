package routes

import (
	"github.com/MatheusMeloAntiquera/api-go/src/controllers"
)

func addBlockRoutes() {
	BlockRoutes := Router.Group("/block")
	{
		BlockRoutes.GET("/", controllers.BlockIndex)
		BlockRoutes.GET("/show", controllers.BlockShow)
		BlockRoutes.POST("/", controllers.BlockCreate)
		BlockRoutes.PUT("/:id", controllers.BlockUpdate)
		//BlockRoutes.DELETE("/:id", controllers.BlockDelete)
	}
}
