package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func ReadFile() (string, error) {
	if _, err := os.Stat(OneBuildConfigFileName); os.IsNotExist(err) {
		return "", errors.New("no '" + OneBuildConfigFileName + "' file found in current directory")
	}
	yamlFile, err := ioutil.ReadFile(OneBuildConfigFileName)
	if err != nil {
		return "", errors.New("error in reading '" + OneBuildConfigFileName + "' configuration file")
	}
	return string(yamlFile), nil
}

func IsConfigFilePresent() bool {
	if _, err := os.Stat(OneBuildConfigFileName); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func WriteConfigFile(configuration OneBuildConfiguration) error {
	yamlData, _ := yaml.Marshal(&configuration)
	content := string(yamlData)
	return ioutil.WriteFile(OneBuildConfigFileName, []byte(content), 0777)
}
