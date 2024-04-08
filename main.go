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
	routes.WorryAPI(router)
	routes.CategoryAPI(router)

	// 서버 시작
	fmt.Println("Server started on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
