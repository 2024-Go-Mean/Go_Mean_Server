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
	"strconv"
)

func init() {
	initMongoClient()
}

func AddAdviceHandler(w http.ResponseWriter, r *http.Request) {
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

func GetAdviceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adviceID, _ := strconv.Atoi(params["advice_id"])

	collection := client.Database("test").Collection("advices")
	cursor, err := collection.Find(context.Background(), bson.M{"advice_id": adviceID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get advices"})
		log.Fatal(err)
		return
	}

	var advices []models.Advices
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var advice models.Advices
		cursor.Decode(&advice)
		advices = append(advices, advice)
	}

	json.NewEncoder(w).Encode(advices)
}

func UpdateAdviceHandler(w http.ResponseWriter, r *http.Request) {
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
