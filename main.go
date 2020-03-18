package main

import (

	// _ "only/src/infrastructure/dao"

	"only/src/userinterface/api/admin"

	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "ponga"})
	})

	app.Post("/admin/group", admin.AddGroup)
	app.Get("/admin/group/{id:int}", admin.GetGroup)
	app.Get("/admin/grouplist", admin.GetGroupList)

	app.Post("/admin/item", admin.AddItem)
	app.Get("/admin/item/{id:int}", admin.GetItem)
	app.Get("/admin/itemlist/{gid:int}/{rid:int}", admin.GetItemList)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":9999"))
}
