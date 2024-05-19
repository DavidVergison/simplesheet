package discordlib

import "encoding/json"

type DiscordResponseDto struct {
	Type int                    `json:"type"` // 4 CHANNEL_MESSAGE_WITH_SOURCE
	Data DiscordDataResponseDto `json:"data"`
}
type DiscordDataResponseDto struct {
	Content string `json:"content"`
}

func GetPingResponse() string {
	json, _ := json.Marshal(
		struct {
			Type int `json:"type"`
		}{
			Type: 1,
		},
	)
	return string(json)
}