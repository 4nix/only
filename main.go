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
	app.Post("/admin/item", admin.AddItem)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":9999"))
}
