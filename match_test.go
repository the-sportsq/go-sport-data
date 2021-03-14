package gsd

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestGetMatches(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	matches, err := client.GetMatches(352)

	if err != nil {
		t.Error(err)
		return
	}

	if matches == nil {
		t.Error(errors.New("Failed to fetch list of matches from API"))
		return
	}
}

func TestGetMatch(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	match, err := client.GetMatch(136906)

	if err != nil {
		t.Error(err)
		return
	}

	if match == nil {
		t.Error(errors.New("Failed to fetch match from API"))
	}
}
