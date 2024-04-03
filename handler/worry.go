package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"noah.io/ark/rest/models"
)

func init() {
	initMongoClient()
}

func AddWorryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var worry models.Worries
	json.NewDecoder(r.Body).Decode(&worry)
	worry.ID = primitive.NewObjectID()

	collection := client.Database("test").Collection("worries")
	_, err := collection.InsertOne(context.Background(), worry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add worry"})
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Worry added successfully"})
}

func GetOneWorryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// URL에서 경로 매개변수 가져오기
	params := mux.Vars(r)
	worryID, err := primitive.ObjectIDFromHex(params["worry_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid worry ID"})
		return
	}

	// MongoDB에서 해당 ID의 데이터 가져오기
	collection := client.Database("test").Collection("worries")
	var worry models.Worries
	err = collection.FindOne(context.Background(), bson.M{"_id": worryID}).Decode(&worry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get worry"})
		return
	}

	// JSON 형식으로 응답 반환
	json.NewEncoder(w).Encode(worry)
}

func GetAllWorriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := client.Database("test").Collection("worries")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get worries"})
		log.Fatal(err)
		return
	}
	defer cursor.Close(context.Background())

	var worries []models.Worries
	for cursor.Next(context.Background()) {
		var worry models.Worries
		cursor.Decode(&worry)
		worries = append(worries, worry)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worries)
}

func UpdateWorryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	worryID, _ := primitive.ObjectIDFromHex(params["worry_id"])

	var worry models.Worries
	json.NewDecoder(r.Body).Decode(&worry)
	worry.ID = worryID

	collection := client.Database("test").Collection("worries")
	_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": worryID}, worry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update worry"})
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Worry updated successfully"})
}

func DeleteWorryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	worryID, _ := primitive.ObjectIDFromHex(params["worry_id"])

	collection := client.Database("test").Collection("worries")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": worryID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to delete worry"})
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
