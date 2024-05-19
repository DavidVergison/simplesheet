package main

import (
	"github.com/DavidVergison/simplesheet/bot/kankalib"
	"github.com/DavidVergison/simplesheet/bot/configlib"
)

type ShowSheet struct {
	KankaClient kankalib.KankaClient
	ServerId string
	ChannelId string
}

func (sheet ShowSheet) execute(playerId string) (string, error) {
	entityId, err := configlib.GetKankaId(playerId, sheet.ServerId, sheet.ChannelId)
	if err != nil {
		return "", err
	}

	// load kanka data into player sheet
	character, err := InitCharacter(entityId, sheet.KankaClient)
	if err != nil {
		return "", err
	}

	// return formated player sheet
	return character.ToString(), nil
}
