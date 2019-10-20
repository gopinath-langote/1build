package fixtures

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureUnsetTestsData() []Test {
	feature := "unset"

	return []Test{
		shouldUnsetTheExistingCommand(feature),
		unsetShouldFailWhenConfigurationFileIsNotFound(feature),
		unsetShouldFailWhenConfigurationFileIsInInvalidFormat(feature),
		unsetShouldFailWhenCommandIsNotFound(feature),
		shouldUnsetMultipleCommands(feature),
		shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound1(feature),
		shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound2(feature),
		shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound3(feature),
		shouldUnsetTheBeforeCommand(feature),
		shouldUnsetTheAfterCommand(feature),
		unsetBeforeShouldFailWhenBeforeIsNotFound(feature),
		unsetAfterShouldFailWhenAfterIsNotFound(feature),
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
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
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
			return assert.Contains(t, actualOutput, "Following command(s) not found: Test")
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

func shouldUnsetMultipleCommands(feature string) Test {

	defaultFileContent := `
project: Sample Project
before: go before
after: go after
commands:
  - build: go build
  - test: go test
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommands",
		CmdArgs: []string{"unset", "build", "test", "before", "after"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound1(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
  - test: go test
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound1",
		CmdArgs: []string{"unset", "build", "test", "missingCmd"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)

			testResult := assert.Contains(t, actualOutput, "Following command(s) not found: missingCmd") &&
				assert.Exactly(t, expectedOutput, string(content))

			return testResult
		},
	}
}

func shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound2(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
  - test: go test
  - lint: go lint
`

	expectedOutput := `project: Sample Project
commands:
  - lint: go lint
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound2",
		CmdArgs: []string{"unset", "build", "missingCmd", "test", "missingCmd2"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)

			testResult := assert.Contains(t, actualOutput, "Following command(s) not found: missingCmd, missingCmd2") &&
				assert.Exactly(t, expectedOutput, string(content))

			return testResult
		},
	}
}

func shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound3(feature string) Test {

	defaultFileContent := `
project: Sample Project
before: go before
after: go after
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound3",
		CmdArgs: []string{"unset", "before", "after", "missingCmd"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)

			testResult := assert.Contains(t, actualOutput, "Following command(s) not found: missingCmd") &&
				assert.Exactly(t, expectedOutput, string(content))

			return testResult
		},
	}
}

func shouldUnsetTheBeforeCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
before: yo
commands: []
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetBeforeAndAfterCommand",
		CmdArgs: []string{"unset", "before"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func shouldUnsetTheAfterCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
after: yo
commands: []
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetTheAfterCommand",
		CmdArgs: []string{"unset", "after"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := filepath.Join(dir, def.ConfigFileName)
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func unsetBeforeShouldFailWhenBeforeIsNotFound(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "unsetBeforeShouldFailWhenBeforeIsNotFound",
		CmdArgs: []string{"unset", "before"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Following command(s) not found: before")
		},
	}
}

func unsetAfterShouldFailWhenAfterIsNotFound(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "unsetAfterShouldFailWhenAfterIsNotFound",
		CmdArgs: []string{"unset", "after"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Following command(s) not found: after")
		},
	}
}
