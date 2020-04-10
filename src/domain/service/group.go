package service

import (
	"only/src/domain/repository"
)

func AddGroup(sName string, sConfig string) int64 {
	groupRepo := repository.GroupRepo{}
	return groupRepo.Add(sName, sConfig)
}

func GetGroup(ID int) interface{} {
	groupRepo := repository.GroupRepo{}

	return groupRepo.Find(ID)
}

func GetGroupList() interface{} {
	groupRepo := repository.GroupRepo{}

	return groupRepo.FindAll()
}

func RemoveGroup(ID int) interface{} {
	groupRepo := repository.GroupRepo{}

	return groupRepo.Remove(ID)
}
