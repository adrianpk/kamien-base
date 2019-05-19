/**
 * Copyright (c) 2018 Adrian K <adrian.git@kuguar.dev>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package boot

import (
	"github.com/adrianpk/kamien/db"
	_ "github.com/golang-migrate/migrate/database/postgres" // Required by github.com/golang-migrate/mien/db"
)

func initDB() {
	dbConfig()
}

func dbConfig() {
	dbConfig := Configuration.GetDBConnectionParameters()
	db.DBConfig.Host = dbConfig["DBHost"]
	db.DBConfig.DB = dbConfig["DBName"]
	db.DBConfig.User = dbConfig["DBUser"]
	db.DBConfig.Pass = dbConfig["DBPass"]
	db.DBConfig.SSL = dbConfig["DBSSL"]
}
