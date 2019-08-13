package constraints

import (
	"context"

	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/pkg/errors"
)

// EntityRetrieverFunc describes function indirection for repositories
type EntityRetrieverFunc func(context.Context) (interface{}, error)

func mustExists(finder EntityRetrieverFunc) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		object, err := finder(ctx)
		if err != nil {
			return errors.Newf(errors.Internal, err, "unable to check object existence")
		}
		if isNil(object) {
			return errors.Newf(errors.NotFound, nil, "object not found")
		}
		return nil
	}
}

// -----------------------------------------------------------------------------

// RealmMustExists specification checks if given squad exists
func RealmMustExists(realms repositories.Realm, id string, entity *models.Realm) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			object, err := realms.Get(ctx, id)
			if object != nil {
				*entity = *object
			}
			return entity, err
		},
	)
}

// UserMustExists specification checks if given user exists
func UserMustExists(users repositories.User, realmID, userID string, entity *models.User) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			object, err := users.Get(ctx, realmID, userID)
			if object != nil {
				*entity = *object
			}
			return entity, err
		},
	)
}
