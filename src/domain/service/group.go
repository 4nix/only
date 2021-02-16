package service

import (
	"only/src/domain/repository"
)

// AddGroup 新增组
func AddGroup(sName string, sConfig string) int64 {
	groupRepo := repository.GroupRepo{}
	return groupRepo.Add(sName, sConfig)
}

// GetGroup 获取一个组信息
func GetGroup(ID int) interface{} {
	groupRepo := repository.GroupRepo{}

	return groupRepo.Find(ID)
}

// GetGroupList 获取组列表
func GetGroupList() interface{} {
	groupRepo := repository.GroupRepo{}

	return groupRepo.FindAll()
}

// RemoveGroup 移除组
func RemoveGroup(ID int) interface{} {
	groupRepo := repository.GroupRepo{}

	return groupRepo.Remove(ID)
}
