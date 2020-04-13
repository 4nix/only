package yinyangshi

import (
	"only/src/domain/service"
	"only/src/userinterface/api"

	"github.com/kataras/iris"
)

func GetXuanshangList(ctx iris.Context) {
	const iGroupID = 26
	const iRelateID = 2

	res := service.GetItemListByRelateID(iGroupID, iRelateID)

	api.Success(ctx, res)
}
