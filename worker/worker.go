package worker

import (
	"fmt"
	"github.com/nurislam03/postoffice/model"
	"github.com/nurislam03/postoffice/repo"
	"github.com/nurislam03/postoffice/services"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Worker struct {
	amqpServer *amqp.Connection
	sRepo      repo.StatusRepo
}

//NewWorker ...
func NewWorker(w *amqp.Connection, sRepo repo.StatusRepo) *Worker {
	return &Worker{
		amqpServer: w,
		sRepo:      sRepo,
	}
}

func (w *Worker) saveData(id string) error {
	sts, err := services.TesterService().GetStatusByID(id)
	if err != nil {
		logrus.Error("Online Status Retrieve Failed", err)
		return err
	}

	pld := &model.Status{
		ID:       fmt.Sprintf("%d", sts.ID),
		Online:   sts.Online,
		LastSeen: time.Now(),
	}

	err = w.sRepo.Upsert(pld)
	if err != nil {
		logrus.Error("Internal Server Error", err)
		return err
	}

	return nil
}

func (w *Worker) Run(count int) {
	name := "get-status"

	ch, err := w.amqpServer.Channel()
	if err != nil {
		logrus.Error("Failed to open RabbitMQ channel", err)
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logrus.Error("Failed to declare queue", err)
	}

	msgs, err := ch.Consume(
		name,
		"",
		false,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		logrus.Error("Failed to register a consumer", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	for i := 1; i <= count; i++ {
		go func(id int) {
			logrus.Info("Starting worker: ", id)
			for m := range msgs {
				logrus.Info("Received a message: ", string(m.Body))
				err := w.saveData(string(m.Body))
				if err != nil {
					logrus.Error("Internal Server Error", err)
					m.Nack(false, true)
				} else {
					m.Ack(false)
				}
			}
		}(i)
	}

	logrus.Info(" [*] Waiting for a message. To exit press CTRL+C")

	<-stop
}
