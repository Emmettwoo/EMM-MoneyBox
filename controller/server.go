package controller

import (
	"fmt"
	"net/http"

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
	r.HandleFunc("/api/cash/outcome", CashSaveOutcome).Methods("POST")
	r.HandleFunc("/api/cash/income", CashSaveIncome).Methods("POST")
	r.HandleFunc("/api/cash/{id}", CashQueryById).Methods("GET")
	r.HandleFunc("/api/cash/date/{date}", CashQueryByDate).Methods("GET")
	r.HandleFunc("/api/cash/{id}", CashDeleteById).Methods("DELETE")
	r.HandleFunc("/api/cash/date/{date}", CashDeleteByDate).Methods("DELETE")
}
