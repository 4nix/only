package darkfood

import (
	"only/src/domain/service"
	"only/src/userinterface/api"

	"github.com/kataras/iris"
)

func GetList(ctx iris.Context) {
	const iGroupID = 27
	const iRelateID = 476

	res := service.GetItemListByRelateID(iGroupID, iRelateID)

	api.Success(ctx, res)
}
