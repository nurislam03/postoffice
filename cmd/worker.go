package cmd

import (
	"github.com/nurislam03/postoffice/config"
	"github.com/nurislam03/postoffice/conn"
	cwrkr "github.com/nurislam03/postoffice/worker"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start worker",
	Long:  `Start the worker server`,
	Run:   worker,
}

func init() {
	RootCmd.AddCommand(workerCmd)
}

func worker(cmd *cobra.Command, args []string) {
	count, err := strconv.Atoi(args[0])
	if err != nil || count < 1 {
		logrus.Error("concurrency should be a positive number", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go startWorker(count)

	<-stop
	log.Println("Worker server shutting down...")

	time.Sleep(5 * time.Second)
}

func startWorker(count int) {
	cfg := config.NewConfig()

	//connection
	amqpServer := conn.AMQPServer(cfg.AMQP)
	wrkr := cwrkr.NewWorker(amqpServer)

	wrkr.Run(count)
}
