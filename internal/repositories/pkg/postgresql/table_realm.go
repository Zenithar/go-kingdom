package postgresql

import (
	"context"
	"strings"

	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/db/adapter/postgresql"

	sq "github.com/Masterminds/squirrel"
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

func (r *pgRealmRepository) Search(ctx context.Context, filter *repositories.RealmSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Realm, int, error) {
	var results []*models.Realm

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	if len(results) == 0 {
		return results, count, db.ErrNoResult
	}

	// Return results and total count
	return results, count, nil
}

func (r *pgRealmRepository) FindByLabel(ctx context.Context, label string) (*models.Realm, error) {
	var entity models.Realm

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"label": label,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

// -----------------------------------------------------------------------------

func (r *pgRealmRepository) buildFilter(filter *repositories.RealmSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.RealmID)) > 0 {
			clauses["realm_id"] = filter.RealmID
		}
		if len(strings.TrimSpace(filter.Label)) > 0 {
			clauses["label"] = filter.Label
		}

		return clauses
	}

	return nil
}
