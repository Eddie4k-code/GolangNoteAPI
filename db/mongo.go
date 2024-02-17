package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var UserCollection *mongo.Collection

/* Initalizes and Connects to the Mongo Database */
func InitDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017")) //set to env

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	DB = client.Database("note_app")

	UserCollection = DB.Collection("users")
}
