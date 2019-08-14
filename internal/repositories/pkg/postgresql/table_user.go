package postgresql

import (
	"context"
	"strings"

	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/kingdom/internal/repositories"
	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/db/adapter/postgresql"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
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
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Create(ctx, entity)
}

func (r *pgUserRepository) Get(ctx context.Context, realmID, id string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"realm_id": realmID,
		"user_id":  id,
	}, &entity); err != nil {
		return nil, err
	}

	spew.Dump(entity)

	return &entity, nil
}

func (r *pgUserRepository) Update(ctx context.Context, entity *models.User) error {
	return r.adapter.Update(ctx, map[string]interface{}{
		"secret": entity.Secret,
	}, map[string]interface{}{
		"realm_id": entity.RealmID,
		"user_id":  entity.ID,
	})
}

func (r *pgUserRepository) Delete(ctx context.Context, realmID, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"realm_id": realmID,
		"user_id":  id,
	})
}

func (r *pgUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.User, int, error) {
	var results []*models.User

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

func (r *pgUserRepository) FindByPrincipal(ctx context.Context, realmID string, principal string) (*models.User, error) {
	var entity models.User

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"realm_id":  realmID,
		"principal": principal,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

// -----------------------------------------------------------------------------

func (r *pgUserRepository) buildFilter(filter *repositories.UserSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.RealmID)) > 0 {
			clauses["realm_id"] = filter.RealmID
		}
		if len(strings.TrimSpace(filter.UserID)) > 0 {
			clauses["user_id"] = filter.UserID
		}
		if len(strings.TrimSpace(filter.Principal)) > 0 {
			clauses["principal"] = filter.Principal
		}

		return clauses
	}

	return nil
}
