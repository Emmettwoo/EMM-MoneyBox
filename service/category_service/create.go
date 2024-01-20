package category_service

import (
	"errors"
	"fmt"

	"github.com/emmettwoo/EMM-MoneyBox/mapper/category_mapper"
	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateService(parentPlainId, categoryName string) error {

	if !isCreateRequiredFiledSatisfied(categoryName) {
		return errors.New("some required fields are empty")
	}

	var categoryEntity = model.CategoryEntity{
		ParentId: primitive.NilObjectID,
		Name:     categoryName,
	}
	if parentPlainId != "" {
		categoryEntity.ParentId = util.Convert2ObjectId(parentPlainId)
	}

	var newCategoryPlainId = category_mapper.INSTANCE.InsertCategoryByEntity(categoryEntity)
	if newCategoryPlainId == "" {
		return errors.New("category create failed")
	}

	var newCategoryEntity = category_mapper.INSTANCE.GetCategoryByObjectId(newCategoryPlainId)
	fmt.Println("category ", 0, ": ", newCategoryEntity.ToString())
	return nil
}

func isCreateRequiredFiledSatisfied(categoryName string) bool {

	return categoryName != ""
}
