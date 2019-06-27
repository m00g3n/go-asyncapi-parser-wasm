package parser

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

var ErrParserEmptyResult = errors.New("parser empty result error")

func toErrorMap(err gojsonschema.ResultError) map[string]interface{} {
	if err == nil {
		return nil
	}
	return NewError(err.Type(), err.Description())
}

func ToErrors(errs []gojsonschema.ResultError) []interface{} {
	var result []interface{}
	for _, err := range errs {
		result = append(result, toErrorMap(err))
	}
	return result
}
