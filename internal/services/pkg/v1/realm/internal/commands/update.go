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

// UpdateCommand handles realm update requests.
var UpdateCommand = func(realms repositories.Realm) HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		// Check request
		req, ok := r.(*realmv1.UpdateRequest)
		if !ok {
			return nil, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		res := &realmv1.UpdateResponse{}
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

		updated := false

		if req.Label != nil {
			if err := constraints.Validate(ctx,
				// Check acceptable label value
				constraints.MustBeALabel(req.Label.Value),
				// Is already used ?
				constraints.RealmLabelMustBeUnique(realms, req.Label.Value),
			); err != nil {
				res.Error = &sysv1.Error{
					Code:    http.StatusConflict,
					Message: err.Error(),
				}
				return res, err
			}
			entity.Label = req.Label.Value
			updated = true
		}

		// Skip operation when no updates
		if updated {
			// Create account in database
			if err := realms.Update(ctx, &entity); err != nil {
				res.Error = &sysv1.Error{
					Code:    http.StatusInternalServerError,
					Message: "Unable to update Realm object",
				}
				return res, err
			}
		}

		// Prepare response
		res.Entity = mapper.FromEntity(&entity)

		// Return expected result
		return res, nil
	}
}
