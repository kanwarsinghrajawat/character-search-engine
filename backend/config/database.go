package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitMongoDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("ERROR: MONGO_URI not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB Connection Error: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB Ping Failed: %v", err)
	}

	fmt.Println("Connected to MongoDB at", mongoURI)

	DB = client.Database("rickmorty")

	createIndexes()
}

func createIndexes() {
	collection := DB.Collection("characters")

	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: "text"}},
		Options: options.Index().SetUnique(false),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create index on 'name' field: %v", err)
	}

	fmt.Println("Index on 'name' field created successfully")
}
