package postgresql

import (
	"github.com/gobuffalo/packr"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	// RealmTableName defines realm entity table's name
	RealmTableName = "realms"
	// UserTableName defines user entity table's name
	UserTableName = "users"
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

	n, err := migrate.Exec(conn.DB, conn.DriverName(), migrations, migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not migrate sql schema, applied %d migrations", n)
	}

	return n, nil
}
