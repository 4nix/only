// Package repository 数据模型仓储层
package repository

import "only/src/infrastructure/dao"

type ItemRepo struct {
}

func (repo ItemRepo) Add(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) int64 {
	itemDao := dao.NewNoseItemDao()
	iID, _ := itemDao.Add(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)

	return iID
}

func (repo ItemRepo) Find(iID int64) interface{} {
	itemDao := dao.NewNoseItemDao()
	oItem := itemDao.Find(iID)

	return oItem
}

func (repo ItemRepo) GetList(iGID int64) interface{} {
	itemDao := dao.NewNoseItemDao()
	oItem := itemDao.FindByGID(iGID)

	return oItem
}

func (repo ItemRepo) Remove(iID int64) bool {
	itemDao := dao.NewNoseItemDao()
	iRowAffected, _ := itemDao.Remove(iID)

	if iRowAffected > 0 {
		return true
	}

	return false
}
