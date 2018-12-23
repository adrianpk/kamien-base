package boot

import (
	"fmt"
	"io/ioutil"
	"path"

	yaml "gopkg.in/yaml.v2"
)

// Initialize AppConfig
func initConfiguration() {
	// Intended only for initial development purposes
	// Allows you to use sample configuration from project source folder,
	// bypassing the need for initial environment variables setup
	// and/or adding '.{{.AppNameLowercase}}/config' user's home folder.
	appName := getAppName()
	dir, err := getAppAssetsDir(appName, Env)
	logErr(err, "Cannot infer app assets directory.")
	AssetsDir = &dir
	err = loadConfig(*AssetsDir, Env)
	checkErr(err, "Cannot load configuration.")
}

func loadConfig(assetsDir, env string) error {
	configFile := fmt.Sprintf("%s.yaml", env)
	fullConfigFilePath := path.Join(assetsDir, "config", configFile)
	return loadConfigurationFromFile(env, fullConfigFilePath)
}

func loadConfigurationFromFile(env, filePath string) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error opening '%s'", filePath)
		return err
	}
	err = yaml.Unmarshal(fileBytes, &Configuration)
	if err != nil {
		return err
	}
	return nil
}
