package gsd

const (
	PREDICTION_CHOICE_HOME = 0
	PREDICTION_CHOICE_AWAY = 1
	PREDICTION_CHOICE_DRAW = 2
)

type Prediction struct {
	UserId   string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	MatchId  int    `json:"match_id,omitempty" bson:"match_id,omitempty"`
	LeagueId int    `json:"league_id,omitempty" bson:"league_id,omitempty"`
	Choice   int    `json:"choice,omitempty" bson:"choice,omitempty"`
}
