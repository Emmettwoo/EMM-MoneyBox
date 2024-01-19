package api

import (
	"net/http"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/gorilla/mux"
)

type RequestBody struct {
	Name string `json:"name"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RegisterDemoRoute(r *mux.Router) {
	r.HandleFunc("/api/demo", helloWorld).Methods("GET")
	r.HandleFunc("/api/demo", helloWorld).Methods("POST")
	r.HandleFunc("/api/demo/{name}", helloWorld).Methods("GET")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		name = "World"
	}

	if r.Method == http.MethodPost {
		// Parse JSON from request body
		var requestBody RequestBody
		err := util.ParseJSONRequest(r, &requestBody)
		// If there is an error, include it in the response
		if err != nil {
			util.ComposeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if requestBody.Name != "" {
			name = requestBody.Name
		}
	}

	// Construct the response
	response := Response{
		Status:  "success",
		Message: "Hello " + name + "!",
	}
	// Use the utility function to write the JSON response
	util.ComposeJSONResponse(w, http.StatusOK, response)
}
