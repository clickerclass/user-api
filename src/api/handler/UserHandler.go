package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func UserFindByIdHandler(w http.ResponseWriter, r *http.Request) {
	var id string = mux.Vars(r)["id"]
	fmt.Println("El id es:", id)
}
