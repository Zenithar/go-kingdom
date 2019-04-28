package cmd

import "github.com/spf13/cobra"

// -----------------------------------------------------------------------------

var configMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database schema using configuration",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
