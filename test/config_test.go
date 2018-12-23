package test

import (
	"fmt"
	"io/ioutil"
	"path"

	"{{.Package}}/boot"
	log "github.com/siddontang/go/log"
	yaml "gopkg.in/yaml.v2"
)

var (
	envConf *boot.EnvConfiguration
	tenv    *Environment
)

// GetConfiguration - Server configuration for S.
func GetConfiguration() *boot.EnvConfiguration {
	appName := getAppName()
	appAssetsDir, err := getAppAssetsDir(appName, currentEnv)
	checkErr(err, "Cannot infer app assets directory.")
	err = loadConfig(appAssetsDir, currentEnv)
	checkErr(err, "Cannot load configuration.")
	log.Infof("Config: %v", envConf)
	return envConf
}

func loadConfig(appAssetsDir, env string) error {
	configFile := fmt.Sprintf("%s.yaml", env)
	fullConfigFilePath := path.Join(appAssetsDir, "config", configFile)
	return loadConfigurationFromFile(env, fullConfigFilePath)
}

func loadConfigurationFromFile(env, filePath string) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Warnf("Error opening '%s'", filePath)
		return err
	}
	err = yaml.Unmarshal(fileBytes, &envConf)
	if err != nil {
		return err
	}
	return nil
}
