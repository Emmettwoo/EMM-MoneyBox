package entity

import (
	"reflect"
	"strconv"

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

func (entity CashFlowEntity) ToString() string {

	return "[ " +
		"Id: " + entity.Id.Hex() +
		", Category: " + entity.Category +
		", Amount: " + strconv.FormatFloat(float64(entity.Amount), 'f', 2, 32) +
		", Desc: " + entity.Desc +
		// ", Remark: " + entity.Remark +
		" ]"
}
