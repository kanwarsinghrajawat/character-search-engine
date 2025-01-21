package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options" 

	"backend/config"
	"backend/models"
)

func GetCharacterByName(c *gin.Context) {
    name := c.Param("name")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit := 10
    skip := (page - 1) * limit

    var characters []models.Character
    filter := bson.M{"name": bson.M{"$regex": fmt.Sprintf(".*%s.*", name), "$options": "i"}}
    opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(skip))

    cursor, err := config.DB.Collection("characters").Find(context.TODO(), filter, opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching characters"})
        return
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var character models.Character
        if err := cursor.Decode(&character); err == nil {
            characters = append(characters, character)
        }
    }

    if len(characters) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No characters found", "count": 0})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "count":      len(characters),
        "characters": characters,
        "page":       page,
    })
}
