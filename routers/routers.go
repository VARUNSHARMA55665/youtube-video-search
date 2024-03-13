package routers

import (
	"video_search/controllers"
	"video_search/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	detailsRoute := r.Group("/details")
	detailsRoute.Use(middlewares.NoAuthMiddleware())
	{
		detailsRoute.GET("/getStoredVideoDetails", controllers.GetStoredVideoDetails)
		detailsRoute.GET("/searchQueryBasedVideo", controllers.SearchQueryBasedVideo)
	}

	return r
}
