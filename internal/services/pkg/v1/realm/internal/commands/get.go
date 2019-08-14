package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/kingdom/internal/services/internal/constraints"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/realm/internal/mapper"
	realmv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/realm/v1"
	sysv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/system/v1"
	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
)

// GetCommand handles realm retrieval from persistence.
var GetCommand = func(realms repositories.Realm) HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		// Check request
		req, ok := r.(*realmv1.GetRequest)
		if !ok {
			return nil, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		res := &realmv1.GetResponse{}

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
		); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: "Unable to validate request",
			}
			return res, err
		}

		// Retrieve Chapter from database
		entity, err := realms.Get(ctx, req.Id)
		if err != nil && err != db.ErrNoResult {
			res.Error = &sysv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to retrieve Realm",
			}
			return res, err
		}
		if entity == nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusNotFound,
				Message: "Realm not found",
			}
			return res, db.ErrNoResult
		}

		// Prepare response
		res.Entity = mapper.FromEntity(entity)

		// Return result
		return res, nil
	}
}
