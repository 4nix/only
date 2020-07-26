package service

import "only/src/domain/repository"

func GetNotice(iGroupID int, iRelateID int) interface{} {
	itemRepo := repository.ItemRepo{}

	return itemRepo.GetListByRelateID(iGroupID, iRelateID)
}
