package cmd

import (
	"os"

	"github.com/ahadiwasti/reacting-auth/cmd/api"

	// "zeus/cmd/migrate"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "Reacting-Auth",
	Short:             "Auth API server",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `Start Auth API server`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	// rootCmd.AddCommand(migrate.MigrateCmd)
}

//Execute : run commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
