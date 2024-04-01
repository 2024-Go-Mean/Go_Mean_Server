package routes

import (
	"github.com/gorilla/mux"
	"noah.io/ark/rest/handler"
)

func WorryAPI(router *mux.Router) {
	// 조언 추가 API 엔드포인트 등록
	router.HandleFunc("/worry", handler.AddWorryHandler).Methods("POST")

	// 조언 가져오기 API 엔드포인트 등록
	router.HandleFunc("/worry/{worry_id}", handler.GetWorryHandler).Methods("GET")

	// 조언 수정하기 API 엔드포인트 등록
	router.HandleFunc("/worry/{worry_id}", handler.UpdateWorryHandler).Methods("PATCH")

	// 조언 삭제하기 API 엔드포인트 등록
	router.HandleFunc("/worry/{worry_id}", handler.DeleteWorryHandler).Methods("DELETE")
}
