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
// GetLeagues fetches a list of leagues the API is subscribed to
// You can also provide a country_id to further filter the list
func (c *Client) GetLeagues(countryID ...int) ([]*League, error) {
	query := Query{
		"subscribed": true,
	}

	if len(countryID) > 0 {
		query = Query{
			"country_id": countryID[0],
		}
	}

	path := getPath("/soccer/leagues", query)
	resp, err := c.MakeRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Received bad status code from API")
	}

	var leagues []*League

	if len(countryID) > 0 {
		// Whenever country_id is provided to API, data is a map instead of array
		// Not sure why they wanted to be weird with it
		type Response struct {
			Leagues map[int]*League `json:"data,omitempty"`
		}

		var response *Response
		decoder := json.NewDecoder(resp.Body)
		if err = decoder.Decode(&response); err != nil {
			return nil, err
		}

		for _, v := range response.Leagues {
			leagues = append(leagues, v)
		}
	} else {
		// Read data as array, as expected
		type Response struct {
			Leagues []*League `json:"data,omitempty"`
		}

		var response *Response
		decoder := json.NewDecoder(resp.Body)
		if err = decoder.Decode(&response); err != nil {
			return nil, err
		}

		leagues = response.Leagues
	}

	return leagues, nil
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
