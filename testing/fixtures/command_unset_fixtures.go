package fixtures

import (
	"github.com/gopinath-langote/1buildgo/testing/def"
	"github.com/gopinath-langote/1buildgo/testing/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func FeatureUnsetTestsData() []Test {
	feature := "unset"

	return []Test{
		shouldUnsetTheExistingCommand(feature),
		UnsetShouldFailWhenConfigurationFileIsNotFound(feature),
		UnsetShouldFailWhenConfigurationFileIsInInvalidFormat(feature),
		UnsetShouldFailWhenCommandIsNotFound(feature),
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

func UnsetShouldFailWhenCommandIsNotFound(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "UnsetShouldFailWhenCommandIsNotFound",
		CmdArgs: []string{"unset", "test"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Command 'test' not found")
		},
	}
}

func UnsetShouldFailWhenConfigurationFileIsNotFound(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "UnsetShouldFailWhenConfigurationFileIsNotFound",
		CmdArgs: []string{"unset", "build"},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found in current directory")
		},
	}
}

func UnsetShouldFailWhenConfigurationFileIsInInvalidFormat(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "UnsetShouldFailWhenConfigurationFileIsInInvalidFormat",
		CmdArgs: []string{"unset", "build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid config content")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Unable to parse '"+def.ConfigFileName+"' config file. Make sure file is in correct format.")
		},
	}
}
