package service

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/clickerclass/user-api/src/api/model"
	"github.com/clickerclass/user-api/src/api/repository"
)

func UserFindByEmail(email string) model.User {
	return repository.UserFindByEmail(email)

}
func UserCreate(user *model.User) interface{} {
	sha512 := sha512.New()
	sha512.Write([]byte(user.Password))
	user.Password = base64.StdEncoding.EncodeToString(sha512.Sum(nil))
	fmt.Println(*user)
	return repository.UserCreate(user)
}

func UserFindByDocTypeAndDoc(docType string, doc string) model.User {
	return repository.UserFindByDocTypeAndDoc(docType, doc)
}
