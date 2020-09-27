package handler

import (
	"encoding/json"
	"github.com/clickerclass/user-api/src/api/model"
	"github.com/clickerclass/user-api/src/api/service"
	"github.com/gorilla/mux"

	"net/http"
)

func UserFindByIdHandler(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	user := service.UserFindByEmail(email)
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(service.UserCreate(&user))
}
func UserFindByDocTypeAndDocHandler(w http.ResponseWriter, r *http.Request) {
	docType := mux.Vars(r)["docType"]
	doc := mux.Vars(r)["doc"]
	json.NewEncoder(w).Encode(service.UserFindByDocTypeAndDoc(docType, doc))
}
