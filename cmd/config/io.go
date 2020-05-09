package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadFile returns OneBuildConfiguration file content as string.
func ReadFile(oneBuildConfigFile string) (string, error) {
	if _, err := os.Stat(oneBuildConfigFile); os.IsNotExist(err) {
		return "", errors.New("no '" + oneBuildConfigFile + "' file found in current directory")
	}
	yamlFile, err := ioutil.ReadFile(oneBuildConfigFile)
	if err != nil {
		return "", errors.New("error in reading '" + oneBuildConfigFile + "' configuration file")
	}
	return string(yamlFile), nil
}

// IsConfigFilePresent return whether the config file present or not
func IsConfigFilePresent(oneBuildConfigFile string) bool {
	if _, err := os.Stat(oneBuildConfigFile); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// WriteConfigFile writes config to the file
// If there is an error, it will be of type *Error.
func WriteConfigFile(configuration OneBuildConfiguration, oneBuildConfigFile string) error {
	yamlData, _ := yaml.Marshal(&configuration)
	content := string(yamlData)
	return ioutil.WriteFile(oneBuildConfigFile, []byte(content), 0777)
}

// DeleteConfigFile deletes the config file
func DeleteConfigFile(oneBuildConfigFile string) error {
	return os.Remove(oneBuildConfigFile)
}
