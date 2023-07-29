package mongodb

import (
    "github.com/emmettwoo/EMM-MoneyBox/entity"
    "github.com/emmettwoo/EMM-MoneyBox/mapper"
    "github.com/emmettwoo/EMM-MoneyBox/util"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryMongoDbMapper struct{}

var categoryMongoDbMapper CategoryMongoDbMapper

func (CategoryMongoDbMapper) GetCategoryByObjectId(objectId primitive.ObjectID) entity.CategoryEntity {

    filter := bson.D{
        primitive.E{Key: "_id", Value: objectId},
    }

    util.OpenMongoDbConnection(mapper.CategoryTableName)
    return convertBsonM2CategoryEntity(util.GetOneInMongoDb(filter))
}

func (CategoryMongoDbMapper) GetCategoryByName(categoryName string) entity.CategoryEntity {

    filter := bson.D{
        primitive.E{Key: "name", Value: categoryName},
    }

    util.OpenMongoDbConnection(mapper.CategoryTableName)

    return convertBsonM2CategoryEntity(util.GetOneInMongoDb(filter))
}

func convertCategoryEntity2BsonD(entity entity.CategoryEntity) bson.D {

    // 为空时自动生成新Id
    if entity.Id == primitive.NilObjectID {
        entity.Id = primitive.NewObjectID()
    }

    return bson.D{
        primitive.E{Key: "_id", Value: entity.Id},
        primitive.E{Key: "parent_id", Value: entity.ParentId},
        primitive.E{Key: "name", Value: entity.Name},
        primitive.E{Key: "remark", Value: entity.Remark},
    }
}

func convertBsonM2CategoryEntity(bsonM bson.M) entity.CategoryEntity {

    var newEntity entity.CategoryEntity
    bsonBytes, _ := bson.Marshal(bsonM)
    err := bson.Unmarshal(bsonBytes, &newEntity)
    if err != nil {
        panic(err)
    }
    return newEntity
}
