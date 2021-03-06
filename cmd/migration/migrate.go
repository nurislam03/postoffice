package migrationcmd

import (
	"github.com/nurislam03/postoffice/config"
	"github.com/nurislam03/postoffice/conn"
	"github.com/nurislam03/postoffice/model"
	"github.com/spf13/cobra"
	"log"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Postgres DB",
	Long:  `Migrate Postgres DB`,
	Run:   migrate,
}

func migrate(cmd *cobra.Command, args []string) {
	cfg := config.NewConfig()
	pgCnn := conn.PostgresServer(cfg.PostgresDB)

	if err := pgCnn.AutoMigrate(model.Status{}); err != nil {
		panic(err)
	}
	log.Println("Migration completed successfully!")
}
