package api

import (
	"net/http"

	"github.com/emmettwoo/EMM-MoneyBox/service/cash_flow_service"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/gorilla/mux"
)

// todo(emmett): move to model layer.
type CashFlowEntityDTO struct {
	BelongsDate  string  `json:"belongs_date"`
	CategoryName string  `json:"category_name"`
	Amount       float64 `json:"amount"`
	Description  string  `json:"description"`
}

// todo(emmett): move all register to route.go file (aloneside root.go).
func RegisterCashRoute(r *mux.Router) {
	r.HandleFunc("/api/cash/outcome", outcome).Methods("POST")
	r.HandleFunc("/api/cash/income", income).Methods("POST")
	r.HandleFunc("/api/cash/{id}", queryById).Methods("GET")
	r.HandleFunc("/api/cash/date/{date}", queryByDate).Methods("GET")
	r.HandleFunc("/api/cash/{id}", deleteById).Methods("DELETE")
	r.HandleFunc("/api/cash/date/{date}", deleteByDate).Methods("DELETE")
}

func queryById(w http.ResponseWriter, r *http.Request) {
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

func queryByDate(w http.ResponseWriter, r *http.Request) {
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

func deleteById(w http.ResponseWriter, r *http.Request) {
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

func deleteByDate(w http.ResponseWriter, r *http.Request) {
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

func outcome(w http.ResponseWriter, r *http.Request) {

	var requestBody CashFlowEntityDTO
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

func income(w http.ResponseWriter, r *http.Request) {

	var requestBody CashFlowEntityDTO
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
