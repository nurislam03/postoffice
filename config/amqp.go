package config

import (
	"errors"
	"github.com/spf13/viper"
	"sync"
)

// AMQP ...
type AMQP struct {
	URI string
}

func (cnf *AMQP) validate() error {
	if cnf.URI == "" {
		return errors.New("AMQP uri cannot be empty")
	}
	return nil
}

var amqpCnf *AMQP
var amqpOnce = &sync.Once{}

func loadamqp() {
	amqpCnf = &AMQP{
		URI: viper.GetString("AMQP_SERVER_URI"),
	}
}

// AMQPCnf ...
func AMQPCnf() *AMQP {
	amqpOnce.Do(func() {
		loadamqp()
		err := amqpCnf.validate(); if err != nil {
			panic(err)
		}
	})
	return amqpCnf
}

