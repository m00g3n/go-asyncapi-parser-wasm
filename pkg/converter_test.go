package parser_test

import (
	. "github.com/onsi/gomega"
	"github.com/xeipuuv/gojsonschema"
	parser "parseasgo/pkg"
	"testing"
)

func TestToErrors(t *testing.T) {
	e1 := gojsonschema.InternalError{
		ResultErrorFields: gojsonschema.ResultErrorFields{},
	}
	e1.SetContext(gojsonschema.NewJsonContext("root", nil))
	e2 := gojsonschema.InternalError{
		ResultErrorFields: gojsonschema.ResultErrorFields{},
	}
	e2.SetContext(gojsonschema.NewJsonContext("root", nil))
	err := parser.ToErrors([]gojsonschema.ResultError{
		&e1,
		&e2,
		nil,
	})
	assert := NewWithT(t)
	assert.Expect(err).NotTo(BeNil())
}
