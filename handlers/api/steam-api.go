package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

	var result map[string]any
	json.Unmarshal(body, &result)

	ctx.JSON(http.StatusOK, result)
}
