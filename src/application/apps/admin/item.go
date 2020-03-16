package admin

import "only/src/domain/service"

// AddItem 添加item
func AddItem(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) int {
	return service.AddItem(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)
}

func GetItem(ID int) interface{} {
	return service.GetItem(ID)
}
