package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	getcollection "noah.io/ark/rest/collection"
	database "noah.io/ark/rest/databases"
	"noah.io/ark/rest/models"
	"time"
)

func UpdateWorry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var worryCollection = getcollection.GetCollection(DB, "Worries")

	worryId := c.Param("Id")
	var worry models.Worries

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(worryId)

	if err := c.BindJSON(&worry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{
		"title":      worry.Title,
		"content":    worry.Content,
		"nickname":   worry.Nickname,
		"aiAdviceId": worry.AiAdviceId,
		"categoryId": worry.CategoryId,
		"createdAt":  worry.CreatedAt,
	}

	result, err := worryCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}
