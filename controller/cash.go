package controller

import (
	"net/http"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/service/cash_flow_service"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/gorilla/mux"
)

func CashQueryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plainId := vars["id"]
	if plainId == "" {
		util.ComposeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "id is empty"})
	}
	cashFlowEntity, err := cash_flow_service.QueryById(plainId)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": err.Error()})
		return
	}
	util.ComposeJSONResponse(w, http.StatusOK, cashFlowEntity)
}

func CashQueryByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	belongsDate := vars["date"]
	if belongsDate == "" {
		util.ComposeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "date is empty"})
	}
	cashFlowEntityList, err := cash_flow_service.QueryByDate(belongsDate)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": err.Error()})
		return
	}
	util.ComposeJSONResponse(w, http.StatusOK, cashFlowEntityList)
}

func CashDeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plainId := vars["id"]
	if plainId == "" {
		util.ComposeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "id is empty"})
	}
	cashFlowEntity, err := cash_flow_service.DeleteById(plainId)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": err.Error()})
		return
	}
	util.ComposeJSONResponse(w, http.StatusOK, cashFlowEntity)
}

func CashDeleteByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	if date == "" {
		util.ComposeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "date is empty"})
	}
	cashFlowEntityList, err := cash_flow_service.DeleteByDate(date)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": err.Error()})
		return
	}
	util.ComposeJSONResponse(w, http.StatusOK, cashFlowEntityList)
}

func CashSaveOutcome(w http.ResponseWriter, r *http.Request) {

	var requestBody model.CashFlowDTO
	err := util.ParseJSONRequest(r, &requestBody)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if !cash_flow_service.IsOutcomeRequiredFiledSatisfied(requestBody.CategoryName, requestBody.Amount) {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": "some required fields are empty"})
		return
	}

	cashFlowEntity, err := cash_flow_service.SaveOutcome(requestBody.BelongsDate, requestBody.CategoryName, requestBody.Amount, requestBody.Description)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": err.Error()})
		return
	}
	util.ComposeJSONResponse(w, http.StatusOK, cashFlowEntity)
}

func CashSaveIncome(w http.ResponseWriter, r *http.Request) {

	var requestBody model.CashFlowDTO
	err := util.ParseJSONRequest(r, &requestBody)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if !cash_flow_service.IsOutcomeRequiredFiledSatisfied(requestBody.CategoryName, requestBody.Amount) {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": "some required fields are empty"})
		return
	}

	cashFlowEntity, err := cash_flow_service.SaveIncome(requestBody.BelongsDate, requestBody.CategoryName, requestBody.Amount, requestBody.Description)
	if err != nil {
		util.ComposeJSONResponse(w, http.StatusOK, map[string]string{"error": err.Error()})
		return
	}
	util.ComposeJSONResponse(w, http.StatusOK, cashFlowEntity)
}
