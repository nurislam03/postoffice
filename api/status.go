package api

import (
	"github.com/sirupsen/logrus"
)

func (a *API) Consumer(name string) {
	ch, err := a.amqpServer.Channel()
	if err != nil {
		logrus.Error("Failed to open a RabbitMQ channel", err)
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

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			logrus.Info("Received a message: %s", m.Body)
		}
	}()

	logrus.Info(" [*] Waiting for a message. To exit press CTRL+C")

	<-forever
}
