package main

import (
	"github.com/DavidVergison/simplesheet/bot/configlib"
	"github.com/DavidVergison/simplesheet/bot/kankalib"
)

type MakeRoll struct {
	KankaClient kankalib.KankaClient
	ServerId string
	ChannelId string
}

func(roll MakeRoll) execute(playerId string, statistic string) (string, error){
	entityId, err := configlib.GetKankaId(playerId, roll.ServerId, roll.ChannelId)
	if err != nil {
		return "", err
	}

	// load kanka data into player sheet
	character, err := InitCharacter(entityId, roll.KankaClient)
	if err != nil {
		return "", err
	}	

	// make the roll & return the result
	return character.Roll(statistic), nil
}