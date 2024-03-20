package main

import (
	"encoding/json"
	"net/http"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// 4. 요청이 들어온 Response Header에 ContentType을 추가하고 전달받은 HandleFunc타입의 함수에 ResponseWriter와 Request를 넘겨준다.
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// 들어오는 요청의 Response Header에 Content-type을 json으로 설정해준다.
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")

		// 전달받은 http.Handler 호출
		next.ServeHTTP(writer, request)
	})
}

func main() {

	// 1. 새로운 mux를 만든다.
	mux := http.NewServeMux()

	// 2. 기존에 만들어놓은 HandleFunc를 HandlerFunc로 변경 ("/users" 삭제)
	// 내가 원하는 경로를 함수와 연결시킬 수 있다.
	userHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet: // 조회
			json.NewEncoder(writer).Encode(users)
		case http.MethodPost: // 삽입
			var user User
			json.NewDecoder(request.Body).Decode(&user) // 디코딩

			users[user.Email] = &user
			json.NewEncoder(writer).Encode(user)
		}
	})

	// 3. 만들어놓은 미들웨어에 파라미터로 넘긴다. ("/users"는 이 때 사용)
	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":8000", mux)
}
