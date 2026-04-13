package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"steam-banner-api/structs"
	"steam-banner-api/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type SteamApiHandler struct {
	Token string
}

func InitSteamApiHandler(token string) *SteamApiHandler {
	return &SteamApiHandler{
		Token: token,
	}
}

func filterFactory(filter string) func(int) string {
	switch filter {
	case "box-art":
		return utils.GetBoxArtUrl
	case "banner":
		return utils.GetHeaderUrl
	default:
		return utils.GetHeaderUrl
	}
}

func processPayload(payload []int, filter func(int) string) []string {
	var banners = []string{}

	channel := make(chan []string)

	segmentSize := int(math.Ceil(float64(len(payload)) / 4))

	for i := range 4 {
		go func() {
			ids := []string{}

			// Use the segment size to calculate the upper range
			upperRange := (i * segmentSize) + segmentSize

			for j := i * segmentSize; j < min(upperRange, len(payload)); j++ {
				ids = append(ids, filter(payload[j]))
			}

			channel <- ids

			fmt.Printf("Obtaining Banners [%d/4]\n", i+1)
		}()
		banners = append(banners, <-channel...)
	}

	return banners
}

func (h *SteamApiHandler) SteamIDSearch(c *gin.Context) {
	steamId := c.Param("sid")

	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", h.Token, steamId)
	ownedGames, ownedGamesError := http.Get(url)

	if ownedGamesError != nil {
		c.JSON(http.StatusInternalServerError, ownedGamesError.Error())
		return
	}

	defer ownedGames.Body.Close()

	body, bodyError := io.ReadAll(ownedGames.Body)
	if bodyError != nil {
		c.JSON(http.StatusInternalServerError, bodyError.Error())
		return
	}

	var result structs.OwnedGameBannerRes

	json.Unmarshal(body, &result)

	gameIds := []int{}

	for _, v := range result.Response.Games {
		gameIds = append(gameIds, v.AppID)
	}

	response := processPayload(gameIds, filterFactory(c.Query("filter")))

	c.JSON(http.StatusOK, response)
}

func (h *SteamApiHandler) GameIdSearch(c *gin.Context) {
	var body structs.GameIDSearchBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := processPayload(body.Games, filterFactory(c.Query("filter")))

	c.JSON(http.StatusOK, response)
}
