package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadFile returns OneBuildConfiguration file content as string.
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

// IsConfigFilePresent return whether the config file present or not
func IsConfigFilePresent() bool {
	if _, err := os.Stat(OneBuildConfigFileName); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// WriteConfigFile writes config to the file
// If there is an error, it will be of type *Error.
func WriteConfigFile(configuration OneBuildConfiguration) error {
	yamlData, _ := yaml.Marshal(&configuration)
	content := string(yamlData)
	return ioutil.WriteFile(OneBuildConfigFileName, []byte(content), 0777)
}

// DeleteConfigFile deletes the config file
func DeleteConfigFile() error {
	return os.Remove(OneBuildConfigFileName)
}
