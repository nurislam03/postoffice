package conn

import (
	"fmt"
	"github.com/nurislam03/postoffice/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var postgresDB *gorm.DB
var postgresOnce = &sync.Once{}

// PostgresServer ...
func PostgresServer(cfg *config.PostgresDB) *gorm.DB {
	postgresOnce.Do(func() {
		var err error
		logrus.Info("Connecting Postgres DB...")
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.DBSSLMode)

		dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logrus.Error("Postgres DB connection failed", err)
		}
		postgresDB = dbConn
		logrus.Info("Postgres DB connected")
	})
	return postgresDB
}
