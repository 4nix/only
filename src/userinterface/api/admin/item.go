package admin

import (
	"only/src/application/apps/admin"
	"only/src/domain/service"
	"only/src/userinterface/api"

	"github.com/kataras/iris"
)

func AddItem(ctx iris.Context) {
	sName := ctx.PostValue("name")
	sContent := ctx.PostValueTrim("content")
	sConfig := ctx.PostValueDefault("config", "")
	iModeID := ctx.PostValueIntDefault("mode_id", 0)
	iGroupID, err := ctx.PostValueInt("group_id")
	iRelateID := ctx.PostValueIntDefault("relate_id", 0)
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
	iRelateID, _ := ctx.Params().GetInt("rid")

	var res interface{}
	if iRelateID >= 0 {
		res = service.GetItemListByRelateID(iGroupID, iRelateID)
	} else {
		res = service.GetItemListByGroupID(iGroupID)
	}

	api.Success(ctx, res)
}
