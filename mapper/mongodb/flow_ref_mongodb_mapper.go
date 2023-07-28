package mongodb

import (
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlowRefMongoDbMapper struct{}

var flowRefMongoDbMapper FlowRefMongoDbMapper

func (FlowRefMongoDbMapper) GetFlowRefByCashFlowId(objectId primitive.ObjectID) entity.FlowRefEntity {

	filter := bson.D{
		primitive.E{Key: "cash_flow_id", Value: objectId},
	}

	// 打开 flowRef 的数据表连线
	util.OpenMongoDbConnection("flowRef")
	return convertBsonM2FlowRefEntity(util.GetOneInMongoDb(filter))
}

func (FlowRefMongoDbMapper) GetFlowRefByDayFlowId(objectId primitive.ObjectID) []entity.FlowRefEntity {

	filter := bson.D{
		primitive.E{Key: "day_flow_id", Value: objectId},
	}

	// 打开 flowRef 的数据表连线
	util.OpenMongoDbConnection("flowRef")
	flowRefBsonArray := util.GetManyInMongoDb(filter)

	var flowRefEntityArray []entity.FlowRefEntity
	for _, flowRef := range flowRefBsonArray {
		flowRefEntityArray = append(flowRefEntityArray, convertBsonM2FlowRefEntity(flowRef))
	}
	return flowRefEntityArray
}

func (FlowRefMongoDbMapper) InsertFlowRefByEntity(entity entity.FlowRefEntity) primitive.ObjectID {

	util.OpenMongoDbConnection("flowRef")
	return util.InsertOneInMongoDb(convertFlowRefEntity2BsonD(entity))
}

func convertFlowRefEntity2BsonD(entity entity.FlowRefEntity) bson.D {

	// 为空时自动生成新Id
	if entity.Id == primitive.NilObjectID {
		entity.Id = primitive.NewObjectID()
	}

	return bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
		primitive.E{Key: "day_flow_id", Value: entity.DayFlowId},
		primitive.E{Key: "cash_flow_id", Value: entity.CashFlowId},
	}
}

func convertBsonM2FlowRefEntity(bsonM bson.M) entity.FlowRefEntity {
	var newEntity entity.FlowRefEntity
	bsonBytes, _ := bson.Marshal(bsonM)
	err := bson.Unmarshal(bsonBytes, &newEntity)
	if err != nil {
		panic(err)
	}
	return newEntity
}
