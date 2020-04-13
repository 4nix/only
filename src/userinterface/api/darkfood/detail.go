package darkfood

import (
	"only/src/domain/service"
	"only/src/userinterface/api"

	"github.com/kataras/iris"
)

func GetDetail(ctx iris.Context) {
	ID, _ := ctx.Params().GetInt("id")

	res := service.GetItem(ID)

	api.Success(ctx, res)
}
