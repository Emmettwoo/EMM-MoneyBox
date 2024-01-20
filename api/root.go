package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(port int32) {
	r := mux.NewRouter()
	RegisterDemoRoute(r)
	RegisterCashRoute(r)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("API server is running on http://localhost%s\n", addr)
	http.ListenAndServe(addr, r)
}
