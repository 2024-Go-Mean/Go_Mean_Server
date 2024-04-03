package routes

import (
	"github.com/gorilla/mux"
	"noah.io/ark/rest/handler"
)

func WorryAPI(router *mux.Router) {
	router.HandleFunc("/worry", handler.AddWorryHandler).Methods("POST")
	router.HandleFunc("/worry/{worry_id}", handler.GetOneWorryHandler).Methods("GET")
	router.HandleFunc("/worry", handler.GetAllWorriesHandler).Methods("GET")
	router.HandleFunc("/worry/{worry_id}", handler.UpdateWorryHandler).Methods("PATCH")
	router.HandleFunc("/worry/{worry_id}", handler.DeleteWorryHandler).Methods("DELETE")
}
