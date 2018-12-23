/**
 * Copyright (c) 2018 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package boot

import (
	"database/sql"
	"strings"

	"github.com/adrianpk/kamien/db"
	"{{.Package}}/seeds"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/romanyx/polluter"
)

func initSeeding() {
	dbConfig()
	if notInProduction(Env) {
		switch Configuration.Seed() {
		case "seed":
			seed()
		case "reset":
			panic("To be implemented")
		case "reset-and-seed":
			panic("To be implemented")
		default:
			// Do nothing
		}
	} else {
		// Todo: 'xxxx' is a CLI command to be implemented.
		log.Warn("You need to explicitly run 'xxxx' to seed in production.")
	}
}

func seed() {
	log.Debug("Seeding...")
	db, err := db.GetDb()
	if err != nil {
		log.Fatalf("Cannot get a database connection: %s", err)
	}
	assets := bindata.Resource(seeds.AssetNames(),
		func(name string) ([]byte, error) {
			return seeds.Asset(name)
		})
	for _, an := range assets.Names {
		data, err := seeds.Asset(an)
		if err != nil {
			log.Debugf("Error loading asset %s", an)
		} else {
			log.Debugf("Processing %s", an)
			seedAsset(db, data)
			log.Debugf("%s processed.", an)
		}
	}
	log.Debug("Seeding completed.")
}

func seedAsset(db *sql.DB, assetData []byte) {
	p := polluter.New(polluter.PostgresEngine(db), polluter.YAMLParser)
	s := string(assetData)
	err := p.Pollute(strings.NewReader(s))
	if err != nil {
		log.Fatalf("Cannot seed the database: %s", err)
	}
}
