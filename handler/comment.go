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
	"time"
)

func init() {
	initMongoClient()
}

func initMongoClient() {
	clientOptions := options.Client().ApplyURI("mongodb://13.125.4.74:3002")
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
	comment.Timestamp = time.Now() // 현재 시간을 설정합니다.

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

	params := mux.Vars(r)
	worryID := params["worry_id"]

	collection := client.Database("test").Collection("comments")

	cursor, err := collection.Find(context.Background(), bson.M{"worry_id": worryID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var comments []models.Comment
	for cursor.Next(context.Background()) {
		var comment models.Comment
		err := cursor.Decode(&comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commentsID, err := primitive.ObjectIDFromHex(params["comments_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var comment models.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment.ID = commentsID

	collection := client.Database("test").Collection("comments")
	result, err := collection.ReplaceOne(context.Background(), bson.M{"_id": commentsID}, comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.ModifiedCount == 0 {
		http.Error(w, "No matching comment found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Comment updated successfully"})
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commentsID, err := primitive.ObjectIDFromHex(params["comments_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("test").Collection("comments")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": commentsID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "No matching comment found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
