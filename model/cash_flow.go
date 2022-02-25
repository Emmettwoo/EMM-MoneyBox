package model

import (
	"reflect"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
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

func GetCashFlowByObjectId(objectId primitive.ObjectID) CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	// 打开cashFlow的数据表连线
	util.OpenConnection("cashFlow")

	return convertBsonM2CashFlowEntity(util.GetOne(filter))
}

func GetCashFlowsByObjectIdArray(objectIdArray []primitive.ObjectID) []CashFlowEntity {

	var entityArray []CashFlowEntity

	// fixme: 未考慮空入參時的情況，(BadValue) $in needs an array
	filter := bson.D{
		primitive.E{Key: "_id", Value: bson.M{"$in": objectIdArray}},
	}

	// 打开cashFlow的数据表连线
	util.OpenConnection("cashFlow")

	// 获取查询结果并转入结构对象
	queryResultArray := util.GetMany(filter)
	for _, queryResult := range queryResultArray {
		entityArray = append(entityArray, convertBsonM2CashFlowEntity(queryResult))
	}

	return entityArray
}

func InsertCashFlowByEntity(entity CashFlowEntity, date time.Time) primitive.ObjectID {

	targetDay := date
	if (targetDay == time.Time{}) {
		targetDay = time.Now()
	}

	util.OpenConnection("cashFlow")
	newCashFlowId := util.InsertOne(convertCashFlowEntity2BsonD(entity))

	// 判断有无dayFlow，无则创建，然後更新cashFlows
	dayFlowEntity := GetDayFlowByDate(targetDay)
	if dayFlowEntity.IsEmpty() {
		dayFlowEntity = GetDayFlowByObjectId(InsertDayFlowByDate(targetDay))
	}
	dayFlowEntity.CashFlows = append(dayFlowEntity.CashFlows, newCashFlowId)
	UpdateDayFlowByEntity(dayFlowEntity)

	return newCashFlowId
}

func UpdateCashFlowByEntity(entity CashFlowEntity) bool {

	if entity.Id == primitive.NilObjectID {
		panic("CashFlow's id can not be nil.")
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
	}

	util.OpenConnection("cashFlow")
	return util.UpdateMany(filter, convertCashFlowEntity2BsonD(entity)) == 1
}

func DeleteCashFlowByObjectId(objectId primitive.ObjectID) CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	entity := GetCashFlowByObjectId(objectId)
	if entity.IsEmpty() {
		panic("CashFlow does not exist!")
	} else {
		util.OpenConnection("cashFlow")
		util.DeleteMany(filter)
		return entity
	}
}

func convertCashFlowEntity2BsonD(entity CashFlowEntity) bson.D {

	// 为空时自动生成新Id
	if entity.Id == primitive.NilObjectID {
		entity.Id = primitive.NewObjectID()
	}

	return bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
		primitive.E{Key: "amount", Value: entity.Amount},
		primitive.E{Key: "category", Value: entity.Category},
		primitive.E{Key: "desc", Value: entity.Desc},
		primitive.E{Key: "remark", Value: entity.Remark},
	}
}

func convertBsonM2CashFlowEntity(bsonM bson.M) CashFlowEntity {
	var entity CashFlowEntity
	bsonBytes, _ := bson.Marshal(bsonM)
	bson.Unmarshal(bsonBytes, &entity)
	return entity
}
