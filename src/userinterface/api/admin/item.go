package admin

import (
	"fmt"
	"only/src/application/apps/admin"
	"only/src/domain/service"
	"only/src/userinterface/api"

	"github.com/kataras/iris"
)

func AddItem(ctx iris.Context) {
	sName := ctx.PostValue("name")
	sContent := ctx.PostValueTrim("content")
	sConfig := ctx.PostValueDefault("config", "")
	iModeID := ctx.PostValueIntDefault("mid", 0)
	iGroupID, err := ctx.PostValueInt("gid")
	iRelateID := ctx.PostValueIntDefault("rid", 0)
	// fmt.Println(id)

	if err != nil {
		return
	}

	admin.AddItem(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)
}

func GetItem(ctx iris.Context) {
	ID, _ := ctx.Params().GetInt("id")
	res := admin.GetItem(ID)

	api.Success(ctx, res)

	// api.Error(ctx, "测试", 22)
}

func GetItemList(ctx iris.Context) {
	iGroupID, _ := ctx.Params().GetInt("gid")
	iRelateID, err := ctx.Params().GetInt("rid")

	fmt.Println(iRelateID)
	fmt.Println(err)

	var res interface{}
	if err == nil {
		res = service.GetItemListByRelateID(iGroupID, iRelateID)
	} else {
		res = service.GetItemListByRelateID(iGroupID, 0)
	}

	api.Success(ctx, res)
}
