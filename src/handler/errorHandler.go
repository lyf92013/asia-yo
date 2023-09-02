package handler

import (
	"asia-yo/src/errors"
	"asia-yo/src/utils/logger"

	"github.com/kataras/iris/v12"
)

type Error interface {
	StatusCode() int
	Error() string
}

func ErrorHandler(ctx iris.Context) {
	defer func() {
		if e := recover(); e != nil {
			logger.Logger().Sugar().Errorf("%+v", e)

			var err, ok = e.(Error)
			if ok {
				ctx.StatusCode(err.StatusCode())
			} else {
				ctx.StatusCode(iris.StatusInternalServerError)
				err = &errors.InternalServerError{}
			}

			ctx.JSON(struct {
				Message string      `json:"msg"`
				Data    interface{} `json:"data,omitempty"`
			}{
				Message: err.Error(),
				Data:    nil,
			})
		}
	}()

	ctx.Next()
}
