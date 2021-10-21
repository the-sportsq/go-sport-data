package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetRounds(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	rounds, err := client.GetRounds(352)

	if err != nil {
		t.Error(err)
		return
	}

	if rounds == nil {
		t.Error(errors.New("Failed to fetch list of rounds from API"))
		return
	}
}

func TestGetRound(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	round, err := client.GetRound(6190)

	if err != nil {
		t.Error(err)
		return
	}

	if round == nil {
		t.Error(errors.New("Failed to fetch round from API"))
	}
}
