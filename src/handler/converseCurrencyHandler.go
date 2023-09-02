package handler

import (
	"asia-yo/src/service"

	"github.com/kataras/iris/v12"
)

func ConverseCurrencyHandler(ctx iris.Context) map[string]string {
	req := struct {
		Source string  `json:"source" validate:"oneof=TWD JPY USD"`
		Target string  `json:"target" validate:"oneof=TWD JPY USD"`
		Amount float64 `json:"amount" validate:"gt=0"`
	}{}
	readQuery(ctx, &req)

	amount := service.ConverseCurrency(req.Source, req.Target, req.Amount)
	return map[string]string{
		"msg":    "success",
		"amount": amount,
	}
}
