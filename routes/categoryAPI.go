package routes

import (
	"github.com/gorilla/mux"
	"noah.io/ark/rest/handler"
)

func CategoryAPI(router *mux.Router) {
	router.HandleFunc("/categories", handler.AddCategoryHandler).Methods("POST")
	router.HandleFunc("/categories/{category_id}", handler.GetOneCategoryHandler).Methods("GET")
	router.HandleFunc("/categories", handler.GetAllCategoriesHandler).Methods("GET")
	router.HandleFunc("/categories/{category_id}", handler.UpdateCategoryHandler).Methods("PATCH")
	router.HandleFunc("/categories/{category_id}", handler.DeleteCategoryHandler).Methods("DELETE")
}
