package configlib

import "fmt"

// guildId-channelId -> player map
type Configuration map[string]ChannelConfiguration

// discord userId -> kanka userId
type ChannelConfiguration map[string]string

func getChannelsConfiguration() Configuration {
	return Configuration{
		"923113480184279090-923113480184279095": ChannelConfiguration{
			"210128576526155777": "5466571",
		},
	}
}

func GetKankaId(discordUserId string, discordServerId string, discordChannelId string) (string, error) {
	conf := getChannelsConfiguration()
	if users, ok := conf[fmt.Sprintf("%s-%s", discordServerId, discordChannelId)]; ok {
		//do something here
		if id, ok := users[discordUserId]; ok {
			return id, nil
		}
	}
	return "", fmt.Errorf("unknown user : %s for server %s amd channel %s", discordUserId, discordServerId, discordChannelId)
}
