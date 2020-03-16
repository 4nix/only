// Package repository 数据模型仓储层
package repository

import "only/src/infrastructure/dao"

type ItemRepo struct {
}

func (repo ItemRepo) Add(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) int {
	itemDao := dao.NoseItem{}
	iID, _ := itemDao.Add(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)

	return int(iID)
}

func (repo ItemRepo) Find(iID int) interface{} {
	itemDao := dao.NoseItem{}
	oItem := itemDao.Find(int64(iID))

	return oItem
}

func (repo ItemRepo) GetList(iGID int) interface{} {
	itemDao := dao.NoseItem{}
	oItem := itemDao.FindByGID(int64(iGID))

	return oItem
}

func (repo ItemRepo) Remove(iID int) bool {
	itemDao := dao.NoseItem{}
	iRowAffected, _ := itemDao.Remove(int64(iID))

	if iRowAffected > 0 {
		return true
	}

	return false
}
