package parser

import (
	"github.com/xeipuuv/gojsonschema"
)

func toErrorMap(err gojsonschema.ResultError) Error {
	return NewError(err.Type(), err.Description())
}

func ToErrors(errs []gojsonschema.ResultError) []Error {
	var result []Error
	for _, err := range errs {
		result = append(result, toErrorMap(err))
	}
	return result
}
