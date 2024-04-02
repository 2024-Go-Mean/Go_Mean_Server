package routes

import (
	"github.com/gorilla/mux"
	"noah.io/ark/rest/handler"
)

func CategoryAPI(router *mux.Router) {
	router.HandleFunc("/category", handler.AddCategoryHandler).Methods("POST")
	router.HandleFunc("/category/{category_id}", handler.GetCategoryHandler).Methods("GET")
	router.HandleFunc("/category/{category_id}", handler.UpdateCategoryHandler).Methods("PATCH")
	router.HandleFunc("/category/{category_id}", handler.DeleteCategoryHandler).Methods("DELETE")
}
