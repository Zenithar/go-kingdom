package constraints

import (
	"context"

	"go.zenithar.org/pkg/errors"
)

// Validable interface used to defines Validation protocol
type Validable interface {
	Validate() error
}

// MustBeValid specification checks that given object is valid
func MustBeValid(validable Validable) func(context.Context) error {
	return func(ctx context.Context) error {
		// Validate request
		if err := validable.Validate(); err != nil {
			return errors.Newf(errors.InvalidArgument, err, "unable to validate object")
		}
		return nil
	}
}
