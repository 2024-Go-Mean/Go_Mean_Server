package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	// CORS 설정
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // 허용할 도메인을 설정
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// CORS 미들웨어로 라우터 감싸기
	handler := c.Handler(router)

	// 서버 시작
	fmt.Println("Server started on port 5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
