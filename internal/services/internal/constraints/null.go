package constraints

import (
	"context"
	"reflect"

	"go.zenithar.org/pkg/errors"
)

func isNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}

// MustNotBeNil specification checks that given object is not nil
func MustNotBeNil(object interface{}, message string) func(context.Context) error {
	return func(ctx context.Context) error {
		if isNil(object) {
			return errors.Newf(errors.InvalidArgument, nil, "%s must not be nil", message)
		}

		// Return no error
		return nil
	}
}
