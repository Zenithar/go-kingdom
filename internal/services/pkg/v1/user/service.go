package user

import (
	"context"

	"go.zenithar.org/kingdom/internal/repositories"
	apiv1 "go.zenithar.org/kingdom/internal/services/pkg/v1"
	userv1 "go.zenithar.org/kingdom/pkg/protocol/kingdom/user/v1"
)

type service struct {
	users repositories.User
}

// New returns a service instance
func New(users repositories.User) apiv1.User {
	return &service{
		users: users,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.SingleResponse, error) {
	res := &userv1.SingleResponse{}

	// Return result
	return res, nil
}

func (s *service) Get(ctx context.Context, req *userv1.GetRequest) (*userv1.SingleResponse, error) {
	res := &userv1.SingleResponse{}

	// Return result
	return res, nil
}

func (s *service) Update(ctx context.Context, req *userv1.UpdateRequest) (*userv1.SingleResponse, error) {
	res := &userv1.SingleResponse{}

	// Return result
	return res, nil
}

func (s *service) Delete(ctx context.Context, req *userv1.GetRequest) (*userv1.SingleResponse, error) {
	res := &userv1.SingleResponse{}

	// Return result
	return res, nil
}

func (s *service) Search(ctx context.Context, req *userv1.SearchRequest) (*userv1.PaginatedResponse, error) {
	res := &userv1.PaginatedResponse{}

	// Return results
	return res, nil
}
