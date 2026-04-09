package utils

import "fmt"

func GetHeaderUrl(gameId int) string {
	return fmt.Sprintf("https://cdn.cloudflare.steamstatic.com/steam/apps/%d/header.jpg", gameId)
}

func GetBoxArtUrl(gameId int) string {
	return fmt.Sprintf("https://cdn.cloudflare.steamstatic.com/steam/apps/%d/library_600x900.jpg", gameId)
}
