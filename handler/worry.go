package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"noah.io/ark/rest/models"
	"os"
	"reflect"
)

func init() {
	initMongoClient()
}

func ChatGptFunc(userInput string) string {
	// OpenAI API 키를 설정합니다.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// API_KEY 환경 변수를 가져옵니다.
	apiKey := os.Getenv("API_KEY")

	// OpenAI 클라이언트를 생성합니다.
	client := openai.NewClient(apiKey)

	baseStr := "위에 대한 대답은 명확하며 반드시 발랄하게 ~요체를 쓰며, 친구같이 진심으로 위로해주며 취업, 학업, 인간관계, 건강, 금전, 개인 카테고리에 대한 고민을 듣고 친구가 친구에게 말하듯이 일상적인 말투로 진지하게 고민에 대한 조언을 스토리텔링으로 뻔하지 않게 답을 해줘"

	// 사용자의 입력을 OpenAI API로 전달하여 응답을 받습니다.
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userInput + baseStr,
				},
			},
		},
	)

	// 오류를 확인하고 출력합니다.
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	// 응답에서 첫 번째 선택의 내용을 출력합니다.
	fmt.Println("GPT-3.5 응답:", resp.Choices[0].Message.Content)
	fmt.Println("타입:", reflect.TypeOf(resp.Choices[0].Message.Content))
	return resp.Choices[0].Message.Content
}

func AddWorryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var worry models.Worries
	json.NewDecoder(r.Body).Decode(&worry)
	worry.ID = primitive.NewObjectID()

	// ChatGPT 응답 추가
	worry.AiAdvice = ChatGptFunc(worry.Content)

	collection := client.Database("test").Collection("worries")
	_, err := collection.InsertOne(context.Background(), worry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to add worry"})
		log.Fatal(err)
		return
	}

	// worry 객체를 반환
	json.NewEncoder(w).Encode(worry)
}

func GetOneWorryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	worryID, err := primitive.ObjectIDFromHex(params["worry_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid worry ID"})
		return
	}

	collection := client.Database("test").Collection("worries")
	var worry models.Worries
	err = collection.FindOne(context.Background(), bson.M{"_id": worryID}).Decode(&worry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get worry"})
		return
	}

	// 카테고리 정보를 가져오는 부분 추가
	categoryCollection := client.Database("test").Collection("categories")
	var category models.Categories
	err = categoryCollection.FindOne(context.Background(), bson.M{"_id": worry.CategoryId}).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to get category"})
		return
	}

	// worry에 category 정보를 추가
	type WorryWithCategory struct {
		models.Worries
		Category string `json:"category"`
	}

	worryWithCategory := WorryWithCategory{
		Worries:  worry,
		Category: category.Category,
	}

	// JSON 형식으로 응답 반환
	json.NewEncoder(w).Encode(worryWithCategory)
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
	worry.AiAdvice = ChatGptFunc(worry.Content)

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
