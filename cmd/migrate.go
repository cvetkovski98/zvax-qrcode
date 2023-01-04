package cmd

import (
	"log"

	"github.com/cvetkovski98/zvax-common/pkg/postgresql"
	"github.com/cvetkovski98/zvax-qrcode/internal/config"
	"github.com/cvetkovski98/zvax-qrcode/internal/model/migrations"
	"github.com/spf13/cobra"
)

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Long:  `Migrate database`,
	RunE:  migrate,
}

func migrate(cmd *cobra.Command, args []string) error {
	cfg := config.GetConfig()
	db, err := postgresql.NewPgDb(&cfg.PostgreSQL)
	if err != nil {
		return err
	}
	if err = postgresql.Migrate(cmd.Context(), db, migrations.Migrations); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
