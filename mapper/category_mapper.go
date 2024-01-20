package mapper

import (
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mysql"
	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
)

var categoryMongoDbMapper CategoryMapper
var categoryMySqlMapper CategoryMapper
var CategoryCommonMapper CategoryMapper

type CategoryMapper interface {
	GetCategoryByObjectId(plainId string) model.CategoryEntity
	GetCategoryByName(categoryName string) model.CategoryEntity
	GetCategoryByParentId(parentPlainId string) []model.CategoryEntity
	InsertCategoryByEntity(newEntity model.CategoryEntity) string
	UpdateCategoryByEntity(plainId string) model.CategoryEntity
	DeleteCategoryByObjectId(plainId string) model.CategoryEntity
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
