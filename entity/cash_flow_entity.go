package entity

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CashFlowEntity struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Amount   float64            `json:"amount" bson:"amount"`
	Category string             `json:"category" bson:"category"`
	Desc     string             `json:"desc" bson:"desc"`
	Remark   string             `json:"remark" bson:"remark"`
}

func (entity CashFlowEntity) IsEmpty() bool {
	return reflect.DeepEqual(entity, DayFlowEntity{})
}