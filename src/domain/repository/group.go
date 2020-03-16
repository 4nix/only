// Package repository 数据分组仓储层, 接口契约
package repository

import (
	"only/src/infrastructure/dao"
)

type GroupRepo struct {
}

func (repository GroupRepo) Add(sName string, sConfig string) int64 {
	groupDao := dao.NewNoseGroupDao()
	// sConfig, _ := json.Marshal(mConfig)

	iID, _ := groupDao.Add(sName, sConfig)

	return iID
}

func (repository GroupRepo) Find(iID int64) interface{} {
	groupDao := dao.NewNoseGroupDao()
	oGroup := groupDao.Find(iID)

	return oGroup
}

func (repository GroupRepo) Remove(iID int64) bool {
	groupDao := dao.NewNoseGroupDao()
	iRowAffected, _ := groupDao.Remove(iID)

	if iRowAffected > 0 {
		return true
	}

	return false
}
