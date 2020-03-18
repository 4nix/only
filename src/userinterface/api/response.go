package api

import "github.com/kataras/iris"

func Success(ctx iris.Context, data interface{}) {
	if data != nil {
		ctx.JSON(data)
	} else {
		res := make(map[string]interface{})
		res["error"] = 0

		ctx.JSON(res)
	}
}

func Error(ctx iris.Context, msg string, err int) {
	res := make(map[string]interface{})

	if err > 0 {
		res["error"] = err
	} else {
		res["error"] = 1
	}

	res["msg"] = msg
	ctx.JSON(res)
}
