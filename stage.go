package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Stage struct {
	StageID int    `json:"stage_id,omitempty" bson:"stage_id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
}

// GetStages fetches a list of stages by season
func (c *Client) GetStages(seasonID int) ([]*Stage, error) {
	type Response struct {
		Stages []*Stage `json:"data,omitempty"`
	}

	path := getPath("/soccer/stages/", Query{
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

	return response.Stages, nil
}

// GetStage returns a single stage by ID
func (c *Client) GetStage(id int) (*Stage, error) {
	type Response struct {
		Stage *Stage `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/stages/%d", id)
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

	return response.Stage, nil
}
