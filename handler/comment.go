package handler

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	json.NewDecoder(r.Body).Decode(&comment)
	comment.ID = primitive.NewObjectID()

	collection := client.Database("test").Collection("comments")
	_, err := collection.InsertOne(context.Background(), comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add comment"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Comment added successfully"})
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	worryID, _ := strconv.Atoi(params["worry_id"])

	collection := client.Database("test").Collection("comments")
	cursor, err := collection.Find(context.Background(), bson.M{"worryid": worryID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get comments"})
		return
	}

	var comments []models.Comment
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var comment models.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}

	json.NewEncoder(w).Encode(comments)
}

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
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
