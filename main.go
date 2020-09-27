package main

import (
	"github.com/clickerclass/user-api/src/api/repository"
	"github.com/clickerclass/user-api/src/api/router"
)

func main() {
	repository.Connect()
	router.UserRouter()
}
