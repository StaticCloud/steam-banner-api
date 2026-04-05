package main

import (
	"os"
	apiHandlers "steam-banner-api/handlers/api"
	"steam-banner-api/router"
)

func main() {
	apiHandler := apiHandlers.InitSteamApiHandler(os.Getenv("KEY"))

	router := router.CreateRouter(apiHandler)

	router.Run()

}
