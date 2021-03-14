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
