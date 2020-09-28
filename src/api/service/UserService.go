package service

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/clickerclass/user-api/src/api/model"
	"github.com/clickerclass/user-api/src/api/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func UserFindByEmail(email string) (model.User, model.HttpResponse) {
	filter := bson.M{"email": email}
	return repository.UserFindByFilters(filter)

}
func UserCreate(user *model.User) model.HttpResponse {
	sha512 := sha512.New()
	sha512.Write([]byte(user.Password))
	user.Password = base64.StdEncoding.EncodeToString(sha512.Sum(nil))
	fmt.Println(*user)
	return repository.UserCreate(user)
}

func UserFindByDocTypeAndDoc(docType string, doc string) (model.User, model.HttpResponse) {
	filter := bson.M{"documentType": docType, "document": doc}
	return repository.UserFindByFilters(filter)
}
