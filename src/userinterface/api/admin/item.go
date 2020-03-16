package admin

import (
	"only/src/application/apps/admin"

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
