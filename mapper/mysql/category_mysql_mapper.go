package mysql

import (
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
)

type CategoryMySqlMapper struct{}

func (CategoryMySqlMapper) GetCategoryByObjectId(plainId string) entity.CategoryEntity {

	util.Logger.Errorln("non-supported yet.")
	return entity.CategoryEntity{}
}

func (CategoryMySqlMapper) GetCategoryByName(categoryName string) entity.CategoryEntity {

	util.Logger.Errorln("non-supported yet.")
	return entity.CategoryEntity{}
}

func (CategoryMySqlMapper) InsertCategoryByEntity(newEntity entity.CategoryEntity) string {

	util.Logger.Errorln("non-supported yet.")
	return ""
}

func (CategoryMySqlMapper) UpdateCategoryByEntity(plainId string) entity.CategoryEntity {

	util.Logger.Errorln("non-supported yet.")
	return entity.CategoryEntity{}
}

func (CategoryMySqlMapper) DeleteCategoryByObjectId(plainId string) entity.CategoryEntity {

	util.Logger.Errorln("non-supported yet.")
	return entity.CategoryEntity{}
}
