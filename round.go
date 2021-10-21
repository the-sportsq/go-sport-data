package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Round struct {
	RoundID    int    `json:"round_id,omitempty" bson:"round_id,omitempty"`
	SeasonID   int    `json:"season_id,omitempty" bson:"season_id,omitempty"`
	LeagueID   int    `json:"league_id,omitempty" bson:"league_id,omitempty"`
	LeagueName string `json:"league_name,omitempty" bson:"league_name,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	IsCurrent  int    `json:"is_current,omitempty" bson:"is_current,omitempty"`
}

// GetRounds returns a list of rounds by season
func (c *Client) GetRounds(seasonID int) ([]*Round, error) {
	type Response struct {
		Rounds []*Round `json:"data,omitempty"`
	}

	path := getPath("/soccer/rounds/", Query{
		"season_id": seasonID,
	})

	resp, err := c.MakeRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Received bad status code from API")
	}

	var response *Response

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response.Rounds, nil
}

// GetRound returns a single round by id
func (c *Client) GetRound(id int) (*Round, error) {
	type Response struct {
		Round *Round `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/rounds/%d", id)
	resp, err := c.MakeRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Received bad status code from API")
	}

	var response *Response

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response.Round, nil
}
