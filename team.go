package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Country struct {
	CountryID   int    `json:"country_id,omitempty" bson:"country_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty" bson:"country_code,omitempty"`
	Continent   string `json:"continent,omitempty" bson:"continent,omitempty"`
}

type Team struct {
	TeamID    int      `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Name      string   `json:"name,omitempty" bson:"name,omitempty"`
	ShortCode string   `json:"short_code,omitempty" bson:"short_code,omitempty"`
	Country   *Country `json:"country,omitempty" bson:"country,omitempty"`
	Logo      string   `json:"logo,omitempty" bson:"logo,omitempty"`
}

type Standings struct {
	SeasonID  int        `json:"season_id,omitempty" bson:"season_id,omitempty"`
	LeagueID  int        `json:"league_id,omitempty" bson:"league_id,omitempty"`
	hasGroups bool       `json:"has_groups,omitempty" bson:"has_groups,omitempty"`
	Standings []Standing `json:"standings,omitempty" bson:"standings,omitempty"`
}

type Standing struct {
	TeamID   int    `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Position int    `json:"position,omitempty" bson:"position,omitempty"`
	Points   int    `json:"points,omitempty" bson:"points,omitempty"`
	Status   string `json:"status,omitempty" bson:"status,omitempty"`
	Results  string `json:"result,omitempty" bson:"result,omitempty"`

	Overall StandingStats `json:"overall,omitempty" bson:"overall,omitempty"`
	Home    StandingStats `json:"home,omitempty" bson:"home,omitempty"`
	Away    StandingStats `json:"away,omitempty" bson:"away,omitempty"`
}

type StandingStats struct {
	GamesPlayed  int `json:"games_played,omitempty" bson:"games_played,omitempty"`
	Won          int `json:"won,omitempty" bson:"won,omitempty"`
	Draw         int `json:"draw,omitempty" bson:"draw,omitempty"`
	Lost         int `json:"lost,omitempty" bson:"lost,omitempty"`
	GoalsDiff    int `json:"goals_diff,omitempty" bson:"goals_diff,omitempty"`
	GoalsScored  int `json:"goals_scored,omitempty" bson:"goals_scored,omitempty"`
	GoalsAgainst int `json:"goals_against,omitempty" bson:"goals_against,omitempty"`
}

// Get list of teams by country_id
func (c *Client) GetTeams(countryID int) ([]*Team, error) {
	type response struct {
		Teams []*Team `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/teams/?country_id=%d", countryID)

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

	return apiResponse.Teams, nil
}

// Get individual team by team_id
func (c *Client) GetTeam(id int) (*Team, error) {
	type response struct {
		Team *Team `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/teams/%d", id)

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

	return apiResponse.Team, nil
}

// Get full league standings by season ID
func (c *Client) GetStandings(seasonID int) (*Standings, error) {
	type response struct {
		Data *Standings `json:"data"`
	}

	path := fmt.Sprintf("/soccer/standings?season_id=%v", seasonID)

	resp, err := c.MakeRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var apiResponse *response

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}
