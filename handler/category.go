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

func AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Categories
	json.NewDecoder(r.Body).Decode(&category)
	category.ID = primitive.NewObjectID()

	collection := client.Database("test").Collection("categories")
	_, err := collection.InsertOne(context.Background(), category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add category"})
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Category added successfully"})
}

func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, _ := strconv.Atoi(params["category_id"])

	collection := client.Database("test").Collection("categories")
	cursor, err := collection.Find(context.Background(), bson.M{"category_id": categoryID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get categories"})
		log.Fatal(err)
		return
	}

	var categories []models.Categories
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var category models.Categories
		cursor.Decode(&category)
		categories = append(categories, category)
	}

	json.NewEncoder(w).Encode(categories)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, _ := primitive.ObjectIDFromHex(params["category_id"])

	var category models.Categories
	json.NewDecoder(r.Body).Decode(&category)
	category.ID = categoryID

	collection := client.Database("test").Collection("categories")
	_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": categoryID}, category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update category"})
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Category updated successfully"})
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, _ := primitive.ObjectIDFromHex(params["category_id"])

	collection := client.Database("test").Collection("categories")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": categoryID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to delete category"})
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
