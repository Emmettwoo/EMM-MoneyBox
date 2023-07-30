package category_service

import (
	"errors"
	"fmt"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
)

func DeleteService(plainId, categoryName string) error {

	if isDeleteFieldsConflicted(plainId, categoryName) {
		return errors.New("should have one and only one delete type")
	}

	if plainId != "" {
		return deleteById(plainId)
	}

	if categoryName != "" {
		return deleteByName(categoryName)
	}

	return errors.New("not supported delete type")
}

func isDeleteFieldsConflicted(plainId, categoryName string) bool {

	// check if already one semi-optional field is filled
	var semiOptionalFieldFilledFlag = false

	// plain_id is not empty
	if plainId != "" {
		semiOptionalFieldFilledFlag = true
	}

	// category_name is not empty
	if categoryName != "" {
		if semiOptionalFieldFilledFlag {
			return true
		}
		semiOptionalFieldFilledFlag = true
	}

	// should have one and only one field filled
	return !semiOptionalFieldFilledFlag
}

func deleteById(plainId string) error {

	var existCategoryEntity = mapper.CategoryCommonMapper.GetCategoryByObjectId(plainId)
	if existCategoryEntity.IsEmpty() {
		fmt.Println("category not found")
		return nil
	}

	existCategoryEntity = mapper.CategoryCommonMapper.DeleteCategoryByObjectId(plainId)
	if existCategoryEntity.IsEmpty() {
		return errors.New("category delete failed")
	}
	fmt.Println("category ", 0, ": ", existCategoryEntity.ToString())
	return nil
}

func deleteByName(categoryName string) error {

	var existCategoryEntity = mapper.CategoryCommonMapper.GetCategoryByName(categoryName)
	if existCategoryEntity.IsEmpty() {
		fmt.Println("category not found")
		return nil
	}
	existCategoryEntity = mapper.CategoryCommonMapper.DeleteCategoryByObjectId(existCategoryEntity.Id.Hex())
	if existCategoryEntity.IsEmpty() {
		return errors.New("category delete failed")
	}
	fmt.Println("category ", 0, ": ", existCategoryEntity.ToString())
	return nil
}
