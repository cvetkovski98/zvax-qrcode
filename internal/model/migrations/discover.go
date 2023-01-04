package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

//go:embed *.sql
var sql embed.FS

func init() {
	if err := Migrations.Discover(sql); err != nil {
		panic(err)
	}
}
