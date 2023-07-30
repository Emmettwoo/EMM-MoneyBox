package category_service

import (
	"errors"
	"fmt"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
)

func QueryService(plainId, categoryName string) error {

	if isQueryFieldsConflicted(plainId, categoryName) {
		return errors.New("should have one and only one query type")
	}

	if plainId != "" {
		queryById(plainId)
		return nil
	}

	if categoryName != "" {
		queryByName(categoryName)
		return nil
	}

	return errors.New("not supported query type")
}

func isQueryFieldsConflicted(plainId, name string) bool {

	// check if already one semi-optional field is filled
	var semiOptionalFieldFilledFlag = false

	// plain_id is not empty
	if plainId != "" {
		semiOptionalFieldFilledFlag = true
	}

	// category name is not empty
	if name != "" {
		if semiOptionalFieldFilledFlag {
			return true
		}
		semiOptionalFieldFilledFlag = true
	}

	// should have one and only one field filled
	return !semiOptionalFieldFilledFlag
}

func queryById(plainId string) {

	categoryEntity := mapper.CategoryCommonMapper.GetCategoryByObjectId(plainId)
	if categoryEntity.IsEmpty() {
		fmt.Println("category not found")
		return
	}
	fmt.Println("category ", 0, ": ", categoryEntity.ToString())
}

func queryByName(categoryName string) {

	categoryEntity := mapper.CategoryCommonMapper.GetCategoryByName(categoryName)
	if categoryEntity.IsEmpty() {
		fmt.Println("category not found")
		return
	}
	fmt.Println("category ", 0, ": ", categoryEntity.ToString())
}
