package mapper

import (
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CategoryTableName = "category"
var categoryMongoDbMapper CategoryMapper

type CategoryMapper interface {
	GetCategoryByObjectId(objectId primitive.ObjectID) entity.CategoryEntity
	GetCategoryByName(categoryName string) entity.CategoryEntity
}

func init() {
	categoryMongoDbMapper = mongodb.CategoryMongoDbMapper{}
}

func GetCategoryMapper() CategoryMapper {

	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		return categoryMongoDbMapper
	case "mysql":
		panic("mysql support is still under dev")
	default:
		panic("database type not supported")
	}
}
