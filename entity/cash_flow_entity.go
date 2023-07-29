package entity

import (
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"reflect"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CashFlowEntity struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	CategoryId  primitive.ObjectID `json:"category_id" bson:"category_id"`
	BelongsDate time.Time          `json:"belongs_date" bson:"belongs_date"`
	Type        string             `json:"type" bson:"type"`
	Amount      float64            `json:"amount" bson:"amount"`
	Desc        string             `json:"desc" bson:"desc"`
	Remark      string             `json:"remark" bson:"remark"`
	CreateTime  time.Time          `json:"create_time" bson:"create_time"`
	ModifyTime  time.Time          `json:"modify_time" bson:"modify_time"`
}

func (entity CashFlowEntity) IsEmpty() bool {
	return reflect.DeepEqual(entity, CashFlowEntity{})
}

func (entity CashFlowEntity) ToString() string {

	// todo: category query from cache like redis would be better.
	return "[ " +
		"Id: " + entity.Id.Hex() +
		", Category: " + mapper.GetCategoryMapper().GetCategoryByObjectId(entity.CategoryId).Name +
		", Date: " + util.FormatDateToString(entity.BelongsDate) +
		", Type: " + entity.Type +
		", Amount: " + strconv.FormatFloat(entity.Amount, 'f', 2, 64) +
		", Desc: " + entity.Desc +
		" ]"
}

func (entity CashFlowEntity) Build(fieldMap map[string]string) CashFlowEntity {
	var newEntity = entity
	for key, value := range fieldMap {
		switch key {
		case "Id":
			objectId, err := primitive.ObjectIDFromHex(value)
			if err != nil {
				util.Logger.Warnln("build cash_flow failed with err: " + err.Error())
			}
			newEntity.Id = objectId
		case "Category":
			// fixme: use name to query id, nil then insert one
			categoryEntity := mapper.GetCategoryMapper().GetCategoryByName(value)
			if categoryEntity.IsEmpty() {
				util.Logger.Warnln("could not find target category: " + value)
			}
			newEntity.CategoryId = categoryEntity.Id
		case "Date":
			newEntity.BelongsDate = util.FormatDateFromString(value)
		case "Type":
			// todo: use enum to check if value available
			newEntity.Type = value
		case "Amount":
			amount, err := strconv.ParseFloat(value, 64)
			if err != nil {
				util.Logger.Warnln("build cash_flow failed with err: " + err.Error())
			}
			newEntity.Amount = amount
		case "Desc":
			newEntity.Desc = value
		case "Remark":
			newEntity.Remark = value
		}
	}
	return newEntity
}
