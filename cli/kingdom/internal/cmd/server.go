package cmd

import (
	"context"

	"github.com/cloudflare/tableflip"
	"github.com/dchest/uniuri"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/kingdom/cli/kingdom/internal/dispatchers/grpc"
	"go.zenithar.org/kingdom/internal/version"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"
)

// -----------------------------------------------------------------------------

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Starts the spotigraph gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Generate an instance identifier
		appID := uniuri.NewLen(64)

		// Initialize config
		initConfig()

		// Prepare logger
		log.Setup(ctx, &log.Options{
			Debug:     conf.Debug.Enable,
			AppName:   "kingdom-grpc",
			AppID:     appID,
			Version:   version.Version,
			Revision:  version.Revision,
			SentryDSN: conf.Instrumentation.Logs.SentryDSN,
		})

		// Starting banner
		log.For(ctx).Info("Starting kingdom gRPC server ...")

		// Start goroutine group
		err := platform.Run(ctx, conf.Debug.Enable, conf.Instrumentation, func(upg *tableflip.Upgrader, group run.Group) {
			ln, err := upg.Fds.Listen(conf.Server.Network, conf.Server.Listen)
			if err != nil {
				log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
			}

			server, err := grpc.New(ctx, conf)
			if err != nil {
				log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
			}

			group.Add(
				func() error {
					log.For(ctx).Info("Starting GRPC server", zap.Stringer("address", ln.Addr()))
					return server.Serve(ln)
				},
				func(e error) {
					log.For(ctx).Info("Shutting GRPC server down")
					server.GracefulStop()
				},
			)
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
