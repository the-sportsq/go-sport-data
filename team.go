package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Country struct {
	CountryId   int    `json:"country_id,omitempty" bson:"country_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty" bson:"country_code,omitempty"`
	Continent   string `json:"continent,omitempty" bson:"continent,omitempty"`
}

type Team struct {
	TeamId    int      `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Name      string   `json:"name,omitempty" bson:"name,omitempty"`
	ShortCode string   `json:"short_code,omitempty" bson:"short_code,omitempty"`
	Country   *Country `json:"country,omitempty" bson:"country,omitempty"`
}

// Get list of teams by country_id
func (c *Client) GetTeams(countryId int) ([]*Team, error) {
	type response struct {
		Teams []*Team `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/teams/?country_id=%d", countryId)

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
