package parser

import (
	"encoding/json"

	asyncparser "github.com/asyncapi/parser/pkg"
)

func NewInternalError(err error) map[string]interface{} {
	return NewError("internal error", err.Error())
}

func NewErrorResult(errors []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"errors": errors,
	}
}

func NewResult(document string) map[string]interface{} {
	return map[string]interface{}{
		"document": document,
	}
}

func NewError(errType string, description string) map[string]interface{} {
	return map[string]interface{}{
		"type":        errType,
		"description": description,
	}
}

type Parser func(document []byte, circularReferences bool) map[string]interface{}

var DefaultParser Parser = func(document []byte, circularReferences bool) map[string]interface{} {
	rawMsg, err := asyncparser.Parse([]byte(document), circularReferences)
	var errors []map[string]interface{}
	if err != nil {
		errors = append(errors, ToErrors(err.ParsingErrors)...)
	}
	if len(errors) > 0 {
		return NewErrorResult(errors)
	}
	doc, err2 := json.MarshalIndent(rawMsg, "", "  ")
	if err2 != nil {
		return NewErrorResult([]map[string]interface{}{
			NewError("internal error", err2.Error()),
		})
	}
	return NewResult(string(doc))
}
