package category_service

import (
	"errors"
	"fmt"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateService(parentPlainId, categoryName string) error {

	if !isCreateRequiredFiledSatisfied(categoryName) {
		return errors.New("some required fields are empty")
	}

	var categoryEntity = entity.CategoryEntity{
		ParentId: primitive.NilObjectID,
		Name:     categoryName,
	}
	if parentPlainId != "" {
		categoryEntity.ParentId = util.Convert2ObjectId(parentPlainId)
	}

	var newCategoryPlainId = mapper.CategoryCommonMapper.InsertCategoryByEntity(categoryEntity)
	if newCategoryPlainId == "" {
		return errors.New("category create failed")
	}

	var newCategoryEntity = mapper.CategoryCommonMapper.GetCategoryByObjectId(newCategoryPlainId)
	fmt.Println("category ", 0, ": ", newCategoryEntity.ToString())
	return nil
}

func isCreateRequiredFiledSatisfied(categoryName string) bool {

	return categoryName != ""
}
