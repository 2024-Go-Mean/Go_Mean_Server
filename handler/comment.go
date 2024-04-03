package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"noah.io/ark/rest/models"
)

func init() {
	initMongoClient()
}

func initMongoClient() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.Background(), clientOptions)
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.ID = primitive.NewObjectID()

	collection := client.Database("test").Collection("comments")
	_, err = collection.InsertOne(context.Background(), comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add comment"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Comment added successfully"})
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// URL에서 게시물 ID를 가져옵니다.
	params := mux.Vars(r)
	worryID := params["worry_id"]

	collection := client.Database("test").Collection("comments")

	// MongoDB에서 해당 게시물 ID에 해당하는 모든 Comment를 검색합니다.
	cursor, err := collection.Find(context.Background(), bson.M{"worry_id": worryID})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var comments []models.Comment
	// 검색된 결과를 Comment 슬라이스에 디코딩합니다.
	for cursor.Next(context.Background()) {
		var comment models.Comment
		err := cursor.Decode(&comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comments = append(comments, comment)
	}

	// 결과를 JSON 형식으로 반환합니다.
	json.NewEncoder(w).Encode(comments)
}

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commentsID, _ := primitive.ObjectIDFromHex(params["comments_id"])

	var comment models.Comment
	json.NewDecoder(r.Body).Decode(&comment)
	comment.ID = commentsID

	collection := client.Database("test").Collection("comments")
	_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": commentsID}, comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update comment"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Comment updated successfully"})
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commentsID, _ := primitive.ObjectIDFromHex(params["comments_id"])

	collection := client.Database("test").Collection("comments")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": commentsID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to delete comment"})
		return
	}

	w.WriteHeader(http.StatusOK)
}
