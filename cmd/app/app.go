// This file holds the setup for RBAC app server

package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	port     string
	loglevel uint8
	mode     string
	cors     bool
	cluster  bool

	AppCmd = &cobra.Command{
		Use:   "app",
		Short: "run api server",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			authorMessage()
			initConfig()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return startengine()
		},
	}
)

func init() {
	AppCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config/in-local.yaml", "Supply config file to setup RBAC app")
	AppCmd.PersistentFlags().StringVarP(&port, "port", "p", "8081", "TCP Port RBAC to listen")
	AppCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "debug", "GIN RUN MODE")
	AppCmd.PersistentFlags().Uint8VarP(&loglevel, "loglevel", "l", 0, "Log Level")
	AppCmd.PersistentFlags().BoolVarP(&cors, "cors", "x", false, "Enable cross origin resouce sharing")
	AppCmd.PersistentFlags().BoolVarP(&cluster, "cluster", "s", false, "cluster alone mode or distributed mode")
}

func authorMessage() {
	author := `Developed by ahadiwasti.com`
	fmt.Println("\n", author)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".reacting-auth" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".reacting-auth")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func setup() {
	// SET LOG LEVEL
	fmt.Println("\n", "SET LOG LEVEL: ", loglevel)
	zerolog.SetGlobalLevel(zerolog.Level(loglevel))

	// SET RUN MODE
	fmt.Println("\n", "SET RUN MODE: ", mode)
	gin.SetMode(mode)

}

func startengine() error {
	engine := gin.Default()
	return engine.Run(":" + port)
}
