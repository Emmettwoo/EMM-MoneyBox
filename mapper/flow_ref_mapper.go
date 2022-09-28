package mapper

import (
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlowRefMapper interface {
	GetFlowRefByCashFlowId(objectId primitive.ObjectID) entity.FlowRefEntity
	GetFlowRefByDayFlowId(objectId primitive.ObjectID) []entity.FlowRefEntity
	InsertFlowRefByEntity(entity entity.FlowRefEntity) primitive.ObjectID
}

func GetFlowRefMapper() FlowRefMapper {
	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		var flowRefMongoDbMapper FlowRefMapper
		flowRefMongoDbMapper = mongodb.FlowRefMongoDbMapper{}
		return flowRefMongoDbMapper
	case "mysql":
		panic("mysql support is still under dev")
	default:
		panic("database type not supported")
	}
}
