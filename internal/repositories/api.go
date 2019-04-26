package repositories

import (
	"context"

	"go.zenithar.org/kingdom/internal/models"
)

// Realm repository defines realm management contract
type Realm interface {
	Create(ctx context.Context, entity *models.Realm) error
	Get(ctx context.Context, id string) (*models.Realm, error)
	Update(ctx context.Context, entity *models.Realm) error
	Delete(ctx context.Context, id string) error
}

// User repository defines user management contract
type User interface {
	Create(ctx context.Context, entity *models.User) error
	Get(ctx context.Context, realmID string, id string) (*models.User, error)
	Update(ctx context.Context, entity *models.User) error
	Delete(ctx context.Context, realmID string, id string) error
	FindByPrincipal(ctx context.Context, realmID string, principal string) (*models.User, error)
}
