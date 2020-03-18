// Package repository 数据分组仓储层, 接口契约
package repository

import (
	"only/src/infrastructure/dao"
)

type GroupRepo struct {
}

func (repository GroupRepo) Add(sName string, sConfig string) int64 {
	groupDao := dao.NoseGroup{}
	// sConfig, _ := json.Marshal(mConfig)

	iID, _ := groupDao.Add(sName, sConfig)

	return iID
}

func (repository GroupRepo) Find(iID int) interface{} {
	groupDao := dao.NoseGroup{}
	oGroup := groupDao.Find(int64(iID))

	return oGroup
}

func (repository GroupRepo) FindAll() interface{} {
	groupDao := dao.NoseGroup{}
	oGroup := groupDao.FindAll()

	return oGroup
}

func (repository GroupRepo) Remove(iID int) bool {
	groupDao := dao.NoseGroup{}
	iRowAffected, _ := groupDao.Remove(int64(iID))

	if iRowAffected > 0 {
		return true
	}

	return false
}
