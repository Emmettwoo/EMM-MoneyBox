package mysql

import (
	"bytes"
	"database/sql"
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/emmettwoo/EMM-MoneyBox/util/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var categoryMySqlMapper CategoryMySqlMapper

type CategoryMySqlMapper struct{}

func (CategoryMySqlMapper) GetCategoryByObjectId(plainId string) entity.CategoryEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, PARENT_ID, NAME FROM ")
	sqlString.WriteString(database.CategoryTableName)
	sqlString.WriteString(" WHERE ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), plainId)
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var categoryEntity entity.CategoryEntity
	for rows.Next() {
		categoryEntity = convertRow2CategoryEntity(rows)
		break
	}
	return categoryEntity
}

func (CategoryMySqlMapper) GetCategoryByName(categoryName string) entity.CategoryEntity {

	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, PARENT_ID, NAME FROM ")
	sqlString.WriteString(database.CategoryTableName)
	sqlString.WriteString(" WHERE NAME = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), categoryName)
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var categoryEntity entity.CategoryEntity
	for rows.Next() {
		categoryEntity = convertRow2CategoryEntity(rows)
		break
	}
	return categoryEntity
}

func (CategoryMySqlMapper) GetCategoryByParentId(parentPlainId string) []entity.CategoryEntity {
	var sqlString bytes.Buffer
	sqlString.WriteString("SELECT ID, PARENT_ID, NAME FROM ")
	sqlString.WriteString(database.CategoryTableName)
	sqlString.WriteString(" WHERE PARENT_ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	rows, err := connection.Query(sqlString.String(), parentPlainId)
	if err != nil {
		util.Logger.Errorw("query failed", "error", err)
	}

	var targetEntityList []entity.CategoryEntity
	for rows.Next() {
		targetEntityList = append(targetEntityList, convertRow2CategoryEntity(rows))
	}
	return targetEntityList
}

func (CategoryMySqlMapper) InsertCategoryByEntity(newEntity entity.CategoryEntity) string {

	var operatingTime = time.Now()
	newEntity.CreateTime = operatingTime
	newEntity.ModifyTime = operatingTime

	var sqlString bytes.Buffer
	sqlString.WriteString("INSERT ")
	sqlString.WriteString(database.CategoryTableName)
	sqlString.WriteString(" SET ID = ?, ")
	sqlString.WriteString(" PARENT_ID = ?, ")
	sqlString.WriteString(" NAME = ?, ")
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
	result, err := statement.Exec(newPlainId, newEntity.ParentId.Hex(), newEntity.Name,
		newEntity.Remark, operatingTime, operatingTime)
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

func (CategoryMySqlMapper) UpdateCategoryByEntity(plainId string) entity.CategoryEntity {

	var targetEntity = categoryMySqlMapper.GetCategoryByObjectId(plainId)
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("category is not exist")
		return entity.CategoryEntity{}
	}

	// todo: update specific fields by passing params (parentId, name)
	targetEntity.ModifyTime = time.Now()

	var sqlString bytes.Buffer
	sqlString.WriteString("UPDATE ")
	sqlString.WriteString(database.CategoryTableName)
	sqlString.WriteString(" SET PARENT_ID = ?, ")
	sqlString.WriteString(" NAME = ?, ")
	sqlString.WriteString(" REMARK = ?, ")
	sqlString.WriteString(" MODIFY_TIME = ? ")
	sqlString.WriteString(" WHERE ID = ? ")

	connection := database.GetMySqlConnection()
	defer database.CloseMySqlConnection()

	statement, err := connection.Prepare(sqlString.String())
	if err != nil {
		util.Logger.Errorw("update failed", "error", err)
	}

	result, err := statement.Exec(targetEntity.ParentId.Hex(), targetEntity.Name, targetEntity.Remark,
		targetEntity.ModifyTime, targetEntity.Id)
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

func (CategoryMySqlMapper) DeleteCategoryByObjectId(plainId string) entity.CategoryEntity {

	var targetEntity = categoryMySqlMapper.GetCategoryByObjectId(plainId)
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("category is not exist")
		return entity.CategoryEntity{}
	}

	// can not delete a category that has referred cash_flows.
	if cashFlowMySqlMapper.CountCashFLowsByCategoryId(targetEntity.Id.Hex()) != 0 {
		util.Logger.Infoln("can not delete a category which has cash_flows refer to")
		return entity.CategoryEntity{}
	}

	// can not delete a category that has referred child-categories.
	if len(categoryMySqlMapper.GetCategoryByParentId(plainId)) != 0 {
		util.Logger.Infoln("can not delete a category which has child-categories refer to")
		return entity.CategoryEntity{}
	}

	var sqlString bytes.Buffer
	sqlString.WriteString("DELETE FROM ")
	sqlString.WriteString(database.CategoryTableName)
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

func convertRow2CategoryEntity(rows *sql.Rows) entity.CategoryEntity {

	var id string
	var parentId string
	var name string

	err := rows.Scan(&id, &parentId, &name)
	if err != nil {
		util.Logger.Errorw("covert into entity failed", "error", err)
	}

	return entity.CategoryEntity{
		Id:       util.Convert2ObjectId(id),
		ParentId: util.Convert2ObjectId(parentId),
		Name:     name,
	}
}
