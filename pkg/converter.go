package parser

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

var ErrParserEmptyResult = errors.New("parser empty result error")

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
