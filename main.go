package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"noah.io/ark/rest/routes"
)

func main() {
	// 라우터 설정
	router := mux.NewRouter()
	routes.CommentAPI(router)

	// 서버 시작
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
