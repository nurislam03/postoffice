package api

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func (a *API) PushToPublisher(name, body string) {
	ch, err := a.amqpServer.Channel()
	if err != nil {
		logrus.Error("Failed to open a RabbitMQ channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
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

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		logrus.Error("Failed to publish the message", err)
	}

	logrus.Info("Successfully Publish Message to Queue")
}
