/**
 * Copyright (c) 2018 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package boot

import (
	"github.com/adrianpk/kamien/db"
	"{{.Package}}/migrations"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // Required by github.com/golang-migrate/mien/db"
	"github.com/golang-migrate/migrate/source/go_bindata"
)

// Steps, Up or Down instead.

func initMigrationOrRollback() {
	dbConfig()
	switch Configuration.Migration() {
	case "migrate":
		initMigration()
	case "migrate_n":
		steps := Configuration.MigrationSteps()
		initMigrateN(steps)
	case "rollback_n":
		steps := Configuration.MigrationSteps()
		initRollbackN(steps)
	case "rollback_all":
		initRollbackAll()
	case "migrate_to_version":
		version := Configuration.MigrationToVersion()
		initToVersion(version)
	case "drop":
		initDrop()
	default:
		// Do nothing
	}
}

func initMigration() {
	log.Debug("Migration init...")
	m := getMigrator()
	err := m.Up()
	checkErr(err)
	log.Debug("Migration completed.")
}

func initMigrateN(steps int) {
	log.Debugf("Migrating %d steps up...", steps)
	if steps < 0 {
		steps = -1 * (steps)
	}
	err := getMigrator().Steps(steps)
	checkErr(err)
	log.Debug("Migration completed.")
}

func initRollbackN(steps int) {
	log.Debugf("Migrating %d steps up...", steps)
	if steps > 0 {
		steps = -1 * (steps)
	}
	err := getMigrator().Steps(steps)
	checkErr(err)
	log.Debug("Migration completed.")
}

func initRollbackAll() {
	log.Debug("Rollback all init...")
	err := getMigrator().Down()
	checkErr(err)
	log.Debug("Rollback completed.")
}

func initToVersion(version int) {
	log.Debug("Migrate to version %d init...", version)
	err := getMigrator().Migrate(uint(version))
	checkErr(err)
	log.Debug("Migration completed.")
}

func initDrop() {
	log.Debug("Dropping database")
	err := getMigrator().Drop()
	checkErr(err)
	log.Debug("Drop completed.")
}

func getMigrator() *migrate.Migrate {
	assets := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	d, _ := bindata.WithInstance(assets)
	m, _ := migrate.NewWithSourceInstance("go-bindata", d, db.GetConnectionString())
	return m
}
