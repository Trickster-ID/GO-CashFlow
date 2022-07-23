package config

import (
	"github.com/gorilla/mux"
)

// var(
// 	cashflowrepo	repository.CashFlowRepo		=
// )

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", nil).Methods("POST")
	// r.HandleFunc("/", nil).Methods("GET")
	// r.HandleFunc("/:id", nil).Methods("PUT")

	// r.HandleFunc("/out", nil).Methods("POST")
	// r.HandleFunc("/out/:id", nil).Methods("PUT")

	// r.HandleFunc("/:id", nil).Methods("DELETE")

	return r
}
