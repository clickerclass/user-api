package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

func FindById(id string) string {
	cursor := collection.FindOne(ctx, bson.M{})
	byteUser, error := cursor.DecodeBytes()
	if error != nil {
		panic(error)
	}
	return byteUser.String()
}
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://192.168.0.188:27017")
	client, err := mongo.NewClient(clientOptions)
	err = client.Connect(ctx)
	//defer client.Disconnect(ctx)
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("clickerclass_user")
	collection = db.Collection("user")
}
