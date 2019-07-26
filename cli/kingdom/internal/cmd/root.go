package cmd

import (
	iconfig "go.zenithar.org/kingdom/cli/kingdom/internal/config"

	"github.com/spf13/cobra"
	"go.zenithar.org/pkg/config"
	cmdcfg "go.zenithar.org/pkg/config/cmd"
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
)

// -----------------------------------------------------------------------------

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "kingdom",
	Short: "Realm and user management microservice",
}

func init() {
	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(cmdcfg.NewConfigCommand(conf, "KNDM"))
	mainCmd.AddCommand(serverCmd)
	mainCmd.AddCommand(clientCmd)
}

// -----------------------------------------------------------------------------

// Execute main command
func Execute() error {
	feature.DefaultMutableGate.AddFlag(mainCmd.Flags())
	return mainCmd.Execute()
}

// -----------------------------------------------------------------------------

var (
	cfgFile string
	conf    = &iconfig.Configuration{}
)

// -----------------------------------------------------------------------------

func initConfig() {
	if err := config.Load(conf, "KNDM", cfgFile); err != nil {
		log.Bg().Fatal("Unable load config", log.Error(err))
	}
}
