package cmd

import (
	"fmt"
	"github.com/nurislam03/postoffice/api"
	"github.com/nurislam03/postoffice/backend"
	"github.com/nurislam03/postoffice/config"
	"github.com/nurislam03/postoffice/conn"
	"github.com/nurislam03/postoffice/model"
	"github.com/nurislam03/postoffice/repo"
	"github.com/nurislam03/postoffice/repo/pgrepo"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Long:  `Start the API server`,
	Run:   serve,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		portStr := viper.GetString("SERVER_PORT")
		lsnr, err := net.Listen("tcp", ":"+portStr)
		if err != nil {
			return fmt.Errorf("Port %s is not available", portStr)
		}
		_ = lsnr.Close()
		return nil
	},
}

func init() {
	// serveCmd.PersistentFlags().IntP("port", "p", 8080, "port on which the server will listen")
	// serveCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
	// viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	RootCmd.AddCommand(serveCmd)
}

func startCron(sRepo repo.StatusRepo) {
	c := cron.New()
	c.AddFunc("@every 30s", func() {
		fmt.Println("Every Thirty Seconds")
		sRepo.Expire(&model.Status{})
	})
	c.Start()
}

func serve(cmd *cobra.Command, args []string) {
	cfg := config.NewConfig()

	pgCnn := conn.PostgresServer(cfg.PostgresDB)
	sRepo := pgrepo.NewStatus(pgCnn)

	startCron(sRepo)

	amqpServer := conn.AMQPServer(cfg.AMQP)

	api := api.NewAPI(cfg, amqpServer, sRepo)

	backend.NewServer(cfg, api).Serve()
}
