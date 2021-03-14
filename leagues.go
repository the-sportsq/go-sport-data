package gsd

import (
	"encoding/json"
)

type League struct {
	LeagueId  int    `json:"league_id,omitempty" bson:"league_id,omitempty"`
	CountryId int    `json:"country_id,omitempty" bson:"country_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Active    bool   `json:"active,omitempty" bson:"active,omitempty"`
}

// Get list of leagues
func (c *Client) GetLeagues() ([]*League, error) {
	req, err := c.NewRequest("GET", "/soccer/leagues", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var leagues []*League

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&leagues); err != nil {
		return nil, err
	}

	return leagues, nil
}
