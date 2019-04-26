package realm

import (
	"context"
	"fmt"
	"net/http"

	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/kingdom/internal/services/internal/constraints"
	apiv1 "go.zenithar.org/kingdom/internal/services/pkg/v1"
	realmv1 "go.zenithar.org/kingdom/pkg/protocol/kingdom/realm/v1"
	sysv1 "go.zenithar.org/kingdom/pkg/protocol/kingdom/system/v1"
	"go.zenithar.org/pkg/db"
)

type service struct {
	realms repositories.Realm
}

// New services instance
func New(realms repositories.Realm) apiv1.Realm {
	return &service{
		realms: realms,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *realmv1.CreateRequest) (*realmv1.SingleResponse, error) {
	res := &realmv1.SingleResponse{}

	// Check request
	if req == nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Label must be unique
		constraints.RealmLabelMustBeUnique(s.realms, req.Label),
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
	if err := s.realms.Create(ctx, entity); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create realm",
		}
		return res, err
	}

	// Prepare response
	res.Entity = FromEntity(entity)

	// Return result
	return res, nil
}

func (s *service) Get(ctx context.Context, req *realmv1.GetRequest) (*realmv1.SingleResponse, error) {
	res := &realmv1.SingleResponse{}

	// Check request
	if req == nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
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

	// Retrieve Chapter from database
	entity, err := s.realms.Get(ctx, req.Id)
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
	res.Entity = FromEntity(entity)

	// Return result
	return res, nil
}

func (s *service) Update(ctx context.Context, req *realmv1.UpdateRequest) (*realmv1.SingleResponse, error) {
	res := &realmv1.SingleResponse{}
	// Prepare expected results
	var entity models.Realm

	// Check request
	if req == nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Chapter must exists
		constraints.RealmMustExists(s.realms, req.Id, &entity),
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
			constraints.RealmLabelMustBeUnique(s.realms, req.Label.Value),
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
		if err := s.realms.Update(ctx, &entity); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to update Realm object",
			}
			return res, err
		}
	}

	// Prepare response
	res.Entity = FromEntity(&entity)

	// Return expected result
	return res, nil
}

func (s *service) Delete(ctx context.Context, req *realmv1.GetRequest) (*realmv1.SingleResponse, error) {
	res := &realmv1.SingleResponse{}

	// Prepare expected results
	var entity models.Realm

	// Check request
	if req == nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Chapter must exists
		constraints.RealmMustExists(s.realms, req.Id, &entity),
	); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	if err := s.realms.Delete(ctx, req.Id); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to delete Realm object",
		}
		return res, err
	}

	// Return expected result
	return res, nil
}

func (s *service) Search(ctx context.Context, req *realmv1.SearchRequest) (*realmv1.PaginatedResponse, error) {
	res := &realmv1.PaginatedResponse{}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
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
	entities, total, err := s.realms.Search(ctx, filter, pagination, sortParams)
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
		res.Members = FromCollection(entities)
	}

	// Return results
	return res, nil
}
