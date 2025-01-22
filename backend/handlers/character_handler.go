package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/config"
	"backend/models"
)

// GetCharacterByName retrieves characters based on name (case-insensitive search)
func GetCharacterByName(c *gin.Context) {
	name := c.Param("name")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10
	skip := (page - 1) * limit

	// MongoDB filter
	filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}}
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(skip))

	// Fetch data
	cursor, err := config.DB.Collection("characters").Find(context.TODO(), filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}
	defer cursor.Close(context.TODO())

	// Decode data
	var characters []models.Character
	if err = cursor.All(context.TODO(), &characters); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}

	// Response
	if len(characters) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No characters found", "count": 0})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":      len(characters),
		"characters": characters,
		"page":       page,
	})
}
