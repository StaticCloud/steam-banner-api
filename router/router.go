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
				profile.GET("/:sid", apiHandler.GetOwnedGameBanners)
			}
		}
	}

	return router
}
