package mapper

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
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

	var dayFlowMongoDbMapper DayFlowMapper
	dayFlowMongoDbMapper = mongodb.DayFlowMongoDbMapper{}
	return dayFlowMongoDbMapper
}
