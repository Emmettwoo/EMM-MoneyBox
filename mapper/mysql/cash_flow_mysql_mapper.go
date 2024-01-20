package mysql

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/emmettwoo/EMM-MoneyBox/util/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var cashFlowMySqlMapper CashFlowMySqlMapper

type CashFlowMySqlMapper struct{}

func (CashFlowMySqlMapper) GetCashFlowByObjectId(plainId string) model.CashFlowEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, CATEGORY_ID, BELONGS_DATE, FLOW_TYPE, AMOUNT, DESCRIPTION FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), plainId)
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var cashFlowEntity model.CashFlowEntity
	for rows.Next() {
		cashFlowEntity = convertRow2CashFlowEntity(rows)
		break
	}
	return cashFlowEntity
}

func (CashFlowMySqlMapper) GetCashFlowsByObjectIdArray(plainIdList []string) []model.CashFlowEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, CATEGORY_ID, BELONGS_DATE, FLOW_TYPE, AMOUNT, DESCRIPTION FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE ID in ")
	// fixme: pass the params by ? instead to avoid SQL inject.
	sqlString.WriteString("(" + util.CombiningWithComma(util.BatchSurroundingWithSingleQuotes(plainIdList)) + ") ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String())
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var targetEntityList = make([]model.CashFlowEntity, len(plainIdList))
	for rows.Next() {
		targetEntityList = append(targetEntityList, convertRow2CashFlowEntity(rows))
	}
	return targetEntityList
}

func (CashFlowMySqlMapper) GetCashFlowsByBelongsDate(belongsDate time.Time) []model.CashFlowEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, CATEGORY_ID, BELONGS_DATE, FLOW_TYPE, AMOUNT, DESCRIPTION FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE BELONGS_DATE = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), util.FormatDateToStringWithDash(belongsDate))
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var targetEntityList []model.CashFlowEntity
	for rows.Next() {
		targetEntityList = append(targetEntityList, convertRow2CashFlowEntity(rows))
	}
	return targetEntityList
}

func (CashFlowMySqlMapper) GetCashFlowsByCategoryId(categoryPlainId string) []model.CashFlowEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, CATEGORY_ID, BELONGS_DATE, FLOW_TYPE, AMOUNT, DESCRIPTION FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE CATEGORY_ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), categoryPlainId)
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var targetEntityList []model.CashFlowEntity
	for rows.Next() {
		targetEntityList = append(targetEntityList, convertRow2CashFlowEntity(rows))
	}
	return targetEntityList
}

func (CashFlowMySqlMapper) GetCashFlowsByCategoryName(categoryName string) []model.CashFlowEntity {

	var targetCategoryEntity = categoryMySqlMapper.GetCategoryByName(categoryName)
	if targetCategoryEntity.IsEmpty() {
		util.Logger.Infoln("category is not exist")
		return nil
	}

	return cashFlowMySqlMapper.GetCashFlowsByCategoryId(targetCategoryEntity.Id.Hex())
}

func (CashFlowMySqlMapper) GetCashFlowsByExactDesc(description string) []model.CashFlowEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, CATEGORY_ID, BELONGS_DATE, FLOW_TYPE, AMOUNT, DESCRIPTION FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE DESCRIPTION = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), description)
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var targetEntityList []model.CashFlowEntity
	for rows.Next() {
		targetEntityList = append(targetEntityList, convertRow2CashFlowEntity(rows))
	}
	return targetEntityList
}

func (CashFlowMySqlMapper) GetCashFlowsByFuzzyDesc(description string) []model.CashFlowEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, CATEGORY_ID, BELONGS_DATE, FLOW_TYPE, AMOUNT, DESCRIPTION FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE DESCRIPTION LIKE ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), "%"+description+"%")
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var targetEntityList []model.CashFlowEntity
	for rows.Next() {
		targetEntityList = append(targetEntityList, convertRow2CashFlowEntity(rows))
	}
	return targetEntityList
}

func (CashFlowMySqlMapper) CountCashFLowsByCategoryId(categoryPlainId string) int64 {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT COUNT(1) FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE CATEGORY_ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), categoryPlainId)
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var rowsAffected int64
	rows.Next()
	if err = rows.Scan(&rowsAffected); err != nil {
		util.Logger.Errorw("parse row affected failed", "error", err)
		return -1
	}
	return rowsAffected
}

