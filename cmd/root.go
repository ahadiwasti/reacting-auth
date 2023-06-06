/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"reacting-auth/cmd/app"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "rbac",
	Short:             "RBAC app server",
	Long:              `The RBAC system ensures secure and scalable management of user roles and permissions. With Gin's powerful routing capabilities and Casbin's flexible access control library, the project delivers efficient authorization management.`,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(app.AppCmd)
}
