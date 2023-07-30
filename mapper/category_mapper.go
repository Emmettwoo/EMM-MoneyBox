package mapper

import (
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mysql"
	"github.com/emmettwoo/EMM-MoneyBox/util"
)

var categoryMongoDbMapper CategoryMapper
var categoryMySqlMapper CategoryMapper
var CategoryCommonMapper CategoryMapper

type CategoryMapper interface {
	GetCategoryByObjectId(plainId string) entity.CategoryEntity
	GetCategoryByName(categoryName string) entity.CategoryEntity
	InsertCategoryByEntity(newEntity entity.CategoryEntity) string
	UpdateCategoryByEntity(plainId string) entity.CategoryEntity
	DeleteCategoryByObjectId(plainId string) entity.CategoryEntity
}

func init() {
	categoryMongoDbMapper = mongodb.CategoryMongoDbMapper{}
	categoryMySqlMapper = mysql.CategoryMySqlMapper{}
	CategoryCommonMapper = GetCategoryMapper()
}

func GetCategoryMapper() CategoryMapper {

	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		return categoryMongoDbMapper
	case "mysql":
		return categoryMySqlMapper
	default:
		panic("database type not supported")
	}
}