func (CashFlowMySqlMapper) InsertCashFlowByEntity(newEntity model.CashFlowEntity) string {

	var operatingTime = time.Now()
	newEntity.CreateTime = operatingTime
	newEntity.ModifyTime = operatingTime

	var sqlString bytes.Buffer
	sqlString.WriteString("INSERT ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" SET ID = ?, ")
	sqlString.WriteString(" CATEGORY_ID = ?, ")
	sqlString.WriteString(" BELONGS_DATE = ?, ")
	sqlString.WriteString(" FLOW_TYPE = ?, ")
	sqlString.WriteString(" AMOUNT = ?, ")
	sqlString.WriteString(" DESCRIPTION = ?, ")
	sqlString.WriteString(" REMARK = ?, ")
	sqlString.WriteString(" CREATE_TIME = ?, ")
	sqlString.WriteString(" MODIFY_TIME = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	statement, err := connection.Prepare(sqlString.String())
	if err != nil {
		util.Logger.Errorw("insert failed", "error", err)
	}

	var newPlainId = primitive.NewObjectID().Hex()
	result, err := statement.Exec(newPlainId, newEntity.CategoryId.Hex(), newEntity.BelongsDate, newEntity.FlowType,
		newEntity.Amount, newEntity.Description, newEntity.Remark, operatingTime, operatingTime)
	if err != nil {
		util.Logger.Errorw("insert failed", "error", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("insert failed", "error", err, "rows_affected", rowsAffected)
	}
	return newPlainId
}

func (CashFlowMySqlMapper) UpdateCashFlowByEntity(plainId string) model.CashFlowEntity {

	var objectId = util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("cash_flow's id is not acceptable")
		return model.CashFlowEntity{}
	}

	var targetEntity = cashFlowMySqlMapper.GetCashFlowByObjectId(plainId)
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("cash_flow is not exist")
		return model.CashFlowEntity{}
	}

	// todo: update specific fields by passing params (category_name, belongs_date, flow_type, amount, description)
	targetEntity.ModifyTime = time.Now()

	var sqlString bytes.Buffer
	sqlString.WriteString("UPDATE ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" SET CATEGORY_ID = ?, ")
	sqlString.WriteString(" BELONGS_DATE = ?, ")
	sqlString.WriteString(" FLOW_TYPE = ?, ")
	sqlString.WriteString(" AMOUNT = ?, ")
	sqlString.WriteString(" DESCRIPTION = ?, ")
	sqlString.WriteString(" REMARK = ?, ")
	sqlString.WriteString(" MODIFY_TIME = ? ")
	sqlString.WriteString(" WHERE ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	statement, err := connection.Prepare(sqlString.String())
	if err != nil {
		util.Logger.Errorw("update failed", "error", err)
	}

	result, err := statement.Exec(targetEntity.CategoryId.Hex(), targetEntity.BelongsDate, targetEntity.FlowType,
		targetEntity.Amount, targetEntity.Description, targetEntity.Remark, targetEntity.ModifyTime, plainId)
	if err != nil {
		util.Logger.Errorw("update failed", "error", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("update failed", "error", err, "rows_affected", rowsAffected)
	}
	return targetEntity
}

func (CashFlowMySqlMapper) DeleteCashFlowByObjectId(plainId string) model.CashFlowEntity {

	var targetEntity = cashFlowMySqlMapper.GetCashFlowByObjectId(plainId)
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("cash_flow is not exist")
		return model.CashFlowEntity{}
	}

	var sqlString bytes.Buffer
	sqlString.WriteString("DELETE FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	statement, err := connection.Prepare(sqlString.String())
	if err != nil {
		util.Logger.Errorw("delete failed", "error", err)
	}

	result, err := statement.Exec(plainId)
	if err != nil {
		util.Logger.Errorw("delete failed", "error", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("delete failed", "error", err, "rows_affected", rowsAffected)
	}
	return targetEntity
}

func (CashFlowMySqlMapper) DeleteCashFlowByBelongsDate(belongsDate time.Time) []model.CashFlowEntity {

	var cashFlowList = cashFlowMySqlMapper.GetCashFlowsByBelongsDate(belongsDate)
	if cashFlowList == nil {
		util.Logger.Infoln("no cash_flow(s) found")
		return cashFlowList
	}

	var sqlString bytes.Buffer
	sqlString.WriteString("DELETE FROM ")
	sqlString.WriteString(database.CashFlowTableName)
	sqlString.WriteString(" WHERE BELONGS_DATE = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	statement, err := connection.Prepare(sqlString.String())
	if err != nil {
		util.Logger.Errorw("delete failed", "error", err)
	}

	result, err := statement.Exec(util.FormatDateToStringWithDash(belongsDate))
	if err != nil {
		util.Logger.Errorw("delete failed", "error", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != int64(len(cashFlowList)) {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("delete failed", "error", err, "rows_affected", rowsAffected)
	}
	return cashFlowList
}

func convertRow2CashFlowEntity(rows *sql.Rows) model.CashFlowEntity {

	var id string
	var categoryId string
	var belongsDate string
	var flowType string
	var amount float64
	var description string

	err := rows.Scan(&id, &categoryId, &belongsDate, &flowType, &amount, &description)
	if err != nil {
		util.Logger.Errorw("covert into entity failed", "error", err)
	}

	return model.CashFlowEntity{
		Id:          util.Convert2ObjectId(id),
		CategoryId:  util.Convert2ObjectId(categoryId),
		BelongsDate: util.FormatDateFromStringWithDash(belongsDate),
		FlowType:    flowType,
		Amount:      amount,
		Description: description,
	}
}
