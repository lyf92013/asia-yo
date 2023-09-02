package middleware

import (
	"github.com/kataras/iris/v12"
)

func CorsMiddleware(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if ctx.Method() == iris.MethodOptions {
		ctx.StatusCode(iris.StatusOK)
		return
	}

	ctx.Next()
}
