package service

import (
	"fmt"
	"only/src/domain/repository"
)

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

func FetchItemListByCursor(groupID int64, cursor int64, limit int, offset int) interface{} {
	itemRepo := repository.ItemRepo{}
	var params []interface{}

	condition := "group_id = ?"
	params = append(params, groupID)
	if cursor > 0 {
		condition += " AND id < ?"
		params = append(params, cursor)
	}

	condition += " ORDER BY id DESC LIMIT ? offset ?"
	params = append(params, limit, offset)

	fmt.Println(cursor)
	fmt.Println(condition)
	fmt.Println(params)

	return itemRepo.FetchAll(condition, params...)
}
