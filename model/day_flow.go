package model

import (
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
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
	return reflect.DeepEqual(entity, DayFlowEntity{})
}

func GetDayFlowByObjectId(objectId primitive.ObjectID) DayFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	// 打开dayFlow的数据表连线
	util.OpenConnection("dayFlow")
	return convertBsonM2DayFlowEntity(util.QueryOne(filter))
}

func GetDayFlowByDate(date time.Time) DayFlowEntity {

	// Golang有趣的日期转换机制，不过转出来是String，数据库是int，所以转一下类型。
	monthInInt, _ := strconv.Atoi(date.Format("01"))

	// 查询条件为指定的年月日，设计上一天只会对应一笔dayFlow数据。
	filter := bson.D{
		primitive.E{Key: "day", Value: date.Day()},
		primitive.E{Key: "month", Value: monthInInt},
		primitive.E{Key: "year", Value: date.Year()},
	}

	util.OpenConnection("dayFlow")
	return convertBsonM2DayFlowEntity(util.QueryOne(filter))
}

func InsertDayFlowByEntity(entity DayFlowEntity) primitive.ObjectID {
	util.OpenConnection("dayFlow")
	return util.InsertOne(convertDayFlowEntity2BsonD(entity))
}

func InsertDayFlowByDate(date time.Time) primitive.ObjectID {

	monthInInt, _ := strconv.Atoi(date.Format("01"))
	entity := DayFlowEntity{
		Day:   date.Day(),
		Month: monthInInt,
		Year:  date.Year(),
	}

	util.OpenConnection("dayFlow")
	return util.InsertOne(convertDayFlowEntity2BsonD(entity))
}

func UpdateDayFlowByEntity(entity DayFlowEntity) bool {

	filter := bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
	}

	updateEntity := bson.D{
		primitive.E{Key: "$set", Value: convertDayFlowEntity2BsonD(entity)},
	}

	util.OpenConnection("dayFlow")
	return util.UpdateMany(filter, updateEntity) == 1
}

func DeleteDayFlowByObjectId(objectId primitive.ObjectID) DayFlowEntity {

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	entity := GetDayFlowByObjectId(objectId)
	if entity.IsEmpty() {
		log.Fatal("DayFlow does not exist!")
	} else {
		util.OpenConnection("dayFlow")
		util.DeleteMany(filter)
		return entity
	}
	return DayFlowEntity{}
}

func DeleteDayFlowByDate(date time.Time) DayFlowEntity {

	monthInInt, _ := strconv.Atoi(date.Format("01"))
	filter := bson.D{
		primitive.E{Key: "day", Value: date.Day()},
		primitive.E{Key: "month", Value: monthInInt},
		primitive.E{Key: "year", Value: date.Year()},
	}

	entity := GetDayFlowByDate(date)
	if entity.IsEmpty() {
		log.Fatal("DayFlow does not exist!")
	} else {
		util.OpenConnection("dayFlow")
		util.DeleteMany(filter)
		return entity
	}
	return DayFlowEntity{}
}

func convertDayFlowEntity2BsonD(entity DayFlowEntity) bson.D {

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

func convertBsonM2DayFlowEntity(bsonM bson.M) DayFlowEntity {
	var entity DayFlowEntity
	bsonBytes, _ := bson.Marshal(bsonM)
	bson.Unmarshal(bsonBytes, &entity)
	return entity
}
