package category_mapper

import (
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/emmettwoo/EMM-MoneyBox/util/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryMongoDbMapper struct{}

func (CategoryMongoDbMapper) GetCategoryByObjectId(plainId string) model.CategoryEntity {

	objectId := util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("category's id is not acceptable")
		return model.CategoryEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CategoryTableName)
	defer database.CloseMongoDbConnection()
	return convertBsonM2CategoryEntity(database.GetOneInMongoDB(filter))
}

func (CategoryMongoDbMapper) GetCategoryByName(categoryName string) model.CategoryEntity {

	filter := bson.D{
		primitive.E{Key: "name", Value: categoryName},
	}

	database.OpenMongoDbConnection(database.CategoryTableName)
	defer database.CloseMongoDbConnection()
	return convertBsonM2CategoryEntity(database.GetOneInMongoDB(filter))
}

func (CategoryMongoDbMapper) GetCategoryByParentId(parentPlainId string) []model.CategoryEntity {

	filter := bson.D{
		primitive.E{Key: "parent_id", Value: parentPlainId},
	}

	database.OpenMongoDbConnection(database.CategoryTableName)
	defer database.CloseMongoDbConnection()

	var targetEntityList []model.CategoryEntity
	queryResultList := database.GetManyInMongoDB(filter)
	for _, queryResult := range queryResultList {
		targetEntityList = append(targetEntityList, convertBsonM2CategoryEntity(queryResult))
	}

	return targetEntityList
}

func (CategoryMongoDbMapper) InsertCategoryByEntity(newEntity model.CategoryEntity) string {

	var operatingTime = time.Now()
	newEntity.CreateTime = operatingTime
	newEntity.ModifyTime = operatingTime

	database.OpenMongoDbConnection(database.CategoryTableName)
	defer database.CloseMongoDbConnection()

	var newCategoryId = database.InsertOneInMongoDB(convertCategoryEntity2BsonD(newEntity))
	return newCategoryId.Hex()
}

func (CategoryMongoDbMapper) UpdateCategoryByEntity(plainId string) model.CategoryEntity {

	var objectId = util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("category's id is not acceptable")
		return model.CategoryEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CategoryTableName)
	defer database.CloseMongoDbConnection()

	var targetEntity = convertBsonM2CategoryEntity(database.GetOneInMongoDB(filter))
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("category is not exist")
		return model.CategoryEntity{}
	}

	// todo: update specific fields by passing params (parentId, name)
	targetEntity.ModifyTime = time.Now()

	var rowsAffected = database.UpdateManyInMongoDB(filter, convertCategoryEntity2BsonD(targetEntity))
	if rowsAffected != 1 {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("update failed", "rows_affected", rowsAffected)
		return model.CategoryEntity{}
	}

	return model.CategoryEntity{}
}

func (CategoryMongoDbMapper) DeleteCategoryByObjectId(plainId string) model.CategoryEntity {

	objectId := util.Convert2ObjectId(plainId)
	if plainId == "" || objectId == primitive.NilObjectID {
		util.Logger.Warnln("category's id is not acceptable")
		return model.CategoryEntity{}
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	database.OpenMongoDbConnection(database.CategoryTableName)
	defer database.CloseMongoDbConnection()

	var targetEntity = convertBsonM2CategoryEntity(database.GetOneInMongoDB(filter))
	if targetEntity.IsEmpty() {
		util.Logger.Infoln("category is not exist")
		return model.CategoryEntity{}
	}

	// can not delete a category that has referred child-categories.
	if len(INSTANCE.GetCategoryByParentId(plainId)) != 0 {
		util.Logger.Infoln("can not delete a category which has child-categories refer to")
		return model.CategoryEntity{}
	}

	var rowsAffected = database.DeleteManyInMongoDB(filter)
	if rowsAffected != 1 {
		// fixme: maybe we should have a rollback here.
		util.Logger.Errorw("delete failed", "rows_affected", rowsAffected)
		return model.CategoryEntity{}
	}
	return targetEntity
}

func convertCategoryEntity2BsonD(entity model.CategoryEntity) bson.D {

	// 为空时自动生成新Id
	if entity.Id == primitive.NilObjectID {
		entity.Id = primitive.NewObjectID()
	}

	return bson.D{
		primitive.E{Key: "_id", Value: entity.Id},
		primitive.E{Key: "parent_id", Value: entity.ParentId},
		primitive.E{Key: "name", Value: entity.Name},
		primitive.E{Key: "remark", Value: entity.Remark},
		primitive.E{Key: "create_time", Value: entity.CreateTime},
		primitive.E{Key: "modify_time", Value: entity.ModifyTime},
	}
}

func convertBsonM2CategoryEntity(bsonM bson.M) model.CategoryEntity {

	var newEntity model.CategoryEntity
	bsonBytes, _ := bson.Marshal(bsonM)
	err := bson.Unmarshal(bsonBytes, &newEntity)
	if err != nil {
		panic(err)
	}
	return newEntity
}
