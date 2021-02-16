package service

import (
	"only/src/domain/repository"
)

// AddItem 添加item
func AddItem(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) int {
	itemRepo := repository.ItemRepo{}
	return itemRepo.Add(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)
}

// GetItem 获取item
func GetItem(ID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.Find(ID)
}

// GetItemListByGroupID 通过组id获取item list
func GetItemListByGroupID(iGroupID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.GetListByGroupID(iGroupID)
}

// GetItemListByRelateID 通过相关id 获取item list
func GetItemListByRelateID(iGroupID int, iRelateID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.GetListByRelateID(iGroupID, iRelateID)
}

// RemoveItem 移除item
func RemoveItem(ID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.Remove(ID)
}

// FetchItemListByCursor 通过组id获取列表，支持游标翻页，默认倒叙
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

	return itemRepo.FetchAll(condition, params...)
}

// FetchItemList 通过组id获取列表
func FetchItemList(groupID int64, relateID int64, limit int, offset int, order string) interface{} {
	itemRepo := repository.ItemRepo{}
	var params []interface{}

	condition := "group_id = ?"
	params = append(params, groupID)

	if relateID > 0 {
		condition += " AND relate_id = ? "
		params = append(params, relateID)
	}

	condition += " ORDER BY " + order + " LIMIT ? offset ?"
	params = append(params, limit, offset)

	return itemRepo.FetchAll(condition, params...)
}

// GetItemCount 获取总数
func GetItemCount(groupID int64, relateID int64) int {
	itemRepo := repository.ItemRepo{}
	var params []interface{}

	condition := "group_id = ?"
	params = append(params, groupID)

	if relateID > 0 {
		condition += " AND relate_id = ? "
		params = append(params, relateID)
	}

	return itemRepo.Count(condition, params...)
}
