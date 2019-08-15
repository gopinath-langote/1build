package fixtures

import (
	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func FeatureSetTestsData() []Test {
	feature := "set"

	return []Test{
		shouldSetNewCommand(feature),
		shouldUpdateExistingCommand(feature),
		shouldFailWhenConfigurationFileIsNotFound(feature),
		shouldFailWhenConfigurationFileIsInInvalidFormat(feature),
	}
}

func shouldSetNewCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build: go build
  - test: go test
`

	return Test{
		Feature: feature,
		Name:    "shouldSetNewCommand",
		CmdArgs: []string{"set", "test", "go test"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldUpdateExistingCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build: go build -o
`

	return Test{
		Feature: feature,
		Name:    "shouldUpdateExistingCommand",
		CmdArgs: []string{"set", "build", "go build -o"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldFailWhenConfigurationFileIsNotFound(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldFailWhenConfigurationFileIsNotFound",
		CmdArgs: []string{"set", "build", "go build -o"},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found in current directory")
		},
	}
}

func shouldFailWhenConfigurationFileIsInInvalidFormat(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldFailWhenConfigurationFileIsInInvalidFormat",
		CmdArgs: []string{"set", "build", "go build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid config content")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Unable to parse '"+def.ConfigFileName+"' config file. Make sure file is in correct format.")
		},
	}
}
