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
	user, http := service.UserFindByEmail(email)
	user.Password = ""
	if http.Error == true {
		w.WriteHeader(http.StatusHttp)
		json.NewEncoder(w).Encode(http)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)

	}
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	http := service.UserCreate(&user)
	if http.Error == true {
		w.WriteHeader(http.StatusHttp)
		json.NewEncoder(w).Encode(http)
	} else {
		w.WriteHeader(201)
	}

}
func UserFindByDocTypeAndDocHandler(w http.ResponseWriter, r *http.Request) {
	docType := mux.Vars(r)["docType"]
	w.Header().Add("Content-Type", "application/json")
	doc := mux.Vars(r)["doc"]
	user, http := service.UserFindByDocTypeAndDoc(docType, doc)
	user.Password = ""
	if http.Error == true {
		w.WriteHeader(http.StatusHttp)
		json.NewEncoder(w).Encode(http)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)

	}
}
func AuthenticationUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	email := mux.Vars(r)["email"]
	user, http := service.UserFindByEmail(email)
	if http.Error == true {
		w.WriteHeader(http.StatusHttp)
		json.NewEncoder(w).Encode(http)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)

	}
}
