package router

import (
	"asia-yo/src/handler"
	"asia-yo/src/middleware"
	"asia-yo/src/utils/config"
	"asia-yo/src/utils/logger"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	iriszap "github.com/iris-contrib/middleware/zap"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func Get() *iris.Application {
	app := iris.New()
	app.Validator = validator.New()

	app.UseRouter(middleware.CorsMiddleware)

	app.Use(iriszap.New(logger.Logger(), time.RFC3339, true))
	app.Use(handler.ErrorHandler)

	app.ConfigureContainer(func(api *iris.APIContainer) {
		v1 := api.Party("/api/v1")
		v1.Get("/ping", hero.Handler(handler.PingHandler))
		v1.Get("/converse-currency", hero.Handler(handler.ConverseCurrencyHandler))
	})

	return app
}

func Run() {
	app := Get()
	app.Run(
		iris.Addr(fmt.Sprintf("%v:%v", config.Get("SERVER_HOST"), config.Get("SERVER_PORT"))))
}
