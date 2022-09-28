package entity

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlowRefEntity struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	DayFlowId  primitive.ObjectID `bson:"day_flow_id,omitempty"`
	CashFlowId primitive.ObjectID `bson:"cash_flow_id,omitempty"`
}

func (entity FlowRefEntity) IsEmpty() bool {
	return reflect.DeepEqual(entity, FlowRefEntity{})
}
