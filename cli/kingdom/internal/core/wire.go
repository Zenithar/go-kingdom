package core

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"io"

	"github.com/google/wire"
	"gocloud.dev/secrets"

	"go.zenithar.org/kingdom/cli/kingdom/internal/config"
	"go.zenithar.org/kingdom/internal/repositories/pkg/postgresql"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/realm"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/user"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/pkg/log"
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

// UserPasswordBlock is the block cipher for password at-rest encryption
func UserPasswordBlock(ctx context.Context, cfg *config.Configuration) (user.PasswordBlock, error) {
	// Open a *secrets.Keeper using the keeperURL.
	keeper, err := secrets.OpenKeeper(ctx, cfg.Security.KeeperURL)
	if err != nil {
		return nil, err
	}
	defer func(closer io.Closer) {
		log.SafeClose(closer, "Unable to close keeper")
	}(keeper)

	// Decode password encryption key
	secret, err := base64.StdEncoding.DecodeString(cfg.Security.PasswordKey)
	if err != nil {
		return nil, err
	}

	// Decrypt password encryption key
	passwordKey, err := keeper.Decrypt(ctx, secret)
	if err != nil {
		return nil, err
	}

	// Return block
	return aes.NewCipher(passwordKey)
}

// -----------------------------------------------------------------------------

var v1ServiceSet = wire.NewSet(
	UserPasswordBlock,
	user.New,
	realm.New,
)

// -----------------------------------------------------------------------------

// LocalPostgreSQLSet initialize the PGSQL Core context
var LocalPostgreSQLSet = wire.NewSet(
	pgRepositorySet,
	v1ServiceSet,
)
