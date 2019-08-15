package config

import (
	"errors"
	"github.com/gopinath-langote/1buildgo/cmd/utils"
	"gopkg.in/yaml.v3"
	"strings"
)

var OneBuildConfigFileName = "1build.yaml"

type OneBuildConfiguration struct {
	Project  string              `yaml:"project"`
	Before   string              `yaml:"before,omitempty"`
	After    string              `yaml:"after,omitempty"`
	Commands []map[string]string `yaml:"commands"`
}

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
