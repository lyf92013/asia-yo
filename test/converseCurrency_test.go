package main

import (
	"testing"

	"asia-yo/src/router"

	"github.com/kataras/iris/v12/httptest"
)

func TestConverseCurrency(t *testing.T) {
	app := router.Get()
	e := httptest.New(t, app)

	// USD to TWD Success
	e.GET("/api/v1/converse-currency").
		WithQuery("source", "USD").
		WithQuery("target", "TWD").
		WithQuery("amount", 100).
		Expect().
		Status(httptest.StatusOK).
		JSON().
		IsEqual(map[string]interface{}{
			"msg":    "success",
			"amount": "$3,044.40",
		})

	// Currency not supported
	e.GET("/api/v1/converse-currency").
		WithQuery("source", "USD").
		WithQuery("target", "CNY").
		WithQuery("amount", 100).
		Expect().
		Status(httptest.StatusBadRequest)
}
