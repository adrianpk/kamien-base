package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Only for package initialization.
)

type (
	// Config - Database configuration parameters.
	Config struct {
		Host, DB, User, Pass, SSL string
	}
)

// DBConfig holds the configuration values from configuration file.
var DBConfig Config

// GetDb - Returns a *sql.DB Database connection pool.
func GetDb() (*sql.DB, error) {
	db, err := GetDbx()
	return db.DB, err
}

// GetDbx - Returns a *sqlx.DB Database connection pool.
func GetDbx() (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", DBConfig.User, DBConfig.Pass, DBConfig.DB, DBConfig.SSL)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
