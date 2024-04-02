package handler

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"noah.io/ark/rest/models"
)

func init() {
	initMongoClient()
}

func AddWorryHandler(w http.ResponseWriter, r *http.Request) {
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

func GetWorryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	worryID, _ := strconv.Atoi(params["worry_id"])

	collection := client.Database("test").Collection("worries")
	cursor, err := collection.Find(context.Background(), bson.M{"worry_id": worryID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get worries"})
		log.Fatal(err)
		return
	}

	var worries []models.Worries
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var worry models.Worries
		cursor.Decode(&worry)
		worries = append(worries, worry)
	}

	json.NewEncoder(w).Encode(worries)
}

func UpdateWorryHandler(w http.ResponseWriter, r *http.Request) {
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
