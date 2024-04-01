package routes

import (
	"github.com/gorilla/mux"
	"noah.io/ark/rest/handler"
)

func AdviceAPI(router *mux.Router) {
	// 조언 추가 API 엔드포인트 등록
	router.HandleFunc("/advice", handler.AddAdviceHandler).Methods("POST")

	// 조언 가져오기 API 엔드포인트 등록
	router.HandleFunc("/advice/{advice_id}", handler.GetAdviceHandler).Methods("GET")

	// 조언 수정하기 API 엔드포인트 등록
	router.HandleFunc("/advice/{advice_id}", handler.UpdateAdviceHandler).Methods("PATCH")

	// 조언 삭제하기 API 엔드포인트 등록
	router.HandleFunc("/advice/{advice_id}", handler.DeleteAdviceHandler).Methods("DELETE")
}
