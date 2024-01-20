package cash_flow_service

import (
	"errors"
	"reflect"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/emmettwoo/EMM-MoneyBox/util"
)

func IsDeleteFieldsConflicted(plainId, belongsDate string) bool {

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

func DeleteById(plainId string) (entity.CashFlowEntity, error) {

	var existCashFlowEntity = mapper.CashFlowCommonMapper.GetCashFlowByObjectId(plainId)
	if existCashFlowEntity.IsEmpty() {
		return entity.CashFlowEntity{}, errors.New("cash_flow not found")
	}

	existCashFlowEntity = mapper.CashFlowCommonMapper.DeleteCashFlowByObjectId(plainId)
	if existCashFlowEntity.IsEmpty() {
		return entity.CashFlowEntity{}, errors.New("cash_flow delete failed")
	}
	return existCashFlowEntity, nil
}

func DeleteByDate(belongsDate string) ([]entity.CashFlowEntity, error) {

	var deleteDate = util.FormatDateFromStringWithoutDash(belongsDate)
	if reflect.DeepEqual(deleteDate, time.Time{}) {
		return []entity.CashFlowEntity{}, errors.New("belongs_date error, try format like 19700101")
	}

	cashFlowList := mapper.CashFlowCommonMapper.DeleteCashFlowByBelongsDate(deleteDate)
	return cashFlowList, nil
}
