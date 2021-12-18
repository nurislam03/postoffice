package conn

import (
	"fmt"
	"github.com/nurislam03/postoffice/config"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

/*
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()
*/

var amqpConn *amqp.Connection
var amqpOnce = &sync.Once{}

func loadAMQPConn() {
	log.Println("Connecting AMQP connection...")
	log.Println("AMQP connection successful")
}

// AMQPServer ...
func AMQPServer(cfg *config.AMQP) *amqp.Connection {
	var err error
	amqpOnce.Do(func() {
		loadAMQPConn()
		amqpConn, err = amqp.Dial(cfg.URI)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to RabbitMQ: %s", err))
		}

		//defer producer.Close()  //ToDo: Need to pass it through go channel.
		println("AMQP Connected to Server")
	})
	return amqpConn
}
