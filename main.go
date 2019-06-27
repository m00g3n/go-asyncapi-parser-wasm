package main

import (
	parser "parseasgo/pkg"
	"strconv"
	"syscall/js"

	"github.com/pkg/errors"
)

func main() {
	c := make(chan struct{}, 0)
	parseFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		document := args[0].String()
		circularReferences, err := strconv.ParseBool(args[1].String())
		if err != nil {
			parser.NewInternalError(errors.Wrapf(err, `invalid argument "circularReferences" %t`, circularReferences))
		}
		result := parser.DefaultParser([]byte(document), circularReferences)
		return js.ValueOf(result)
	})
	defer parseFunc.Release()
	js.Global().Set("parse", parseFunc)
	<-c
}
