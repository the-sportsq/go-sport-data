package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Odds struct {
	Away string `json:"away,omitempty" bson:"away,omitempty"`
	Draw string `json:"draw,omitempty" bson:"draw,omitempty"`
	Home string `json:"home,omitempty" bson:"home,omitempty"`
}

type Bookmaker struct {
	BookmakerID int        `json:"bookmaker_id,omitempty" bson:"bookmaker_id,omitempty"`
	Name        string     `json:"bookmaker_name,omitempty" bson:"bookmaker_name,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty" bson:"last_updated,omitempty"`
	Odds        *Odds      `json:"odds_data,omitempty" bson:"odds_data,omitempty"`
}

// Get list of bookmakers by match_id
func (c *Client) GetOdds(matchID int, oddsType ...string) ([]*Bookmaker, error) {
	// Have to use a janky map here because api returns JSON with commas in the keys
	type response struct {
		Data map[string]struct {
			Bookmakers []*Bookmaker `json:"bookmakers,omitempty"`
		} `json:"data,omitempty"`
	}

	query := Query{}
	if len(oddsType) > 0 {
		query["type"] = oddsType[0]
	} else {
		query["type"] = "prematch"
	}

	path := getPath(fmt.Sprintf("/soccer/odds/%d", matchID), query)

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

	return apiResponse.Data["1X2, Full Time Result"].Bookmakers, nil
}
