package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/kingdom/internal/services/internal/constraints"
	realmv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/realm/v1"
	sysv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/system/v1"
	"go.zenithar.org/pkg/errors"
)

// DeleteCommand handles realm deletion from persistence.
var DeleteCommand = func(realms repositories.Realm) HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		// Check request
		req, ok := r.(*realmv1.DeleteRequest)
		if !ok {
			return nil, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		res := &realmv1.DeleteResponse{}

		// Prepare expected results
		var entity models.Realm

		// Check request
		if req == nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusBadRequest,
				Message: "request must not be nil",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request must not be nil")
		}

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Chapter must exists
			constraints.RealmMustExists(realms, req.Id, &entity),
		); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		if err := realms.Delete(ctx, req.Id); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to delete Realm object",
			}
			return res, errors.Newf(errors.Internal, err, "unable to delete entity")
		}

		// Return expected result
		return res, nil
	}
}
