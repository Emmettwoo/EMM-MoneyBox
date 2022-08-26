package entity

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DayFlowEntity struct {
	Id        primitive.ObjectID   `bson:"_id,omitempty"`
	CashFlows []primitive.ObjectID `json:"cashFlows" bson:"cashFlows"`
	Day       int                  `json:"day" bson:"day"`
	Month     int                  `json:"month" bson:"month"`
	Year      int                  `json:"year" bson:"year"`
}

func (entity DayFlowEntity) IsEmpty() bool {
	if reflect.DeepEqual(entity, DayFlowEntity{}) {
		return true
	}
	return len(entity.CashFlows) == 0
}
