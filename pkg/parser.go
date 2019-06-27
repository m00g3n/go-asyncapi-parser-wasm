package parser

import (
	asyncparser "github.com/asyncapi/parser/pkg"
)

func NewInternalError(err error) Error {
	return NewError("internal error", err.Error())
}

type Error struct {
	Type        string
	Description string
}

func NewError(errType string, description string) Error {
	return Error{
		Type:        errType,
		Description: description,
	}
}

type Result struct {
	Errors   []Error
	Document string
}

func NewErrorResult(errors []Error) Result {
	return Result{
		Errors: errors,
	}
}

func NewResult(document string) Result {
	return Result{
		Document: document,
	}
}

type Parser func(document []byte, circularReferences bool) Result

var DefaultParser Parser = func(document []byte, circularReferences bool) Result {
	rawMsg, err := asyncparser.Parse([]byte(document), circularReferences)
	if err != nil {
		return NewErrorResult(ToErrors(err.ParsingErrors))
	}
	return NewResult(string(rawMsg))
}
