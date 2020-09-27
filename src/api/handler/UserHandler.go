package handler

import (
	"encoding/json"
	"github.com/clickerclass/user-api/src/api/service"
	"github.com/gorilla/mux"
	"net/http"
)

func UserFindByIdHandler(w http.ResponseWriter, r *http.Request) {
	var id string = mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(service.UserFindById(id))
}
