package admin

import (
	"only/src/application/apps/admin"
	"only/src/domain/service"
	"only/src/userinterface/api"

	"github.com/kataras/iris"
)

func AddGroup(ctx iris.Context) {
	name := ctx.PostValue("name")
	config := ctx.PostValue("config")

	s := map[string]string{"aaa": "ccc"}

	vals := make(map[string]interface{}, len(s))

	for i, v := range s {
		vals[i] = v
	}

	admin.AddGroup(name, config)
}

func GetGroup(ctx iris.Context) {
	ID, _ := ctx.Params().GetInt("id")
	res := admin.GetGroup(ID)

	// ctx.JSON(res)
	api.Success(ctx, res)
}

func GetGroupList(ctx iris.Context) {
	res := service.GetGroupList()

	api.Success(ctx, res)
}

func RemoveGroup(ctx iris.Context) {
	ID, _ := ctx.Params().GetInt("id")
	res := service.RemoveGroup(ID)

	api.Success(ctx, res)
}
