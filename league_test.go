package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetLeagues(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	leagues, err := client.GetLeagues()

	if err != nil {
		t.Error(err)
		return
	}

	if leagues == nil {
		t.Error(errors.New("Failed to fetch list of leagues from API"))
		return
	}

	league := leagues[0]

	if league.Name == "" {
		t.Error(errors.New("Failed to properly unmarshal league data"))
	}
}

func TestGetLeaguesByCountry(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	countryID := 42 // England
	leagues, err := client.GetLeagues(countryID)

	if err != nil {
		t.Error(err)
		return
	}

	if leagues == nil {
		t.Error(errors.New("Failed to fetch list of leagues from API"))
		return
	}

	league := leagues[0]

	if league.Name == "" {
		t.Error(errors.New("Failed to properly unmarshal league data"))
	}
}

func TestGetLeague(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	league, err := client.GetLeague(314)

	if err != nil {
		t.Error(err)
		return
	}

	if league == nil {
		t.Error(errors.New("Failed to fetch league from API"))
	}
}
