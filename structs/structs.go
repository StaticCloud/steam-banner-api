package structs

type GameIDSearchBody struct {
	Games []int `json:"games"`
}

type OwnedGameBannerRes struct {
	Response struct {
		GameCount int           `json:"game_count"`
		Games     []GameInfoRes `json:"games"`
	} `json:"response"`
}

type GameInfoRes struct {
	AppID                 int `json:"appid"`
	PlaytimeDeckForever   int `json:"playtime_deck_forever"`
	PlaytimeDisconnected  int `json:"playtime_disconnected"`
	PlaytimeForever       int `json:"playtime_forever"`
	PlaytimeLinuxForever  int `json:"playtime_linux_forever"`
	PlaytimeMacForever    int `json:"playtime_mac_forever"`
	PlaytimeWindowForever int `json:"playtime_windows_forever"`
	RtimeLastPlayed       int `json:"rtime_last_played"`
}
