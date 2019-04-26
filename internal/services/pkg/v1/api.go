package v1

import (
	"context"

	realmv1 "go.zenithar.org/kingdom/pkg/protocol/kingdom/realm/v1"
	userv1 "go.zenithar.org/kingdom/pkg/protocol/kingdom/user/v1"
)

// Realm management service contract
type Realm interface {
	Create(ctx context.Context, req *realmv1.CreateRequest) (res *realmv1.SingleResponse, err error)
	Get(ctx context.Context, req *realmv1.GetRequest) (res *realmv1.SingleResponse, err error)
	Update(ctx context.Context, req *realmv1.UpdateRequest) (res *realmv1.SingleResponse, err error)
	Delete(ctx context.Context, req *realmv1.GetRequest) (res *realmv1.SingleResponse, err error)
}

// User management service contract
type User interface {
	Create(ctx context.Context, req *userv1.CreateRequest) (res *userv1.SingleResponse, err error)
	Get(ctx context.Context, req *userv1.GetRequest) (res *userv1.SingleResponse, err error)
	Update(ctx context.Context, req *userv1.UpdateRequest) (res *userv1.SingleResponse, err error)
	Delete(ctx context.Context, req *userv1.GetRequest) (res *userv1.SingleResponse, err error)
}
