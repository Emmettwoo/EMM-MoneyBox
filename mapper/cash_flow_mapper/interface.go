package cash_flow_mapper

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
)

var INSTANCE CashFlowMapper

type CashFlowMapper interface {
	GetCashFlowByObjectId(plainId string) model.CashFlowEntity
	GetCashFlowsByObjectIdArray(plainIdList []string) []model.CashFlowEntity
	GetCashFlowsByBelongsDate(belongsDate time.Time) []model.CashFlowEntity
	GetCashFlowsByCategoryId(categoryPlainId string) []model.CashFlowEntity
	GetCashFlowsByExactDesc(description string) []model.CashFlowEntity
	GetCashFlowsByFuzzyDesc(description string) []model.CashFlowEntity
	CountCashFLowsByCategoryId(categoryPlainId string) int64
	InsertCashFlowByEntity(newEntity model.CashFlowEntity) string
	UpdateCashFlowByEntity(plainId string) model.CashFlowEntity
	DeleteCashFlowByObjectId(plainId string) model.CashFlowEntity
	DeleteCashFlowByBelongsDate(belongsDate time.Time) []model.CashFlowEntity
}

func init() {
	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		INSTANCE = CashFlowMongoDbMapper{}
	case "mysql":
		INSTANCE = CashFlowMySqlMapper{}
	default:
		panic("database type not supported")
	}
}
