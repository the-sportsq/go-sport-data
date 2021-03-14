package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	MATCH_NOT_STARTED      = 0
	MATCH_INPLAY           = 1
	MATCH_HALF_TIME        = 11
	MATCH_EXTRA_TIME       = 12
	MATCH_PENALTIES        = 13
	MATCH_BREAK_TIME       = 14
	MATCH_AWARDING         = 15
	MATCH_UPDATE_LATER     = 2
	MATCH_ENDED            = 3
	MATCH_AFTER_PENALTIES  = 31
	MATCH_AFTER_EXTRA_TIME = 32
	MATCH_POSTPONED        = 4
	MATCH_CANCELLED        = 5
	MATCH_ABANDONED        = 6
	MATCH_INTERRUPTED      = 7
	MATCH_SUSPENDED        = 8
	MATCH_AWARDED          = 9
	MATCH_DELAYED          = 10
	MATCH_TO_BE_ANNOUNCED  = 17
)

type Stats struct {
	HomeScore int    `json:"home_score,omitempty" bson:"home_score,omitempty"`
	AwayScore int    `json:"away_score,omitempty" bson:"away_score,omitempty"`
	HTScore   string `json:"ht_score,omitempty" bson:"ht_score,omitempty"`
	FTScore   string `json:"ft_score,omitempty" bson:"ft_score,omitempty"`
	ETScore   string `json:"et_score,omitempty" bson:"et_score,omitempty"`
	PSScore   string `json:"ps_score,omitempty" bson:"ps_score,omitempty"`
}

type Venue struct {
	VenueId   int    `json:"venue_id,omitempty" bson:"venue_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Capacity  int    `json:"capacity,omitempty" bson:"capacity,omitempty"`
	City      string `json:"city,omitempty" bson:"city,omitempty"`
	CountryId int    `json:"country_id,omitempty" bson:"country_id,omitempty"`
}

type Match struct {
	MatchId    int    `json:"match_id,omitempty" bson:"match_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty" bson:"status_code,omitempty"`
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
	MatchStart string `json:"match_start,omitempty" bson:"match_start,omitempty"`
	LeagueId   int    `json:"league_id,omitempty" bson:"league_id,omitempty"`
	SeasonId   int    `json:"season_id,omitempty" bson:"season_id,omitempty"`
	HomeTeam   *Team  `json:"home_team,omitempty" bson:"home_team,omitempty"`
	AwayTeam   *Team  `json:"away_team,omitempty" bson:"away_team,omitempty"`
	Stats      *Stats `json:"stats,omitempty" bson:"stats,omitempty"`
	Venue      *Venue `json:"venue,omitempty" bson:"venue,omitempty"`
}

// Get list of matches by season_id
func (c *Client) GetMatches(seasonId int) ([]*Match, error) {
	type response struct {
		Matches []*Match `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/matches/?season_id=%d", seasonId)

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

	return apiResponse.Matches, nil
}

// Get individual match by match_id
func (c *Client) GetMatch(id int) (*Match, error) {
	type response struct {
		Match *Match `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/matches/%d", id)

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

	return apiResponse.Match, nil
}
