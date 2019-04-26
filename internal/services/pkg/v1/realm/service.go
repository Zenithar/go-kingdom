package realm

import (
	"context"

	"go.zenithar.org/kingdom/internal/repositories"
	apiv1 "go.zenithar.org/kingdom/internal/services/pkg/v1"
	realmv1 "go.zenithar.org/kingdom/pkg/protocol/kingdom/realm/v1"
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

	// Return result
	return &res, nil
}

func (s *service) Get(ctx context.Context, req *realmv1.GetRequest) (*realmv1.SingleResponse, error) {
	res := &realmv1.SingleResponse{}

	// Return result
	return &res, nil
}

func (s *service) Update(ctx context.Context, req *realmv1.UpdateRequest) (*realmv1.SingleResponse, error) {
	res := &realmv1.SingleResponse{}

	// Return result
	return &res, nil
}

func (s *service) Delete(ctx context.Context, req *realmv1.GetRequest) (*realmv1.SingleResponse, error) {
	res := &realmv1.SingleResponse{}

	// Return result
	return &res, nil
}

func (s *service) Search(ctx context.Context, req *realmv1.SearchReq) (*realmv1.PaginatedResponse, error) {
	res := &realmv1.PaginatedResponse{}

	// Return results
	return res, nil
}
