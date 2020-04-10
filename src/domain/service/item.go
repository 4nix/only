package service

import "only/src/domain/repository"

func AddItem(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) int {
	itemRepo := repository.ItemRepo{}
	return itemRepo.Add(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)
}

func GetItem(ID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.Find(ID)
}

func GetItemListByGroupID(iGroupID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.GetListByGroupID(iGroupID)
}

func GetItemListByRelateID(iGroupID int, iRelateID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.GetListByRelateID(iGroupID, iRelateID)
}

func RemoveItem(ID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.Remove(ID)
}
