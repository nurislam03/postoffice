package worker

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
)

type Worker struct {
	amqpServer *amqp.Connection
}

//NewWorker ...
func NewWorker(w *amqp.Connection) *Worker {
	return &Worker{
		amqpServer: w,
	}
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
		true,
		false,
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
				logrus.Info("Received a message: ",  string(m.Body))
				//Todo : database write
			}
		}(i)
	}

	logrus.Info(" [*] Waiting for a message. To exit press CTRL+C")

	<-stop
}
