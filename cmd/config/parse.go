package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/viper"

	"gopkg.in/yaml.v3"
)

const (
	// BeforeCommand contains before command definition
	BeforeCommand = "before"

	// AfterCommand contains after command definition
	AfterCommand = "after"
)

// OneBuildConfigFileName one global declaration of config file name
var OneBuildConfigFileName = "1build.yaml"

// OneBuildConfiguration is a representation of yaml configuration as struct
type OneBuildConfiguration struct {
	Project  string              `yaml:"project"`
	Before   string              `yaml:"before,omitempty"`
	After    string              `yaml:"after,omitempty"`
	Commands []map[string]string `yaml:"commands"`
}

// GetConfigFile returns the 1build configuration file from root file flag or global file variable
func GetConfigFile() string {
	fileFlag := viper.GetString("file")
	if fileFlag == "" {
		return OneBuildConfigFileName
	}

	return fileFlag
}

// LoadOneBuildConfiguration returns the config from file as struct.
// If there is an error, it will be of type *Error.
func LoadOneBuildConfiguration() (OneBuildConfiguration, error) {
	var configuration OneBuildConfiguration
	fileContent, err := ReadFile()
	if err != nil {
		return OneBuildConfiguration{}, err
	}
	yamlError := yaml.Unmarshal([]byte(fileContent), &configuration)
	if yamlError != nil {
		message :=
			`Sample format is:

------------------------------------------------------------------------
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
------------------------------------------------------------------------
`
		message = "Unable to parse '" + OneBuildConfigFileName + "' config file. Make sure file is in correct format.\n" +
			message
		return configuration, errors.New(message)
	}
	return configuration, nil
}

// GetCommand return command by name
func (oneBuildConfiguration *OneBuildConfiguration) GetCommand(name string) (value string) {
	for _, command := range oneBuildConfiguration.Commands {
		for k, v := range command {
			if k == name {
				return v
			}
		}
	}
	return
}

// Print prints the configuration to the console
func (oneBuildConfiguration *OneBuildConfiguration) Print() {
	fmt.Println(utils.Dash() + "\nproject: " + oneBuildConfiguration.Project)
	if oneBuildConfiguration.Before != "" {
		fmt.Println("before: " + oneBuildConfiguration.Before)
	}
	if oneBuildConfiguration.After != "" {
		fmt.Println("after: " + oneBuildConfiguration.After)
	}
	fmt.Println("commands:")

	for _, command := range oneBuildConfiguration.Commands {
		for k, v := range command {
			fmt.Println(strings.TrimSpace(k) + " | " + strings.TrimSpace(v))
		}
	}
	fmt.Println(utils.Dash())
}
