package service

import "only/src/domain/repository"

func AddGroup(sName string, sConfig string) int64 {
	groupRepo := repository.GroupRepo{}
	return groupRepo.Add(sName, sConfig)
}

func GetGroup(ID int) interface{} {
	groupRepo := repository.GroupRepo{}

	return groupRepo.Find(ID)
}

func AddItem(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) int {
	itemRepo := repository.ItemRepo{}
	return itemRepo.Add(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)
}

func GetItem(ID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.Find(ID)
}
