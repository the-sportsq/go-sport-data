package gsd

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PREDICTION_CHOICE_HOME = 0
	PREDICTION_CHOICE_AWAY = 1
	PREDICTION_CHOICE_DRAW = 2
)

type Prediction struct {
	UserID   primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	MatchID  int                `json:"match_id,omitempty" bson:"match_id,omitempty"`
	LeagueID int                `json:"league_id,omitempty" bson:"league_id,omitempty"`
	Choice   int                `json:"choice" bson:"choice,omitempty"`
}
