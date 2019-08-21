package config

import (
	"errors"
	"strings"

	"github.com/gopinath-langote/1build/cmd/utils"
	"gopkg.in/yaml.v3"
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

--------------------------------------------------
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
--------------------------------------------------
`
		message = "Unable to parse '" + OneBuildConfigFileName + "' config file. Make sure file is in correct format.\n" +
			message
		return configuration, errors.New(message)
	}
	return configuration, nil
}

// PrintConfiguration prints the configuration to the console
func PrintConfiguration(oneBuildConfiguration OneBuildConfiguration) {
	utils.Println(utils.DASH())
	utils.Println("project: " + oneBuildConfiguration.Project)
	if oneBuildConfiguration.Before != "" {
		utils.Println("before: " + oneBuildConfiguration.Before)
	}
	if oneBuildConfiguration.After != "" {
		utils.Println("after: " + oneBuildConfiguration.After)
	}
	utils.Println("commands:")

	for _, command := range oneBuildConfiguration.Commands {
		for k, v := range command {
			utils.Println(strings.TrimSpace(k) + " | " + strings.TrimSpace(v))
		}
	}

	utils.Println(utils.DASH())
}
