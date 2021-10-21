package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
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

var MatchInProgressStatuses []int = []int{
	MATCH_INPLAY,
	MATCH_HALF_TIME,
	MATCH_EXTRA_TIME,
	MATCH_PENALTIES,
	MATCH_BREAK_TIME,
	MATCH_AWARDING,
}

var MatchEndedStatuses []int = []int{
	MATCH_ENDED,
	MATCH_AFTER_PENALTIES,
	MATCH_AFTER_EXTRA_TIME,
}

type Stats struct {
	HomeScore int    `json:"home_score,omitempty" bson:"home_score,omitempty"`
	AwayScore int    `json:"away_score,omitempty" bson:"away_score,omitempty"`
	HTScore   string `json:"ht_score,omitempty" bson:"ht_score,omitempty"`
	FTScore   string `json:"ft_score,omitempty" bson:"ft_score,omitempty"`
	ETScore   string `json:"et_score,omitempty" bson:"et_score,omitempty"`
	PSScore   string `json:"ps_score,omitempty" bson:"ps_score,omitempty"`
}

type Match struct {
	MatchID    int    `json:"match_id,omitempty" bson:"match_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty" bson:"status_code,omitempty"`
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
	MatchStart string `json:"match_start,omitempty" bson:"match_start,omitempty"`
	Timestamp  uint64 `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Minute     int    `json:"minute,omitempty" bson:"minute,omitempty"`
	LeagueID   int    `json:"league_id,omitempty" bson:"league_id,omitempty"`
	SeasonID   int    `json:"season_id,omitempty" bson:"season_id,omitempty"`
	HomeTeam   *Team  `json:"home_team,omitempty" bson:"home_team,omitempty"`
	AwayTeam   *Team  `json:"away_team,omitempty" bson:"away_team,omitempty"`
	Stats      *Stats `json:"stats,omitempty" bson:"stats,omitempty"`
	Venue      *Venue `json:"venue,omitempty" bson:"venue,omitempty"`

	// SportsQ Properties
	MatchEnd     string `json:"match_end,omitempty" bson:"match_end,omitempty"`
	RewardsGiven bool   `json:"rewards_given,omitempty" bson:"rewards_given,omitempty"`
}

// Get list of matches by season_id
func (c *Client) GetMatches(seasonID int) ([]*Match, error) {
	type response struct {
		Matches []*Match `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/matches/?season_id=%d", seasonID)

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

func getDateString(d time.Time) string {
	year, month, day := d.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

// Fetch list of matches by season id and date range
func (c *Client) GetMatchesByDateRange(seasonID int, dateFrom time.Time, dateTo time.Time) ([]*Match, error) {
	dateFromStr := getDateString(dateFrom)
	dateToStr := getDateString(dateTo)

	type response struct {
		Matches []*Match `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/matches/?season_id=%d&date_from=%s&date_to=%s", seasonID, dateFromStr, dateToStr)

	resp, err := c.MakeRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 403 {
			// API returns 403 when there's no results for some reason
			return make([]*Match, 0), nil
		}

		return nil, errors.New(fmt.Sprintf("Received bad status code from API (%v)", resp.StatusCode))
	}

	var apiResponse *response

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Matches, nil
}

// Helper function for fetching list of matches by date range
// From yesterday to tomorrow
func (c *Client) GetMatchesForToday(seasonID int) ([]*Match, error) {
	dateFrom := time.Now().AddDate(0, 0, -1)
	dateTo := time.Now().AddDate(0, 0, 1)
	return c.GetMatchesByDateRange(seasonID, dateFrom, dateTo)
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
