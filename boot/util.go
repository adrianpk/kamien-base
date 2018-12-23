/**
 * Copyright (c) 2018 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package boot

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/adrianpk/kamien"
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
	log.Info("Looking for app assets in directory pointed by environment var.")
	appAssetsDir, err = getAppAssetsDirEnv(appName)
	if err != nil {
		// As a second try to get them from user's home: ~/<.appname>
		return getAppAssetsDirUser(appName)
	}
	return appAssetsDir, nil
}

func getAppAssetsDirDev() (appAssetsDir string) {
	log.Info("App is running on development mode.")
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(b))
	path := path.Join(basepath, "resources")
	return path
}

func getAppAssetsDirEnv(appName string) (appAssetsDir string, err error) {
	appEnvVar := fmt.Sprintf("%s_HOME", strings.ToUpper(appName))
	log.Info("App home environment var name: '%s'", appEnvVar)
	appAssetsDir, err = kamien.Env.GetEnvValue(appEnvVar)
	logErr(err, fmt.Sprintf("'%s' environment var not set.", appEnvVar))
	return appAssetsDir, err
}

func getAppAssetsDirUser(appName string) (appAssetsDir string, err error) {
	log.Info("Looking for app assets in user's home app assets directory.")
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

func logWarn(err error, message ...string) {
	if err != nil {
		if message[0] != "" {
			log.Printf("WARN[%s] %s:\n %v", time.Now(), message[0], err)
			return
		}
		log.Printf("WARN[%s] %s", time.Now(), err.Error())
	}
}

func logErr(err error, message ...string) {
	if err != nil {
		if message[0] != "" {
			log.Printf("ERRO[%s] %s:\n %v", time.Now(), message[0], err)
			return
		}
		log.Printf("ERRO[%s] %s", time.Now(), err.Error())
	}
}

func checkErr(err error, message ...string) {
	if err != nil {
		if message[0] != "" {
			log.Fatalf("ERRO[%s] %s:\n %v", time.Now(), message[0], err)
			return
		}
		log.Fatalf("ERRO[%s] %s", time.Now(), err.Error())
	}
}
