package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetOdds(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	bookmakers, err := client.GetOdds(120423)

	if err != nil {
		t.Error(err)
		return
	}

	if bookmakers == nil {
		t.Error(errors.New("Failed to fetch odds from API"))
	}
}

func TestGetOddsInPlay(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	bookmakers, err := client.GetOdds(120423, "inplay")

	if err != nil {
		t.Error(err)
		return
	}

	if bookmakers == nil {
		t.Error(errors.New("Failed to fetch odds from API"))
	}
}
