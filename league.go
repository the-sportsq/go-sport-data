package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type League struct {
	LeagueID  int    `json:"league_id,omitempty" bson:"league_id,omitempty"`
	CountryID int    `json:"country_id,omitempty" bson:"country_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Active    bool   `json:"active,omitempty" bson:"active,omitempty"`
}

// Get list of leagues
func (c *Client) GetLeagues() ([]*League, error) {
	type response struct {
		Leagues []*League `json:"data,omitempty"`
	}

	resp, err := c.MakeRequest("GET", "/soccer/leagues", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Received bad status code from API")
	}

	var apiResponse *response

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Leagues, nil
}

// Get a single league by league_id
func (c *Client) GetLeague(id int) (*League, error) {
	type response struct {
		League *League `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/leagues/%d", id)
	resp, err := c.MakeRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Received bad status code from API")
	}

	var apiResponse *response

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.League, nil
}
