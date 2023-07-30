package mysql

import (
	"bytes"
	"database/sql"
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/emmettwoo/EMM-MoneyBox/util/database"
	"time"
)

type CashFlowMySqlMapper struct{}

func (CashFlowMySqlMapper) GetCashFlowByObjectId(plainId string) entity.CashFlowEntity {

	// 打开數據庫连线
	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, CATEGORY_ID, BELONGS_DATE, FLOW_TYPE, AMOUNT, DESCRIPTION FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE ID = ")
	sqlString.WriteString("'" + plainId + "' ")

	rows, err := connection.Query(sqlString.String())
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var cashFlowEntity entity.CashFlowEntity
	for rows.Next() {
		cashFlowEntity = convertRow2CashFlowEntity(rows)
		break
	}
	return cashFlowEntity
}

func (CashFlowMySqlMapper) GetCashFlowsByObjectIdArray(plainIdList []string) []entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return nil
}

func (CashFlowMySqlMapper) GetCashFlowsByBelongsDate(belongsDate time.Time) []entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return nil
}

func (CashFlowMySqlMapper) GetCashFlowsByCategoryId(categoryPlainId string) []entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return nil
}

func (CashFlowMySqlMapper) GetCashFlowsByCategoryName(categoryName string) []entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return nil
}

func (CashFlowMySqlMapper) GetCashFlowsByExactDesc(description string) []entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return nil
}

func (CashFlowMySqlMapper) GetCashFlowsByFuzzyDesc(description string) []entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return nil
}

func (CashFlowMySqlMapper) CountCashFLowsByCategoryId(categoryPlainId string) int64 {

	util.Logger.Errorln("non-supported yet.")
	return 0
}

func (CashFlowMySqlMapper) InsertCashFlowByEntity(newEntity entity.CashFlowEntity) string {

	util.Logger.Errorln("non-supported yet.")
	return ""
}

func (CashFlowMySqlMapper) UpdateCashFlowByEntity(plainId string) entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return entity.CashFlowEntity{}
}

func (CashFlowMySqlMapper) DeleteCashFlowByObjectId(plainId string) entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return entity.CashFlowEntity{}
}

func (CashFlowMySqlMapper) DeleteCashFlowByBelongsDate(belongsDate time.Time) []entity.CashFlowEntity {

	util.Logger.Errorln("non-supported yet.")
	return nil
}

func convertRow2CashFlowEntity(rows *sql.Rows) entity.CashFlowEntity {

	var id string
	var categoryId string
	var belongsDate time.Time
	var flowType string
	var amount float64
	var description string

	err := rows.Scan(&id, &categoryId, &belongsDate, &flowType, &amount, &description)
	if err != nil {
		util.Logger.Errorw("covert into entity failed", "error", err)
	}

	return entity.CashFlowEntity{
		Id:          util.Convert2ObjectId(id),
		CategoryId:  util.Convert2ObjectId(categoryId),
		BelongsDate: belongsDate,
		FlowType:    flowType,
		Amount:      amount,
		Description: description,
	}
}
