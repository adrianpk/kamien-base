/**
 * Copyright (c) 2018 Adrian K <adrian.git@kuguar.dev>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package boot

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	testfixtures "gopkg.in/testfixtures.v2"
)

type (
	// EnvInterface - Environment interface
	EnvInterface interface {
		GetServerAddressPort() string
		GetServerAddress() string
		GetServerPort() string
		GetServerSSLPort() string
		GetDomains() []string
		GetDBConnectionParameters() map[string]string // DBServer, Database, DBUser, DBPass, DBSSL string
		GetBaseDir() string
		GetResourcesDir() string
		GetPublicDir() string
		GetLogFile() string
		GetLogLevel() int
		IsAutoreloadOn() bool
		Migration() string
		Seed() string
		// Used for testing
		Start(m *testing.M)
		PrepareTestDatabase()
		AuthorizeRequest(req *http.Request, user, username, role string)
		GetDBConnString() string
		GetDBInstance() *sql.DB
		GetFixtures() *testfixtures.Context
		GetServerInstance() *httptest.Server
		GetReader() io.Reader //Ignore this for now
		GetAPIPath() string
		GetAPIVersion() string
		GetAPIServerURL() string
	}

	// EnvConfiguration - Environment configuration
	EnvConfiguration struct {
		Server struct {
			Address      string   `yaml:"address"`
			Port         string   `yaml:"port"`
			SSLPort      string   `yaml:"SSLPort"`
			ResourcesDir string   `yaml:"resourcesDir"`
			PublicDir    string   `yaml:"publicDir"`
			LogFile      string   `yaml:"logFile"`
			LogLevel     string   `yaml:"logLevel"`
			Autoreload   bool     `yaml:"autoReload"`
			Domains      []string `yaml:"domains"`
		}
		Database struct {
			Address  string `yaml:"address"`
			Port     string `yaml:"port"`
			Name     string `yaml:"name"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			SSL      string `yaml:"SSL"`
		}
		Tasks struct {
			Migration struct {
				MigrationType string `yaml:"mgType"`
				Steps         int    `yaml:"steps"`
				ToVersion     int    `yaml:"migrateToVersion"`
			}
			Seed string `yaml:"seed"`
		}
	}
)

func (ec EnvConfiguration) GetServerAddressPort() string {
	return fmt.Sprintf("%s:%s", ec.Server.Address, ec.Server.SSLPort)
}

func (ec EnvConfiguration) GetServerAddress() string {
	return ec.Server.Address
}

func (ec EnvConfiguration) GetServerPort() string {
	return ec.Server.Port
}

func (ec EnvConfiguration) GetServerSSLPort() string {
	return ec.Server.SSLPort
}

func (ec EnvConfiguration) GetDomains() []string {
	return ec.Server.Domains
}

func (ec EnvConfiguration) GetDBConnectionParameters() map[string]string {
	dbConf := make(map[string]string)
	dbConf["DBHost"] = ec.Database.Address
	dbConf["DBPort"] = ec.Database.Port
	dbConf["DBName"] = ec.Database.Name
	dbConf["DBUser"] = ec.Database.User
	dbConf["DBPass"] = ec.Database.Password
	dbConf["DBSSL"] = ec.Database.SSL
	return dbConf
}

// GetDBConnectionString - Returns the database connection string.
func (ec EnvConfiguration) GetDBConnectionString() string {
	return fmt.Sprintf(
		"postgresql://%s/%s?user=%s&password=%s&dbname=%s&sslmode=%s",
		ec.Database.Address,
		ec.Database.Port,
		ec.Database.User,
		ec.Database.Password,
		ec.Database.Name,
		ec.Database.SSL)
}

func (ec EnvConfiguration) GetResourcesDir() string {
	return ec.Server.ResourcesDir
}

func (ec EnvConfiguration) GetLogFile() string {
	return ec.Server.LogFile
}

func (ec EnvConfiguration) GetLogLevel() string {
	return ec.Server.LogLevel
}

func (ec EnvConfiguration) IsAutoreloadOn() bool {
	return ec.Server.Autoreload
}

func (ec EnvConfiguration) Migration() string {
	return ec.Tasks.Migration.MigrationType
}

func (ec EnvConfiguration) MigrationSteps() int {
	return ec.Tasks.Migration.Steps
}

func (ec EnvConfiguration) MigrationToVersion() int {
	return ec.Tasks.Migration.ToVersion
}

func (ec EnvConfiguration) Seed() string {
	return ec.Tasks.Seed
}
