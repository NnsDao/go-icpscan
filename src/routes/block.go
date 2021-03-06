package routes

import (
	"icpscan/src/controllers"
)

func addBlockRoutes() {
	BlockRoutes := Router.Group("/api/block")
	{
		//BlockRoutes.GET("/", controllers.BlockIndex)
		BlockRoutes.GET("/show", controllers.BlockShow)
		BlockRoutes.GET("/test", controllers.BlockShowPpp)
		BlockRoutes.GET("/newList", controllers.BlockNew)
		//BlockRoutes.PUT("/:id", controllers.BlockUpdate)
		BlockRoutes.GET("/detail", controllers.BlockDetail)
		BlockRoutes.GET("/search", controllers.SearchDetail)
		BlockRoutes.GET("/searchDetail", controllers.GetAccountDealDetail)
		BlockRoutes.GET("/accountBalanceCurve", controllers.GetAccountBalanceCurve)
		BlockRoutes.GET("/accountDAU", controllers.GetDAUAccount)
		BlockRoutes.GET("/transationCount", controllers.GetDayTransationCount)
		BlockRoutes.GET("/blockCount", controllers.GetDayBlockCount)
		BlockRoutes.GET("/transationAmount", controllers.GetDayTransationAmount)
		BlockRoutes.GET("/destroyAmount", controllers.GetDayDestroyAmount)
		BlockRoutes.GET("/mintAmount", controllers.GetDayMintAmount)

	}
}
