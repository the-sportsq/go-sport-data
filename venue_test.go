package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetVenues(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	venues, err := client.GetVenues(48)

	if err != nil {
		t.Error(err)
		return
	}

	if venues == nil {
		t.Error(errors.New("Failed to fetch list of venues from API"))
		return
	}
}

func TestGetVenue(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	venue, err := client.GetVenue(813)

	if err != nil {
		t.Error(err)
		return
	}

	if venue == nil {
		t.Error(errors.New("Failed to fetch venue from API"))
	}
}
