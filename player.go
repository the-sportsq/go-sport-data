package gsd

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Player struct {
	PlayerID int `json:"player_id,omitempty" bson:"player_id,omitempty"`

	FirstName string `json:"firstname,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"last_name,omitempty"`
	Birthday  string `json:"birthday,omitempty" bson:"birthday,omitempty"`

	Age    int `json:"age,omitempty" bson:"age,omitempty"`
	Weight int `json:"weight,omitempty" bson:"weight,omitempty"`
	Height int `json:"height,omitempty" bson:"height,omitempty"`

	Image string `json:"img,omitempty" bson:"img,omitempty"`

	Country *Country `json:"country,omitempty" bson:"country,omitempty"`
}

// GetPlayers returns a list of players by country
func (c *Client) GetPlayers(countryID int) ([]*Player, error) {
	type Response struct {
		Players []*Player `json:"data,omitempty"`
	}

	path := getPath("/soccer/players/", Query{
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

	return response.Players, nil
}

// GetPlayer returns a player by id
func (c *Client) GetPlayer(id int) (*Player, error) {
	type Response struct {
		Player *Player `json:"data,omitempty"`
	}

	path := fmt.Sprintf("/soccer/players/%d", id)

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

	return response.Player, nil
}
