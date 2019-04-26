package constraints

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.zenithar.org/kingdom/internal/helpers"
)

// MustBeAnIdentifier returns a ID constraint validator
func MustBeAnIdentifier(value string) func(context.Context) error {
	return func(ctx context.Context) error {
		return validation.Validate(value, helpers.IDValidationRules...)
	}
}

// MustBeALabel returns a label constraint validator
func MustBeALabel(value string) func(context.Context) error {
	return func(ctx context.Context) error {
		return validation.Validate(value, validation.Required, is.PrintableASCII, validation.Length(3, 50))
	}
}
