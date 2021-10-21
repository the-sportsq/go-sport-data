package gsd

import (
	"errors"
	"os"
	"testing"
)

func TestGetStages(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	stages, err := client.GetStages(352)

	if err != nil {
		t.Error(err)
		return
	}

	if stages == nil {
		t.Error(errors.New("Failed to fetch list of stages from API"))
		return
	}
}

func TestGetStage(t *testing.T) {
	apiKey := os.Getenv("SPORT_DATA_API_KEY")
	client := NewClient(apiKey)

	stage, err := client.GetStage(1)

	if err != nil {
		t.Error(err)
		return
	}

	if stage == nil {
		t.Error(errors.New("Failed to fetch list of stages from API"))
		return
	}
}
