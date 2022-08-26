package mongodb

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CashFlowMongoDbMapper struct{}

var cashFlowMongoDbMapper CashFlowMongoDbMapper

func (CashFlowMongoDbMapper) GetCashFlowByObjectId(objectId primitive.ObjectID) entity.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	// 打开cashFlow的数据表连线
	util.OpenMongoDbConnection("cashFlow")

	return convertBsonM2CashFlowEntity(util.GetOneInMongoDb(filter))
}

func (CashFlowMongoDbMapper) GetCashFlowsByObjectIdArray(objectIdArray []primitive.ObjectID) []entity.CashFlowEntity {

	var entityArray []entity.CashFlowEntity

	// fixme: 未考慮空入參時的情況，(BadValue) $in needs an array
	filter := bson.D{
		primitive.E{Key: "_id", Value: bson.M{"$in": objectIdArray}},
	}

	// 打开cashFlow的数据表连线
	util.OpenMongoDbConnection("cashFlow")

	// 获取查询结果并转入结构对象
	queryResultArray := util.GetManyInMongoDb(filter)
	for _, queryResult := range queryResultArray {
		entityArray = append(entityArray, convertBsonM2CashFlowEntity(queryResult))
	}

	return entityArray
}

func (CashFlowMongoDbMapper) GetCashFlowsByExactDesc(desc string) []entity.CashFlowEntity {

	var entityArray []entity.CashFlowEntity

	filter := bson.D{
		primitive.E{Key: "desc", Value: desc},
	}

	// 打开cashFlow的数据表连线
	util.OpenMongoDbConnection("cashFlow")

	// 获取查询结果并转入结构对象
	queryResultArray := util.GetManyInMongoDb(filter)
	for _, queryResult := range queryResultArray {
		entityArray = append(entityArray, convertBsonM2CashFlowEntity(queryResult))
	}

	return entityArray
}

func (CashFlowMongoDbMapper) GetCashFlowsByFuzzyDesc(desc string) []entity.CashFlowEntity {

	var entityArray []entity.CashFlowEntity

	// Options i for disable case sensitive.
	filter := bson.D{
		primitive.E{Key: "desc", Value: primitive.Regex{
			Pattern: desc,
			Options: "i",
		}},
	}

	// 打开cashFlow的数据表连线
	util.OpenMongoDbConnection("cashFlow")

	// 获取查询结果并转入结构对象
	queryResultArray := util.GetManyInMongoDb(filter)
	for _, queryResult := range queryResultArray {
		entityArray = append(entityArray, convertBsonM2CashFlowEntity(queryResult))
	}

	return entityArray
}

func (CashFlowMongoDbMapper) InsertCashFlowByEntity(entity entity.CashFlowEntity, date time.Time) primitive.ObjectID {

	targetDay := date
	if (targetDay == time.Time{}) {
		targetDay = time.Now()
	}

	util.OpenMongoDbConnection("cashFlow")
	newCashFlowId := util.InsertOneInMongoDb(convertCashFlowEntity2BsonD(entity))

	// 判断有无dayFlow，无则创建，然後更新cashFlows
	dayFlowEntity := dayFlowMongoDbMapper.GetDayFlowByDate(targetDay)
	if dayFlowEntity.IsEmpty() {
		dayFlowEntity = dayFlowMongoDbMapper.GetDayFlowByObjectId(
			dayFlowMongoDbMapper.InsertDayFlowByDate(targetDay))
	}
	dayFlowEntity.CashFlows = append(dayFlowEntity.CashFlows, newCashFlowId)
	dayFlowMongoDbMapper.UpdateDayFlowByEntity(dayFlowEntity)

	return newCashFlowId
}

func (CashFlowMongoDbMapper) UpdateCashFlowByEntity(entity entity.CashFlowEntity) bool {

	if entity.Id == primitive.NilObjectID {
		panic("CashFlow's id can not be nil.")
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
	}

	util.OpenMongoDbConnection("cashFlow")
	return util.UpdateManyInMongoDb(filter, convertCashFlowEntity2BsonD(entity)) == 1
}

func (CashFlowMongoDbMapper) DeleteCashFlowByObjectId(objectId primitive.ObjectID) entity.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	entity := cashFlowMongoDbMapper.GetCashFlowByObjectId(objectId)
	if entity.IsEmpty() {
		panic("CashFlow does not exist!")
	} else {
		util.OpenMongoDbConnection("cashFlow")
		util.DeleteManyInMongoDb(filter)
		return entity
	}
}

func convertCashFlowEntity2BsonD(entity entity.CashFlowEntity) bson.D {

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

func convertBsonM2CashFlowEntity(bsonM bson.M) entity.CashFlowEntity {
	var entity entity.CashFlowEntity
	bsonBytes, _ := bson.Marshal(bsonM)
	bson.Unmarshal(bsonBytes, &entity)
	return entity
}
