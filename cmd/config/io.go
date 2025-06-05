package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// ReadFile returns OneBuildConfiguration file content as string.
func ReadFile() (string, error) {
	oneBuildConfigFile := getConfigFile()
	if _, err := os.Stat(oneBuildConfigFile); os.IsNotExist(err) {
		return "", errors.New("no '" + oneBuildConfigFile + "' file found")
	}
	yamlFile, err := ioutil.ReadFile(oneBuildConfigFile)
	if err != nil {
		return "", errors.New("error in reading '" + oneBuildConfigFile + "' configuration file")
	}
	return string(yamlFile), nil
}

// IsConfigFilePresent return whether the config file present or not
func IsConfigFilePresent() bool {
	oneBuildConfigFile := getConfigFile()
	if _, err := os.Stat(oneBuildConfigFile); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

// WriteConfigFile writes config to the file
// If there is an error, it will be of type *Error.
func WriteConfigFile(configuration OneBuildConfiguration) error {
	oneBuildConfigFile := getConfigFile()

	// Deep copy configuration and remove all raw fields
	cleanConfig := configuration
	for i, cmdMap := range cleanConfig.Commands {
		for name, def := range cmdMap {
			cleanConfig.Commands[i][name] = def
		}
	}

	yamlData, _ := yaml.Marshal(&cleanConfig)
	content := string(yamlData)
	return ioutil.WriteFile(oneBuildConfigFile, []byte(content), 0750)
}

// DeleteConfigFile deletes the config file
func DeleteConfigFile() error {
	oneBuildConfigFile := getConfigFile()
	return os.Remove(oneBuildConfigFile)
}

// GetAbsoluteDirPathOfConfigFile gets the base directory from the configuration file location
func GetAbsoluteDirPathOfConfigFile() (string, error) {
	oneBuildConfigFile := getConfigFile()
	abs, err := filepath.Abs(oneBuildConfigFile)
	if err != nil {
		return "", errors.New("error in resolving file path for '" + oneBuildConfigFile + "' configuration file.")
	}
	baseDirFromAbs := filepath.Dir(abs)
	return baseDirFromAbs, nil
}

// getConfigFile returns the 1build configuration file from root file flag or global file variable
func getConfigFile() string {
	fileFlag := viper.GetString("file")
	if fileFlag == "" {
		return OneBuildConfigFileName
	}

	return fileFlag
}
