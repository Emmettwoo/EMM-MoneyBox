package entity

import (
	"github.com/emmettwoo/EMM-MoneyBox/util"
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
	return reflect.DeepEqual(entity, CashFlowEntity{})
}

func (entity CashFlowEntity) ToString() string {

	return "[ " +
		"Id: " + entity.Id.Hex() +
		", Category: " + entity.Category +
		", Amount: " + strconv.FormatFloat(entity.Amount, 'f', 2, 64) +
		", Desc: " + entity.Desc +
		// ", Remark: " + entity.Remark +
		" ]"
}

func (entity CashFlowEntity) Build(fieldMap map[string]string) CashFlowEntity {
	for key, value := range fieldMap {
		switch key {
		case "Id":
			objectId, err := primitive.ObjectIDFromHex(value)
			entity.Id = objectId
			if err != nil {
				util.Logger.Warn("build cash_flow failed with err: " + err.Error())
			}
		case "Amount":
			amount, err := strconv.ParseFloat(value, 64)
			entity.Amount = amount
			if err != nil {
				util.Logger.Warn("build cash_flow failed with err: " + err.Error())
			}
		case "Category":
			entity.Category = value
		case "Desc":
			entity.Desc = value
		case "Remark":
			entity.Remark = value
		}
	}
	return entity
}
