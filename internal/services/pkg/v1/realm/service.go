package realm

import (
	"context"

	"go.zenithar.org/kingdom/internal/repositories"
	apiv1 "go.zenithar.org/kingdom/internal/services/pkg/v1"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/realm/internal/commands"
	realmv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/realm/v1"
)

type service struct {
	createCmd commands.HandlerFunc
	getCmd    commands.HandlerFunc
	updateCmd commands.HandlerFunc
	deleteCmd commands.HandlerFunc
	searchCmd commands.HandlerFunc
}

// New services instance
func New(realms repositories.Realm) apiv1.Realm {
	return &service{
		createCmd: commands.CreateCommand(realms),
		getCmd:    commands.GetCommand(realms),
		updateCmd: commands.UpdateCommand(realms),
		deleteCmd: commands.DeleteCommand(realms),
		searchCmd: commands.SearchCommand(realms),
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *realmv1.CreateRequest) (*realmv1.CreateResponse, error) {
	res, err := s.createCmd(ctx, req)
	return res.(*realmv1.CreateResponse), err
}

func (s *service) Get(ctx context.Context, req *realmv1.GetRequest) (*realmv1.GetResponse, error) {
	res, err := s.getCmd(ctx, req)
	return res.(*realmv1.GetResponse), err
}

func (s *service) Update(ctx context.Context, req *realmv1.UpdateRequest) (*realmv1.UpdateResponse, error) {
	res, err := s.updateCmd(ctx, req)
	return res.(*realmv1.UpdateResponse), err
}

func (s *service) Delete(ctx context.Context, req *realmv1.DeleteRequest) (*realmv1.DeleteResponse, error) {
	res, err := s.deleteCmd(ctx, req)
	return res.(*realmv1.DeleteResponse), err
}

func (s *service) Search(ctx context.Context, req *realmv1.SearchRequest) (*realmv1.SearchResponse, error) {
	res, err := s.searchCmd(ctx, req)
	return res.(*realmv1.SearchResponse), err
}
