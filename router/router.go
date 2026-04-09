package router

import (
	"fmt"
	apiHandlers "steam-banner-api/handlers/api"

	"github.com/gin-gonic/gin"
)

func CreateRouter(apiHandler *apiHandlers.SteamApiHandler) *gin.Engine {
	router := gin.Default()

	fmt.Printf("Generating routes...")

	apiGroup := router.Group("/api")
	{
		v1 := apiGroup.Group("/v1")
		{
			profile := v1.Group("/profile")
			{
				sid := profile.Group("/:sid")
				{
					sid.GET("/headers", apiHandler.GetGameHeaders)
					sid.GET("/box-art", apiHandler.GetGameBoxart)
				}
			}
		}
	}

	return router
}
