package main

import (
	"fmt"
	"log"
	"net/http"
	_ "noah.io/ark/rest/handler"
	"noah.io/ark/rest/routes"

	"github.com/gorilla/mux"
)

func main() {
	// 라우터 설정
	router := mux.NewRouter()

	// 댓글 추가 API 엔드포인트 등록
	router.HandleFunc("/comment", routes.AddCommentHandler).Methods("POST")

	// 댓글 불러오기 API 엔드포인트 등록
	router.HandleFunc("/comment/{worry_id}", routes.GetCommentsHandler).Methods("GET")

	// 댓글 수정하기 API 엔드포인트 등록
	router.HandleFunc("/comment/{comments_id}", routes.UpdateCommentHandler).Methods("PATCH")

	// 댓글 삭제하기 API 엔드포인트 등록
	router.HandleFunc("/comment/{comments_id}", routes.DeleteCommentHandler).Methods("DELETE")

	// 서버 시작
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
