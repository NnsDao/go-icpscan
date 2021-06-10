package routes

import (
	"github.com/MatheusMeloAntiquera/api-go/src/controllers"
)

func addBlockRoutes() {
	BlockRoutes := Router.Group("/api/block")
	{
		BlockRoutes.GET("/", controllers.BlockIndex)
		BlockRoutes.GET("/show", controllers.BlockShow)
		BlockRoutes.GET("/test", controllers.BlockShowPpp)
		BlockRoutes.GET("/newList", controllers.BlockNew)
		BlockRoutes.PUT("/:id", controllers.BlockUpdate)
		//BlockRoutes.DELETE("/:id", controllers.BlockDelete)
	}
}
