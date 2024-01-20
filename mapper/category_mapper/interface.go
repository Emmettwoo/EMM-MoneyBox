package category_mapper

import (
	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
)

var INSTANCE CategoryMapper

type CategoryMapper interface {
	GetCategoryByObjectId(plainId string) model.CategoryEntity
	GetCategoryByName(categoryName string) model.CategoryEntity
	GetCategoryByParentId(parentPlainId string) []model.CategoryEntity
	InsertCategoryByEntity(newEntity model.CategoryEntity) string
	UpdateCategoryByEntity(plainId string) model.CategoryEntity
	DeleteCategoryByObjectId(plainId string) model.CategoryEntity
}

func init() {
	switch util.GetConfigByKey("db.type") {
	case "mongodb":
		INSTANCE = CategoryMongoDbMapper{}
	case "mysql":
		INSTANCE = CategoryMySqlMapper{}
	default:
		panic("database type not supported")
	}
}
