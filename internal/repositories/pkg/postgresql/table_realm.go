package postgresql

import (
	"context"
	
	"go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgRealmRepository struct {
	adapter *postgresql.Default
}

// NewRealmRepository returns a Postgresql realm management repository instance
func NewRealmRepository(session *sqlx.DB) repositories.Realm {
	// Default columns to retrieve
	defaultColumns := []string{"realm_id", "label", "creation_date"}

	// Sortable columns for criteria
	sortableColumns := []string{"realm_id", "label", "creation_date"}

	// Initialize repository
	return &pgRealmRepository{
		adapter: postgresql.NewCRUDTable(session, "", RealmTableName, defaultColumns, sortableColumns),
	}
}

// -----------------------------------------------------------------------------

func (r *pgRealmRepository) Create(ctx context.Context, entity *models.Realm) error {
	return r.adapter.Create(ctx, entity)
}

func (r *pgRealmRepository) Get(ctx context.Context, id string) (*models.Realm, error) {
	var entity models.Realm

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"realm_id": id,
	}, &entity); err != nil {
		return nil, err
	}
	
	return &entity, nil
}

func (r *pgRealmRepository) Update(ctx context.Context, entity *models.Realm) error {
	return r.adapter.Update(ctx, map[string]interface{}{
		"label": entity.Label,
	}, map[string]interface{}{
		"realm_id": entity.ID,
	})
}

func (r *pgRealmRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"realm_id": id,
	})
}
