package fixtures

import (
	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func featureUnsetTestsData() []Test {
	feature := "unset"

	return []Test{
		shouldUnsetTheExistingCommand(feature),
		unsetShouldFailWhenConfigurationFileIsNotFound(feature),
		unsetShouldFailWhenConfigurationFileIsInInvalidFormat(feature),
		unsetShouldFailWhenCommandIsNotFound(feature),
	}
}

func shouldUnsetTheExistingCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetTheExistingCommand",
		CmdArgs: []string{"unset", "build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func unsetShouldFailWhenCommandIsNotFound(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "unsetShouldFailWhenCommandIsNotFound",
		CmdArgs: []string{"unset", "Test"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Command 'Test' not found")
		},
	}
}

func unsetShouldFailWhenConfigurationFileIsNotFound(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "unsetShouldFailWhenConfigurationFileIsNotFound",
		CmdArgs: []string{"unset", "build"},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found in current directory")
		},
	}
}

func unsetShouldFailWhenConfigurationFileIsInInvalidFormat(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "unsetShouldFailWhenConfigurationFileIsInInvalidFormat",
		CmdArgs: []string{"unset", "build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid config content")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Unable to parse '"+def.ConfigFileName+"' config file. Make sure file is in correct format.")
		},
	}
}
