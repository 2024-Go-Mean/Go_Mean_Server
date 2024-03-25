package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Comment 모델 정의
type Comment struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	WorryID  int                `json:"worry_id"`
	Nickname string             `json:"nickname,omitempty"`
	Comment  string             `json:"comment"`
}

// MongoDB 클라이언트
var client *mongo.Client

// MongoDB 연결 초기화
func initMongoClient() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// MongoDB 클라이언트 생성
	newClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// MongoDB 연결 테스트
	err = newClient.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	// 기존 클라이언트를 새로운 클라이언트로 업데이트
	client = newClient

	fmt.Println("Connected to MongoDB!")
	return nil
}

// 댓글 추가 API
func addCommentHandler(w http.ResponseWriter, r *http.Request) {
	// MongoDB 클라이언트가 nil이면 초기화 시도
	if client == nil {
		if err := initMongoClient(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Failed to connect to MongoDB"})
			return
		}
	}

	var comment Comment
	json.NewDecoder(r.Body).Decode(&comment)

	// 새로운 ObjectID 생성
	comment.ID = primitive.NewObjectID()

	// MongoDB에 댓글 추가
	collection := client.Database("test").Collection("comments")
	_, err := collection.InsertOne(context.Background(), comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add comment"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Comment added successfully"})
}

// 댓글 불러오기 API
func getCommentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	worryID, _ := strconv.Atoi(params["worry_id"])

	// MongoDB에서 해당 쓰레기의 댓글 불러오기
	collection := client.Database("test").Collection("comments")
	cursor, err := collection.Find(context.Background(), bson.M{"worryid": worryID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get comments"})
		return
	}

	var comments []Comment
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var comment Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}

	json.NewEncoder(w).Encode(comments)
}

// 댓글 수정하기 API
func updateCommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentsID, _ := primitive.ObjectIDFromHex(params["comments_id"])

	var comment Comment
	json.NewDecoder(r.Body).Decode(&comment)
	comment.ID = commentsID // URL에서 가져온 ID로 설정

	// MongoDB에서 해당 ID의 댓글을 업데이트
	collection := client.Database("test").Collection("comments")
	_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": commentsID}, comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update comment"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Comment updated successfully"})
}

// 댓글 삭제하기 API
func deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentsID, _ := primitive.ObjectIDFromHex(params["comments_id"])

	// MongoDB에서 해당 ID의 댓글 삭제
	collection := client.Database("test").Collection("comments")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": commentsID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to delete comment"})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	// MongoDB 클라이언트 초기화
	if err := initMongoClient(); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// 라우터 설정
	router := mux.NewRouter()

	// 댓글 추가 API 엔드포인트 등록
	router.HandleFunc("/comment", addCommentHandler).Methods("POST")

	// 댓글 불러오기 API 엔드포인트 등록
	router.HandleFunc("/comment/{worry_id}", getCommentsHandler).Methods("GET")

	// 댓글 수정하기 API 엔드포인트 등록
	router.HandleFunc("/comment/{comments_id}", updateCommentHandler).Methods("PATCH")

	// 댓글 삭제하기 API 엔드포인트 등록
	router.HandleFunc("/comment/{comments_id}", deleteCommentHandler).Methods("DELETE")

	// 서버 시작
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
