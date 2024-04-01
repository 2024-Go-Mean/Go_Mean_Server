package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB 연결 정보
const (
	uri        = "mongodb://localhost:27017"
	database   = "gomean"
	collection = "gomean"
)

// Person 구조체 정의
type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	app := fiber.New()

	//run database
	//configs.ConnectDB()

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// 데이터베이스 및 컬렉션 선택
	db := client.Database(database)
	col := db.Collection(collection)

	// 데이터 삽입 예제
	person := Person{"John", 30, "New York"}
	insertResult, err := col.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// 데이터 조회 예제
	var result Person
	err = col.FindOne(ctx, map[string]string{"name": "John"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found a single document: ", result)

	app.Listen(":8000")
}
