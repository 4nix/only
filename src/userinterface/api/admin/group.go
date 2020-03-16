package admin

import (
	"only/src/application/apps/admin"

	"github.com/kataras/iris"
)

func AddGroup(ctx iris.Context) {
	name := ctx.PostValue("name")
	config := ctx.PostValue("config")
	// fmt.Println(id)

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

	ctx.JSON(res)
}
