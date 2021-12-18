package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
	"time"
)

// PostgresDB ..
type PostgresDB struct {
	Host       string
	Port       int64
	Name       string
	User       string
	Password   string
	DBSSLMode  string
	DBTimeZone string
	TimeOut    time.Duration
}

func (cnf *PostgresDB) validate() {
	if cnf.Host == "" {
		logrus.Error("host", "required")
	}
	if cnf.Port == 0 {
		logrus.Error("port", "required")
	}
	if cnf.Name == "" {
		logrus.Error("name", "required")
	}
	if cnf.User == "" {
		logrus.Error("user", "required")
	}
	if cnf.Password == "" {
		logrus.Error("host", "required")
	}
}

var postgresCnf *PostgresDB
var postgresOnce = &sync.Once{}

func loadMongo() {
	postgresCnf = &PostgresDB{
		Host:      viper.GetString("DB_HOST"),
		Port:      viper.GetInt64("DB_PORT"),
		Name:      viper.GetString("DB_NAME"),
		User:      viper.GetString("DB_USER"),
		Password:  viper.GetString("DB_PASSWORD"),
		DBSSLMode: viper.GetString("DB_DBSSLMODE"),
	}
}

// PostgresCnf ...
func PostgresCnf() *PostgresDB {
	postgresOnce.Do(func() {
		loadMongo()
		postgresCnf.validate()
	})
	return postgresCnf
}
