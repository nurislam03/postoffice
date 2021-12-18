package cmd

import (
	migrationcmd "github.com/nurislam03/postoffice/cmd/migration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// cfgFile store the configuration file name
	cfgFile string

	// RootCmd is the root command of the server
	RootCmd = &cobra.Command{
		Use:   "postoffice",
		Short: "postoffice is a API server",
		Long:  "postoffice is a API server",
	}
)

func init() {
	// cobra.OnInitialize(initConfig)
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
	viper.AutomaticEnv()
	RootCmd.AddCommand(migrationcmd.MigrateCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
}
