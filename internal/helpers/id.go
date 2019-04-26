package helpers

import (
	"github.com/dchest/uniuri"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	// IDLen defines identitier length
	IDLen = 32
)

// IDGeneratorFunc is used to generate an identifier
var IDGeneratorFunc = func() string {
	return uniuri.NewLen(IDLen)
}

var (
	IDValidationRules = []validation.Rule{
		validation.Required,
		validation.Length(IDLen, IDLen),
		is.PrintableASCII,
	}
)
