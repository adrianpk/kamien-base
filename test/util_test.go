package test

import (
	"fmt"
	"path"
	"strings"

	"github.com/adrianpk/kamien"
	log "github.com/siddontang/go/log"
)

const (
	developmentEnv = "development"
	testEnv        = "test"
	productionEnv  = "production"
)

func getAppName() string {
	return "{{.AppNameLowercase}}"
}

func getAppAssetsDir(appName, env string) (appAssetsDir string, err error) {
	// In dev mode user project values.
	if isDevelopmentMode(env) {
		return getAppAssetsDirDev(), nil
	}
	// In test or production mode.
	// First try to get app assets dir from environment var.
	log.Infof("Looking for app assets in directory pointed by environment var.")
	appAssetsDir, err = getAppAssetsDirEnv(appName)
	if err != nil {
		// As a second try to get them from user's home: ~/<.appname>
		return getAppAssetsDirUser(appName)
	}
	return appAssetsDir, nil
}

func getAppAssetsDirDev() (appAssetsDir string) {
	log.Info("App is running on development mode.")
	return path.Join(".", "resources")
}

func getAppAssetsDirEnv(appName string) (appAssetsDir string, err error) {
	appEnvVar := fmt.Sprintf("%s_HOME", strings.ToUpper(appName))
	log.Infof("App home environment var name: '%s'", appEnvVar)
	appAssetsDir, err = kamien.Env.GetEnvValue(appEnvVar)
	logErr(err, fmt.Sprintf("'%s' environment var not set.", appEnvVar))
	return appAssetsDir, err
}

func getAppAssetsDirUser(appName string) (appAssetsDir string, err error) {
	log.Infof("Looking for app assets in user's home app assets directory.")
	userHome, err := kamien.Env.GetHomePath()
	logErr(err, fmt.Sprintf("Cannot determine user's home directory."))
	appAssetsDir = fmt.Sprintf("%s/.%s", userHome, appName)
	return appAssetsDir, err
}

func isDevelopmentMode(env string) bool {
	return env == developmentEnv
}

func notInProduction(env string) bool {
	return env != productionEnv
}

func checkErr(err error, message ...string) {
	if err != nil {
		log.Info("Test execution not completed.")
		if message[0] != "" {
			log.Fatalf("%s:\n %v", message[0], err)
			return
		}
		log.Fatalf("%s", err.Error())
	}
}

func logErr(err error, message ...string) {
	if err != nil {
		if message[0] != "" {
			log.Debugf("%s:\n %v", message[0], err)
			return
		}
		log.Debugf("%s", err.Error())
	}
}
