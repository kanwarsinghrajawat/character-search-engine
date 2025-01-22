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

// Load environment variables
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}
}

// Get MongoDB connection URI from environment variables
func getMongoURI() string {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB")
	authSource := os.Getenv("MONGO_AUTH_SOURCE")

	if username == "" || password == "" || host == "" || port == "" || dbName == "" || authSource == "" {
		log.Fatal("ERROR: Missing MongoDB environment variables")
	}

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		username, password, host, port, dbName, authSource)
}

// Initialize MongoDB connection
func InitMongoDB() {
	LoadEnv()
	mongoURI := getMongoURI()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB Connection Error: %v", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("MongoDB Ping Failed: %v", err)
	}

	log.Println("Connected to MongoDB")
	DB = client.Database(os.Getenv("MONGO_DB"))

	createIndexes()
}

// Ensure indexes are created
func createIndexes() {
	collection := DB.Collection("characters")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: "text"}},
		Options: options.Index().SetUnique(false),
	}

	if _, err := collection.Indexes().CreateOne(context.TODO(), indexModel); err != nil {
		log.Fatalf("Failed to create index on 'name' field: %v", err)
	}

	log.Println("Index on 'name' field created successfully")
}
