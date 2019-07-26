package v1

import (
	"context"

	realmv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/realm/v1"
	userv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/user/v1"
)

// Realm management service contract
type Realm interface {
	Create(ctx context.Context, req *realmv1.CreateRequest) (res *realmv1.CreateResponse, err error)
	Get(ctx context.Context, req *realmv1.GetRequest) (res *realmv1.GetResponse, err error)
	Update(ctx context.Context, req *realmv1.UpdateRequest) (res *realmv1.UpdateResponse, err error)
	Delete(ctx context.Context, req *realmv1.DeleteRequest) (res *realmv1.DeleteResponse, err error)
	Search(ctx context.Context, req *realmv1.SearchRequest) (res *realmv1.SearchResponse, err error)
}

// User management service contract
type User interface {
	Create(ctx context.Context, req *userv1.CreateRequest) (res *userv1.CreateResponse, err error)
	Get(ctx context.Context, req *userv1.GetRequest) (res *userv1.GetResponse, err error)
	Update(ctx context.Context, req *userv1.UpdateRequest) (res *userv1.UpdateResponse, err error)
	Delete(ctx context.Context, req *userv1.DeleteRequest) (res *userv1.DeleteResponse, err error)
	Search(ctx context.Context, req *userv1.SearchRequest) (res *userv1.SearchResponse, err error)
	Authenticate(ctx context.Context, req *userv1.AuthenticateRequest) (res *userv1.AuthenticateResponse, err error)
}
