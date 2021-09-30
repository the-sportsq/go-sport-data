package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetTeams(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	teams, err := client.GetTeams(48)

	if err != nil {
		t.Error(err)
		return
	}

	if teams == nil {
		t.Error(errors.New("Failed to fetch list of teams from API"))
	}
}

func TestGetTeam(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	team, err := client.GetTeam(4066)

	if err != nil {
		t.Error(err)
		return
	}

	if team == nil {
		t.Error(errors.New("Failed to fetch team from API"))
	}
}

func TestGetStandings(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	data, err := client.GetStandings(352)
	if err != nil {
		t.Error(err)
		return
	}

	if data == nil {
		t.Error(errors.New("Failed to fetch standings from API"))
		return
	}

	standing := data.Standings[0]

	if standing.TeamId != 12400 {
		t.Error(errors.New("Unexpected data returned from API for standings"))
	}
}
