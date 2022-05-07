package mapper

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DayFlowMapper interface {
	GetDayFlowByObjectId(objectId primitive.ObjectID) entity.DayFlowEntity
	GetDayFlowByDate(date time.Time) entity.DayFlowEntity
	InsertDayFlowByEntity(entity entity.DayFlowEntity) primitive.ObjectID
	UpdateDayFlowByEntity(entity entity.DayFlowEntity) bool
	DeleteDayFlowByObjectId(objectId primitive.ObjectID) entity.DayFlowEntity
	DeleteDayFlowByDate(date time.Time) entity.DayFlowEntity
}

func GetDayFlowMapper() DayFlowMapper {
	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		var dayFlowMongoDbMapper DayFlowMapper
		dayFlowMongoDbMapper = mongodb.DayFlowMongoDbMapper{}
		return dayFlowMongoDbMapper
	case "mysql":
		panic("mysql support is still under dev")
	default:
		panic("database type not supported")
	}
}
