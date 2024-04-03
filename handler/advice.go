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

func AddAdviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var advice models.Advices
	json.NewDecoder(r.Body).Decode(&advice)
	advice.ID = primitive.NewObjectID()

	collection := client.Database("test").Collection("advices")
	_, err := collection.InsertOne(context.Background(), advice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add advice"})
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Advice added successfully"})
}

func GetOneAdviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// URL에서 경로 매개변수 가져오기
	params := mux.Vars(r)
	adviceID, err := primitive.ObjectIDFromHex(params["advice_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid advice ID"})
		return
	}

	// MongoDB에서 해당 ID의 데이터 가져오기
	collection := client.Database("test").Collection("advices")
	var advice models.Advices
	err = collection.FindOne(context.Background(), bson.M{"_id": adviceID}).Decode(&advice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get advice"})
		return
	}

	// JSON 형식으로 응답 반환
	json.NewEncoder(w).Encode(advice)
}

func GetAllAdvicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := client.Database("test").Collection("advices")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get advices"})
		log.Fatal(err)
		return
	}
	defer cursor.Close(context.Background())

	var advices []models.Advices
	for cursor.Next(context.Background()) {
		var advice models.Advices
		cursor.Decode(&advice)
		advices = append(advices, advice)
	}

	json.NewEncoder(w).Encode(advices)
}

func UpdateAdviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	adviceID, _ := primitive.ObjectIDFromHex(params["advice_id"])

	var advice models.Advices
	json.NewDecoder(r.Body).Decode(&advice)
	advice.ID = adviceID

	collection := client.Database("test").Collection("advices")
	_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": adviceID}, advice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update advice"})
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Advice updated successfully"})
}

func DeleteAdviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	adviceID, _ := primitive.ObjectIDFromHex(params["advice_id"])

	collection := client.Database("test").Collection("advices")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": adviceID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to delete advice"})
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
