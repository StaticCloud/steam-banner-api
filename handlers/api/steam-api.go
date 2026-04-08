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

func (h *SteamApiHandler) GetOwnedGameBanners(ctx *gin.Context) {
	steamId := ctx.Param("sid")

	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", h.Token, steamId)
	res, err := http.Get(url)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
				ids = append(ids, utils.GetHeaderUrl(games[j].AppID))
			}

			gameIdChan <- ids
		}()
		gameIds = append(gameIds, <-gameIdChan...)
	}

	ctx.JSON(http.StatusOK, gameIds)
}
