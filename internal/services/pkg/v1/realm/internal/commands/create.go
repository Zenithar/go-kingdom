package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/kingdom/internal/services/internal/constraints"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/realm/internal/mapper"
	realmv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/realm/v1"
	sysv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/system/v1"
	"go.zenithar.org/pkg/errors"
)

// CreateCommand handle realm creation in persistence.
var CreateCommand = func(realms repositories.Realm) HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		// Check request
		req, ok := r.(*realmv1.CreateRequest)
		if !ok {
			return nil, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		res := &realmv1.CreateResponse{}

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Label must be unique
			constraints.RealmLabelMustBeUnique(realms, req.Label),
		); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		// Prepare entity creation
		entity := models.NewRealm(req.Label)

		// Create entity in database
		if err := realms.Create(ctx, entity); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to create realm",
			}
			return res, errors.Newf(errors.Internal, err, "unable to create entity")
		}

		// Prepare response
		res.Entity = mapper.FromEntity(entity)

		// Return result
		return res, nil
	}
}
