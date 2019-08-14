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

// SearchCommand handles realm search requests.
var SearchCommand = func(realms repositories.Realm) HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		// Check request
		req, ok := r.(*realmv1.SearchRequest)
		if !ok {
			return nil, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		res := &realmv1.SearchResponse{}

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
		); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: "Unable to validate request",
			}
			return res, err
		}

		// Prepare request parameters
		sortParams := db.SortConverter(req.Sorts)
		pagination := db.NewPaginator(uint(req.Page), uint(req.PerPage))

		// Build search filter
		filter := &repositories.RealmSearchFilter{}
		if req.Id != nil {
			filter.RealmID = req.Id.Value
		}
		if req.Label != nil {
			filter.Label = req.Label.Value
		}

		// Do the search
		entities, total, err := realms.Search(ctx, filter, pagination, sortParams)
		if err != nil && err != db.ErrNoResult {
			res.Error = &sysv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to process request",
			}
			return res, err
		}

		// Set pagination total for paging calcul
		pagination.SetTotal(uint(total))
		res.Total = uint32(pagination.Total())
		res.Count = uint32(pagination.CurrentPageCount())
		res.PerPage = uint32(pagination.PerPage)
		res.CurrentPage = uint32(pagination.Page)

		// If no result back to first page
		if err != db.ErrNoResult {
			res.Members = mapper.FromCollection(entities)
		}

		// Return results
		return res, nil
	}
}
