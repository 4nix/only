package main

import (

	// _ "only/src/infrastructure/dao"

	"net/http"
	"only/src/userinterface/api/admin"
	"only/src/userinterface/api/darkfood"
	"only/src/userinterface/api/scp"
	"only/src/userinterface/api/yinyangshi"

	"github.com/kataras/iris"
	"github.com/rs/cors"
)

func main() {
	app := iris.Default()
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}

	corsWrapper := cors.New(corsOptions).ServeHTTP

	app.WrapRouter(corsWrapper)
	// app.Use(Cors)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "ponga"})
	})

	app.Post("/admin/group", admin.AddGroup)
	app.Get("/admin/group/{id:int}", admin.GetGroup)
	app.Delete("/admin/group/{id:int}", admin.RemoveGroup)

	app.Post("/admin/item", admin.AddItem)
	app.Get("/admin/item/{id:int}", admin.GetItem)
	app.Put("/admin/item/{id:int}", admin.EditItem)
	app.Delete("/admin/item/{id:int}", admin.RemoveItem)

	app.Get("/admin/list", admin.GetGroupList)
	app.Get("/admin/list/{gid:int}", admin.GetItemList)
	app.Get("/admin/list/{gid:int}/{rid:int}", admin.GetItemList)

	app.Get("/yinyangshi/wenda", yinyangshi.GetWenDaList)
	app.Get("/yinyangshi/xuanshang", yinyangshi.GetXuanshangList)
	app.Get("/yinyangshi/shenmi", yinyangshi.GetShenmiList)

	app.Get("/darkfood/list", darkfood.GetList)
	app.Get("/darkfood/food/{id:int}", darkfood.GetDetail)

	app.Get("/scp/list", scp.GetList)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":9999"))
}
