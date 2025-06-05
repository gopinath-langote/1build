package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gopinath-langote/1build/cmd/utils"
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

// CommandDefinition supports both inline and nested command definitions
type CommandDefinition struct {
	Before  string `yaml:"before,omitempty"`
	Command string `yaml:"command,omitempty"`
	After   string `yaml:"after,omitempty"`
}

// Custom unmarshal to support both string and map
func (c *CommandDefinition) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode {
		// Inline command: build: "go build"
		c.Command = value.Value
		return nil
	}
	if value.Kind == yaml.MappingNode {
		// Nested command: build: { before: "...", command: "...", after: "..." }
		type Alias CommandDefinition
		var a Alias
		if err := value.Decode(&a); err != nil {
			return err
		}
		*c = CommandDefinition(a)
		return nil
	}
	return nil
}

// OneBuildConfiguration is a representation of yaml configuration as struct
type OneBuildConfiguration struct {
	Project   string                         `yaml:"project"`
	BeforeAll string                         `yaml:"beforeAll,omitempty"`
	AfterAll  string                         `yaml:"afterAll,omitempty"`
	Commands  []map[string]CommandDefinition `yaml:"commands"`
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
  - build:
      before: echo "before"
      command: npm run build
      after: echo "after"
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
				return v.Command
			}
		}
	}
	return
}

// Print prints the configuration to the console
func (oneBuildConfiguration *OneBuildConfiguration) Print() {
	fmt.Println(utils.Dash() + "\nproject: " + oneBuildConfiguration.Project)
	if oneBuildConfiguration.BeforeAll != "" {
		fmt.Println("beforeAll: " + oneBuildConfiguration.BeforeAll)
	}
	if oneBuildConfiguration.AfterAll != "" {
		fmt.Println("afterAll: " + oneBuildConfiguration.AfterAll)
	}
	fmt.Println("commands:")

	for _, command := range oneBuildConfiguration.Commands {
		for k, v := range command {
			fmt.Println(strings.TrimSpace(k) + " | " + strings.TrimSpace(v.Command))
		}
	}
	fmt.Println(utils.Dash())
}
