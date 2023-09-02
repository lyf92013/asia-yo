package handler

import (
	"asia-yo/src/errors"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type response struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func readQuery(ctx iris.Context, output interface{}) {
	err := ctx.ReadQuery(output)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		panic(&errors.ValidationError{
			Reason: validationErrors,
		})
	}
}
