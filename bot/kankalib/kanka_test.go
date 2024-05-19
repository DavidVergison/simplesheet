package kankalib

import (
	"os"
	"testing"
)

func TestCharacterAttributes(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		res, err := KankaClient{
			Token: os.Getenv("SHEET_DISCORD_TOKEN"),
			Campain: "217781",
		}.CharacterAttributes("5466571")

		if err != nil {
			t.Fatal(err.Error())
		}
		t.Fatal(res)
	})

}