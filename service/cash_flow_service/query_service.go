package cash_flow_service

import (
	"errors"
	"fmt"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"reflect"
	"time"
)

func QueryService(plainId string, belongsDate string, exactDescription string, fuzzyDescription string) error {

	if isQueryFieldsConflicted(plainId, belongsDate, exactDescription, fuzzyDescription) {
		return errors.New("should have one and only one query type")
	}

	if plainId != "" {
		queryById(plainId)
		return nil
	}
	if belongsDate != "" {
		return queryByDate(belongsDate)
	}
	if exactDescription != "" {
		queryByExactDescription(exactDescription)
		return nil
	}
	if fuzzyDescription != "" {
		queryByFuzzyDescription(fuzzyDescription)
		return nil
	}

	return errors.New("not supported query type")
}

func isQueryFieldsConflicted(plainId string, belongsDate string, exactDescription string, fuzzyDescription string) bool {

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

	// exact_description is not empty
	if exactDescription != "" {
		if semiOptionalFieldFilledFlag {
			return true
		}
		semiOptionalFieldFilledFlag = true
	}

	// fuzzy_description is not empty
	if fuzzyDescription != "" {
		if semiOptionalFieldFilledFlag {
			return true
		}
		semiOptionalFieldFilledFlag = true
	}

	// should have one and only one field filled
	return !semiOptionalFieldFilledFlag
}

func queryById(plainId string) {

	cashFlowEntity := mapper.CashFlowCommonMapper.GetCashFlowByObjectId(plainId)
	if cashFlowEntity.IsEmpty() {
		fmt.Println("cash_flow not found")
		return
	}
	fmt.Println("cash_flow ", 0, ": ", cashFlowEntity.ToString())
}

func queryByDate(belongsDate string) error {

	var queryDate = util.FormatDateFromString(belongsDate)
	if reflect.DeepEqual(queryDate, time.Time{}) {
		return errors.New("belongs_date error, try format like 19700101")
	}

	matchedCashFlowList := mapper.CashFlowCommonMapper.GetCashFlowsByBelongsDate(queryDate)
	if len(matchedCashFlowList) == 0 {
		fmt.Println("the day's flow is empty")
		return nil
	}

	for index, cashFlowEntity := range matchedCashFlowList {
		fmt.Println("cash_flow ", index, ": ", cashFlowEntity.ToString())
	}
	return nil
}

func queryByExactDescription(exactDescription string) {

	matchedCashFlowList := mapper.CashFlowCommonMapper.GetCashFlowsByExactDesc(exactDescription)
	if len(matchedCashFlowList) == 0 {
		fmt.Println("no matched cash_flows")
		return
	}

	for index, cashFlowEntity := range matchedCashFlowList {
		fmt.Println("cash_flow ", index, ": ", cashFlowEntity.ToString())
	}
}

func queryByFuzzyDescription(fuzzyDescription string) {

	matchedCashFlowList := mapper.CashFlowCommonMapper.GetCashFlowsByFuzzyDesc(fuzzyDescription)
	if len(matchedCashFlowList) == 0 {
		fmt.Println("no matched cash_flows")
		return
	}

	for index, cashFlowEntity := range matchedCashFlowList {
		fmt.Println("cash_flow ", index, ": ", cashFlowEntity.ToString())
	}
}
