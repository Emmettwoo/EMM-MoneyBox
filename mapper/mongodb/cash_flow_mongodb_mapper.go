package mongodb

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/emmettwoo/EMM-MoneyBox/util/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var cashFlowMongoDbMapper CashFlowMongoDbMapper

type CashFlowMongoDbMapper struct{}

func (CashFlowMongoDbMapper) GetCashFlowByObjectId(plainId string) model.CashFlowEntity {

	objectId := util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("cash_flow's id is not acceptable")
		return model.CashFlowEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()
	return convertBsonM2CashFlowEntity(database.GetOneInMongoDB(filter))
}

func (CashFlowMongoDbMapper) GetCashFlowsByObjectIdArray(plainIdList []string) []model.CashFlowEntity {

	var objectIdArray = make([]primitive.ObjectID, len(plainIdList))
	for _, plainId := range plainIdList {
		objectId := util.Convert2ObjectId(plainId)
		objectIdArray = append(objectIdArray, objectId)
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: bson.M{"$in": objectIdArray}},
	}

	// 打开cashFlow的数据表连线
	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	// 获取查询结果并转入结构对象
	var targetEntityList []model.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}
	return targetEntityList
}

func (CashFlowMongoDbMapper) GetCashFlowsByBelongsDate(belongsDate time.Time) []model.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "belongs_date", Value: belongsDate},
	}

	// 打开cashFlow的数据表连线
	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	// 获取查询结果并转入结构对象
	var targetEntityList []model.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}
	return targetEntityList
}

func (CashFlowMongoDbMapper) GetCashFlowsByCategoryId(categoryPlainId string) []model.CashFlowEntity {

	categoryObjectId := util.Convert2ObjectId(categoryPlainId)
	if categoryPlainId == "" || categoryObjectId == primitive.NilObjectID {
		util.Logger.Warnln("category's id is not acceptable")
		return nil
	}

	filter := bson.D{
		primitive.E{Key: "category_id", Value: categoryObjectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	var targetEntityList []model.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}
	return targetEntityList
}

func (CashFlowMongoDbMapper) CountCashFLowsByCategoryId(categoryPlainId string) int64 {

	categoryObjectId := util.Convert2ObjectId(categoryPlainId)
	if categoryPlainId == "" || categoryObjectId == primitive.NilObjectID {
		util.Logger.Warnln("category's id is not acceptable")
		return 0
	}

	filter := bson.D{
		primitive.E{Key: "category_id", Value: categoryObjectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	return database.CountInMongoDB(filter)
}

func (CashFlowMongoDbMapper) GetCashFlowsByCategoryName(categoryName string) []model.CashFlowEntity {

	var categoryEntity = categoryMongoDbMapper.GetCategoryByName(categoryName)
	if categoryEntity.IsEmpty() {
		util.Logger.Warnln("category name not existed")
		return nil
	}

	return cashFlowMongoDbMapper.GetCashFlowsByCategoryId(categoryEntity.Id.Hex())
}

func (CashFlowMongoDbMapper) GetCashFlowsByExactDesc(description string) []model.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "description", Value: description},
	}

	// 打开cashFlow的数据表连线
	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	// 获取查询结果并转入结构对象
	var targetEntityList []model.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}

	return targetEntityList
}

func (CashFlowMongoDbMapper) GetCashFlowsByFuzzyDesc(description string) []model.CashFlowEntity {

	// Options i for disable case sensitive.
	filter := bson.D{
		primitive.E{Key: "description", Value: primitive.Regex{
			Pattern: description,
			Options: "i",
		}},
	}

	// 打开 cash_flow 的数据表连线
	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	// 获取查询结果并转入结构对象
	var targetEntityList []model.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}
	return targetEntityList
}

func (CashFlowMongoDbMapper) InsertCashFlowByEntity(newEntity model.CashFlowEntity) string {

	var operatingTime = time.Now()
	newEntity.CreateTime = operatingTime
	newEntity.ModifyTime = operatingTime

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	newCashFlowId := database.InsertOneInMongoDB(convertCashFlowEntity2BsonD(newEntity))
	return newCashFlowId.Hex()
}

func (CashFlowMongoDbMapper) UpdateCashFlowByEntity(plainId string) model.CashFlowEntity {

	var objectId = util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("cash_flow's id is not acceptable")
		return model.CashFlowEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	var targetEntity = convertBsonM2CashFlowEntity(database.GetOneInMongoDB(filter))
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("cash_flow is not exist")
		return model.CashFlowEntity{}
	}

	// todo: update specific fields by passing params (category_name, belongs_date, flow_type, amount, description)
	targetEntity.ModifyTime = time.Now()

	var rowsAffected = database.UpdateManyInMongoDB(filter, convertCashFlowEntity2BsonD(targetEntity))
	if rowsAffected != 1 {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("update failed", "rows_affected", rowsAffected)
		return model.CashFlowEntity{}
	}

	return targetEntity
}

func (CashFlowMongoDbMapper) DeleteCashFlowByObjectId(plainId string) model.CashFlowEntity {

	objectId := util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("cash_flow's id is not acceptable")
		return model.CashFlowEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()
	var targetEntity = convertBsonM2CashFlowEntity(database.GetOneInMongoDB(filter))
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("cash_flow is not exist")
		return model.CashFlowEntity{}
	}
	var rowsAffected = database.DeleteManyInMongoDB(filter)
	if rowsAffected != 1 {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("delete failed", "rows_affected", rowsAffected)
		return model.CashFlowEntity{}
	}
	return targetEntity
}

func (CashFlowMongoDbMapper) DeleteCashFlowByBelongsDate(belongsDate time.Time) []model.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "belongs_date", Value: belongsDate},
	}

	var cashFlowList = cashFlowMongoDbMapper.GetCashFlowsByBelongsDate(belongsDate)
	if cashFlowList == nil {
		util.Logger.Infoln("no cash_flow(s) found")
		return []model.CashFlowEntity{}
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	var rowsAffected = database.DeleteManyInMongoDB(filter)
	if rowsAffected != int64(len(cashFlowList)) {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("delete failed", "rows_affected", rowsAffected)
	}
	return cashFlowList
}

func convertCashFlowEntity2BsonD(entity model.CashFlowEntity) bson.D {

	// 为空时自动生成新Id
	if entity.Id == primitive.NilObjectID {
		entity.Id = primitive.NewObjectID()
	}

	return bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
		primitive.E{Key: "category_id", Value: entity.CategoryId},
		primitive.E{Key: "belongs_date", Value: entity.BelongsDate},
		primitive.E{Key: "flow_type", Value: entity.FlowType},
		primitive.E{Key: "amount", Value: entity.Amount},
		primitive.E{Key: "description", Value: entity.Description},
		primitive.E{Key: "remark", Value: entity.Remark},
		primitive.E{Key: "create_time", Value: entity.CreateTime},
		primitive.E{Key: "modify_time", Value: entity.ModifyTime},
	}
}

func convertBsonM2CashFlowEntity(bsonM bson.M) model.CashFlowEntity {

	var newEntity model.CashFlowEntity
	bsonBytes, err := bson.Marshal(bsonM)
	if err != nil {
		util.Logger.Errorln(err)
		panic(err)
	}
	if err = bson.Unmarshal(bsonBytes, &newEntity); err != nil {
		util.Logger.Errorln(err)
		panic(err)
	}
	return newEntity
}
