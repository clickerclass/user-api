package service

import "github.com/clickerclass/user-api/src/api/repository"

func UserFindById(id string) string {
	return repository.FindById(id)

}
