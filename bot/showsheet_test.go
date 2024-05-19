package main

import (
	"os"
	"testing"

	"github.com/DavidVergison/simplesheet/bot/kankalib"
)

func TestCharacterAttributes(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		kankaClient := kankalib.KankaClient{
			Token: os.Getenv("SHEET_DISCORD_TOKEN"),
			Campain: "217781",
		}
		res, err := ShowSheet{
			KankaClient: kankaClient,
		}.execute("5466571")
		if err != nil {
			t.Fatal(err.Error())
		}
		println(res)

		t.Fatal("boom")
	})

}

func TestCharacterRoll(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		kankaClient := kankalib.KankaClient{
			Token: os.Getenv("SHEET_DISCORD_TOKEN"),
			Campain: "217781",
		}
		res, err := MakeRoll{
			KankaClient: kankaClient,
		}.execute("5466571", "Perception")
		if err != nil {
			t.Fatal(err.Error())
		}
		println(res)

		t.Fatal("boom")
	})

}