package repository

import (
	"context"
	"github.com/clickerclass/user-api/src/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

func UserFindByFilters(filters bson.M) (model.User, model.HttpResponse) {
	cursor := collection.FindOne(ctx, filters)
	var user model.User
	var httpResponse model.HttpResponse
	error := cursor.Decode(&user)
	if error != nil {
		httpResponse.MessageError = error.Error()
		httpResponse.StatusHttp = 404
		httpResponse.Error = true
	}
	return user, httpResponse
}
func UserCreate(user *model.User) model.HttpResponse {
	_, error := collection.InsertOne(ctx, user)
	var httpResponse model.HttpResponse
	if error != nil {
		httpResponse.MessageError = error.Error()
		httpResponse.StatusHttp = 409
		httpResponse.Error = true
	}
	return httpResponse
}
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://192.168.0.188:27017")
	client, err := mongo.NewClient(clientOptions)
	err = client.Connect(ctx)
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("clickerclass_user")
	collection = db.Collection("user")
}
