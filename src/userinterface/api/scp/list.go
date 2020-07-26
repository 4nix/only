package scp

import (
	"only/src/domain/service"
	"only/src/userinterface/api"
	"strconv"

	"github.com/kataras/iris"
)

func GetList(ctx iris.Context) {
	const groupID = 29
	ID, _ := strconv.ParseInt(ctx.Request().URL.Query().Get("id"), 10, 64)

	// res := service.GetItemListByGroupID(iGroupID)
	res := service.FetchItemListByCursor(groupID, ID, 20, 0)

	api.Success(ctx, res)
}
