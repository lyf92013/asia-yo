package handler

import (
	"github.com/kataras/iris/v12"
)

func PingHandler(ctx iris.Context) string {
	return "pong"
}
