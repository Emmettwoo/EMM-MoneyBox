package cash_flow_service

import (
	"errors"
	"fmt"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"reflect"
	"time"
)

func DeleteService(plainId string, belongsDate string) error {

	if isDeleteFieldsConflicted(plainId, belongsDate) {
		return errors.New("should have one and only one delete type")
	}

	if plainId != "" {
		return deleteById(plainId)
	}

	if belongsDate != "" {
		return deleteByDate(belongsDate)
	}

	return errors.New("not supported delete type")
}

func isDeleteFieldsConflicted(plainId string, belongsDate string) bool {

	// check if already one semi-optional field is filled
	var semiOptionalFieldFilledFlag = false

	// plain_id is not empty
	if plainId != "" {
		semiOptionalFieldFilledFlag = true
	}

	// belongs_date is not empty
	if belongsDate != "" {
		if semiOptionalFieldFilledFlag {
			return true
		}
		semiOptionalFieldFilledFlag = true
	}

	// should have one and only one field filled
	return !semiOptionalFieldFilledFlag
}

func deleteById(plainId string) error {

	var existCashFlowEntity = mapper.CashFlowCommonMapper.GetCashFlowByObjectId(plainId)
	if existCashFlowEntity.IsEmpty() {
		fmt.Println("cash_flow not found")
		return nil
	}

	existCashFlowEntity = mapper.CashFlowCommonMapper.DeleteCashFlowByObjectId(plainId)
	if existCashFlowEntity.IsEmpty() {
		return errors.New("cash_flow delete failed")
	}
	fmt.Println("cash_flow ", 0, ": ", existCashFlowEntity.ToString())
	return nil
}

func deleteByDate(belongsDate string) error {

	var deleteDate = util.FormatDateFromStringWithoutDash(belongsDate)
	if reflect.DeepEqual(deleteDate, time.Time{}) {
		return errors.New("belongs_date error, try format like 19700101")
	}

	cashFlowList := mapper.CashFlowCommonMapper.DeleteCashFlowByBelongsDate(deleteDate)
	if len(cashFlowList) == 0 {
		fmt.Println("the day's flow is empty")
		return nil
	}

	for index, cashFlow := range cashFlowList {
		fmt.Println("cash_flow ", index, ": ", cashFlow.ToString())
	}
	return nil
}
