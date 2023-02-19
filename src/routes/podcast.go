package routes

import "icpscan/src/controllers"

func addPodcastRoutes() {
	PodcastRoutes := Router.Group("/api/podcast")
	{
		PodcastRoutes.GET("info", controllers.GetPodcastInfo)
		PodcastRoutes.POST("record_play", controllers.RecordPlay)
	}
}
