package admin

import (
	"only/src/application/apps/admin"
	"only/src/domain/service"
	"only/src/userinterface/api"
	"strconv"

	"github.com/kataras/iris"
)

func AddItem(ctx iris.Context) {
	sName := ctx.PostValue("name")
	sContent := ctx.PostValueTrim("content")
	sConfig := ctx.PostValueDefault("config", "")
	iModeID := ctx.PostValueIntDefault("model_id", 0)
	iGroupID, err := ctx.PostValueInt("gid")
	iRelateID := ctx.PostValueIntDefault("rid", 0)
	// fmt.Println(id)

	if err != nil {
		return
	}

	admin.AddItem(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)
}

func EditItem(ctx iris.Context) {
	// id, _ := ctx.Params().GetInt("id")
	sName := ctx.PostValue("name")
	sContent := ctx.PostValueTrim("content")
	sConfig := ctx.PostValueDefault("config", "")
	iModeID := ctx.PostValueIntDefault("model_id", 0)
	iGroupID, _ := ctx.PostValueInt("gid")
	iRelateID := ctx.PostValueIntDefault("rid", 0)
	// fmt.Println(id)

	// if err != nil {
	// 	return
	// }

	admin.AddItem(sName, sContent, sConfig, iModeID, iGroupID, iRelateID)
	// admin.EditItem()
}

func GetItem(ctx iris.Context) {
	ID, _ := ctx.Params().GetInt("id")
	res := admin.GetItem(ID)

	api.Success(ctx, res)

	// api.Error(ctx, "测试", 22)
}

func GetItemList(ctx iris.Context) {
	groupID, _ := ctx.Params().GetInt("gid")
	relateID := ctx.Params().GetIntDefault("rid", 0)
	page, _ := strconv.Atoi(ctx.Request().URL.Query().Get("page"))

	limit := 20
	offset := 0
	if page > 0 {
		offset = (page - 1) * 20
	}

	res := make(map[string]interface{})

	res["list"] = service.FetchItemList(int64(groupID), int64(relateID), limit, offset, "id asc")
	res["total"] = service.GetItemCount(int64(groupID), int64(relateID))

	api.Success(ctx, res)
}

func RemoveItem(ctx iris.Context) {
	ID, _ := ctx.Params().GetInt("id")
	res := service.RemoveItem(ID)

	api.Success(ctx, res)
}
