package constraints

import (
	"context"

	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
)

// mustBeUnique specification checks if the given name already exists
func mustBeUnique(finder EntityRetrieverFunc, attribute string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// Retrieve object from repository
		object, err := finder(ctx)
		if err != nil && err != db.ErrNoResult {
			return errors.Newf(errors.Internal, err, "unable to check uniqueness property")
		}
		if !isNil(object) {
			return errors.Newf(errors.AlreadyExists, nil, "%s is already used", attribute)
		}

		return nil
	}
}

// -----------------------------------------------------------------------------

// RealmLabelMustBeUnique returns specification for realm label uniqueness
func RealmLabelMustBeUnique(realms repositories.Realm, label string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return realms.FindByLabel(ctx, label)
		}, "Realm label")
}

// UserPrincipalMustBeUnique returns specification for user principal uniqueness
func UserPrincipalMustBeUnique(users repositories.User, realmID, principal string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return users.FindByPrincipal(ctx, realmID, principal)
		}, "User principal")
}
