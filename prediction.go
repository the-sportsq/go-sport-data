package gsd

const (
	PREDICTION_CHOICE_HOME = 0
	PREDICTION_CHOICE_AWAY = 1
	PREDICTION_CHOICE_DRAW = 2
)

type Prediction struct {
	MatchId  int `json:"match_id,omitempty" bson:"match_id,omitempty"`
	LeagueId int `json:"league_id,omitempty" bson:"league_id,omitempty"`
	Choice   int `json:"choice,omitempty" bson:"choice,omitempty"`
}
