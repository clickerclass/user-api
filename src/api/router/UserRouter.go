package router

import (
	"github.com/clickerclass/user-api/src/api/handler"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func UserRouter() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/user/{email}", handler.UserFindByIdHandler).Methods("GET")
	router.HandleFunc("/api/user", handler.UserCreateHandler).Methods("POST")
	router.HandleFunc("/api/user/{docType}/{doc}", handler.UserFindByDocTypeAndDocHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
