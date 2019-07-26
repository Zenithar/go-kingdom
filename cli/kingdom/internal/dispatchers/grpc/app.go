package grpc

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"

	"go.zenithar.org/kingdom/cli/kingdom/internal/config"
	"go.zenithar.org/pkg/log"
)

type application struct {
	cfg    *config.Configuration
	server *grpc.Server
}

var (
	app  *application
	once sync.Once
)

// -----------------------------------------------------------------------------

// New initialize the application
func New(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	var err error

	once.Do(func() {
		// Initialize application
		app = &application{}

		// Apply configuration
		if err := app.ApplyConfiguration(cfg); err != nil {
			log.For(ctx).Fatal("Unable to initialize server settings", log.Error(err))
		}

		// Initialize Core components
		switch cfg.Core.Type {
		case "postgresql":
			app.server, err = setupLocalPostgreSQL(ctx, cfg)
		}
	})

	// Return server
	return app.server, err
}

// -----------------------------------------------------------------------------

// ApplyConfiguration apply the configuration after checking it
func (s *application) ApplyConfiguration(cfg interface{}) error {
	// Check configuration validity
	if err := s.checkConfiguration(cfg); err != nil {
		return err
	}

	// Apply to current component (type assertion done if check)
	s.cfg, _ = cfg.(*config.Configuration)

	// No error
	return nil
}

// -----------------------------------------------------------------------------

func (s *application) checkConfiguration(cfg interface{}) error {
	// Check via type assertion
	config, ok := cfg.(*config.Configuration)
	if !ok {
		return fmt.Errorf("server: invalid configuration")
	}

	switch config.Core.Type {
	case "postgresql":
		if config.Core.Hosts == "" {
			return fmt.Errorf("server: database hosts list is mandatory")
		}
	default:
		return fmt.Errorf("server: invalid type (postgresql)")
	}

	// No error
	return nil
}
