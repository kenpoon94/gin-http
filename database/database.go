package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/gin-http/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
) 

var host = utils.GetEnvVariable("MONGODB_HOST")
var port = utils.GetEnvVariable("MONGODB_PORT")
var database = utils.GetEnvVariable("MONGODB_DATABASE")

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())


	// Check the connection
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongodDB!")

	return &DB{
		client: client,
	}
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string    `bson:"name,omitempty"`
	Jobtitle string    `bson:"jobtitle,omitempty"`
	Age      int       `bson:"age,omitempty"`
	City     string    `bson:"city,omitempty"`
	Hobbies	 []string  `bson:"hobbies,omitempty"`
}

func (db* DB) Find()  []User{
	collection := db.client.Database(database).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	if err = cur.All(ctx, &users); err != nil {
		log.Fatal(err)
	}
	return users
}