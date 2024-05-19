package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/DavidVergison/simplesheet/bot/discordlib"
	"github.com/DavidVergison/simplesheet/bot/kankalib"
	"github.com/aws/aws-lambda-go/events"
)

func readOption(options discordlib.DataOptions) map[string]string {
	res := map[string]string{}
	for _, o := range options {
		res[o.Name] = fmt.Sprintf("%v", o.Value)
	}
	return res
}



func discordSlashCommandHandler(discordRequest discordlib.DiscordRequestDto) (events.APIGatewayProxyResponse, error) {

	slashCommand := discordRequest.Data.Name
	nickname := discordRequest.Member.Nick
	playerId := discordRequest.Member.User.Id
	serverId := discordRequest.GuildId
	channelId := discordRequest.ChannelId

	text := "ok"
	switch slashCommand {
	case "me":
		return formatResponse(fmt.Sprintf("Vous êtes %s -> %s", nickname, playerId))
	case "sheet":
		token := os.Getenv("SHEET_DISCORD_TOKEN")
		if token == "" {
			log.Fatal("Error : Empty TOKEN")
		}
		kankaClient := kankalib.KankaClient{
			Token:   token,
			Campain: "217781",
		}
		resp, err := ShowSheet{
			KankaClient: kankaClient,
			ServerId: serverId,
			ChannelId: channelId,
		}.execute(playerId)
		if err != nil {
			log.Fatalf("Error : %v", err)
		}

		return formatResponse(resp)
	case "statistic":
		options := readOption(discordRequest.Data.Options)
		optionsJSON, err := json.Marshal(options)
		if err != nil {
			log.Fatalf("Erreur lors de l'encodage en JSON : %v", err)
		}
		log.Println(string(optionsJSON))
		token := os.Getenv("SHEET_DISCORD_TOKEN")
		if token == "" {
			log.Fatal("Error : Empty TOKEN")
		}
		kankaClient := kankalib.KankaClient{
			Token:   token,
			Campain: "217781",
		}
		resp, err := MakeRoll{
			KankaClient: kankaClient,
			ServerId: serverId,
			ChannelId: channelId,
		}.execute(playerId, options["statistic"])
		if err != nil {
			log.Fatalf("Erreur lors du jet : %v", err)
		}
		return formatResponse(resp)
	}
	return formatResponse(text)
}
