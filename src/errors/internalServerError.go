package errors

import (
	"github.com/kataras/iris/v12"
)

type InternalServerError struct {
}

func (e *InternalServerError) StatusCode() int {
	return iris.StatusInternalServerError
}

func (e *InternalServerError) Error() string {
	return "Internal Server Error"
}
