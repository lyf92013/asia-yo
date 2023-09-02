package errors

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type ValidationError struct {
	Reason interface{}
}

func (e *ValidationError) StatusCode() int {
	return iris.StatusBadRequest
}

func (e *ValidationError) Error() string {
	result := ""

	switch err := e.Reason.(type) {
	default:
		result = fmt.Sprintf("%v", err)
	case validator.ValidationErrors:
		for _, reason := range err {
			result += fmt.Sprintf("Field `%s` has rule '%s=%s', but got '%v';",
				reason.Field(),
				reason.Tag(),
				reason.Param(),
				reason.Value(),
			)
		}
	case *json.UnmarshalTypeError:
		result = err.Field + " is not a " + err.Type.String()
	}
	return result
}
