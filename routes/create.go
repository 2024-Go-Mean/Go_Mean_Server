package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	getcollection "noah.io/ark/rest/collection"
	database "noah.io/ark/rest/databases"
	"noah.io/ark/rest/models"
	"time"
)

func CreateWorry(c *gin.Context) {
	var DB = database.ConnectDB()
	var worryCollection = getcollection.GetCollection(DB, "Worries")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	worry := new(models.Worries)
	defer cancel()

	if err := c.BindJSON(&worry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	worryPayload := models.Worries{
		Id:         primitive.NewObjectID(),
		Title:      worry.Title,
		Content:    worry.Content,
		Nickname:   worry.Nickname,
		AiAdviceId: worry.AiAdviceId,
		CategoryId: worry.CategoryId,
		CreatedAt:  worry.CreatedAt,
	}

	result, err := worryCollection.InsertOne(ctx, worryPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Worry Posted successfully", "Data": map[string]interface{}{"data": result}})
}
