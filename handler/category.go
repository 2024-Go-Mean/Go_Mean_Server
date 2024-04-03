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

func GetOneCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// URL에서 경로 매개변수 가져오기
	params := mux.Vars(r)
	categoryID, err := primitive.ObjectIDFromHex(params["category_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid category ID"})
		return
	}

	// MongoDB에서 해당 ID의 데이터 가져오기
	collection := client.Database("test").Collection("categories")
	var category models.Categories
	err = collection.FindOne(context.Background(), bson.M{"_id": categoryID}).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get category"})
		return
	}

	// JSON 형식으로 응답 반환
	json.NewEncoder(w).Encode(category)
}

func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("test").Collection("categories")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get categories"})
		log.Fatal(err)
		return
	}
	defer cursor.Close(context.Background())

	var categories []models.Categories
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
