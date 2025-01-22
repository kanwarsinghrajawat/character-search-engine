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

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB")
	authSource := os.Getenv("MONGO_AUTH_SOURCE")

	if username == "" || password == "" || host == "" || port == "" || dbName == "" || authSource == "" {
		log.Fatal("ERROR: Missing MongoDB environment variables")
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		username, password, host, port, dbName, authSource)

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB Connection Error: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB Ping Failed: %v", err)
	}

	fmt.Println("Connected to MongoDB at", host)

	DB = client.Database(dbName)

	createIndexes()
}

func createIndexes() {
	collection := DB.Collection("characters")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: "text"}}, 
		Options: options.Index().SetUnique(false),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create index on 'name' field: %v", err)
	}

	fmt.Println("Index on 'name' field created successfully")
}
