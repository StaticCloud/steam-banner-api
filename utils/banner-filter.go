package utils

import "fmt"

func GetHeaderUrl(gameId int) string {
	return fmt.Sprintf("https://cdn.cloudflare.steamstatic.com/steam/apps/%d/header.jpg", gameId)
}
