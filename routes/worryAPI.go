package routes

import (
	"github.com/gorilla/mux"
	"noah.io/ark/rest/handler"
)

func WorryAPI(router *mux.Router) {
	router.HandleFunc("/worries", handler.AddWorryHandler).Methods("POST")
	router.HandleFunc("/worries/{worry_id}", handler.GetOneWorryHandler).Methods("GET")
	router.HandleFunc("/worries", handler.GetAllWorriesHandler).Methods("GET")
	router.HandleFunc("/worries/{worry_id}", handler.UpdateWorryHandler).Methods("PATCH")
	router.HandleFunc("/worries/{worry_id}", handler.DeleteWorryHandler).Methods("DELETE")
}
