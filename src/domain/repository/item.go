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

func (repo ItemRepo) GetListByGroupID(iGroupID int) interface{} {
	itemDao := dao.NoseItem{}
	oItem := itemDao.FindByGroupID(int64(iGroupID))

	return oItem
}

func (repo ItemRepo) GetListByRelateID(iGroupID int, iRelateID int) interface{} {
	itemDao := dao.NoseItem{}
	oItem := itemDao.FindByRelateID(int64(iGroupID), int64(iRelateID))

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

func (repo ItemRepo) FetchAll(condition string, params ...interface{}) interface{} {
	itemDao := dao.NoseItem{}

	return itemDao.FetchAll(condition, params...)
}
