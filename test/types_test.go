package test

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"{{.Package}}/app"
	"{{.Package}}/boot"
	"{{.Package}}/migrations"
	"github.com/golang-migrate/migrate"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/jmoiron/sqlx"
	testfixtures "gopkg.in/testfixtures.v2"
)

type (
	// Environment - Todo: compelete comment
	Environment struct {
		*boot.EnvConfiguration
		// DBConnString   string
		DBxInstance    *sqlx.DB
		Fixtures       *testfixtures.Context
		ServerInstance *httptest.Server
		// Reader         io.Reader //Ignore this for now
	}
)

// MakeEnvironment - Todo: complete comment.
func MakeEnvironment(ec *boot.EnvConfiguration) *Environment {
	return &Environment{ec, nil, nil, nil}
}

// Start - Todo: complete comment.
func (e *Environment) Start(m *testing.M) {
	e.prepareDatabaseInstance()
	// e.clearDatabaseData()
	// e.processFixtures()
	app.InitServer()
}

// PrepareDBInstance - Todo: complete comment.
func (e *Environment) prepareDatabaseInstance() {
	var err error
	e.DBxInstance, err = sqlx.Open("postgres", e.GetDBConnString())
	checkErr(err, "Cannot get a connection to database.")
}

func (e *Environment) clearDatabase() {
	e.clearDatabaseData()
}

func (e *Environment) resetDatabase() {
	e.clearDatabaseData()
	e.processFixtures()
}

func (e *Environment) clearDatabaseData() {
	cs := e.GetDBConnectionString()
	err := getMigrator(cs).Down()
	checkErr(err, "Cannot clean database data.")
	err = getMigrator(cs).Up()
	checkErr(err, "Cannot migrate database.")
}

func (e *Environment) AuthorizeRequest(req *http.Request, user, username, role string) {
	token, _ := boot.GenerateJWT(user, username, role)
	var bearer = "Bearer " + token
	req.Header.Add("authorization", bearer)
}

func getMigrator(connectionString string) *migrate.Migrate {
	assets := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	d, err := bindata.WithInstance(assets)
	checkErr(err, "Cannot obtain migration assets.")
	m, err := migrate.NewWithSourceInstance("go-bindata", d, connectionString)
	checkErr(err, "Cannot get migrator.")
	return m
}

// GetDBConnString - Todo: complete comment.
func (e *Environment) GetDBConnString() string {
	dbcp := e.GetDBConnectionParameters()
	return fmt.Sprintf("postgresql://%s:%s/%s?user=%s&password=%s&sslmode=%s", dbcp["DBHost"], dbcp["DBPort"], dbcp["DBName"], dbcp["DBUser"], dbcp["DBPass"], dbcp["DBSSL"])
}

// GetDBInstance - Todo: complete comment.
func (e *Environment) GetDBInstance() *sql.DB {
	return e.DBxInstance.DB
}

// GetDBxInstance - Todo: complete comment.
func (e *Environment) GetDBxInstance() *sqlx.DB {
	return e.DBxInstance
}

// GetDBxTx - Todo: complete comment.
func (e *Environment) GetDBxTx() *sqlx.Tx {
	return e.DBxInstance.MustBegin()
}

// GetFixturesDir - Todo: complete comment.
// func (e *Environment) GetFixturesDir() string {
// 	return e.FixturesDir
// }

// GetFixtures - Todo: complete comment.
func (e *Environment) GetFixtures() *testfixtures.Context {
	return e.Fixtures
}

// GetServerInstance - Todo: complete comment.
func (e *Environment) GetServerInstance() *httptest.Server {
	return e.ServerInstance
}

// GetReader - Todo: complete comment.
// func (e *Environment) GetReader() io.Reader {
// 	return e.Reader
// }

// GetAPIPath - Todo: complete comment.
// func (e *Environment) GetAPIPath() string {
// 	return e.APIPath
// }

// GetAPIVersion - Todo: complete comment.
// func (e *Environment) GetAPIVersion() string {
// 	return e.APIVersion
// }

// // GetAPIServerURL - Todo: complete comment.
// // func (e *Environment) GetAPIServerURL() string {
// 	return e.APIServerURL
// }

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [DEBUG] " + string(bytes))
}
