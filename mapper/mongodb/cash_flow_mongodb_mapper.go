package mongodb

import (
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/emmettwoo/EMM-MoneyBox/util/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var cashFlowMongoDbMapper CashFlowMongoDbMapper

type CashFlowMongoDbMapper struct{}

func (CashFlowMongoDbMapper) GetCashFlowByObjectId(plainId string) entity.CashFlowEntity {

	objectId := util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("cash_flow's id is not acceptable")
		return entity.CashFlowEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()
	return convertBsonM2CashFlowEntity(database.GetOneInMongoDB(filter))
}

func (CashFlowMongoDbMapper) GetCashFlowsByObjectIdArray(plainIdList []string) []entity.CashFlowEntity {

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
	var targetEntityList []entity.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}
	return targetEntityList
}

func (CashFlowMongoDbMapper) GetCashFlowsByBelongsDate(belongsDate time.Time) []entity.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "belongs_date", Value: belongsDate},
	}

	// 打开cashFlow的数据表连线
	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	// 获取查询结果并转入结构对象
	var targetEntityList []entity.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}
	return targetEntityList
}

func (CashFlowMongoDbMapper) GetCashFlowsByCategoryId(categoryPlainId string) []entity.CashFlowEntity {

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

	var targetEntityList []entity.CashFlowEntity
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

func (CashFlowMongoDbMapper) GetCashFlowsByCategoryName(categoryName string) []entity.CashFlowEntity {

	var categoryEntity = categoryMongoDbMapper.GetCategoryByName(categoryName)
	if categoryEntity.IsEmpty() {
		util.Logger.Warnln("category name not existed")
		return nil
	}

	return cashFlowMongoDbMapper.GetCashFlowsByCategoryId(categoryEntity.Id.Hex())
}

func (CashFlowMongoDbMapper) GetCashFlowsByExactDesc(description string) []entity.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "description", Value: description},
	}

	// 打开cashFlow的数据表连线
	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	// 获取查询结果并转入结构对象
	var targetEntityList []entity.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}

	return targetEntityList
}

func (CashFlowMongoDbMapper) GetCashFlowsByFuzzyDesc(description string) []entity.CashFlowEntity {

	// Options i for disable case sensitive.
	filter := bson.D{
		primitive.E{Key: "description", Value: primitive.Regex{
			Pattern: description,
			Options: "i",
		}},
	}

	// 打开cashFlow的数据表连线
	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	// 获取查询结果并转入结构对象
	var targetEntityList []entity.CashFlowEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CashFlowEntity(queryResult))
	}
	return targetEntityList
}

func (CashFlowMongoDbMapper) InsertCashFlowByEntity(newEntity entity.CashFlowEntity) string {

	var operatingTime = time.Now()
	newEntity.CreateTime = operatingTime
	newEntity.ModifyTime = operatingTime

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	newCashFlowId := database.InsertOneInMongoDB(convertCashFlowEntity2BsonD(newEntity))
	return newCashFlowId.Hex()
}

func (CashFlowMongoDbMapper) UpdateCashFlowByEntity(plainId string) entity.CashFlowEntity {

	var objectId = util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("cash_flow's id is not acceptable")
		return entity.CashFlowEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()

	var targetEntity = convertBsonM2CashFlowEntity(database.GetOneInMongoDB(filter))
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("cash_flow is not exist")
		return entity.CashFlowEntity{}
	}

	// todo: update specific fields by passing params (category_name, belongs_date, flow_type, amount, description)

	targetEntity.ModifyTime = time.Now()
	if database.UpdateManyInMongoDB(filter, convertCashFlowEntity2BsonD(targetEntity)) == 0 {
		util.Logger.Errorln("cash_flow update failed")
		return entity.CashFlowEntity{}
	}

	return targetEntity
}

func (CashFlowMongoDbMapper) DeleteCashFlowByObjectId(plainId string) entity.CashFlowEntity {

	objectId := util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("cash_flow's id is not acceptable")
		return entity.CashFlowEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()
	var targetEntity = convertBsonM2CashFlowEntity(database.GetOneInMongoDB(filter))
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("cash_flow is not exist")
		return entity.CashFlowEntity{}
	}
	if database.DeleteManyInMongoDB(filter) == 0 {
		util.Logger.Errorln("cash_flow delete failed")
		return entity.CashFlowEntity{}
	}
	return targetEntity
}

func (CashFlowMongoDbMapper) DeleteCashFlowByBelongsDate(belongsDate time.Time) []entity.CashFlowEntity {

	filter := bson.D{
		primitive.E{Key: "belongs_date", Value: belongsDate},
	}
	var cashFlowList = cashFlowMongoDbMapper.GetCashFlowsByBelongsDate(belongsDate)

	database.OpenMongoDbConnection(database.CashFlowTableName)
	defer database.CloseMongoDbConnection()
	database.DeleteManyInMongoDB(filter)
	return cashFlowList
}

func convertCashFlowEntity2BsonD(entity entity.CashFlowEntity) bson.D {

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

func convertBsonM2CashFlowEntity(bsonM bson.M) entity.CashFlowEntity {

	var newEntity entity.CashFlowEntity
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
