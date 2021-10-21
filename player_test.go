package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	players, err := client.GetPlayers(48)

	if err != nil {
		t.Error(err)
		return
	}

	if players == nil {
		t.Error(errors.New("Failed to fetch list of players from API"))
		return
	}

	player := players[0]

	if player.FirstName == "" {
		t.Errorf("Failed to properly parse player data: %v", player)
	}
}

func TestGetPlayer(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	player, err := client.GetPlayer(580)

	if err != nil {
		t.Error(err)
		return
	}

	if player == nil {
		t.Error(errors.New("Failed to fetch player from API"))
		return
	}
}
