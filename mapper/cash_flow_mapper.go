package mapper

import (
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mysql"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"time"
)

var cashFlowMongoDbMapper CashFlowMapper
var cashFlowMySqlMapper CashFlowMapper
var CashFlowCommonMapper CashFlowMapper

type CashFlowMapper interface {
	GetCashFlowByObjectId(plainId string) entity.CashFlowEntity
	GetCashFlowsByObjectIdArray(plainIdList []string) []entity.CashFlowEntity
	GetCashFlowsByBelongsDate(belongsDate time.Time) []entity.CashFlowEntity
	GetCashFlowsByCategoryId(categoryPlainId string) []entity.CashFlowEntity
	GetCashFlowsByCategoryName(categoryName string) []entity.CashFlowEntity
	GetCashFlowsByExactDesc(description string) []entity.CashFlowEntity
	GetCashFlowsByFuzzyDesc(description string) []entity.CashFlowEntity
	CountCashFLowsByCategoryId(categoryPlainId string) int64
	InsertCashFlowByEntity(newEntity entity.CashFlowEntity) string
	UpdateCashFlowByEntity(plainId string) entity.CashFlowEntity
	DeleteCashFlowByObjectId(plainId string) entity.CashFlowEntity
	DeleteCashFlowByBelongsDate(belongsDate time.Time) []entity.CashFlowEntity
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
