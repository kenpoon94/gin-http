// package utils
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
) 

type User struct {
	Name string 
	JobTitle string
	Age int
	City string
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env files")
	}
	return os.Getenv(key)
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + getEnvVariable("MONGODB_HOST") + ":" + getEnvVariable("MONGODB_PORT"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	// Check the connection
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongodDB!")

	database := client.Database(getEnvVariable("MONGODB_DATABASE"))
	users_collection := database.Collection("users")

	ash := User{Name: "Ash", JobTitle: "Accountant", Age: 23, City: "Singapore"}
	result, err := users_collection.InsertOne(context.TODO(), ash )

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}