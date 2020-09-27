package router

import (
	"github.com/clickerclass/user-api/src/api/handler"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func UserRouter() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/user/{id}", handler.UserFindByIdHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
