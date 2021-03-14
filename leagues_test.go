package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetLeagues(t *testing.T) {
	apiKey = os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	leagues, err := client.GetLeagues()
	if err != nil {
		t.Error(err)
		return
	}

	if leagues == nil {
		t.Error(errors.New("Failed to fetch list of leagues from API"))
	}
}
