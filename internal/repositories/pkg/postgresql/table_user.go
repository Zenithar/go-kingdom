package postgresql

import (
	"context"
	
	"go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgUserRepository struct {
	adapter *postgresql.Default
}

// NewUserRepository returns a Postgresql user management repository instance
func NewUserRepository(session *sqlx.DB) repositories.User {
	// Default columns to retrieve
	defaultColumns := []string{"realm_id", "user_id", "principal", "secret", "creation_date"}

	// Sortable columns for criteria
	sortableColumns := []string{"realm_id", "user_id", "principal", "creation_date"}

	// Initialize repository
	return &pgUserRepository{
		adapter: postgresql.NewCRUDTable(session, "", UserTableName, defaultColumns, sortableColumns),
	}
}

// -----------------------------------------------------------------------------

func (r *pgUserRepository) Create(ctx context.Context, entity *models.User) error {
	return r.adapter.Create(ctx, entity)
}

func (r *pgUserRepository) Get(ctx context.Context, realmID, id string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"realm_id": realmID,
		"user_id": id,
	}, &entity); err != nil {
		return nil, err
	}
	
	return &entity, nil
}

func (r *pgUserRepository) Update(ctx context.Context, entity *models.User) error {
	return r.adapter.Update(ctx, map[string]interface{}{
		"secret": entity.Secret,
	}, map[string]interface{}{
		"realm_id": entity.RealmID,
		"user_id": entity.ID,
	})
}

func (r *pgUserRepository) Delete(ctx context.Context, realmID, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"realm_id": realmID,
		"user_id": id,
	})
}

func (r *pgUserRepository) FindByPrincipal(ctx context.Context, realmID string, principal string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"realm_id": realmID,
		"principal": principal,
	}, &entity); err != nil {
		return nil, err
	}
	
	return &entity, nil
}