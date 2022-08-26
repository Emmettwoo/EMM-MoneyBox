package mapper

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CashFlowMapper interface {
	GetCashFlowByObjectId(objectId primitive.ObjectID) entity.CashFlowEntity
	GetCashFlowsByObjectIdArray(objectIdArray []primitive.ObjectID) []entity.CashFlowEntity
	GetCashFlowsByExactDesc(desc string) []entity.CashFlowEntity
	GetCashFlowsByFuzzyDesc(desc string) []entity.CashFlowEntity
	InsertCashFlowByEntity(entity entity.CashFlowEntity, date time.Time) primitive.ObjectID
	UpdateCashFlowByEntity(entity entity.CashFlowEntity) bool
	DeleteCashFlowByObjectId(objectId primitive.ObjectID) entity.CashFlowEntity
}

func GetCashFlowMapper() CashFlowMapper {

	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		var cashFlowMongoDbMapper CashFlowMapper
		cashFlowMongoDbMapper = mongodb.CashFlowMongoDbMapper{}
		return cashFlowMongoDbMapper
	case "mysql":
		panic("mysql support is still under dev")
	default:
		panic("database type not supported")
	}
}
