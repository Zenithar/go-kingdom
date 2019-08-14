package user

import (
	"context"
	"crypto/cipher"
	"encoding/base64"
	"net/http"

	"go.zenithar.org/kingdom/internal/helpers"
	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/kingdom/internal/services/internal/constraints"
	"go.zenithar.org/kingdom/internal/services/internal/value"
	aestransformer "go.zenithar.org/kingdom/internal/services/internal/value/encrypt/aes"
	apiv1 "go.zenithar.org/kingdom/internal/services/pkg/v1"
	sysv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/system/v1"
	userv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/user/v1"
	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
)

type service struct {
	users       repositories.User
	transformer value.Transformer
}

// PasswordBlock is an alias to cipher.Block for wiring
type PasswordBlock cipher.Block

// New services instance
func New(users repositories.User, passwordBlock PasswordBlock) apiv1.User {
	// Initialize transformer
	return &service{
		users:       users,
		transformer: aestransformer.NewCBCTransformer(passwordBlock),
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	res := &userv1.CreateResponse{}

	// Check request
	if req == nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, errors.Newf(errors.InvalidArgument, nil, "request must not be nil")
	}

	// Hash principal
	principal := helpers.PrincipalHashFunc(req.Principal)

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Principal must be unique
		constraints.UserPrincipalMustBeUnique(s.users, req.RealmId, principal),
	); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	// Prepare entity creation
	entity := models.NewUser(req.RealmId)

	// Update attributes
	entity.Principal = principal

	secret, err := helpers.DerivePasswordFunc(req.Secret)
	if err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to hash password",
		}
		return res, errors.Newf(errors.Internal, err, "unable to hash password")
	}

	// Encrypt secret (at-rest encryption)
	out, err := s.transformer.TransformToStorage([]byte(secret))
	if err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to transform password",
		}
		return res, errors.Newf(errors.Internal, err, "unable to transform password")
	}

	// Assign secret
	entity.Secret = base64.RawStdEncoding.EncodeToString(out)

	// Create entity in database
	if err := s.users.Create(ctx, entity); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create user",
		}
		return res, errors.Newf(errors.Internal, err, "unable to create entity")
	}

	// Prepare response
	res.Entity = FromEntity(entity)

	// Return result
	return res, nil
}

func (s *service) Get(ctx context.Context, req *userv1.GetRequest) (*userv1.GetResponse, error) {
	res := &userv1.GetResponse{}

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

	// Retrieve user from database
	entity, err := s.users.Get(ctx, req.RealmId, req.UserId)
	if err != nil && err != db.ErrNoResult {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve User",
		}
		return res, errors.Newf(errors.Internal, err, "unable to retrieve entity")
	}
	if entity == nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
		return res, errors.Newf(errors.NotFound, err, "entity not found")
	}

	// Prepare response
	res.Entity = FromEntity(entity)

	// Return result
	return res, nil
}

func (s *service) Update(ctx context.Context, req *userv1.UpdateRequest) (*userv1.UpdateResponse, error) {
	res := &userv1.UpdateResponse{}
	// Prepare expected results
	var entity models.User

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
		// User must exists
		constraints.UserMustExists(s.users, req.RealmId, req.UserId, &entity),
	); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	updated := false

	// Skip operation when no updates
	if updated {
		// Create account in database
		if err := s.users.Update(ctx, &entity); err != nil {
			res.Error = &sysv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to update User object",
			}
			return res, errors.Newf(errors.Internal, err, "unable to retrieve entity")
		}
	}

	// Prepare response
	res.Entity = FromEntity(&entity)

	// Return expected result
	return res, nil
}

func (s *service) Delete(ctx context.Context, req *userv1.DeleteRequest) (*userv1.DeleteResponse, error) {
	res := &userv1.DeleteResponse{}

	// Prepare expected results
	var entity models.User

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
		constraints.UserMustExists(s.users, req.RealmId, req.UserId, &entity),
	); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	if err := s.users.Delete(ctx, req.RealmId, req.UserId); err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to delete User object",
		}
		return res, errors.Newf(errors.Internal, err, "unable to delete entity")
	}

	// Return expected result
	return res, nil
}

func (s *service) Search(ctx context.Context, req *userv1.SearchRequest) (*userv1.SearchResponse, error) {
	res := &userv1.SearchResponse{}

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
	filter := &repositories.UserSearchFilter{}
	if req.UserId != nil {
		filter.UserID = req.UserId.Value
	}
	if req.RealmId != nil {
		filter.RealmID = req.RealmId.Value
	}
	if req.Principal != nil {
		filter.Principal = helpers.PrincipalHashFunc(req.Principal.Value)
	}

	// Do the search
	entities, total, err := s.users.Search(ctx, filter, pagination, sortParams)
	if err != nil && err != db.ErrNoResult {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to process request",
		}
		return res, errors.Newf(errors.Internal, err, "unable to delete entity")
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

func (s *service) Authenticate(ctx context.Context, req *userv1.AuthenticateRequest) (*userv1.AuthenticateResponse, error) {
	res := &userv1.AuthenticateResponse{}

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

	// Retrieve user from database
	entity, err := s.users.FindByPrincipal(ctx, req.RealmId, helpers.PrincipalHashFunc(req.Principal))
	if err != nil && err != db.ErrNoResult {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve User",
		}
		return res, errors.Newf(errors.Internal, err, "unable to retrieve entity")
	}
	if entity == nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
		return res, errors.Newf(errors.NotFound, err, "entity not found")
	}

	// Assign secret
	secret, err := base64.RawStdEncoding.DecodeString(entity.Secret)
	if err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to decode password",
		}
		return res, errors.Newf(errors.Internal, err, "unable to transform password")
	}

	// Decrypt secret (at-rest encryption)
	out, _, err := s.transformer.TransformFromStorage(secret)
	if err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to transform password",
		}
		return res, errors.Newf(errors.Internal, err, "unable to transform password")
	}

	// Check password
	valid, err := helpers.CheckPasswordFunc(string(out), req.Secret)
	if err != nil {
		res.Error = &sysv1.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve User",
		}
		return res, errors.Newf(errors.Internal, err, "unable to retrieve entity")
	}
	if !valid {
		res.Error = &sysv1.Error{
			Code:    http.StatusUnauthorized,
			Message: "Authentication failed",
		}
		return res, errors.Newf(errors.Unauthenticated, nil, "authentication failed")
	}

	// Return user
	res.Identity = FromEntity(entity)

	// Return result
	return res, nil
}
