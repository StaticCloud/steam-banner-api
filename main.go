package main

import "steam-banner-api/router"

func main() {

	router := router.CreateRouter()

	router.Run()

}
