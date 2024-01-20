package controller

import (
	"fmt"
	"net/http"

	"github.com/emmettwoo/EMM-MoneyBox/controller/cash_flow_controller"
	"github.com/gorilla/mux"
)

func StartServer(port int32) {
	r := mux.NewRouter()
	registerCashRoute(r)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("API server is running on http://localhost%s\n", addr)
	http.ListenAndServe(addr, r)
}

func registerCashRoute(r *mux.Router) {
	r.HandleFunc("/api/cash/outcome", cash_flow_controller.CreateOutcome).Methods("POST")
	r.HandleFunc("/api/cash/income", cash_flow_controller.CreateIncome).Methods("POST")
	r.HandleFunc("/api/cash/{id}", cash_flow_controller.QueryById).Methods("GET")
	r.HandleFunc("/api/cash/date/{date}", cash_flow_controller.QueryByDate).Methods("GET")
	r.HandleFunc("/api/cash/{id}", cash_flow_controller.DeleteById).Methods("DELETE")
	r.HandleFunc("/api/cash/date/{date}", cash_flow_controller.DeleteByDate).Methods("DELETE")
}
