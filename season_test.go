package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetSeasons(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	seasons, err := client.GetSeasons(314)

	if err != nil {
		t.Error(err)
		return
	}

	if seasons == nil {
		t.Error(errors.New("Failed to fetch list of seasons from API"))
		return
	}
}

func TestGetSeason(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	season, err := client.GetSeason(496)

	if err != nil {
		t.Error(err)
		return
	}

	if season == nil {
		t.Error(errors.New("Failed to fetch season from API"))
	}
}
