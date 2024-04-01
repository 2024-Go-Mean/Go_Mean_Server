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

func ReadOneWorry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var worryCollection = getcollection.GetCollection(DB, "Worries")

	worryId := c.Param("Id")
	var result models.Worries

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(worryId)

	err := worryCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}
