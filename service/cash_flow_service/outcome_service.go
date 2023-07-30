package cash_flow_service

import (
	"errors"
	"fmt"
	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/shopspring/decimal"
	"time"
)

func OutcomeService(belongsDate string, categoryName string, amount float64, description string) error {

	if !isOutcomeRequiredFiledSatisfied(categoryName, amount) {
		return errors.New("some required fields are empty")
	}

	// 取小數點後兩位
	amount, _ = decimal.NewFromFloat(amount).Round(2).Float64()

	// 必填參數: 類別
	categoryEntity := mapper.CategoryCommonMapper.GetCategoryByName(categoryName)
	if categoryEntity.IsEmpty() {
		fmt.Println("category does not exist")
		return nil
	}

	// 選填參數: 日期（默認當天）
	date := util.FormatDateFromString(util.FormatDateToString(time.Now()))
	if belongsDate != "" {
		util.FormatDateFromString(belongsDate)
	}

	newCashFlowId := mapper.CashFlowCommonMapper.InsertCashFlowByEntity(entity.CashFlowEntity{
		CategoryId:  categoryEntity.Id,
		BelongsDate: date,
		FlowType:    "OUTCOME",
		Amount:      amount,
		Description: description,
	})
	if newCashFlowId == "" {
		return errors.New("cash_flow create failed")
	}

	var newCashFlow = mapper.CashFlowCommonMapper.GetCashFlowByObjectId(newCashFlowId)
	fmt.Println("cash_flow ", 0, ": ", newCashFlow.ToString())
	return nil
}

func isOutcomeRequiredFiledSatisfied(categoryName string, amount float64) bool {

	if categoryName == "" {
		return false
	}
	if amount == 0 {
		return false
	}

	return true
}
