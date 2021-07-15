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
		//BlockRoutes.PUT("/:id", controllers.BlockUpdate)
		BlockRoutes.GET("/detail", controllers.BlockDetail)
		BlockRoutes.GET("/search", controllers.SearchDetail)
		BlockRoutes.GET("/searchDetail", controllers.GetAccountDealDetail)
		BlockRoutes.GET("/accountBalanceCurve", controllers.GetAccountBalanceCurve)

	}
}
