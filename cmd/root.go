package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// cfgFile store the configuration file name
	cfgFile string

	// RootCmd is the root command of the server
	RootCmd = &cobra.Command{
		Use:   "kafkapusher",
		Short: "kafkapusher is a API server",
		Long:  "kafkapusher is a API server",
	}
)

func init() {
	// cobra.OnInitialize(initConfig)
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
	viper.AutomaticEnv()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
}
