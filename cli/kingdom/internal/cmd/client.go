package cmd

import (
	"github.com/spf13/cobra"

	realmv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/realm/v1"
	userv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/user/v1"
)

// -----------------------------------------------------------------------------

var clientCmd = &cobra.Command{
	Use:     "client",
	Aliases: []string{"c", "cli"},
	Short:   "Query the gRPC server",
}

func init() {
	clientCmd.AddCommand(
		realmv1.RealmAPIClientCommand,
		userv1.UserAPIClientCommand,
	)
}
