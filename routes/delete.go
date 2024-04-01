package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	getcollection "noah.io/ark/rest/collection"
	database "noah.io/ark/rest/databases"
	"time"
)

func DeleteWorry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	worryId := c.Param("Id")

	var worryCollection = getcollection.GetCollection(DB, "Worries")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(worryId)
	result, err := worryCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Worry deleted successfully", "Data": res})
}
