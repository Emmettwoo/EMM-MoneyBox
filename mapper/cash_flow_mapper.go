package mapper

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mysql"
	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
)

var cashFlowMongoDbMapper CashFlowMapper
var cashFlowMySqlMapper CashFlowMapper
var CashFlowCommonMapper CashFlowMapper

type CashFlowMapper interface {
	GetCashFlowByObjectId(plainId string) model.CashFlowEntity
	GetCashFlowsByObjectIdArray(plainIdList []string) []model.CashFlowEntity
	GetCashFlowsByBelongsDate(belongsDate time.Time) []model.CashFlowEntity
	GetCashFlowsByCategoryId(categoryPlainId string) []model.CashFlowEntity
	GetCashFlowsByCategoryName(categoryName string) []model.CashFlowEntity
	GetCashFlowsByExactDesc(description string) []model.CashFlowEntity
	GetCashFlowsByFuzzyDesc(description string) []model.CashFlowEntity
	CountCashFLowsByCategoryId(categoryPlainId string) int64
	InsertCashFlowByEntity(newEntity model.CashFlowEntity) string
	UpdateCashFlowByEntity(plainId string) model.CashFlowEntity
	DeleteCashFlowByObjectId(plainId string) model.CashFlowEntity
	DeleteCashFlowByBelongsDate(belongsDate time.Time) []model.CashFlowEntity
}

func init() {
	cashFlowMongoDbMapper = mongodb.CashFlowMongoDbMapper{}
	cashFlowMySqlMapper = mysql.CashFlowMySqlMapper{}
	CashFlowCommonMapper = GetCashFlowMapper()
}

func GetCashFlowMapper() CashFlowMapper {

	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		return cashFlowMongoDbMapper
	case "mysql":
		return cashFlowMySqlMapper
	default:
		panic("database type not supported")
	}
}
