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

func (h *SteamApiHandler) SteamIDSearch(c *gin.Context) {
	filter := func(gameId int) string {
		switch c.Query("filter") {
		case "box-art":
			return utils.GetBoxArtUrl(gameId)
		case "banner":
			return utils.GetHeaderUrl(gameId)
		default:
			return utils.GetHeaderUrl(gameId)
		}
	}

	steamId := c.Param("sid")

	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", h.Token, steamId)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result structs.OwnedGameBannerRes

	json.Unmarshal(body, &result)

	var games []structs.GameInfoRes = result.Response.Games

	gameIdChan := make(chan []string)

	gameIds := []string{}

	// Segment Size =/= # of Segments
	segmentSize := int(math.Ceil(float64(len(games)) / 4))

	for i := range 4 {
		go func() {
			fmt.Printf("Obtaining AppIDs [%d/4]\n", i+1)

			ids := []string{}

			// Use the segment size to calculate the upper range
			upperRange := (i * segmentSize) + segmentSize

			for j := i * segmentSize; j < min(upperRange, len(games)); j++ {
				ids = append(ids, filter(games[j].AppID))
			}

			gameIdChan <- ids
		}()
		gameIds = append(gameIds, <-gameIdChan...)
	}

	c.JSON(http.StatusOK, gameIds)
}

func (h *SteamApiHandler) GameIdSearch(c *gin.Context) {
	var body structs.GameIDSearchBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(body)

	filter := func(gameId int) string {
		switch c.Query("filter") {
		case "box-art":
			return utils.GetBoxArtUrl(gameId)
		case "banner":
			return utils.GetHeaderUrl(gameId)
		default:
			return utils.GetHeaderUrl(gameId)
		}
	}

	gameIdChan := make(chan []string)

	gameIds := []string{}

	// Segment Size =/= # of Segments
	segmentSize := int(math.Ceil(float64(len(body.Games)) / 4))

	for i := range 4 {
		go func() {
			fmt.Printf("Obtaining AppIDs [%d/4]\n", i+1)

			ids := []string{}

			// Use the segment size to calculate the upper range
			upperRange := (i * segmentSize) + segmentSize

			for j := i * segmentSize; j < min(upperRange, len(body.Games)); j++ {
				ids = append(ids, filter(body.Games[j]))
			}

			gameIdChan <- ids
		}()
		gameIds = append(gameIds, <-gameIdChan...)
	}

	c.JSON(http.StatusOK, gameIds)
}
