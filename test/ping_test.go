package main

import (
	"testing"

	"asia-yo/src/router"

	"github.com/kataras/iris/v12/httptest"
)

func TestPing(t *testing.T) {
	app := router.Get()
	e := httptest.New(t, app)

	e.GET("/api/v1/ping").Expect().Status(httptest.StatusOK).Body().IsEqual("pong")
}
