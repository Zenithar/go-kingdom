package postgresql

import (
	"context"

	"github.com/gobuffalo/packr"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"

	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/pkg/log"
)

const (
	// RealmTableName defines realm entity table's name
	RealmTableName = "realms"
	// UserTableName defines user entity table's name
	UserTableName = "users"
)

// ----------------------------------------------------------

// RepositorySet exposes Google Wire providers
var RepositorySet = wire.NewSet(
	AutoMigrate,
	NewUserRepository,
	NewRealmRepository,
)

// ----------------------------------------------------------

//go:generate packr

// migrations contains all schema migrations
var migrations = &migrate.PackrMigrationSource{
	Box: packr.NewBox("./migrations"),
}

// CreateSchemas create or updates the current database schema
func CreateSchemas(conn *sqlx.DB) (int, error) {
	// Migrate schema
	migrate.SetTable("schema_migration")

	n, err := migrate.Exec(conn.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not migrate sql schema, applied %d migrations", n)
	}

	return n, nil
}

// AutoMigrate provider for auto schema migration feature
func AutoMigrate(ctx context.Context, cfg *db.Configuration) (*sqlx.DB, error) {
	// Initialize database connection
	conn, err := db.Connection(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.AutoMigrate {
		log.For(ctx).Info("Migrating database schema ...")

		_, err := CreateSchemas(conn)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to migrate database schema")
		}
	}

	// No error
	return conn, nil
}
