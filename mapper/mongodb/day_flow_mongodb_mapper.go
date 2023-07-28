package mongodb

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DayFlowMongoDbMapper struct{}

var dayFlowMongoDbMapper DayFlowMongoDbMapper

func (DayFlowMongoDbMapper) GetDayFlowByObjectId(objectId primitive.ObjectID) entity.DayFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	// 打开dayFlow的数据表连线
	util.OpenMongoDbConnection("dayFlow")
	return convertBsonM2DayFlowEntity(util.GetOneInMongoDb(filter))
}

func (DayFlowMongoDbMapper) GetDayFlowByDate(date time.Time) entity.DayFlowEntity {

	// Golang有趣的日期转换机制，不过转出来是String，数据库是int，所以转一下类型。
	monthInInt := util.ToInteger(date.Format("01"))

	// 查询条件为指定的年月日，设计上一天只会对应一笔dayFlow数据。
	filter := bson.D{
		primitive.E{Key: "day", Value: date.Day()},
		primitive.E{Key: "month", Value: monthInInt},
		primitive.E{Key: "year", Value: date.Year()},
	}

	util.OpenMongoDbConnection("dayFlow")
	return convertBsonM2DayFlowEntity(util.GetOneInMongoDb(filter))
}

func (DayFlowMongoDbMapper) InsertDayFlowByEntity(entity entity.DayFlowEntity) primitive.ObjectID {
	util.OpenMongoDbConnection("dayFlow")
	return util.InsertOneInMongoDb(convertDayFlowEntity2BsonD(entity))
}

func (DayFlowMongoDbMapper) InsertDayFlowByDate(date time.Time) primitive.ObjectID {

	monthInInt := util.ToInteger(date.Format("01"))
	newEntity := entity.DayFlowEntity{
		Day:   date.Day(),
		Month: monthInInt,
		Year:  date.Year(),
	}

	util.OpenMongoDbConnection("dayFlow")
	return util.InsertOneInMongoDb(convertDayFlowEntity2BsonD(newEntity))
}

func (DayFlowMongoDbMapper) UpdateDayFlowByEntity(entity entity.DayFlowEntity) bool {

	if entity.Id == primitive.NilObjectID {
		panic("DayFlow's id can not be nil.")
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
	}

	util.OpenMongoDbConnection("dayFlow")
	return util.UpdateManyInMongoDb(filter, convertDayFlowEntity2BsonD(entity)) == 1
}

func (DayFlowMongoDbMapper) DeleteDayFlowByObjectId(objectId primitive.ObjectID) entity.DayFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	//todo: 還需刪除 flow_ref, cash_flow --20221202
	targetEntity := dayFlowMongoDbMapper.GetDayFlowByObjectId(objectId)
	if targetEntity.IsEmpty() {
		panic("DayFlow does not exist!")
	} else {
		util.OpenMongoDbConnection("dayFlow")
		util.DeleteManyInMongoDb(filter)
		return targetEntity
	}
}

func (DayFlowMongoDbMapper) DeleteDayFlowByDate(date time.Time) entity.DayFlowEntity {

	monthInInt := util.ToInteger(date.Format("01"))
	filter := bson.D{
		primitive.E{Key: "day", Value: date.Day()},
		primitive.E{Key: "month", Value: monthInInt},
		primitive.E{Key: "year", Value: date.Year()},
	}

	//todo: 還需刪除 flow_ref, cash_flow --20221202
	dayFlow := dayFlowMongoDbMapper.GetDayFlowByDate(date)
	if dayFlow.IsEmpty() {
		util.Logger.Infow("day_flow to be delete does not exist",
			"date", util.FormatDateToString(date))
	} else {
		util.OpenMongoDbConnection("dayFlow")
		util.DeleteManyInMongoDb(filter)
	}
	return dayFlow
}

func convertDayFlowEntity2BsonD(entity entity.DayFlowEntity) bson.D {

	// 为空时自动生成新Id
	if entity.Id == primitive.NilObjectID {
		entity.Id = primitive.NewObjectIDFromTimestamp(time.Now())
	}

	return bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
		primitive.E{Key: "cashFlows", Value: entity.CashFlows},
		primitive.E{Key: "day", Value: entity.Day},
		primitive.E{Key: "month", Value: entity.Month},
		primitive.E{Key: "year", Value: entity.Year},
	}
}

func convertBsonM2DayFlowEntity(bsonM bson.M) entity.DayFlowEntity {
	var entity entity.DayFlowEntity
	bsonBytes, _ := bson.Marshal(bsonM)
	err := bson.Unmarshal(bsonBytes, &entity)
	if err != nil {
		panic(err)
	}
	return entity
}
