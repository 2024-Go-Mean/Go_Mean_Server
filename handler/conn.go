package handler

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	InitMongoClient()
}

func InitMongoClient() {
	clientOptions := options.Client().ApplyURI("mongodb://13.125.4.74:3002")
	client, _ = mongo.Connect(context.Background(), clientOptions)
	if err := client.Ping(context.Background(), nil); err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
	}
}
