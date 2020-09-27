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

func UserFindByEmail(email string) model.User {
	cursor := collection.FindOne(ctx, bson.M{"email": email})
	var user model.User
	error := cursor.Decode(&user)
	if error != nil {
		panic(error)
	}
	return user
}
func UserCreate(user *model.User) interface{} {
	insertResult, error := collection.InsertOne(ctx, user)
	if error != nil {
		panic(error)
	}
	return insertResult.InsertedID
}
func UserFindByDocTypeAndDoc(docType string, doc string) model.User {
	cursor := collection.FindOne(ctx, bson.M{"documentType": docType, "document": doc})
	var user model.User
	error := cursor.Decode(&user)
	if error != nil {
		panic(error)
	}
	return user
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
