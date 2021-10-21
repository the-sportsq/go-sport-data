package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Venue struct {
	VenueID   int    `json:"venue_id,omitempty" bson:"venue_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Capacity  int    `json:"capacity,omitempty" bson:"capacity,omitempty"`
	City      string `json:"city,omitempty" bson:"city,omitempty"`
	CountryID int    `json:"country_id,omitempty" bson:"country_id,omitempty"`
}

// GetVenues returns a list of venues by country
func (c *Client) GetVenues(countryID int) ([]*Venue, error) {
	type Response struct {
		Venues []*Venue `json:"data,omitempty"`
	}

	path := getPath("/soccer/venues/", Query{
		"country_id": countryID,
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

	return response.Venues, nil
}

// GetVenue returns a single venue by id
func (c *Client) GetVenue(id int) (*Venue, error) {
	type Response struct {
		Venue *Venue `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/venues/%d", id)
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

	return response.Venue, nil
}
