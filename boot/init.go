/**
 * Copyright (c) 2018 Adrian K <adrian.git@kuguar.dev>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package boot

import (
	"github.com/adrianpk/kamien" // Initialize logger
	"github.com/adrianpk/kamien/boot"
)

const (
	developmentEnv = "development"
	testEnv        = "test"
	productionEnv  = "production"
)

var (
	// Log - App main logger
	log = kamien.Log
	// AppHomeName - App home name
	AppHomeName = "{{.AppNameUppercase}}_HOME"
	// Env - Current environment. Default: development
	Env = kamien.Env.GetEnvType("{{.AppNameUppercase}}_HOME", developmentEnv) // development, test, production.
	// Env = testEnv
	// Configuration keeps the configuration values from config.yaml file
	Configuration EnvConfiguration
	// {{.AppNamePascalCase}}Context - {{.AppNamePascalCase}} application context.
	{{.AppNamePascalCase}}Context *boot.AppContext
	// AssetsDir - {{.AppNamePascalCase}} assets directory.
	AssetsDir *string
)

func init() {
	initConfiguration()
	initLogger()
	initKeys()
}

// Init - Init
func Init() {
	initDB()
	initMigrationOrRollback()
	initSeeding()
}

// InitForTest - Init triggered by tests
func InitForTest(configuration *EnvConfiguration) {
	Configuration = *configuration
	initLogger()
	initDB()
	initAppContext()
}
