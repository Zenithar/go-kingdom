package core

import (
	"github.com/google/wire"

	"go.zenithar.org/kingdom/cli/kingdom/internal/config"
	"go.zenithar.org/kingdom/internal/repositories/pkg/postgresql"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/realm"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/user"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"
)

// -----------------------------------------------------------------------------

// PosgreSQLConfig declares a Database configuration provider for Wire
func PosgreSQLConfig(cfg *config.Configuration) *pgdb.Configuration {
	return &pgdb.Configuration{
		AutoMigrate:      cfg.Core.AutoMigrate,
		ConnectionString: cfg.Core.Hosts,
		Username:         cfg.Core.Username,
		Password:         cfg.Core.Password,
	}
}

var pgRepositorySet = wire.NewSet(
	PosgreSQLConfig,
	postgresql.RepositorySet,
)

// -----------------------------------------------------------------------------

var v1ServiceSet = wire.NewSet(
	user.New,
	realm.New,
)

// -----------------------------------------------------------------------------

// LocalPostgreSQLSet initialize the PGSQL Core context
var LocalPostgreSQLSet = wire.NewSet(
	pgRepositorySet,
	v1ServiceSet,
)
