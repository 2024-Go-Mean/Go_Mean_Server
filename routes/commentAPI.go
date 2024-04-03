package routes

import (
	"github.com/gorilla/mux"
	"noah.io/ark/rest/handler"
)

// 주어진 라우터에 댓글 API 엔드포인트 등록
func CommentAPI(router *mux.Router) {
	// 댓글 추가 API 엔드포인트 등록
	router.HandleFunc("/comments", handler.AddCommentHandler).Methods("POST")

	// 댓글 불러오기 API 엔드포인트 등록
	router.HandleFunc("/comments/{worry_id}", handler.GetCommentsHandler).Methods("GET")

	// 댓글 수정하기 API 엔드포인트 등록
	router.HandleFunc("/comments/{comments_id}", handler.UpdateCommentHandler).Methods("PATCH")

	// 댓글 삭제하기 API 엔드포인트 등록
	router.HandleFunc("/comments/{comments_id}", handler.DeleteCommentHandler).Methods("DELETE")
}
