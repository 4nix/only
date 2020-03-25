package main

import (

	// _ "only/src/infrastructure/dao"

	"only/src/userinterface/api/admin"

	"github.com/rs/cors"

	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}

	corsWrapper := cors.New(corsOptions).ServeHTTP

	app.WrapRouter(corsWrapper)

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

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
