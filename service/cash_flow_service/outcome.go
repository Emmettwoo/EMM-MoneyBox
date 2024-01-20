package cash_flow_service

import (
	"errors"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/mapper/cash_flow_mapper"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/category_mapper"
	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/shopspring/decimal"
)

func SaveOutcome(belongsDate, categoryName string, amount float64, description string) (model.CashFlowEntity, error) {

	// 取小數點後兩位
	amount, _ = decimal.NewFromFloat(amount).Round(2).Float64()

	// 必填參數: 類別
	categoryEntity := category_mapper.INSTANCE.GetCategoryByName(categoryName)
	if categoryEntity.IsEmpty() {
		return model.CashFlowEntity{}, errors.New("category does not exist")
	}

	// 選填參數: 日期（默認當天）
	date := util.FormatDateFromStringWithoutDash(util.FormatDateToStringWithoutDash(time.Now()))
	if belongsDate != "" {
		date = util.FormatDateFromStringWithoutDash(belongsDate)
	}

	newCashFlowId := cash_flow_mapper.INSTANCE.InsertCashFlowByEntity(model.CashFlowEntity{
		CategoryId:  categoryEntity.Id,
		BelongsDate: date,
		FlowType:    "OUTCOME",
		Amount:      amount,
		Description: description,
	})
	if newCashFlowId == "" {
		return model.CashFlowEntity{}, errors.New("cash_flow create failed")
	}

	var newCashFlow = cash_flow_mapper.INSTANCE.GetCashFlowByObjectId(newCashFlowId)
	return newCashFlow, nil
}

func IsOutcomeRequiredFiledSatisfied(categoryName string, amount float64) bool {

	if categoryName == "" {
		return false
	}
	if amount == 0 {
		return false
	}

	return true
}
