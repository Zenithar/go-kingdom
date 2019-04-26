package repositories

import (
	"context"

	"go.zenithar.org/kingdom/internal/models"
	"go.zenithar.org/pkg/db"
)

//go:generate mockgen -destination test/mock/realm.gen.go -package mock go.zenithar.org/kingdom/internal/repositories Realm

// RealmSearchFilter represents realm entity collection search criteria
type RealmSearchFilter struct {
	RealmID string
	Label   string
}

// Realm repository defines realm management contract
type Realm interface {
	Create(ctx context.Context, entity *models.Realm) error
	Get(ctx context.Context, id string) (*models.Realm, error)
	Update(ctx context.Context, entity *models.Realm) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filter *RealmSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Realm, int, error)
	FindByLabel(ctx context.Context, label string) (*models.Realm, error)
}

//go:generate mockgen -destination test/mock/user.gen.go -package mock go.zenithar.org/kingdom/internal/repositories User

// UserSearchFilter represents user entity collection search criteria
type UserSearchFilter struct {
	RealmID   string
	UserID    string
	Principal string
}

// User repository defines user management contract
type User interface {
	Create(ctx context.Context, entity *models.User) error
	Get(ctx context.Context, realmID string, id string) (*models.User, error)
	Update(ctx context.Context, entity *models.User) error
	Delete(ctx context.Context, realmID string, id string) error
	Search(ctx context.Context, filter *UserSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.User, int, error)
	FindByPrincipal(ctx context.Context, realmID string, principal string) (*models.User, error)
}
