package yinyangshi

import (
	"only/src/domain/service"
	"only/src/userinterface/api"

	"github.com/kataras/iris"
)

func GetShenmiList(ctx iris.Context) {
	const iGroupID = 26
	const iRelateID = 24

	res := service.GetItemListByRelateID(iGroupID, iRelateID)

	api.Success(ctx, res)
}
