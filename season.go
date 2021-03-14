package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Season struct {
	SeasonId  int    `json:"season_id,omitempty" bson:"season_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	IsCurrent int8   `json:"is_current,omitempty" bson:"is_current,omitempty"`
	CountryId int    `json:"country_id,omitempty" bson:"country_id,omitempty"`
	LeagueId  int    `json:"league_id,omitempty" bson:"league_id,omitempty"`
	StartDate string `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty" bson:"end_date,omitempty"`
}

// Get list of seasons by league_id
func (c *Client) GetSeasons(leagueId int) ([]*Season, error) {
	type response struct {
		Seasons []*Season `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/seasons/?league_id=%d", leagueId)

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

	return apiResponse.Seasons, nil
}

// Get individual season by season_id
func (c *Client) GetSeason(id int) (*Season, error) {
	type response struct {
		Season *Season `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/seasons/%d", id)

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

	return apiResponse.Season, nil
}
