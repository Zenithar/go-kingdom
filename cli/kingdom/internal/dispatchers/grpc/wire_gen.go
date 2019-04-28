// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package grpc

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"go.zenithar.org/kingdom/cli/kingdom/internal/config"
	"go.zenithar.org/kingdom/cli/kingdom/internal/core"
	"go.zenithar.org/kingdom/internal/repositories/pkg/postgresql"
	"go.zenithar.org/kingdom/internal/services/pkg/v1"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/realm"
	"go.zenithar.org/kingdom/internal/services/pkg/v1/user"
	"go.zenithar.org/kingdom/pkg/protocol/kingdom/realm/v1"
	"go.zenithar.org/kingdom/pkg/protocol/kingdom/user/v1"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/tlsconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Injectors from wire.go:

func setupLocalPostgreSQL(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	configuration := core.PosgreSQLConfig(cfg)
	db, err := postgresql.AutoMigrate(ctx, configuration)
	if err != nil {
		return nil, err
	}
	repositoriesUser := postgresql.NewUserRepository(db)
	v1User := user.New(repositoriesUser)
	repositoriesRealm := postgresql.NewRealmRepository(db)
	v1Realm := realm.New(repositoriesRealm)
	server, err := grpcServer(ctx, cfg, v1User, v1Realm)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// wire.go:

func grpcServer(ctx context.Context, cfg *config.Configuration, users v1.User, realms v1.Realm) (*grpc.Server, error) {
	sopts := []grpc.ServerOption{}
	grpc_zap.ReplaceGrpcLogger(zap.L())

	sopts = append(sopts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_zap.StreamServerInterceptor(zap.L()), grpc_recovery.StreamServerInterceptor())), grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_recovery.UnaryServerInterceptor(), grpc_zap.UnaryServerInterceptor(zap.L()))), grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)

	if cfg.Server.UseTLS {

		clientAuth := tls.VerifyClientCertIfGiven
		if cfg.Server.TLS.ClientAuthenticationRequired {
			clientAuth = tls.RequireAndVerifyClientCert
		}

		tlsConfig, err := tlsconfig.Server(tlsconfig.Options{
			KeyFile:    cfg.Server.TLS.PrivateKeyPath,
			CertFile:   cfg.Server.TLS.CertificatePath,
			CAFile:     cfg.Server.TLS.CACertificatePath,
			ClientAuth: clientAuth,
		})
		if err != nil {
			log.For(ctx).Error("Unable to build TLS configuration from settings", zap.Error(err))
			return nil, err
		}

		sopts = append(sopts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	} else {
		log.For(ctx).Info("No transport authentication enabled")
	}

	server := grpc.NewServer(sopts...)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	realmv1.RegisterRealmAPIServer(server, realms)
	userv1.RegisterUserAPIServer(server, users)
	reflection.Register(server)

	err := view.Register(ochttp.ServerRequestCountView, ochttp.ServerRequestBytesView, ochttp.ServerResponseBytesView, ochttp.ServerLatencyView, ochttp.ServerRequestCountByMethod, ochttp.ServerResponseCountByStatusCode)
	if err != nil {
		log.For(ctx).Fatal("Unable to register HTTP stat views", zap.Error(err))
	}

	err = view.Register(ocgrpc.DefaultServerViews...)
	if err != nil {
		log.For(ctx).Fatal("Unable to register gRPC stat views", zap.Error(err))
	}

	return server, nil
}
